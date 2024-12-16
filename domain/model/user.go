package model

import (
	"github.com/google/uuid"
)

type User struct {
	ID             uuid.UUID
	Email          string
	Username       string
	ProfileImgPath string
}

func CreateUser(id uuid.UUID, email string, username string, profileImg string) User {
	return User{
		ID:             id,
		Email:          email,
		Username:       username,
		ProfileImgPath: profileImg,
	}
}

// UpdateNotEmpty updates the user with the given data. If data is empty, it will not be updated.
func (u *User) UpdateNotEmpty(email string, username string, profileImg string) {
	if email != "" {
		u.Email = email
	}
	if username != "" {
		u.Username = username
	}
	if profileImg != "" {
		u.ProfileImgPath = profileImg
	}
}
