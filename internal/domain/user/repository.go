package user

import "context"

type Repository interface {
	// FindByID finds a user by id
	FindByID(ctx context.Context, id string) (*User, error)
	// Save saves a user to the repository
	Save(ctx context.Context, user *User) (*User, error)
	// ExistsByFirstNameAndLastName checks if a user exists by first name and last name
	ExistsByFirstNameAndLastName(ctx context.Context, firstName string, lastName string) bool
	// ExistsByFirstNameAndLastNameAndIDNot checks if a user exists by first name and last name but not by id
	ExistsByFirstNameAndLastNameAndIDNot(ctx context.Context, firstName string, lastName string, id string) bool
}
