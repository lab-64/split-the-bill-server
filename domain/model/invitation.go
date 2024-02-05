package model

import (
	"github.com/google/uuid"
	"time"
)

// GroupInvitationModel represents an invitation to a group used to generalize between the user inputs and the database.
// TODO: maybe add UserID to GroupInvitationModel as a foreignkey
type GroupInvitationModel struct {
	ID      uuid.UUID
	Date    time.Time
	Group   GroupModel
	Invitee UserModel
}
