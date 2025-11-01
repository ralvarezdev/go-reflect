package goreflect

import (
	"errors"
)

const (
	ErrFunctionParameterCountMismatch = "function parameter count mismatch, expected %d, got %d"
	ErrFunctionParameterTypeMismatch  = "function parameter type mismatch on index %d, expected %s, got %s"
	ErrExpectedStructDestination      = "expected struct destination but got %T"
	ErrExpectedMapValue               = "expected map value but got %T"
	ErrExpectedMapDestination         = "expected map destination but got %T"
	ErrExpectedMapValueForStruct      = "expected map value for struct but got %T"
	ErrExpectedPointerDestination     = "expected pointer destination but got %T"
	ErrExpectedSliceValue             = "expected slice value but got %v"
	ErrExpectedSliceDestination       = "expected slice destination but got %T"
	ErrExpectedSliceOrArrayValue      = "expected slice or array value but got %T"
	ErrExpectedArrayDestination       = "expected array destination but got %T"
	ErrConversionFailed               = "cannot convert %v to %v"
	ErrUnexpectedInterface            = "unexpected interface type: %T, cannot be converted due to unknown underlying type"
)

var (
	ErrNotAFunction                  = errors.New("not a function")
	ErrNilFunctionValue              = errors.New("nil function value")
	ErrFailedToMapToStructNotAStruct = errors.New("failed to map to struct: destination is not a struct")
	ErrNilMap                        = errors.New("nil map")
)
