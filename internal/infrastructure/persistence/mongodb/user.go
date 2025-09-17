package mongodb

import userEntity "github.com/surajswarnapuri/ps-tag-onboarding-go/internal/domain/user"

type user struct {
	ID        string `bson:"id"`
	FirstName string `bson:"first_name"`
	LastName  string `bson:"last_name"`
	Email     string `bson:"email"`
	Age       int    `bson:"age"`
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
