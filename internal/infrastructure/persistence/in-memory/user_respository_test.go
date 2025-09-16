package inmemory

import (
	"context"
	"testing"

	"github.com/surajswarnapuri/ps-tag-onboarding-go/internal/domain/user"
)

func TestRepository_FindByID(t *testing.T) {
	tests := []struct {
		name          string
		existingUsers map[string]*user.User
		searchID      string
		expectedUser  *user.User
		expectedError bool
	}{
		{
			name: "user exists",
			existingUsers: map[string]*user.User{
				"1": {ID: "1", FirstName: "John", LastName: "Doe", Email: "john@example.com", Age: 25},
				"2": {ID: "2", FirstName: "Jane", LastName: "Smith", Email: "jane@example.com", Age: 30},
			},
			searchID:      "1",
			expectedUser:  &user.User{ID: "1", FirstName: "John", LastName: "Doe", Email: "john@example.com", Age: 25},
			expectedError: false,
		},
		{
			name: "user does not exist",
			existingUsers: map[string]*user.User{
				"1": {ID: "1", FirstName: "John", LastName: "Doe", Email: "john@example.com", Age: 25},
			},
			searchID:      "999",
			expectedUser:  nil,
			expectedError: true,
		},
		{
			name:          "empty repository",
			existingUsers: map[string]*user.User{},
			searchID:      "1",
			expectedUser:  nil,
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &repository{users: tt.existingUsers}

			result, err := repo.FindByID(context.Background(), tt.searchID)

			if tt.expectedError {
				if err == nil {
					t.Errorf("FindByID() expected error, got nil")
				}
				return
			}

			if err != nil {
				t.Errorf("FindByID() unexpected error: %v", err)
				return
			}

			if result == nil {
				t.Errorf("FindByID() expected user, got nil")
				return
			}

			if result.ID != tt.expectedUser.ID {
				t.Errorf("FindByID() ID = %v, want %v", result.ID, tt.expectedUser.ID)
			}
			if result.FirstName != tt.expectedUser.FirstName {
				t.Errorf("FindByID() FirstName = %v, want %v", result.FirstName, tt.expectedUser.FirstName)
			}
			if result.LastName != tt.expectedUser.LastName {
				t.Errorf("FindByID() LastName = %v, want %v", result.LastName, tt.expectedUser.LastName)
			}
			if result.Email != tt.expectedUser.Email {
				t.Errorf("FindByID() Email = %v, want %v", result.Email, tt.expectedUser.Email)
			}
			if result.Age != tt.expectedUser.Age {
				t.Errorf("FindByID() Age = %v, want %v", result.Age, tt.expectedUser.Age)
			}
		})
	}
}

func TestRepository_Save(t *testing.T) {
	tests := []struct {
		name          string
		existingUsers map[string]*user.User
		userToSave    *user.User
		expectedError bool
	}{
		{
			name:          "save new user",
			existingUsers: map[string]*user.User{},
			userToSave:    &user.User{ID: "1", FirstName: "John", LastName: "Doe", Email: "john@example.com", Age: 25},
			expectedError: false,
		},
		{
			name: "save user with existing ID (overwrite)",
			existingUsers: map[string]*user.User{
				"1": {ID: "1", FirstName: "Old", LastName: "Name", Email: "old@example.com", Age: 20},
			},
			userToSave:    &user.User{ID: "1", FirstName: "John", LastName: "Doe", Email: "john@example.com", Age: 25},
			expectedError: false,
		},
		{
			name: "save user to non-empty repository",
			existingUsers: map[string]*user.User{
				"1": {ID: "1", FirstName: "Jane", LastName: "Smith", Email: "jane@example.com", Age: 30},
			},
			userToSave:    &user.User{ID: "2", FirstName: "John", LastName: "Doe", Email: "john@example.com", Age: 25},
			expectedError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &repository{users: tt.existingUsers}

			err := repo.Save(context.Background(), tt.userToSave)

			if tt.expectedError && err == nil {
				t.Errorf("Save() expected error, got nil")
			}
			if !tt.expectedError && err != nil {
				t.Errorf("Save() unexpected error: %v", err)
			}

			// Verify user was saved
			savedUser, err := repo.FindByID(context.Background(), tt.userToSave.ID)
			if err != nil {
				t.Errorf("Save() user not found after save: %v", err)
			}
			if savedUser.ID != tt.userToSave.ID {
				t.Errorf("Save() saved user ID = %v, want %v", savedUser.ID, tt.userToSave.ID)
			}
		})
	}
}

func TestRepository_ExistsByFirstNameAndLastName(t *testing.T) {
	tests := []struct {
		name          string
		existingUsers map[string]*user.User
		firstName     string
		lastName      string
		expected      bool
		expectedError bool
	}{
		{
			name: "user exists with exact match",
			existingUsers: map[string]*user.User{
				"1": {ID: "1", FirstName: "John", LastName: "Doe", Email: "john@example.com", Age: 25},
				"2": {ID: "2", FirstName: "Jane", LastName: "Smith", Email: "jane@example.com", Age: 30},
			},
			firstName:     "John",
			lastName:      "Doe",
			expected:      true,
			expectedError: false,
		},
		{
			name: "user does not exist",
			existingUsers: map[string]*user.User{
				"1": {ID: "1", FirstName: "John", LastName: "Doe", Email: "john@example.com", Age: 25},
			},
			firstName:     "Jane",
			lastName:      "Smith",
			expected:      false,
			expectedError: false,
		},
		{
			name: "partial match - first name only",
			existingUsers: map[string]*user.User{
				"1": {ID: "1", FirstName: "John", LastName: "Doe", Email: "john@example.com", Age: 25},
			},
			firstName:     "John",
			lastName:      "Smith",
			expected:      false,
			expectedError: false,
		},
		{
			name: "partial match - last name only",
			existingUsers: map[string]*user.User{
				"1": {ID: "1", FirstName: "John", LastName: "Doe", Email: "john@example.com", Age: 25},
			},
			firstName:     "Jane",
			lastName:      "Doe",
			expected:      false,
			expectedError: false,
		},
		{
			name:          "empty repository",
			existingUsers: map[string]*user.User{},
			firstName:     "John",
			lastName:      "Doe",
			expected:      false,
			expectedError: false,
		},
		{
			name: "case sensitive match",
			existingUsers: map[string]*user.User{
				"1": {ID: "1", FirstName: "John", LastName: "Doe", Email: "john@example.com", Age: 25},
			},
			firstName:     "john",
			lastName:      "doe",
			expected:      false,
			expectedError: false,
		},
		{
			name: "multiple users with same name",
			existingUsers: map[string]*user.User{
				"1": {ID: "1", FirstName: "John", LastName: "Doe", Email: "john1@example.com", Age: 25},
				"2": {ID: "2", FirstName: "John", LastName: "Doe", Email: "john2@example.com", Age: 30},
			},
			firstName:     "John",
			lastName:      "Doe",
			expected:      true,
			expectedError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &repository{users: tt.existingUsers}

			result, err := repo.ExistsByFirstNameAndLastName(context.Background(), tt.firstName, tt.lastName)

			if tt.expectedError && err == nil {
				t.Errorf("ExistsByFirstNameAndLastName() expected error, got nil")
			}
			if !tt.expectedError && err != nil {
				t.Errorf("ExistsByFirstNameAndLastName() unexpected error: %v", err)
			}
			if result != tt.expected {
				t.Errorf("ExistsByFirstNameAndLastName() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestRepository_ExistsByFirstNameAndLastNameAndIDNot(t *testing.T) {
	tests := []struct {
		name          string
		existingUsers map[string]*user.User
		firstName     string
		lastName      string
		excludeID     string
		expected      bool
		expectedError bool
	}{
		{
			name: "user exists with different ID",
			existingUsers: map[string]*user.User{
				"1": {ID: "1", FirstName: "John", LastName: "Doe", Email: "john1@example.com", Age: 25},
				"2": {ID: "2", FirstName: "John", LastName: "Doe", Email: "john2@example.com", Age: 30},
			},
			firstName:     "John",
			lastName:      "Doe",
			excludeID:     "1",
			expected:      true,
			expectedError: false,
		},
		{
			name: "user exists but same ID (should be excluded)",
			existingUsers: map[string]*user.User{
				"1": {ID: "1", FirstName: "John", LastName: "Doe", Email: "john@example.com", Age: 25},
			},
			firstName:     "John",
			lastName:      "Doe",
			excludeID:     "1",
			expected:      false,
			expectedError: false,
		},
		{
			name: "no user exists with that name",
			existingUsers: map[string]*user.User{
				"1": {ID: "1", FirstName: "John", LastName: "Doe", Email: "john@example.com", Age: 25},
			},
			firstName:     "Jane",
			lastName:      "Smith",
			excludeID:     "1",
			expected:      false,
			expectedError: false,
		},
		{
			name:          "empty repository",
			existingUsers: map[string]*user.User{},
			firstName:     "John",
			lastName:      "Doe",
			excludeID:     "1",
			expected:      false,
			expectedError: false,
		},
		{
			name: "multiple users, one excluded by ID",
			existingUsers: map[string]*user.User{
				"1": {ID: "1", FirstName: "John", LastName: "Doe", Email: "john1@example.com", Age: 25},
				"2": {ID: "2", FirstName: "John", LastName: "Doe", Email: "john2@example.com", Age: 30},
				"3": {ID: "3", FirstName: "Jane", LastName: "Smith", Email: "jane@example.com", Age: 28},
			},
			firstName:     "John",
			lastName:      "Doe",
			excludeID:     "1",
			expected:      true,
			expectedError: false,
		},
		{
			name: "case sensitive match",
			existingUsers: map[string]*user.User{
				"1": {ID: "1", FirstName: "John", LastName: "Doe", Email: "john@example.com", Age: 25},
			},
			firstName:     "john",
			lastName:      "doe",
			excludeID:     "2",
			expected:      false,
			expectedError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &repository{users: tt.existingUsers}

			result, err := repo.ExistsByFirstNameAndLastNameAndIDNot(context.Background(), tt.firstName, tt.lastName, tt.excludeID)

			if tt.expectedError && err == nil {
				t.Errorf("ExistsByFirstNameAndLastNameAndIDNot() expected error, got nil")
			}
			if !tt.expectedError && err != nil {
				t.Errorf("ExistsByFirstNameAndLastNameAndIDNot() unexpected error: %v", err)
			}
			if result != tt.expected {
				t.Errorf("ExistsByFirstNameAndLastNameAndIDNot() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestRepository_EdgeCases(t *testing.T) {
	repo := NewRepository()

	t.Run("save user with empty ID", func(t *testing.T) {
		user := &user.User{
			ID:        "",
			FirstName: "Test",
			LastName:  "User",
			Email:     "test@example.com",
			Age:       25,
		}

		err := repo.Save(context.Background(), user)
		if err != nil {
			t.Errorf("Save() with empty ID should not error: %v", err)
		}

		// Should be able to find by empty ID
		found, err := repo.FindByID(context.Background(), "")
		if err != nil {
			t.Errorf("FindByID() with empty ID should not error: %v", err)
		}
		if found == nil {
			t.Errorf("FindByID() with empty ID should return user")
		}
	})

	t.Run("exists with empty names", func(t *testing.T) {
		// Add a user with empty names
		user := &user.User{
			ID:        "1",
			FirstName: "",
			LastName:  "",
			Email:     "test@example.com",
			Age:       25,
		}
		repo.Save(context.Background(), user)

		// Check if empty names exist
		exists, err := repo.ExistsByFirstNameAndLastName(context.Background(), "", "")
		if err != nil {
			t.Errorf("ExistsByFirstNameAndLastName() with empty names should not error: %v", err)
		}
		if !exists {
			t.Errorf("ExistsByFirstNameAndLastName() with empty names should return true")
		}
	})
}
