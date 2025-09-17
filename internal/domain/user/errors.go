package user

import "github.com/surajswarnapuri/ps-tag-onboarding-go/internal/domain/shared"

const (
	ErrorAgeMinimum    = "AGE_MINIMUM"
	ErrorEmailFormat   = "EMAIL_FORMAT"
	ErrorEmailRequired = "EMAIL_REQUIRED"
	ErrorNameRequired  = "NAME_REQUIRED"
)

// Error Constructors
// NewAgeMinimumError creates a new age minimum error
func NewAgeMinimumError() shared.ValidationError {
	return shared.ValidationError{
		Code:    ErrorAgeMinimum,
		Message: "User does not meet minimum age requirement",
	}
}

// NewEmailFormatError creates a new email format error
func NewEmailFormatError() shared.ValidationError {
	return shared.ValidationError{
		Code:    ErrorEmailFormat,
		Message: "User email must be properly formatted",
	}
}

// NewEmailRequiredError creates a new email required error
func NewEmailRequiredError() shared.ValidationError {
	return shared.ValidationError{
		Code:    ErrorEmailRequired,
		Message: "User email is required",
	}
}

// NewNameRequiredError creates a new name required error
func NewNameRequiredError() shared.ValidationError {
	return shared.ValidationError{
		Code:    ErrorNameRequired,
		Message: "User first/last name is required",
	}
}
