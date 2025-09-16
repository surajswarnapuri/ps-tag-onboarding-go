package user

import (
	"errors"
	"strings"
)

type ValidationService struct{}

func NewValidationService() *ValidationService {
	return &ValidationService{}
}

func (s *ValidationService) ValidateUser(user User) error {
	return errors.Join(
		s.validateAge(user),
		s.validateEmail(user),
		s.validateName(user),
	)
}

func (s *ValidationService) validateAge(user User) error {
	if user.Age < 18 {
		return NewAgeMinimumError()
	}
	return nil
}

func (s *ValidationService) validateEmail(user User) error {
	if user.Email == "" {
		return NewEmailRequiredError()
	}
	if !strings.Contains(user.Email, "@") {
		return NewEmailFormatError()
	}
	return nil
}

func (s *ValidationService) validateName(user User) error {
	if user.FirstName == "" || user.LastName == "" {
		return NewNameRequiredError()
	}
	return nil
}
