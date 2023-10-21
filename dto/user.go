package dto

import (
	"github.com/google/uuid"
	"split-the-bill-server/types"
)

type UserInputDTO struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
}

type UserOutputDTO struct {
	ID          uuid.UUID        `json:"id"`
	Username    string           `json:"username"`
	Email       string           `json:"email"`
	Groups      []GroupOutputDTO `json:"groups"`
	Invitations []uuid.UUID      `json:"invitations"`
}

func (r UserInputDTO) ToUser() types.User {
	return types.CreateUser(r.Username, r.Email)
}

func ToUserDTO(u *types.User) UserOutputDTO {
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
		Username:    u.Username,
		Email:       u.Email,
		Groups:      groupsDTO,
		Invitations: invitations,
	}
}
