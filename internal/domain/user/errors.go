package user

import "github.com/surajswarnapuri/ps-tag-onboarding-go/internal/domain/shared"

const (
	ErrorAgeMinimum    = "AGE_MINIMUM"
	ErrorEmailFormat   = "EMAIL_FORMAT"
	ErrorEmailRequired = "EMAIL_REQUIRED"
	ErrorNameRequired  = "NAME_REQUIRED"
)

// Error Constructors
func NewAgeMinimumError() shared.ValidationError {
	return shared.ValidationError{
		Code:    ErrorAgeMinimum,
		Message: "User does not meet minimum age requirement",
	}
}

func NewEmailFormatError() shared.ValidationError {
	return shared.ValidationError{
		Code:    ErrorEmailFormat,
		Message: "User email must be properly formatted",
	}
}

func NewEmailRequiredError() shared.ValidationError {
	return shared.ValidationError{
		Code:    ErrorEmailRequired,
		Message: "User email is required",
	}
}

func NewNameRequiredError() shared.ValidationError {
	return shared.ValidationError{
		Code:    ErrorNameRequired,
		Message: "User first/last name is required",
	}
}
