package goreflect

import (
	"errors"
)

const (
	ErrFunctionParameterCountMismatch = "function parameter count mismatch, expected %d, got %d"
	ErrFunctionParameterTypeMismatch  = "function parameter type mismatch on index %d, expected %s, got %s"
)

var (
	ErrNotAFunction                  = errors.New("not a function")
	ErrNilFunctionValue              = errors.New("nil function value")
	ErrFailedToMapToStructNotAStruct = errors.New("failed to map to struct: destination is not a struct")
	ErrNilMap 						= errors.New("nil map")
	ErrNilDestination 				= errors.New("nil destination")
)
