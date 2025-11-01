package goreflect

import (
	"reflect"
)

// UniqueTypeReference returns a unique string representation of the type of the given any
//
// Parameters:
//
//   - i: The any to get the unique type reference from
//
// Returns:
//
//   - string: The unique type reference in the format "package.TypeName"
func UniqueTypeReference(i any) string {
	t := reflect.TypeOf(i)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	return t.PkgPath() + "." + t.Name()
}

// MapToStruct maps a map[string]any to a struct
//
// Parameters:
//
//   - m: The map to map from
//   - dest: The destination struct to map to
//
// Returns:
//
//   - error: The error if any
func MapToStruct(m map[string]any, dest any) error {
	// Check if the map is nil
	if m == nil {
		return ErrNilMap
	}
	
	// Check if the destination is nil
	if dest == nil {
		return ErrNilDestination
	}
	
	// Dereference the destination
	reflectValue := GetDereferencedValue(dest)

	// Ensure the destination is a struct
	reflectType := reflectValue.Type()
	if reflectType.Kind() != reflect.Struct {
		return ErrFailedToMapToStructNotAStruct
	}

	// Map the fields
	for i := 0; i < reflectType.NumField(); i++ {
		// Get the field and its type
		fieldValue := reflectValue.Field(i)
		fieldType := reflectType.Field(i)
		fieldName := fieldType.Name

		// Check if the field exists in the map and is settable
		value, ok := m[fieldName]
		if !ok || !fieldValue.CanSet() {
			continue
		}

		// Set the field value based on its kind
		switch fieldValue.Kind() {
		case reflect.Struct:
			// Handle nested structs
			nestedMap, nestedOk := value.(map[string]any)
			if nestedOk {
				if err := MapToStruct(
					nestedMap,
					fieldValue.Addr().Interface(),
				); err != nil {
					return err
				}
			}
		default:
			fieldValue.Set(reflect.ValueOf(value).Convert(fieldType.Type))
		}
	}
	return nil
}

// IsStructFieldExported checks if a struct field is exported
// 
// Parameters:
// 
//   - field: The reflect.StructField to check
// 
// Returns:
// 
//  - bool: True if the field is exported, false otherwise
func IsStructFieldExported(field reflect.StructField) bool {
	return field.PkgPath == ""
}