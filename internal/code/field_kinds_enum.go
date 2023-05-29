// Code generated by go-enum DO NOT EDIT.
// Version: 0.5.6
// Revision: 97611fddaa414f53713597918c5e954646cb8623
// Build Date: 2023-03-26T21:38:06Z
// Built By: goreleaser

package code

import (
	"errors"
	"fmt"
)

const (
	// FieldKindNative is a FieldKind of type Native.
	FieldKindNative FieldKind = "Native"
	// FieldKindStdLib is a FieldKind of type StdLib.
	FieldKindStdLib FieldKind = "StdLib"
	// FieldKindStringMap is a FieldKind of type StringMap.
	FieldKindStringMap FieldKind = "StringMap"
	// FieldKindTCloudTag is a FieldKind of type TCloudTag.
	FieldKindTCloudTag FieldKind = "TCloudTag"
	// FieldKindStruct is a FieldKind of type Struct.
	FieldKindStruct FieldKind = "Struct"
)

var ErrInvalidFieldKind = errors.New("not a valid FieldKind")

// String implements the Stringer interface.
func (x FieldKind) String() string {
	return string(x)
}

// IsValid provides a quick way to determine if the typed value is
// part of the allowed enumerated values
func (x FieldKind) IsValid() bool {
	_, err := ParseFieldKind(string(x))
	return err == nil
}

var _FieldKindValue = map[string]FieldKind{
	"Native":    FieldKindNative,
	"StdLib":    FieldKindStdLib,
	"StringMap": FieldKindStringMap,
	"TCloudTag": FieldKindTCloudTag,
	"Struct":    FieldKindStruct,
}

// ParseFieldKind attempts to convert a string to a FieldKind.
func ParseFieldKind(name string) (FieldKind, error) {
	if x, ok := _FieldKindValue[name]; ok {
		return x, nil
	}
	return FieldKind(""), fmt.Errorf("%s is %w", name, ErrInvalidFieldKind)
}
