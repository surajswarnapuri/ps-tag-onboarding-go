package user

import userDomain "github.com/surajswarnapuri/ps-tag-onboarding-go/internal/domain/user"

// UserDTO is a user data transfer object
type UserDTO struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Age       int    `json:"age"`
}

// ToEntity converts a UserDTO to a userDomain.User
func (u *UserDTO) ToEntity() *userDomain.User {
	return &userDomain.User{
		ID:        u.ID,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
		Age:       u.Age,
	}
}

// FromEntity converts a userDomain.User to a UserDTO
func (u *UserDTO) FromEntity(user *userDomain.User) {
	u.ID = user.ID
	u.FirstName = user.FirstName
	u.LastName = user.LastName
	u.Email = user.Email
	u.Age = user.Age
}
