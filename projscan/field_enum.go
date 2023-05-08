// Code generated by go-enum DO NOT EDIT.
// Version: 0.5.6
// Revision: 97611fddaa414f53713597918c5e954646cb8623
// Build Date: 2023-03-26T21:38:06Z
// Built By: goreleaser

package projscan

import (
	"errors"
	"fmt"
)

const (
	// FieldTypeString is a FieldType of type string.
	FieldTypeString FieldType = "string"
	// FieldTypeBool is a FieldType of type bool.
	FieldTypeBool FieldType = "bool"
	// FieldTypeInt is a FieldType of type int.
	FieldTypeInt FieldType = "int"
	// FieldTypeInt8 is a FieldType of type int8.
	FieldTypeInt8 FieldType = "int8"
	// FieldTypeInt16 is a FieldType of type int16.
	FieldTypeInt16 FieldType = "int16"
	// FieldTypeInt32 is a FieldType of type int32.
	FieldTypeInt32 FieldType = "int32"
	// FieldTypeInt64 is a FieldType of type int64.
	FieldTypeInt64 FieldType = "int64"
	// FieldTypeFloat32 is a FieldType of type float32.
	FieldTypeFloat32 FieldType = "float32"
	// FieldTypeFloat64 is a FieldType of type float64.
	FieldTypeFloat64 FieldType = "float64"
)

var ErrInvalidFieldType = errors.New("not a valid FieldType")

// String implements the Stringer interface.
func (x FieldType) String() string {
	return string(x)
}

// IsValid provides a quick way to determine if the typed value is
// part of the allowed enumerated values
func (x FieldType) IsValid() bool {
	_, err := ParseFieldType(string(x))
	return err == nil
}

var _FieldTypeValue = map[string]FieldType{
	"string":  FieldTypeString,
	"bool":    FieldTypeBool,
	"int":     FieldTypeInt,
	"int8":    FieldTypeInt8,
	"int16":   FieldTypeInt16,
	"int32":   FieldTypeInt32,
	"int64":   FieldTypeInt64,
	"float32": FieldTypeFloat32,
	"float64": FieldTypeFloat64,
}

// ParseFieldType attempts to convert a string to a FieldType.
func ParseFieldType(name string) (FieldType, error) {
	if x, ok := _FieldTypeValue[name]; ok {
		return x, nil
	}
	return FieldType(""), fmt.Errorf("%s is %w", name, ErrInvalidFieldType)
}