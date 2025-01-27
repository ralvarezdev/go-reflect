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
	paramValues := make([]reflect.Value, len(params))
	for i, param := range params {
		paramValues[i] = reflect.ValueOf(param)
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
	for i, paramValue := range paramValues {
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

	return &fnValue, &paramValues, nil
}

// CallFunction dynamically calls a function with some typed parameters
func CallFunction(fnValue *reflect.Value, paramValues ...reflect.Value) (
	[]interface{},
	error,
) {
	// Call the function with the parameter
	results := fnValue.Call(paramValues)

	// Convert the results to an interface slice
	interfaceResults := make([]interface{}, len(results))
	for i, result := range results {
		interfaceResults[i] = result.Interface()
	}

	return interfaceResults, nil
}

// CheckAndCallFunction checks if a function is valid and calls it with some typed parameters
func CheckAndCallFunction(fn interface{}, params ...interface{}) (
	[]interface{},
	error,
) {
	// Check if the function is valid
	fnValue, paramValues, err := CheckFunction(fn, params...)
	if err != nil {
		return nil, err
	}

	// Call the function with the parameter
	return CallFunction(fnValue, *paramValues...)
}
