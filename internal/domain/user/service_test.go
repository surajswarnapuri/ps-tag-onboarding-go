package user

import (
	"errors"
	"testing"
)

func TestValidateUser(t *testing.T) {
	tests := []struct {
		name string
		user User
		want error
	}{
		{
			name: "valid user",
			user: User{
				ID:        "1",
				FirstName: "John",
				LastName:  "Doe",
				Email:     "john.doe@example.com",
				Age:       20,
			},
			want: nil,
		},
		{
			name: "user with invalid age",
			user: User{
				ID:        "1",
				FirstName: "John",
				LastName:  "Doe",
				Email:     "john.doe@example.com",
				Age:       17,
			},
			want: NewAgeMinimumError(),
		},
		{
			name: "user without an email",
			user: User{
				ID:        "1",
				FirstName: "John",
				LastName:  "Doe",
				Email:     "",
				Age:       20,
			},
			want: NewEmailRequiredError(),
		},
		{
			name: "user with invalid email",
			user: User{
				ID:        "1",
				FirstName: "John",
				LastName:  "Doe",
				Email:     "john.doe",
				Age:       20,
			},
			want: NewEmailFormatError(),
		},
		{
			name: "user with invalid first name",
			user: User{
				ID:        "1",
				FirstName: "",
				LastName:  "Doe",
				Email:     "john.doe@example.com",
				Age:       20,
			},
			want: NewNameRequiredError(),
		},
		{
			name: "user with invalid last name",
			user: User{
				ID:        "1",
				FirstName: "John",
				LastName:  "",
				Email:     "john.doe@example.com",
				Age:       20,
			},
			want: NewNameRequiredError(),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := NewValidationService().ValidateUser(test.user)
			if !errors.Is(got, test.want) {
				t.Errorf("ValidateUser(%v) = %v, want %v", test.user, got, test.want)
			}
		})
	}
}
