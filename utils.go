package go_reflect

import (
	"reflect"
)

// UniqueTypeReference returns a unique string representation of the type of the given interface{}
//
// Parameters:
//
//   - i: The interface{} to get the unique type reference from
//
// Returns:
//
//   - string: The unique type reference in the format "package.TypeName"
func UniqueTypeReference(i interface{}) string {
	t := reflect.TypeOf(i)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	return t.PkgPath() + "." + t.Name()
}

// MapToStruct maps a map[string]interface{} to a struct
//
// Parameters:
//
//   - m: The map to map from
//   - dest: The destination struct to map to
//
// Returns:
//
//   - error: The error if any
func MapToStruct(m map[string]interface{}, dest interface{}) error {
	// Dereference the destination
	v := reflect.ValueOf(dest)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	// Ensure the destination is a struct
	t := v.Type()
	if t.Kind() != reflect.Struct {
		return ErrFailedToMapToStructNotAStruct
	}

	// Map the fields
	for i := 0; i < v.NumField(); i++ {
		// Get the field and its type
		field := v.Field(i)
		fieldType := t.Field(i)

		// Check if the field exists in the map and is settable
		key := fieldType.Name
		val, ok := m[key]
		if !ok || !field.CanSet() {
			continue
		}

		// Set the field value based on its kind
		switch field.Kind() {
		case reflect.Struct:
			// Handle nested structs
			nestedMap, ok := val.(map[string]interface{})
			if ok {
				if err := MapToStruct(
					nestedMap,
					field.Addr().Interface(),
				); err != nil {
					return err
				}
			}
		default:
			field.Set(reflect.ValueOf(val).Convert(field.Type()))
		}
	}
	return nil
}
