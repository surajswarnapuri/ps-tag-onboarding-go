// Package shared contains shared code for the domain layer
package shared

// ValidationError is a validation error
type ValidationError struct {
	Code    string
	Message string
}

// Error returns the error message
func (e ValidationError) Error() string {
	return e.Message
}
