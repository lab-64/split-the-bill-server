package dto

import (
	"github.com/google/uuid"
	"split-the-bill-server/domain/model"
)

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Input Section
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type UserInputDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserUpdateDTO struct {
	Username string `json:"username"`
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Output Section
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// TODO: maybe use only one single DTO for the user output
type UserCoreOutputDTO struct {
	ID       uuid.UUID `json:"id"`
	Email    string    `json:"email"`
	Username string    `json:"username"`
}

func ConvertToUserCoreDTO(u *model.UserModel) UserCoreOutputDTO {
	return UserCoreOutputDTO{
		ID:       u.ID,
		Email:    u.Email,
		Username: u.Username,
	}
}

func ConvertToUserCoreDTOs(users []model.UserModel) []UserCoreOutputDTO {
	usersDTO := make([]UserCoreOutputDTO, len(users))

	for i, user := range users {
		usersDTO[i] = ConvertToUserCoreDTO(&user)
	}
	return usersDTO
}

type UserDetailedOutputDTO struct {
	ID       uuid.UUID `json:"id"`
	Email    string    `json:"email"`
	Username string    `json:"username"`
}

func ConvertToUserDetailedDTO(u *model.UserModel) UserDetailedOutputDTO {

	return UserDetailedOutputDTO{
		ID:       u.ID,
		Email:    u.Email,
		Username: u.Username,
	}
}
