package main

import (
	"fmt"
)

// BasicError represents an error with a single message
type BasicError struct {
	Message string
}

// Error implements the error interface{}
func (b BasicError) Error() string {
	return b.Message
}

// RaiseError is a factory that creates a BasicError
func RaiseError(message string) error {
	return BasicError{message}
}

// RaiseErrorf is a factory that creates a BasicError from a format string
func RaiseErrorf(format string, a ...interface{}) error {
	return BasicError{fmt.Sprintf(format, a...)}
}
