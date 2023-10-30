package types

import (
	"github.com/google/uuid"
	"time"
)

// TODO: maybe add UserID to GroupInvitation as a foreignkey
// GroupInvitation represents an invitation to a group used to generalize between the user inputs and the database.
// During Create (From User Input to Database) the GroupID is used to specify the group. Group is ignored.
// During Read (From Database to User Output) Group is used to specify the group.
type GroupInvitation struct {
	ID      uuid.UUID
	Date    time.Time
	Group   Group
	Invitee User
}

// TODO: check usages, function has changed
func CreateGroupInvitation(groupID uuid.UUID, inviteeID uuid.UUID) GroupInvitation {
	return GroupInvitation{
		ID:      uuid.New(),
		Date:    time.Now(),
		Group:   Group{ID: groupID},
		Invitee: User{ID: inviteeID},
	}
}
