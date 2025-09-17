package mongodb

import (
	"context"

	userEntity "github.com/surajswarnapuri/ps-tag-onboarding-go/internal/domain/user"
)

type repository struct {
	client *MongoDBClient
}

func NewRepository(client *MongoDBClient) userEntity.Repository {
	return &repository{client: client}
}

func (r *repository) FindByID(ctx context.Context, id string) (*userEntity.User, error) {
	return nil, nil
}

func (r *repository) Save(ctx context.Context, user *userEntity.User) (*userEntity.User, error) {
	return nil, nil
}

func (r *repository) ExistsByFirstNameAndLastName(ctx context.Context, firstName string, lastName string) bool {
	return false
}

func (r *repository) ExistsByFirstNameAndLastNameAndIDNot(ctx context.Context, firstName string, lastName string, id string) bool {
	return false
}
