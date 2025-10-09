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
