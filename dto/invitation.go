package dto

import "github.com/google/uuid"

type InvitationInputDTO struct {
	Type   string    `json:"type"`
	ID     uuid.UUID `json:"id"`
	Accept bool      `json:"accept"`
}
