package dto

import (
	"github.com/google/uuid"
	. "split-the-bill-server/domain/model"
)

type UserInputDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserOutputDTO struct {
	ID          uuid.UUID        `json:"id"`
	Email       string           `json:"email"`
	Groups      []GroupOutputDTO `json:"groups"`
	Invitations []uuid.UUID      `json:"invitationIDs"`
}

func ToUserModel(r UserInputDTO) UserModel {
	return CreateUserModel(r.Email)
}

func ToUserDTO(u *UserModel) UserOutputDTO {
	groupsDTO := make([]GroupOutputDTO, len(u.Groups))

	for i, group := range u.Groups {
		groupsDTO[i] = ToGroupDTO(group)
	}

	invitations := make([]uuid.UUID, len(u.PendingGroupInvitations))

	for i, inv := range u.PendingGroupInvitations {
		invitations[i] = inv.ID
	}

	return UserOutputDTO{
		ID:          u.ID,
		Email:       u.Email,
		Groups:      groupsDTO,
		Invitations: invitations,
	}
}
