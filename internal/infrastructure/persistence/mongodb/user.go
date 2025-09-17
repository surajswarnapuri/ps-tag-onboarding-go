package mongodb

import userEntity "github.com/surajswarnapuri/ps-tag-onboarding-go/internal/domain/user"

type user struct {
	ID        string `bson:"_id,omitempty"`
	FirstName string `bson:"first_name,omitempty"`
	LastName  string `bson:"last_name,omitempty"`
	Email     string `bson:"email,omitempty"`
	Age       int    `bson:"age,omitempty"`
}

func (u *user) ToEntity() *userEntity.User {
	return &userEntity.User{
		ID:        u.ID,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
		Age:       u.Age,
	}
}

func (u *user) FromEntity(user *userEntity.User) {
	u.ID = user.ID
	u.FirstName = user.FirstName
	u.LastName = user.LastName
	u.Email = user.Email
	u.Age = user.Age
}
