//go:build integration

package mongodb

import (
	"context"
	"testing"

	userEntity "github.com/surajswarnapuri/ps-tag-onboarding-go/internal/domain/user"
)

// wipeCollection removes all documents from the test collection
func wipeCollection(ctx context.Context, client *MongoDBClient) error {
	collection := client.GetCollection()
	_, err := collection.DeleteMany(ctx, map[string]interface{}{})
	return err
}

// setupTestEnvironment creates a MongoDB client and wipes the collection for testing
func setupTestEnvironment(t *testing.T) (*MongoDBClient, userEntity.Repository) {
	ctx := context.Background()
	client, err := NewMongoDBClient(ctx, "mongodb://root:password@localhost:27017", "user_integration_test")
	if err != nil {
		t.Fatalf("Failed to create MongoDB client: %v", err)
	}

	// Wipe the collection before running the test
	if err := wipeCollection(ctx, client); err != nil {
		t.Fatalf("Failed to wipe collection: %v", err)
	}

	userRepository := NewRepository(client)
	return client, userRepository
}

func TestUserRepository_Integration_SaveUserAndFindByID(t *testing.T) {
	ctx := context.Background()
	client, userRepository := setupTestEnvironment(t)
	defer client.Close(ctx)
	user := &userEntity.User{
		ID:        "1",
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john@example.com",
		Age:       25,
	}
	userRepository.Save(ctx, user)
	foundUser, err := userRepository.FindByID(ctx, "1")
	if err != nil {
		t.Fatalf("Failed to find user: %v", err)
	}
	if foundUser.ID != user.ID {
		t.Fatalf("Found user ID = %v, want %v", foundUser.ID, user.ID)
	}
	if foundUser.FirstName != user.FirstName {
		t.Fatalf("Found user FirstName = %v, want %v", foundUser.FirstName, user.FirstName)
	}
	if foundUser.LastName != user.LastName {
		t.Fatalf("Found user LastName = %v, want %v", foundUser.LastName, user.LastName)
	}
	if foundUser.Email != user.Email {
		t.Fatalf("Found user Email = %v, want %v", foundUser.Email, user.Email)
	}
	if foundUser.Age != user.Age {
		t.Fatalf("Found user Age = %v, want %v", foundUser.Age, user.Age)
	}
}

func TestUserRepository_Integration_SaveUserAndCheckExistsByFirstNameAndLastName(t *testing.T) {
	ctx := context.Background()
	client, userRepository := setupTestEnvironment(t)
	defer client.Close(ctx)
	user := &userEntity.User{
		ID:        "2",
		FirstName: "Jane",
		LastName:  "Doe",
		Email:     "jane@example.com",
		Age:       25,
	}
	_, err := userRepository.Save(ctx, user)
	if err != nil {
		t.Fatalf("Failed to save user: %v", err)
	}

	exists := userRepository.ExistsByFirstNameAndLastName(ctx, "Jane", "Doe")
	if !exists {
		t.Fatalf("User should exist")
	}
}

func TestUserRepository_Integration_SaveUserAndCheckExistsByFirstNameAndLastNameAndIDNot(t *testing.T) {
	ctx := context.Background()
	client, userRepository := setupTestEnvironment(t)
	defer client.Close(ctx)
	user1 := &userEntity.User{
		ID:        "3",
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john@example.com",
		Age:       25,
	}
	_, err := userRepository.Save(ctx, user1)
	if err != nil {
		t.Fatalf("Failed to save user: %v", err)
	}
	user2 := &userEntity.User{
		ID:        "4",
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john2@example.com",
		Age:       25,
	}
	_, err = userRepository.Save(ctx, user2)
	if err != nil {
		t.Fatalf("Failed to save user: %v", err)
	}

	exists := userRepository.ExistsByFirstNameAndLastNameAndIDNot(ctx, "John", "Doe", "3")
	if !exists {
		t.Fatalf("User should exist")
	}
}
