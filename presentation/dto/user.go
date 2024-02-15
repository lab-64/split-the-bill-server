package dto

import (
	"github.com/google/uuid"
)

type UserInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserUpdate struct {
	Username string `json:"username" form:"username"`
}

type UserCoreOutput struct {
	ID             uuid.UUID `json:"id"`
	Email          string    `json:"email"`
	Username       string    `json:"username"`
	ProfileImgPath string    `json:"profileImgPath"`
}
