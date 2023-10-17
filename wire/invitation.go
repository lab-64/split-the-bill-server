package wire

import "github.com/google/uuid"

type InvitationReply struct {
	Type   string    `json:"type"`
	ID     uuid.UUID `json:"id"`
	Accept bool      `json:"accept"`
}
