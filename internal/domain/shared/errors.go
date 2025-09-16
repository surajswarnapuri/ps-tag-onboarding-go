// Package shared contains shared code for the domain layer
package shared

import "fmt"

type ValidationError struct {
	Code    string
	Message string
}

func (e ValidationError) Error() string {
	return e.Message
}

type WrappedError struct {
	Message string
	Err     error
}

func (e WrappedError) Error() string {
	return fmt.Sprintf("%s: %v", e.Message, e.Err)
}

func NewWrappedError(message string, err error) WrappedError {
	return WrappedError{
		Message: message,
		Err:     err,
	}
}

func (e WrappedError) Unwrap() error {
	return e.Err
}
