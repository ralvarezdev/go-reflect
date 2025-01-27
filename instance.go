package go_reflect

import (
	"reflect"
)

// GetTypeOf returns the type of the given bodyType
func GetTypeOf(bodyType interface{}) reflect.Type {
	return reflect.TypeOf(bodyType)
}

// NewInstance creates a new instance of the given type
func NewInstance(bodyType interface{}) interface{} {
	// Get the reflect.Type of the bodyType
	t := reflect.TypeOf(bodyType)

	// Create a new instance of the type and return a pointer to it
	return reflect.New(t).Interface()
}

// NewInstanceFromType creates a new instance of the given type
func NewInstanceFromType(t reflect.Type) interface{} {
	return reflect.New(t).Interface()
}
