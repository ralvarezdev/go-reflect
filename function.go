package go_reflect

import (
	"fmt"
	"reflect"
)

// CallFunction dynamically calls a function with some typed parameters
func CallFunction(fn interface{}, params ...interface{}) (
	[]interface{},
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
		return nil, ErrNotAFunction
	}

	// Check if the function has the correct number of parameters
	paramsCount := len(params)
	fnParamsCount := fnValue.Type().NumIn()
	if paramsCount != fnParamsCount {
		return nil, fmt.Errorf(
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
			return nil, fmt.Errorf(
				ErrFunctionParameterTypeMismatch,
				i,
				fnParamType,
				paramType,
			)
		}
	}

	// Call the function with the parameter
	results := fnValue.Call(paramValues)

	// Convert the results to an interface slice
	interfaceResults := make([]interface{}, len(results))
	for i, result := range results {
		interfaceResults[i] = result.Interface()
	}

	return interfaceResults, nil
}
