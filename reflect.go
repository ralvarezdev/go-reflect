package go_reflect

import (
	"reflect"
)

type (
	// Reflection struct to hold reflection data
	Reflection struct {
		instance          interface{}
		reflectedValue    reflect.Value
		reflectedType     reflect.Type
		reflectedTypeName string
	}
)

// GetValue returns the value reflection
//
// Parameters:
//
// - instance: the instance to reflect
//
// Returns:
//
// - reflect.Value: the value reflection
func GetValue(instance interface{}) reflect.Value {
	// Check if the instance is nil
	if instance == nil {
		return reflect.Value{}
	}
	return reflect.ValueOf(instance)
}

// GetDereferencedValue returns the dereferenced value reflection
//
// Parameters:
//
// - instance: the instance to reflect
//
// Returns:
//
// - reflect.Value: the dereferenced value reflection
func GetDereferencedValue(instance interface{}) reflect.Value {
	// Reflect data
	valueReflection := GetValue(instance)

	// If data is a pointer, dereference it
	if valueReflection.Kind() == reflect.Ptr {
		valueReflection = valueReflection.Elem()
	}
	return valueReflection
}

// GetType returns the type reflection
//
// Parameters:
//
// - instance: the instance to reflect
//
// Returns:
//
// - reflect.Type: the type reflection
func GetType(instance interface{}) reflect.Type {
	// Check if the instance is nil
	if instance == nil {
		return nil
	}
	return reflect.TypeOf(instance)
}

// GetDereferencedType returns the dereferenced type reflection
//
// Parameters:
//
// - instance: the instance to reflect
//
// Returns:
//
// - reflect.Type: the dereferenced type reflection
func GetDereferencedType(instance interface{}) reflect.Type {
	// Reflect data
	typeReflection := GetType(instance)

	// If data is a pointer, dereference it
	if typeReflection.Kind() == reflect.Ptr {
		typeReflection = typeReflection.Elem()
	}
	return typeReflection
}

// GetTypeName returns the type name
//
// Parameters:
//
// - typeReflection: the type reflection
//
// Returns:
//
// - string: the type name
func GetTypeName(typeReflection reflect.Type) string {
	return typeReflection.Name()
}

// NewReflection creates a new reflection from an instance
//
// Parameters:
//
// - instance: the instance to reflect
//
// Returns:
//
// - *Reflection: the reflection instance
func NewReflection(instance interface{}) *Reflection {
	// Reflect data
	reflectedValue := GetValue(instance)
	reflectedType := GetType(instance)
	reflectedTypeName := GetTypeName(reflectedType)

	return &Reflection{
		instance,
		reflectedValue,
		reflectedType,
		reflectedTypeName,
	}
}

// NewDereferencedReflection creates a new reflection from a dereferenced instance
//
// Parameters:
//
// - instance: the instance to reflect
//
// Returns:
//
// - *Reflection: the reflection instance
func NewDereferencedReflection(instance interface{}) *Reflection {
	// Reflect data
	reflectedValue := GetDereferencedValue(instance)
	reflectedType := GetDereferencedType(instance)
	reflectedTypeName := GetTypeName(reflectedType)

	return &Reflection{
		instance,
		reflectedValue,
		reflectedType,
		reflectedTypeName,
	}
}

// GetInstance returns the instance
//
// Returns:
//
// - interface{}: the instance
func (r Reflection) GetInstance() interface{} {
	return r.instance
}

// GetReflectedValue returns the reflected value
//
// Returns:
//
// - reflect.Value: the reflected value
func (r Reflection) GetReflectedValue() reflect.Value {
	return r.reflectedValue
}

// GetReflectedType returns the reflected type
//
// Returns:
//
// - reflect.Type: the reflected type
func (r Reflection) GetReflectedType() reflect.Type {
	return r.reflectedType
}

// GetReflectedTypeName returns the reflected type name
//
// Returns:
//
// - string: the reflected type name
func (r Reflection) GetReflectedTypeName() string {
	return r.reflectedTypeName
}
