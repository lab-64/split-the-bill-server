package dto

import "github.com/google/uuid"

type InvitationInputDTO struct {
	ID     uuid.UUID `json:"id"`
	User   uuid.UUID `json:"user"`
	Type   string    `json:"type"`
	Accept bool      `json:"accept"`
}
