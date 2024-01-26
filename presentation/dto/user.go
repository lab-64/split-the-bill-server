package dto

import (
	"github.com/google/uuid"
	. "split-the-bill-server/domain/model"
)

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Input Section
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type UserInputDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func ToUserModel(r UserInputDTO) UserModel {
	return CreateUserModel(r.Email)
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Output Section
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type UserCoreOutputDTO struct {
	ID    uuid.UUID `json:"id"`
	Email string    `json:"email"`
}

func ToUserCoreDTO(u *UserModel) UserCoreOutputDTO {
	return UserCoreOutputDTO{
		ID:    u.ID,
		Email: u.Email,
	}
}

func ToUserCoreDTOs(users []UserModel) []UserCoreOutputDTO {
	usersDTO := make([]UserCoreOutputDTO, len(users))

	for i, user := range users {
		usersDTO[i] = ToUserCoreDTO(&user)
	}
	return usersDTO
}

type UserDetailedOutputDTO struct {
	ID          uuid.UUID            `json:"id"`
	Email       string               `json:"email"`
	Groups      []GroupCoreOutputDTO `json:"groups"`
	Invitations []uuid.UUID          `json:"invitationIDs"`
}

func ToUserDetailedDTO(u *UserModel) UserDetailedOutputDTO {
	groupsDTO := make([]GroupCoreOutputDTO, len(u.Groups))

	for i, group := range u.Groups {
		groupsDTO[i] = ToGroupCoreDTO(group)
	}

	invitations := make([]uuid.UUID, len(u.PendingGroupInvitations))

	for i, inv := range u.PendingGroupInvitations {
		invitations[i] = inv.ID
	}

	return UserDetailedOutputDTO{
		ID:          u.ID,
		Email:       u.Email,
		Groups:      groupsDTO,
		Invitations: invitations,
	}
}

// TODO: maybe delete and use UserCoreOutputDTO instead
type UserPublicOutputDTO struct {
	ID    uuid.UUID `json:"id"`
	Email string    `json:"email"`
}

func ToUserPublicDTO(u UserModel) UserPublicOutputDTO {
	return UserPublicOutputDTO{
		ID:    u.ID,
		Email: u.Email,
	}
}
