// Package user contains the logic for the User application service.
package user

import (
	"context"
	"fmt"

	"github.com/surajswarnapuri/ps-tag-onboarding-go/internal/domain/user"
)

type userValidationService interface {
	ValidateUser(user user.User) error
}

type service struct {
	userValidationService userValidationService
	userRepository        user.Repository
}

func NewService(userValidationService userValidationService, userRepository user.Repository) *service {
	return &service{
		userValidationService: userValidationService,
		userRepository:        userRepository,
	}
}

func (s *service) Find(ctx context.Context, id string) (*user.User, error) {
	user, err := s.userRepository.FindByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("service: failed to find user by ID %q: %w", id, err)
	}
	return user, nil
}

func (s *service) Save(ctx context.Context, user *user.User) (*user.User, error) {
	err := s.userValidationService.ValidateUser(*user)
	if err != nil {
		return nil, fmt.Errorf("service: failed to validate user: %w", err)
	}

	if s.nameCombinationExists(ctx, user) {
		return nil, fmt.Errorf("service: name combination already exists")
	}

	return s.userRepository.Save(ctx, user)
}

func (s *service) nameCombinationExists(ctx context.Context, user *user.User) bool {
	if user.ID == "" {
		return s.userRepository.ExistsByFirstNameAndLastName(ctx, user.FirstName, user.LastName)
	}
	return s.userRepository.ExistsByFirstNameAndLastNameAndIDNot(ctx, user.FirstName, user.LastName, user.ID)
}
