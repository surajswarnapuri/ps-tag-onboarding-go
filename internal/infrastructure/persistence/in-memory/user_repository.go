// Package inmemory contains the implementation of the User repository in memory.
package inmemory

import (
	"context"
	"fmt"

	"github.com/surajswarnapuri/ps-tag-onboarding-go/internal/domain/user"
)

type repository struct {
	users map[string]*user.User
}

func NewRepository() user.Repository {
	return &repository{
		users: make(map[string]*user.User),
	}
}

func (r *repository) FindByID(ctx context.Context, id string) (*user.User, error) {
	user, ok := r.users[id]
	if !ok {
		return nil, fmt.Errorf("user not found in memory")
	}
	return user, nil
}

func (r *repository) Save(ctx context.Context, user *user.User) (*user.User, error) {
	r.users[user.ID] = user
	return user, nil
}

func (r *repository) ExistsByFirstNameAndLastName(ctx context.Context, firstName string, lastName string) bool {
	for _, user := range r.users {
		if user.FirstName == firstName && user.LastName == lastName {
			return true
		}
	}
	return false
}

func (r *repository) ExistsByFirstNameAndLastNameAndIDNot(ctx context.Context, firstName string, lastName string, id string) bool {
	for _, user := range r.users {
		if user.FirstName == firstName && user.LastName == lastName && user.ID != id {
			return true
		}
	}
	return false
}
