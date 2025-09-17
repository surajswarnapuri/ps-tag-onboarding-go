// Package user contains the logic for the User domain.
package user

// User is a user entity
type User struct {
	ID        string
	FirstName string
	LastName  string
	Email     string
	Age       int
}
