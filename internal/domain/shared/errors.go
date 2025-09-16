// Package shared contains shared code for the domain layer
package shared

type ValidationError struct {
	Code    string
	Message string
}

func (e ValidationError) Error() string {
	return e.Message
}
