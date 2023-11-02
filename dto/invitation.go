package dto

import "github.com/google/uuid"

type GroupInvitationDTO struct {
	Issuer   uuid.UUID   `json:"issuer"`
	GroupID  uuid.UUID   `json:"groupID"`
	Invitees []uuid.UUID `json:"invitees"`
}

type InvitationInputDTO struct {
	Type   string    `json:"type"`
	ID     uuid.UUID `json:"id"`
	Accept bool      `json:"accept"`
}
