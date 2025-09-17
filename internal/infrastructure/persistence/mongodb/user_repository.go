package mongodb

import (
	"context"
	"fmt"

	userEntity "github.com/surajswarnapuri/ps-tag-onboarding-go/internal/domain/user"
	"go.mongodb.org/mongo-driver/bson"
)

type repository struct {
	client *MongoDBClient
}

// NewRepository creates a new repository using a MongoDB client
func NewRepository(client *MongoDBClient) userEntity.Repository {
	return &repository{client: client}
}

func (r *repository) FindByID(ctx context.Context, id string) (*userEntity.User, error) {
	// query filter by id
	filter := bson.M{"_id": id}

	var userDTO user
	err := r.client.GetCollection().FindOne(ctx, filter).Decode(&userDTO)
	if err != nil {
		return nil, fmt.Errorf("mongodb: failed to find user by ID %q: %w", id, err)
	}

	return userDTO.ToEntity(), nil
}

func (r *repository) Save(ctx context.Context, userEntity *userEntity.User) (*userEntity.User, error) {
	var userDTO user
	userDTO.FromEntity(userEntity)

	_, err := r.client.GetCollection().InsertOne(ctx, userDTO)
	if err != nil {
		return nil, fmt.Errorf("mongodb: failed to save user: %w", err)
	}

	return userDTO.ToEntity(), nil

}

func (r *repository) ExistsByFirstNameAndLastName(ctx context.Context, firstName string, lastName string) bool {
	filter := bson.M{"first_name": firstName, "last_name": lastName}

	var userDTO user
	err := r.client.GetCollection().FindOne(ctx, filter).Decode(&userDTO)
	// if there is no error then user exists
	return err == nil
}

func (r *repository) ExistsByFirstNameAndLastNameAndIDNot(ctx context.Context, firstName string, lastName string, id string) bool {
	filter := bson.M{"first_name": firstName, "last_name": lastName, "_id": bson.M{"$ne": id}}

	var userDTO user
	err := r.client.GetCollection().FindOne(ctx, filter).Decode(&userDTO)

	// if there is no error then user exists
	return err == nil
}
