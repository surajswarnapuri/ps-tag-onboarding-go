package user

import (
	"errors"
	"strings"
)

type validationService struct{}

func NewValidationService() *validationService {
	return &validationService{}
}

func (s *validationService) ValidateUser(user User) error {
	return errors.Join(
		s.validateAge(user),
		s.validateEmail(user),
		s.validateName(user),
	)
}

func (s *validationService) validateAge(user User) error {
	if user.Age < 18 {
		return NewAgeMinimumError()
	}
	return nil
}

func (s *validationService) validateEmail(user User) error {
	if user.Email == "" {
		return NewEmailRequiredError()
	}
	if !strings.Contains(user.Email, "@") {
		return NewEmailFormatError()
	}
	return nil
}

func (s *validationService) validateName(user User) error {
	if user.FirstName == "" || user.LastName == "" {
		return NewNameRequiredError()
	}
	return nil
}
