// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1beta1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in ExtraValue) DeepCopyInto(out *ExtraValue) {
	{
		in := &in
		*out = make(ExtraValue, len(*in))
		copy(*out, *in)
		return
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExtraValue.
func (in ExtraValue) DeepCopy() ExtraValue {
	if in == nil {
		return nil
	}
	out := new(ExtraValue)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalSubjectAccessReview) DeepCopyInto(out *LocalSubjectAccessReview) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalSubjectAccessReview.
func (in *LocalSubjectAccessReview) DeepCopy() *LocalSubjectAccessReview {
	if in == nil {
		return nil
	}
	out := new(LocalSubjectAccessReview)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LocalSubjectAccessReview) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NonResourceAttributes) DeepCopyInto(out *NonResourceAttributes) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NonResourceAttributes.
func (in *NonResourceAttributes) DeepCopy() *NonResourceAttributes {
	if in == nil {
		return nil
	}
	out := new(NonResourceAttributes)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NonResourceRule) DeepCopyInto(out *NonResourceRule) {
	*out = *in
	if in.Verbs != nil {
		in, out := &in.Verbs, &out.Verbs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.NonResourceURLs != nil {
		in, out := &in.NonResourceURLs, &out.NonResourceURLs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NonResourceRule.
func (in *NonResourceRule) DeepCopy() *NonResourceRule {
	if in == nil {
		return nil
	}
	out := new(NonResourceRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceAttributes) DeepCopyInto(out *ResourceAttributes) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceAttributes.
func (in *ResourceAttributes) DeepCopy() *ResourceAttributes {
	if in == nil {
		return nil
	}
	out := new(ResourceAttributes)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceRule) DeepCopyInto(out *ResourceRule) {
	*out = *in
	if in.Verbs != nil {
		in, out := &in.Verbs, &out.Verbs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.APIGroups != nil {
		in, out := &in.APIGroups, &out.APIGroups
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ResourceNames != nil {
		in, out := &in.ResourceNames, &out.ResourceNames
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceRule.
func (in *ResourceRule) DeepCopy() *ResourceRule {
	if in == nil {
		return nil
	}
	out := new(ResourceRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SelfSubjectAccessReview) DeepCopyInto(out *SelfSubjectAccessReview) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SelfSubjectAccessReview.
func (in *SelfSubjectAccessReview) DeepCopy() *SelfSubjectAccessReview {
	if in == nil {
		return nil
	}
	out := new(SelfSubjectAccessReview)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SelfSubjectAccessReview) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SelfSubjectAccessReviewSpec) DeepCopyInto(out *SelfSubjectAccessReviewSpec) {
	*out = *in
	if in.ResourceAttributes != nil {
		in, out := &in.ResourceAttributes, &out.ResourceAttributes
		*out = new(ResourceAttributes)
		**out = **in
	}
	if in.NonResourceAttributes != nil {
		in, out := &in.NonResourceAttributes, &out.NonResourceAttributes
		*out = new(NonResourceAttributes)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SelfSubjectAccessReviewSpec.
func (in *SelfSubjectAccessReviewSpec) DeepCopy() *SelfSubjectAccessReviewSpec {
	if in == nil {
		return nil
	}
	out := new(SelfSubjectAccessReviewSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SelfSubjectRulesReview) DeepCopyInto(out *SelfSubjectRulesReview) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SelfSubjectRulesReview.
func (in *SelfSubjectRulesReview) DeepCopy() *SelfSubjectRulesReview {
	if in == nil {
		return nil
	}
	out := new(SelfSubjectRulesReview)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SelfSubjectRulesReview) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SelfSubjectRulesReviewSpec) DeepCopyInto(out *SelfSubjectRulesReviewSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SelfSubjectRulesReviewSpec.
func (in *SelfSubjectRulesReviewSpec) DeepCopy() *SelfSubjectRulesReviewSpec {
	if in == nil {
		return nil
	}
	out := new(SelfSubjectRulesReviewSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubjectAccessReview) DeepCopyInto(out *SubjectAccessReview) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubjectAccessReview.
func (in *SubjectAccessReview) DeepCopy() *SubjectAccessReview {
	if in == nil {
		return nil
	}
	out := new(SubjectAccessReview)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SubjectAccessReview) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubjectAccessReviewSpec) DeepCopyInto(out *SubjectAccessReviewSpec) {
	*out = *in
	if in.ResourceAttributes != nil {
		in, out := &in.ResourceAttributes, &out.ResourceAttributes
		*out = new(ResourceAttributes)
		**out = **in
	}
	if in.NonResourceAttributes != nil {
		in, out := &in.NonResourceAttributes, &out.NonResourceAttributes
		*out = new(NonResourceAttributes)
		**out = **in
	}
	if in.Groups != nil {
		in, out := &in.Groups, &out.Groups
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Extra != nil {
		in, out := &in.Extra, &out.Extra
		*out = make(map[string]ExtraValue, len(*in))
		for key, val := range *in {
			var outVal []string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make(ExtraValue, len(*in))
				copy(*out, *in)
			}
			(*out)[key] = outVal
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubjectAccessReviewSpec.
func (in *SubjectAccessReviewSpec) DeepCopy() *SubjectAccessReviewSpec {
	if in == nil {
		return nil
	}
	out := new(SubjectAccessReviewSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubjectAccessReviewStatus) DeepCopyInto(out *SubjectAccessReviewStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubjectAccessReviewStatus.
func (in *SubjectAccessReviewStatus) DeepCopy() *SubjectAccessReviewStatus {
	if in == nil {
		return nil
	}
	out := new(SubjectAccessReviewStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubjectRulesReviewStatus) DeepCopyInto(out *SubjectRulesReviewStatus) {
	*out = *in
	if in.ResourceRules != nil {
		in, out := &in.ResourceRules, &out.ResourceRules
		*out = make([]ResourceRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.NonResourceRules != nil {
		in, out := &in.NonResourceRules, &out.NonResourceRules
		*out = make([]NonResourceRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubjectRulesReviewStatus.
func (in *SubjectRulesReviewStatus) DeepCopy() *SubjectRulesReviewStatus {
	if in == nil {
		return nil
	}
	out := new(SubjectRulesReviewStatus)
	in.DeepCopyInto(out)
	return out
}
