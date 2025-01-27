package go_reflect

import (
	"errors"
)

var (
	ErrNotAFunction                   = errors.New("not a function")
	ErrFunctionParameterCountMismatch = "function parameter count mismatch, expected %d, got %d"
	ErrFunctionParameterTypeMismatch  = "function parameter type mismatch on index %d, expected %s, got %s"
)
