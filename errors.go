package go_reflect

import (
	"errors"
)

const (
	ErrFunctionParameterCountMismatch = "function parameter count mismatch, expected %d, got %d"
	ErrFunctionParameterTypeMismatch  = "function parameter type mismatch on index %d, expected %s, got %s"
)

var (
	ErrNotAFunction     = errors.New("not a function")
	ErrNilFunctionValue = errors.New("nil function value")
)
