package projscan

import "strings"

// Field represents a field of a struct.
type Field struct {
	Name      string  // Name of the field
	Type      string  // Type of the field
	Doc       string  // Documentation for the field
	StructRef *Struct // Reference to the struct that contains this field
	Pointer   bool    // Indicates whether the field is a pointer type or not
	Array     bool    // Indicates whether the field is an array type or not
}

// FromStandardLibrary returns true if the field's containing struct is part of the Go standard library.
func (s *Field) FromStandardLibrary() bool {
	if s.StructRef == nil {
		return true
	}
	return strings.HasPrefix(s.StructRef.Package.Path, "std/")
}

// FieldFinder provides a way to find all fields of a struct.
type FieldFinder interface {
	FindFieldsByStruct(st *Struct) ([]*Field, error)
}
