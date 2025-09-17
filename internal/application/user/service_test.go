package user

import (
	"context"
	"fmt"
	"strings"
	"testing"

	userDomain "github.com/surajswarnapuri/ps-tag-onboarding-go/internal/domain/user"
)

type mockUserValidationService struct {
	ValidateUserFunc func(user userDomain.User) error
}

type mockUserRepository struct {
	FindByIDFunc                             func(ctx context.Context, id string) (*userDomain.User, error)
	SaveFunc                                 func(ctx context.Context, user *userDomain.User) (*userDomain.User, error)
	ExistsByFirstNameAndLastNameFunc         func(ctx context.Context, firstName string, lastName string) bool
	ExistsByFirstNameAndLastNameAndIDNotFunc func(ctx context.Context, firstName string, lastName string, id string) bool
}

func (m *mockUserRepository) FindByID(ctx context.Context, id string) (*userDomain.User, error) {
	return m.FindByIDFunc(ctx, id)
}
func (m *mockUserRepository) Save(ctx context.Context, user *userDomain.User) (*userDomain.User, error) {
	return m.SaveFunc(ctx, user)
}
func (m *mockUserRepository) ExistsByFirstNameAndLastName(ctx context.Context, firstName string, lastName string) bool {
	return m.ExistsByFirstNameAndLastNameFunc(ctx, firstName, lastName)
}
func (m *mockUserRepository) ExistsByFirstNameAndLastNameAndIDNot(ctx context.Context, firstName string, lastName string, id string) bool {
	return m.ExistsByFirstNameAndLastNameAndIDNotFunc(ctx, firstName, lastName, id)
}
func (m *mockUserValidationService) ValidateUser(user userDomain.User) error {
	return m.ValidateUserFunc(user)
}

func TestService_Find(t *testing.T) {
	tests := []struct {
		name                      string
		userID                    string
		mockUserRepository        *mockUserRepository
		mockUserValidationService *mockUserValidationService
		expectedUser              *userDomain.User
		expectedError             bool
		errorContains             string
	}{
		{
			name:   "user found",
			userID: "1",
			mockUserRepository: &mockUserRepository{
				FindByIDFunc: func(ctx context.Context, id string) (*userDomain.User, error) {
					return &userDomain.User{ID: "1", FirstName: "John", LastName: "Doe", Email: "john@example.com", Age: 25}, nil
				},
			},
			mockUserValidationService: &mockUserValidationService{
				ValidateUserFunc: func(user userDomain.User) error {
					return nil
				},
			},
			expectedUser:  &userDomain.User{ID: "1", FirstName: "John", LastName: "Doe", Email: "john@example.com", Age: 25},
			expectedError: false,
		},
		{
			name:   "user not found",
			userID: "2",
			mockUserRepository: &mockUserRepository{
				FindByIDFunc: func(ctx context.Context, id string) (*userDomain.User, error) {
					return nil, fmt.Errorf("user not found")
				},
			},
			mockUserValidationService: &mockUserValidationService{
				ValidateUserFunc: func(user userDomain.User) error {
					return nil
				},
			},
			expectedUser:  nil,
			expectedError: true,
			errorContains: "failed to find user by ID",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			service := NewService(test.mockUserValidationService, test.mockUserRepository)
			user, err := service.Find(context.Background(), test.userID)

			// Check error
			if test.expectedError {
				if err == nil {
					t.Errorf("Find() expected error, got nil")
					return
				}
				if test.errorContains != "" && !strings.Contains(err.Error(), test.errorContains) {
					t.Errorf("Find() error = %v, want error containing %q", err, test.errorContains)
				}
				return
			}

			if err != nil {
				t.Errorf("Find() unexpected error: %v", err)
				return
			}

			// Check user
			if test.expectedUser == nil {
				if user != nil {
					t.Errorf("Find() expected nil user, got %v", user)
				}
				return
			}

			if user == nil {
				t.Errorf("Find() expected user, got nil")
				return
			}

			if user.ID != test.expectedUser.ID {
				t.Errorf("Find() user ID = %v, want %v", user.ID, test.expectedUser.ID)
			}
			if user.FirstName != test.expectedUser.FirstName {
				t.Errorf("Find() user FirstName = %v, want %v", user.FirstName, test.expectedUser.FirstName)
			}
			if user.LastName != test.expectedUser.LastName {
				t.Errorf("Find() user LastName = %v, want %v", user.LastName, test.expectedUser.LastName)
			}
			if user.Email != test.expectedUser.Email {
				t.Errorf("Find() user Email = %v, want %v", user.Email, test.expectedUser.Email)
			}
			if user.Age != test.expectedUser.Age {
				t.Errorf("Find() user Age = %v, want %v", user.Age, test.expectedUser.Age)
			}
		})
	}
}

func TestService_Save(t *testing.T) {

	tests := []struct {
		name                      string
		user                      userDomain.User
		expectedError             bool
		errorContains             string
		mockUserValidationService *mockUserValidationService
		mockUserRepository        *mockUserRepository
	}{
		{
			name:          "save a valid user",
			user:          userDomain.User{ID: "1", FirstName: "John", LastName: "Doe", Email: "john@example.com", Age: 25},
			expectedError: false,
			errorContains: "",
			mockUserValidationService: &mockUserValidationService{
				ValidateUserFunc: func(user userDomain.User) error {
					return nil
				},
			},
			mockUserRepository: &mockUserRepository{
				SaveFunc: func(ctx context.Context, user *userDomain.User) (*userDomain.User, error) {
					return nil, nil
				},
				ExistsByFirstNameAndLastNameAndIDNotFunc: func(ctx context.Context, firstName string, lastName string, id string) bool {
					return false
				},
			},
		},
		{
			name:          "save a user with existing name combination",
			user:          userDomain.User{ID: "2", FirstName: "John", LastName: "Doe", Email: "john@example.com", Age: 25},
			expectedError: true,
			errorContains: "name combination already exists",
			mockUserValidationService: &mockUserValidationService{
				ValidateUserFunc: func(user userDomain.User) error {
					return nil
				},
			},
			mockUserRepository: &mockUserRepository{
				ExistsByFirstNameAndLastNameAndIDNotFunc: func(ctx context.Context, firstName string, lastName string, id string) bool {
					return true
				},
			},
		},
		{
			name:          "save a user with existing name combination but no ID",
			user:          userDomain.User{ID: "", FirstName: "John", LastName: "Doe", Email: "john@example.com", Age: 25},
			expectedError: true,
			errorContains: "name combination already exists",
			mockUserValidationService: &mockUserValidationService{
				ValidateUserFunc: func(user userDomain.User) error {
					return nil
				},
			},
			mockUserRepository: &mockUserRepository{
				ExistsByFirstNameAndLastNameFunc: func(ctx context.Context, firstName string, lastName string) bool {
					return true
				},
			},
		},
		{
			name:          "save a user with invalid validation",
			user:          userDomain.User{ID: "3", FirstName: "John", LastName: "Doe", Email: "john@example.com", Age: 25},
			expectedError: true,
			errorContains: "invalid validation",
			mockUserValidationService: &mockUserValidationService{
				ValidateUserFunc: func(user userDomain.User) error {
					return fmt.Errorf("invalid validation")
				},
			},
			mockUserRepository: &mockUserRepository{
				ExistsByFirstNameAndLastNameAndIDNotFunc: func(ctx context.Context, firstName string, lastName string, id string) bool {
					return false
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			service := NewService(test.mockUserValidationService, test.mockUserRepository)
			_, err := service.Save(context.Background(), &test.user)

			if test.expectedError {
				if err == nil {
					t.Errorf("Save() expected error, got nil")
					return
				}
				if test.errorContains != "" && !strings.Contains(err.Error(), test.errorContains) {
					t.Errorf("Save() error = %v, want error containing %q", err, test.errorContains)
				}
				return
			}

			if err != nil {
				t.Errorf("Save() unexpected error: %v", err)
			}
		})
	}
}
