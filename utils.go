package goreflect

import (
	"fmt"
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

// ValueToReflectType maps a value to a type using reflection
//
// Parameters:
//
//   - value: The value to map from
//   - destValue: The destination value to map to
//   - skip: Whether to skip fields that cannot be set
//
// Returns:
//
// - error: The error if any
func ValueToReflectType(value any, destValue reflect.Value, skip bool) error {
	// Get the destination type
	destType := destValue.Type()

	// Handle complex types
	switch destValue.Kind() {
	case reflect.Interface:
		if !skip {
			return fmt.Errorf(ErrUnexpectedInterface, destType)
		}
	case reflect.Ptr:
		return ValueToReflectPtr(value, destValue, skip)
	case reflect.Struct:
		// Parse the value as a map
		parsedMap, ok := value.(map[string]any)
		if !ok {
			if !skip {
				return fmt.Errorf(ErrExpectedMapValueForStruct, destType)
			}
			return nil
		}
		return ValueToReflectStruct(parsedMap, destValue, skip)
	case reflect.Slice:
		return ValueToReflectSlice(value, destValue, skip)
	case reflect.Map:
		return ValueToReflectMap(value, destValue, skip)
	case reflect.Array:
		return ValueToReflectArray(value, destValue, skip)
	default:
		// Get the reflection of the value
		reflectValue := reflect.ValueOf(value)

		// Set the value if convertible
		if reflectValue.Type().ConvertibleTo(destType) {
			destValue.Set(reflectValue.Convert(destType))
		} else if !skip {
			return fmt.Errorf(ErrConversionFailed, reflectValue.Type(), destType)
		}
	}
	return nil
}

// ValueToReflectPtr maps a value to a pointer using reflection
//
// Parameters:
//
//   - value: The value to map from
//   - destValue: The destination value to map to
//   - skip: Whether to skip fields that cannot be set
//
// Returns:
//
//   - error: The error if any
func ValueToReflectPtr(value any, destValue reflect.Value, skip bool) error {
	// Check if the destination value is a pointer
	if destValue.Kind() != reflect.Ptr {
		if !skip {
			return fmt.Errorf(ErrExpectedPointerDestination, destValue.Kind())
		}
		return nil
	}

	// Dereference the pointer value
	destValueElem := destValue.Elem()

	// Initialize the pointer if it's nil
	if destValueElem.Kind() == reflect.Invalid {
		// Get the element type
		destTypeElem := destValue.Type().Elem()

		// Set a new value
		destValue.Set(reflect.New(destTypeElem))
		destValueElem = destValue.Elem()
	}

	// Set the value if convertible
	return ValueToReflectType(value, destValueElem, skip)
}

// ValueToReflectSlice maps a value to a slice using reflection
//
// Parameters:
//
// - value: The value to map from
// - destValue: The destination value to map to
// - skip: Whether to skip fields that cannot be set
//
// Returns:
//
// - error: The error if any
func ValueToReflectSlice(value any, destValue reflect.Value, skip bool) error {
	// Get the reflection of the value
	reflectValue := reflect.ValueOf(value)

	// Skip if the value is not a slice
	if reflectValue.Kind() != reflect.Slice {
		if !skip {
			return fmt.Errorf(ErrExpectedSliceValue, reflectValue.Kind())
		}
		return nil
	}

	// Check if the destination type is a slice
	if destValue.Kind() != reflect.Slice {
		if !skip {
			return fmt.Errorf(ErrExpectedSliceDestination, destValue.Kind())
		}
		return nil
	}

	// Get the element type of the slice
	destTypeElem := destValue.Type().Elem()

	// Initialize a new slice
	newSlice := reflect.MakeSlice(destTypeElem, reflectValue.Len(), reflectValue.Len())

	// Map each element
	for j := 0; j < reflectValue.Len(); j++ {
		reflectValueElem := reflectValue.Index(j)
		destValueElem := newSlice.Index(j)

		// Set the element value
		if err := ValueToReflectType(reflectValueElem.Interface(), destValueElem, skip); err != nil {
			return err
		}
	}

	// Set the destination value
	destValue.Set(newSlice)
	return nil
}

// ValueToReflectMap maps a value to a map using reflection
//
// Parameters:
//
// - value: The value to map from
// - destValue: The destination value to map to
// - skip: Whether to skip fields that cannot be set
//
// Returns:
//
// - error: The error if any
func ValueToReflectMap(value any, destValue reflect.Value, skip bool) error {
	// Get the reflection of the value
	reflectValue := reflect.ValueOf(value)

	// Check if the value is a map, if not, skip or return an error
	if reflectValue.Kind() != reflect.Map {
		if !skip {
			return fmt.Errorf(ErrExpectedMapValue, reflectValue.Kind())
		}
		return nil
	}

	// Check if the destination type is a map
	if destValue.Kind() != reflect.Map {
		if !skip {
			return fmt.Errorf(ErrExpectedMapDestination, destValue.Kind())
		}
		return nil
	}

	// Get the destination type and element type
	destType := destValue.Type()
	destTypeElem := destType.Elem()

	// Initialize a new map
	newMap := reflect.MakeMap(destType)

	// Map each key/value pair
	for _, newKey := range reflectValue.MapKeys() {
		// Get the new element
		valueNewElem := reflectValue.MapIndex(newKey)

		// Create a new value for the element
		destNewElem := reflect.New(destTypeElem).Elem()

		// Map the element value
		if err := ValueToReflectType(valueNewElem.Interface(), destNewElem, skip); err != nil {
			return err
		}

		// Set the key/value pair in the new map
		newMap.SetMapIndex(newKey, destNewElem)
	}

	// Set the destination value
	destValue.Set(newMap)
	return nil
}

// ValueToReflectArray maps a value to an array using reflection
//
// Parameters:
//
// - value: The value to map from
// - destValue: The destination value to map to
// - skip: Whether to skip fields that cannot be set
//
// Returns:
//
// - error: The error if any
func ValueToReflectArray(value any, destValue reflect.Value, skip bool) error {
	// Get the reflection of the value
	reflectValue := reflect.ValueOf(value)

	// Check if the value is an array, if not, skip or return an error
	if reflectValue.Kind() != reflect.Array && reflectValue.Kind() != reflect.Slice {
		if !skip {
			return fmt.Errorf(ErrExpectedSliceOrArrayValue, reflectValue.Kind())
		}
		return nil
	}

	// Check if the destination type is an array
	if destValue.Kind() != reflect.Array {
		if !skip {
			return fmt.Errorf(ErrExpectedArrayDestination, destValue.Kind())
		}
		return nil
	}

	// Get the length of the array
	arrayLen := destValue.Len()

	// Map each element up to the length of the array
	for j := 0; j < arrayLen; j++ {
		reflectValueElem := reflectValue.Index(j)
		destValueElem := destValue.Index(j)

		// Set the element value
		if err := ValueToReflectType(reflectValueElem.Interface(), destValueElem, skip); err != nil {
			return err
		}
	}
	return nil
}

// ValueToReflectStruct maps a map to a struct using reflection
//
// Parameters:
//
//   - m: The map to map from
//   - destValue: The destination struct value to map to
//   - skip: Whether to skip fields that cannot be set
//
// Returns:
//
//   - error: The error if any
func ValueToReflectStruct(m map[string]any, destValue reflect.Value, skip bool) error {
	// Check if the map is nil
	if m == nil {
		return ErrNilMap
	}

	// Check if the destination value is a struct
	if destValue.Kind() != reflect.Struct {
		if !skip {
			return fmt.Errorf(ErrExpectedStructDestination, destValue.Kind())
		}
		return nil
	}

	// Get the destination type
	destType := destValue.Type()

	// Map the fields
	for i := 0; i < destType.NumField(); i++ {
		// Get the field and its type
		structField := destType.Field(i)
		fieldValue := destValue.Field(i)
		fieldType := structField.Type
		fieldName := structField.Name

		// Check if the field exists in the map and is settable
		value, ok := m[fieldName]
		if !ok || !fieldValue.CanSet() {
			continue
		}

		// Initialize pointer fields if they are nil
		if fieldValue.IsNil() {
			fieldValue.Set(reflect.New(fieldType.Elem()))
		}

		// Set the field value based on its kind
		if err := ValueToReflectType(value, fieldValue, skip); err != nil {
			return err
		}
	}
	return nil
}

// IsStructFieldExported checks if a struct field is exported
//
// Parameters:
//
//   - field: The reflect.StructField to check, if nil, returns false
//
// Returns:
//
//   - bool: True if the field is exported, false otherwise
func IsStructFieldExported(field *reflect.StructField) bool {
	// Check if the field is nil, if so, return false
	if field == nil {
		return false
	}
	return field.PkgPath == ""
}
