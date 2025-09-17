package user

import "context"

type Repository interface {
	FindByID(ctx context.Context, id string) (*User, error)
	Save(ctx context.Context, user *User) error
	ExistsByFirstNameAndLastName(ctx context.Context, firstName string, lastName string) bool
	ExistsByFirstNameAndLastNameAndIDNot(ctx context.Context, firstName string, lastName string, id string) bool
}
