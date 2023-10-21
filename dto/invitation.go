package dto

import "github.com/google/uuid"

type InvitationReplyDTO struct {
	Type   string    `json:"type"`
	ID     uuid.UUID `json:"id"`
	Accept bool      `json:"accept"`
}
