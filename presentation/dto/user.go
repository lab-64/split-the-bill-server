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
	ID            uuid.UUID            `json:"id"`
	Email         string               `json:"email"`
	Username      string               `json:"username"`
	Groups        []GroupCoreOutputDTO `json:"groups"`
	InvitationIDs []uuid.UUID          `json:"invitationIDs"`
}

func ConvertToUserDetailedDTO(u *model.UserModel) UserDetailedOutputDTO {
	groupsDTO := make([]GroupCoreOutputDTO, len(u.Groups))

	for i, group := range u.Groups {
		groupsDTO[i] = ConvertToGroupCoreDTO(group)
	}

	invitations := make([]uuid.UUID, len(u.PendingGroupInvitations))

	for i, inv := range u.PendingGroupInvitations {
		invitations[i] = inv.ID
	}

	return UserDetailedOutputDTO{
		ID:            u.ID,
		Email:         u.Email,
		Username:      u.Username,
		Groups:        groupsDTO,
		InvitationIDs: invitations,
	}
}
