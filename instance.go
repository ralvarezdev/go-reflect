package go_reflect

import (
	"reflect"
)

// GetTypeOf returns the type of the given bodyType
func GetTypeOf(bodyType interface{}) reflect.Type {
	return reflect.TypeOf(bodyType)
}

// CreateNewInstance creates a new instance of the given type
func CreateNewInstance(bodyType interface{}) interface{} {
	// Get the reflect.Type of the bodyType
	t := reflect.TypeOf(bodyType)

	// Create a new instance of the type and return a pointer to it
	return reflect.New(t).Interface()
}

// CreateNewInstanceFromType creates a new instance of the given type
func CreateNewInstanceFromType(t reflect.Type) interface{} {
	return reflect.New(t).Interface()
}
