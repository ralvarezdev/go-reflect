package go_reflect

import (
	"fmt"
	"reflect"
)

// CheckFunction checks if a function is valid
func CheckFunction(fn interface{}, params ...interface{}) (
	*reflect.Value,
	*[]reflect.Value,
	error,
) {
	// Get the function and its parameters
	fnValue := reflect.ValueOf(fn)
	paramsValues := make([]reflect.Value, len(params))
	for i, param := range params {
		paramsValues[i] = reflect.ValueOf(param)
	}

	// Check if the function is valid
	if fnValue.Kind() != reflect.Func {
		return nil, nil, ErrNotAFunction
	}

	// Check if the function has the correct number of parameters
	paramsCount := len(params)
	fnParamsCount := fnValue.Type().NumIn()
	if paramsCount != fnParamsCount {
		return nil, nil, fmt.Errorf(
			ErrFunctionParameterCountMismatch,
			fnParamsCount,
			paramsCount,
		)
	}

	// Check if the parameter type matches the function's parameter type
	var paramType, fnParamType reflect.Type
	for i, paramValue := range paramsValues {
		paramType = paramValue.Type()
		fnParamType = fnValue.Type().In(i)

		if paramType != fnParamType {
			return nil, nil, fmt.Errorf(
				ErrFunctionParameterTypeMismatch,
				i,
				fnParamType,
				paramType,
			)
		}
	}

	return &fnValue, &paramsValues, nil
}

// UnsafeCallFunction calls a function with some typed parameters without checking if the function is valid
func UnsafeCallFunction(fnValue *reflect.Value, paramsValues ...reflect.Value) (
	[]interface{},
	error,
) {
	// Check if the function or the parameters values are nil
	if fnValue == nil {
		return nil, ErrNilFunctionValue
	}
	if paramsValues == nil {
		paramsValues = make([]reflect.Value, 0)
	}

	// Call the function with the parameter
	results := fnValue.Call(paramsValues)

	// Convert the results to an interface slice
	interfaceResults := make([]interface{}, len(results))
	for i, result := range results {
		interfaceResults[i] = result.Interface()
	}

	return interfaceResults, nil
}

// SafeCallFunction calls a function with some typed parameters after checking if the function is valid
func SafeCallFunction(fn interface{}, params ...interface{}) (
	[]interface{},
	error,
) {
	// Check if the function is valid
	fnValue, paramsValues, err := CheckFunction(fn, params...)
	if err != nil {
		return nil, err
	}

	// Call the function with the parameter (now, we are sure that the function is valid)
	return UnsafeCallFunction(fnValue, *paramsValues...)
}
