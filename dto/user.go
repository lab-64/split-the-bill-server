package dto

import (
	"github.com/google/uuid"
	"split-the-bill-server/types"
)

type UserCreateDTO struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
}

type UserDTO struct {
	ID          uuid.UUID   `json:"id"`
	Username    string      `json:"username"`
	Email       string      `json:"email"`
	Groups      []GroupDTO  `json:"groups"`
	Invitations []uuid.UUID `json:"invitations"`
}

func (r UserCreateDTO) ToUser() types.User {
	return types.CreateUser(r.Username, r.Email)
}

func ToUserDTO(u *types.User) UserDTO {
	groupsDTO := make([]GroupDTO, len(u.Groups))

	for i, group := range u.Groups {
		groupsDTO[i] = ToGroupDTO(group)
	}

	invitations := make([]uuid.UUID, len(u.PendingGroupInvitations))

	for i, inv := range u.PendingGroupInvitations {
		invitations[i] = inv.ID
	}

	return UserDTO{
		ID:          u.ID,
		Username:    u.Username,
		Email:       u.Email,
		Groups:      groupsDTO,
		Invitations: invitations,
	}
}
