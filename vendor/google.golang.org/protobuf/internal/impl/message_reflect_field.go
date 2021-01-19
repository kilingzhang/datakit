// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package impl

import (
	"fmt"
	"math"
	"reflect"
	"sync"

	"google.golang.org/protobuf/internal/flags"
	pref "google.golang.org/protobuf/reflect/protoreflect"
	preg "google.golang.org/protobuf/reflect/protoregistry"
)

type fieldInfo struct {
	fieldDesc pref.FieldDescriptor

	// These fields are used for protobuf reflection support.
	has        func(pointer) bool
	clear      func(pointer)
	get        func(pointer) pref.Value
	set        func(pointer, pref.Value)
	mutable    func(pointer) pref.Value
	newMessage func() pref.Message
	newField   func() pref.Value
}

func fieldInfoForOneof(fd pref.FieldDescriptor, fs reflect.StructField, x exporter, ot reflect.Type) fieldInfo {
	ft := fs.Type
	if ft.Kind() != reflect.Interface {
		panic(fmt.Sprintf("field %v has invalid type: got %v, want interface kind", fd.FullName(), ft))
	}
	if ot.Kind() != reflect.Struct {
		panic(fmt.Sprintf("field %v has invalid type: got %v, want struct kind", fd.FullName(), ot))
	}
	if !reflect.PtrTo(ot).Implements(ft) {
		panic(fmt.Sprintf("field %v has invalid type: %v does not implement %v", fd.FullName(), ot, ft))
	}
	conv := NewConverter(ot.Field(0).Type, fd)
	isMessage := fd.Message() != nil

	// TODO: Implement unsafe fast path?
	fieldOffset := offsetOf(fs, x)
	return fieldInfo{
		// NOTE: The logic below intentionally assumes that oneof fields are
		// well-formatted. That is, the oneof interface never contains a
		// typed nil pointer to one of the wrapper structs.

		fieldDesc: fd,
		has: func(p pointer) bool {
			if p.IsNil() {
				return false
			}
			rv := p.Apply(fieldOffset).AsValueOf(fs.Type).Elem()
			if rv.IsNil() || rv.Elem().Type().Elem() != ot || rv.Elem().IsNil() {
				return false
			}
			return true
		},
		clear: func(p pointer) {
			rv := p.Apply(fieldOffset).AsValueOf(fs.Type).Elem()
			if rv.IsNil() || rv.Elem().Type().Elem() != ot {
				// NOTE: We intentionally don't check for rv.Elem().IsNil()
				// so that (*OneofWrapperType)(nil) gets cleared to nil.
				return
			}
			rv.Set(reflect.Zero(rv.Type()))
		},
		get: func(p pointer) pref.Value {
			if p.IsNil() {
				return conv.Zero()
			}
			rv := p.Apply(fieldOffset).AsValueOf(fs.Type).Elem()
			if rv.IsNil() || rv.Elem().Type().Elem() != ot || rv.Elem().IsNil() {
				return conv.Zero()
			}
			rv = rv.Elem().Elem().Field(0)
			return conv.PBValueOf(rv)
		},
		set: func(p pointer, v pref.Value) {
			rv := p.Apply(fieldOffset).AsValueOf(fs.Type).Elem()
			if rv.IsNil() || rv.Elem().Type().Elem() != ot || rv.Elem().IsNil() {
				rv.Set(reflect.New(ot))
			}
			rv = rv.Elem().Elem().Field(0)
			rv.Set(conv.GoValueOf(v))
		},
		mutable: func(p pointer) pref.Value {
			if !isMessage {
				panic(fmt.Sprintf("field %v with invalid Mutable call on field with non-composite type", fd.FullName()))
			}
			rv := p.Apply(fieldOffset).AsValueOf(fs.Type).Elem()
			if rv.IsNil() || rv.Elem().Type().Elem() != ot || rv.Elem().IsNil() {
				rv.Set(reflect.New(ot))
			}
			rv = rv.Elem().Elem().Field(0)
			if rv.IsNil() {
				rv.Set(conv.GoValueOf(pref.ValueOfMessage(conv.New().Message())))
			}
			return conv.PBValueOf(rv)
		},
		newMessage: func() pref.Message {
			return conv.New().Message()
		},
		newField: func() pref.Value {
			return conv.New()
		},
	}
}

func fieldInfoForMap(fd pref.FieldDescriptor, fs reflect.StructField, x exporter) fieldInfo {
	ft := fs.Type
	if ft.Kind() != reflect.Map {
		panic(fmt.Sprintf("field %v has invalid type: got %v, want map kind", fd.FullName(), ft))
	}
	conv := NewConverter(ft, fd)

	// TODO: Implement unsafe fast path?
	fieldOffset := offsetOf(fs, x)
	return fieldInfo{
		fieldDesc: fd,
		has: func(p pointer) bool {
			if p.IsNil() {
				return false
			}
			rv := p.Apply(fieldOffset).AsValueOf(fs.Type).Elem()
			return rv.Len() > 0
		},
		clear: func(p pointer) {
			rv := p.Apply(fieldOffset).AsValueOf(fs.Type).Elem()
			rv.Set(reflect.Zero(rv.Type()))
		},
		get: func(p pointer) pref.Value {
			if p.IsNil() {
				return conv.Zero()
			}
			rv := p.Apply(fieldOffset).AsValueOf(fs.Type).Elem()
			if rv.Len() == 0 {
				return conv.Zero()
			}
			return conv.PBValueOf(rv)
		},
		set: func(p pointer, v pref.Value) {
			rv := p.Apply(fieldOffset).AsValueOf(fs.Type).Elem()
			pv := conv.GoValueOf(v)
			if pv.IsNil() {
				panic(fmt.Sprintf("map field %v cannot be set with read-only value", fd.FullName()))
			}
			rv.Set(pv)
		},
		mutable: func(p pointer) pref.Value {
			v := p.Apply(fieldOffset).AsValueOf(fs.Type).Elem()
			if v.IsNil() {
				v.Set(reflect.MakeMap(fs.Type))
			}
			return conv.PBValueOf(v)
		},
		newField: func() pref.Value {
			return conv.New()
		},
	}
}

func fieldInfoForList(fd pref.FieldDescriptor, fs reflect.StructField, x exporter) fieldInfo {
	ft := fs.Type
	if ft.Kind() != reflect.Slice {
		panic(fmt.Sprintf("field %v has invalid type: got %v, want slice kind", fd.FullName(), ft))
	}
	conv := NewConverter(reflect.PtrTo(ft), fd)

	// TODO: Implement unsafe fast path?
	fieldOffset := offsetOf(fs, x)
	return fieldInfo{
		fieldDesc: fd,
		has: func(p pointer) bool {
			if p.IsNil() {
				return false
			}
			rv := p.Apply(fieldOffset).AsValueOf(fs.Type).Elem()
			return rv.Len() > 0
		},
		clear: func(p pointer) {
			rv := p.Apply(fieldOffset).AsValueOf(fs.Type).Elem()
			rv.Set(reflect.Zero(rv.Type()))
		},
		get: func(p pointer) pref.Value {
			if p.IsNil() {
				return conv.Zero()
			}
			rv := p.Apply(fieldOffset).AsValueOf(fs.Type)
			if rv.Elem().Len() == 0 {
				return conv.Zero()
			}
			return conv.PBValueOf(rv)
		},
		set: func(p pointer, v pref.Value) {
			rv := p.Apply(fieldOffset).AsValueOf(fs.Type).Elem()
			pv := conv.GoValueOf(v)
			if pv.IsNil() {
				panic(fmt.Sprintf("list field %v cannot be set with read-only value", fd.FullName()))
			}
			rv.Set(pv.Elem())
		},
		mutable: func(p pointer) pref.Value {
			v := p.Apply(fieldOffset).AsValueOf(fs.Type)
			return conv.PBValueOf(v)
		},
		newField: func() pref.Value {
			return conv.New()
		},
	}
}

var (
	nilBytes   = reflect.ValueOf([]byte(nil))
	emptyBytes = reflect.ValueOf([]byte{})
)

func fieldInfoForScalar(fd pref.FieldDescriptor, fs reflect.StructField, x exporter) fieldInfo {
	ft := fs.Type
	nullable := fd.HasPresence()
	isBytes := ft.Kind() == reflect.Slice && ft.Elem().Kind() == reflect.Uint8
	if nullable {
		if ft.Kind() != reflect.Ptr && ft.Kind() != reflect.Slice {
			panic(fmt.Sprintf("field %v has invalid type: got %v, want pointer", fd.FullName(), ft))
		}
		if ft.Kind() == reflect.Ptr {
			ft = ft.Elem()
		}
	}
	conv := NewConverter(ft, fd)

	// TODO: Implement unsafe fast path?
	fieldOffset := offsetOf(fs, x)
	return fieldInfo{
		fieldDesc: fd,
		has: func(p pointer) bool {
			if p.IsNil() {
				return false
			}
			rv := p.Apply(fieldOffset).AsValueOf(fs.Type).Elem()
			if nullable {
				return !rv.IsNil()
			}
			switch rv.Kind() {
			case reflect.Bool:
				return rv.Bool()
			case reflect.Int32, reflect.Int64:
				return rv.Int() != 0
			case reflect.Uint32, reflect.Uint64:
				return rv.Uint() != 0
			case reflect.Float32, reflect.Float64:
				return rv.Float() != 0 || math.Signbit(rv.Float())
			case reflect.String, reflect.Slice:
				return rv.Len() > 0
			default:
				panic(fmt.Sprintf("field %v has invalid type: %v", fd.FullName(), rv.Type())) // should never happen
			}
		},
		clear: func(p pointer) {
			rv := p.Apply(fieldOffset).AsValueOf(fs.Type).Elem()
			rv.Set(reflect.Zero(rv.Type()))
		},
		get: func(p pointer) pref.Value {
			if p.IsNil() {
				return conv.Zero()
			}
			rv := p.Apply(fieldOffset).AsValueOf(fs.Type).Elem()
			if nullable {
				if rv.IsNil() {
					return conv.Zero()
				}
				if rv.Kind() == reflect.Ptr {
					rv = rv.Elem()
				}
			}
			return conv.PBValueOf(rv)
		},
		set: func(p pointer, v pref.Value) {
			rv := p.Apply(fieldOffset).AsValueOf(fs.Type).Elem()
			if nullable && rv.Kind() == reflect.Ptr {
				if rv.IsNil() {
					rv.Set(reflect.New(ft))
				}
				rv = rv.Elem()
			}
			rv.Set(conv.GoValueOf(v))
			if isBytes && rv.Len() == 0 {
				if nullable {
					rv.Set(emptyBytes) // preserve presence
				} else {
					rv.Set(nilBytes) // do not preserve presence
				}
			}
		},
		newField: func() pref.Value {
			return conv.New()
		},
	}
}

func fieldInfoForWeakMessage(fd pref.FieldDescriptor, weakOffset offset) fieldInfo {
	if !flags.ProtoLegacy {
		panic("no support for proto1 weak fields")
	}

	var once sync.Once
	var messageType pref.MessageType
	lazyInit := func() {
		once.Do(func() {
			messageName := fd.Message().FullName()
			messageType, _ = preg.GlobalTypes.FindMessageByName(messageName)
			if messageType == nil {
				panic(fmt.Sprintf("weak message %v for field %v is not linked in", messageName, fd.FullName()))
			}
		})
	}

	num := fd.Number()
	return fieldInfo{
		fieldDesc: fd,
		has: func(p pointer) bool {
			if p.IsNil() {
				return false
			}
			_, ok := p.Apply(weakOffset).WeakFields().get(num)
			return ok
		},
		clear: func(p pointer) {
			p.Apply(weakOffset).WeakFields().clear(num)
		},
		get: func(p pointer) pref.Value {
			lazyInit()
			if p.IsNil() {
				return pref.ValueOfMessage(messageType.Zero())
			}
			m, ok := p.Apply(weakOffset).WeakFields().get(num)
			if !ok {
				return pref.ValueOfMessage(messageType.Zero())
			}
			return pref.ValueOfMessage(m.ProtoReflect())
		},
		set: func(p pointer, v pref.Value) {
			lazyInit()
			m := v.Message()
			if m.Descriptor() != messageType.Descriptor() {
				if got, want := m.Descriptor().FullName(), messageType.Descriptor().FullName(); got != want {
					panic(fmt.Sprintf("field %v has mismatching message descriptor: got %v, want %v", fd.FullName(), got, want))
				}
				panic(fmt.Sprintf("field %v has mismatching message descriptor: %v", fd.FullName(), m.Descriptor().FullName()))
			}
			p.Apply(weakOffset).WeakFields().set(num, m.Interface())
		},
		mutable: func(p pointer) pref.Value {
			lazyInit()
			fs := p.Apply(weakOffset).WeakFields()
			m, ok := fs.get(num)
			if !ok {
				m = messageType.New().Interface()
				fs.set(num, m)
			}
			return pref.ValueOfMessage(m.ProtoReflect())
		},
		newMessage: func() pref.Message {
			lazyInit()
			return messageType.New()
		},
		newField: func() pref.Value {
			lazyInit()
			return pref.ValueOfMessage(messageType.New())
		},
	}
}

func fieldInfoForMessage(fd pref.FieldDescriptor, fs reflect.StructField, x exporter) fieldInfo {
	ft := fs.Type
	conv := NewConverter(ft, fd)

	// TODO: Implement unsafe fast path?
	fieldOffset := offsetOf(fs, x)
	return fieldInfo{
		fieldDesc: fd,
		has: func(p pointer) bool {
			if p.IsNil() {
				return false
			}
			rv := p.Apply(fieldOffset).AsValueOf(fs.Type).Elem()
			return !rv.IsNil()
		},
		clear: func(p pointer) {
			rv := p.Apply(fieldOffset).AsValueOf(fs.Type).Elem()
			rv.Set(reflect.Zero(rv.Type()))
		},
		get: func(p pointer) pref.Value {
			if p.IsNil() {
				return conv.Zero()
			}
			rv := p.Apply(fieldOffset).AsValueOf(fs.Type).Elem()
			return conv.PBValueOf(rv)
		},
		set: func(p pointer, v pref.Value) {
			rv := p.Apply(fieldOffset).AsValueOf(fs.Type).Elem()
			rv.Set(conv.GoValueOf(v))
			if rv.IsNil() {
				panic(fmt.Sprintf("field %v has invalid nil pointer", fd.FullName()))
			}
		},
		mutable: func(p pointer) pref.Value {
			rv := p.Apply(fieldOffset).AsValueOf(fs.Type).Elem()
			if rv.IsNil() {
				rv.Set(conv.GoValueOf(conv.New()))
			}
			return conv.PBValueOf(rv)
		},
		newMessage: func() pref.Message {
			return conv.New().Message()
		},
		newField: func() pref.Value {
			return conv.New()
		},
	}
}

type oneofInfo struct {
	oneofDesc pref.OneofDescriptor
	which     func(pointer) pref.FieldNumber
}

func makeOneofInfo(od pref.OneofDescriptor, si structInfo, x exporter) *oneofInfo {
	oi := &oneofInfo{oneofDesc: od}
	if od.IsSynthetic() {
		fs := si.fieldsByNumber[od.Fields().Get(0).Number()]
		fieldOffset := offsetOf(fs, x)
		oi.which = func(p pointer) pref.FieldNumber {
			if p.IsNil() {
				return 0
			}
			rv := p.Apply(fieldOffset).AsValueOf(fs.Type).Elem()
			if rv.IsNil() { // valid on either *T or []byte
				return 0
			}
			return od.Fields().Get(0).Number()
		}
	} else {
		fs := si.oneofsByName[od.Name()]
		fieldOffset := offsetOf(fs, x)
		oi.which = func(p pointer) pref.FieldNumber {
			if p.IsNil() {
				return 0
			}
			rv := p.Apply(fieldOffset).AsValueOf(fs.Type).Elem()
			if rv.IsNil() {
				return 0
			}
			rv = rv.Elem()
			if rv.IsNil() {
				return 0
			}
			return si.oneofWrappersByType[rv.Type().Elem()]
		}
	}
	return oi
}
