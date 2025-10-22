package goreflect

import (
	"reflect"
)

// GetTypeOf returns the type of the given bodyType
//
// Parameters:
//
//   - bodyType: The body type to get the type of
//
// Returns:
//
//   - reflect.Type: The type of the given bodyType
func GetTypeOf(bodyType any) reflect.Type {
	return reflect.TypeOf(bodyType)
}

// NewInstance creates a new instance of the given type
//
// Parameters:
//
// - bodyType: The body type to create a new instance of
//
// Returns:
//
// - any: A pointer to a new instance of the given type
func NewInstance(bodyType any) any {
	// Get the reflect.Type of the bodyType
	t := reflect.TypeOf(bodyType)

	// Create a new instance of the type and return a pointer to it
	return reflect.New(t).Interface()
}

// NewInstanceFromType creates a new instance of the given type
//
// Parameters:
//
// - t: The reflect.Type to create a new instance of
//
// Returns:
//
// - any: A pointer to a new instance of the given type
func NewInstanceFromType(t reflect.Type) any {
	return reflect.New(t).Interface()
}
