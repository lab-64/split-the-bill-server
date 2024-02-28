package dto

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
)

var allowedImageTypes = map[string]string{
	"image/jpeg": "jpeg",
	"image/png":  "png",
	"image/gif":  "gif",
	"image/jpg":  "jpg",
}
var allowedImageTypesString = "jpg, png, gif, jpeg"

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Input/Output DTOs
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

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

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Validators
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (u UserInput) ValidateInputs() error {
	if u.Email == "" {
		return ErrEmailRequired
	}
	if u.Password == "" {
		return ErrPasswordRequired
	}
	return nil
}

func (u UserUpdate) ValidateInputs(contentType string) error {
	if _, ok := allowedImageTypes[contentType]; !ok {
		return ErrWrongImageType
	}
	return nil
}

var ErrEmailRequired = errors.New("email is required")
var ErrPasswordRequired = errors.New("password is required")
var ErrWrongImageType = fmt.Errorf("uploaded image has wrong type. Allowed types: %s", allowedImageTypesString)
