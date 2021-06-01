package storage

// Code generated by cdproto-gen. DO NOT EDIT.

import (
	"errors"

	"github.com/mailru/easyjson"
	"github.com/mailru/easyjson/jlexer"
	"github.com/mailru/easyjson/jwriter"
)

// Type enum of possible storage types.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Storage#type-StorageType
type Type string

// String returns the Type as string value.
func (t Type) String() string {
	return string(t)
}

// Type values.
const (
	TypeAppcache       Type = "appcache"
	TypeCookies        Type = "cookies"
	TypeFileSystems    Type = "file_systems"
	TypeIndexeddb      Type = "indexeddb"
	TypeLocalStorage   Type = "local_storage"
	TypeShaderCache    Type = "shader_cache"
	TypeWebsql         Type = "websql"
	TypeServiceWorkers Type = "service_workers"
	TypeCacheStorage   Type = "cache_storage"
	TypeAll            Type = "all"
	TypeOther          Type = "other"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t Type) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t Type) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *Type) UnmarshalEasyJSON(in *jlexer.Lexer) {
	switch Type(in.String()) {
	case TypeAppcache:
		*t = TypeAppcache
	case TypeCookies:
		*t = TypeCookies
	case TypeFileSystems:
		*t = TypeFileSystems
	case TypeIndexeddb:
		*t = TypeIndexeddb
	case TypeLocalStorage:
		*t = TypeLocalStorage
	case TypeShaderCache:
		*t = TypeShaderCache
	case TypeWebsql:
		*t = TypeWebsql
	case TypeServiceWorkers:
		*t = TypeServiceWorkers
	case TypeCacheStorage:
		*t = TypeCacheStorage
	case TypeAll:
		*t = TypeAll
	case TypeOther:
		*t = TypeOther

	default:
		in.AddError(errors.New("unknown Type value"))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *Type) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}

// UsageForType usage for a storage type.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Storage#type-UsageForType
type UsageForType struct {
	StorageType Type    `json:"storageType"` // Name of storage type.
	Usage       float64 `json:"usage"`       // Storage usage (bytes).
}

// TrustTokens pair of issuer origin and number of available (signed, but not
// used) Trust Tokens from that issuer.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Storage#type-TrustTokens
type TrustTokens struct {
	IssuerOrigin string  `json:"issuerOrigin"`
	Count        float64 `json:"count"`
}
