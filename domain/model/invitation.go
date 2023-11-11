package model

import (
	"github.com/google/uuid"
	"time"
)

// TODO: maybe add UserID to GroupInvitationModel as a foreignkey
// GroupInvitationModel represents an invitation to a group used to generalize between the user inputs and the database.
// During Create (From UserModel Input to Database) the GroupID is used to specify the group. GroupModel is ignored.
// During Read (From Database to UserModel Output) GroupModel is used to specify the group.
type GroupInvitationModel struct {
	ID      uuid.UUID
	Date    time.Time
	Group   GroupModel
	Invitee UserModel
}

// TODO: check usages, function has changed
func CreateGroupInvitation(groupID uuid.UUID, inviteeID uuid.UUID) GroupInvitationModel {
	return GroupInvitationModel{
		ID:      uuid.New(),
		Date:    time.Now(),
		Group:   GroupModel{ID: groupID},
		Invitee: UserModel{ID: inviteeID},
	}
}
