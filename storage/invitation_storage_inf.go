package storage

import (
	"github.com/google/uuid"
	"split-the-bill-server/types"
)

type IInvitationStorage interface {
	// AddGroupInvitation adds the given group invitation to the storage.
	AddGroupInvitation(invitation types.GroupInvitation) error
	// DeleteGroupInvitation deletes the group invitation with the given ID from the storage, if it exists.
	DeleteGroupInvitation(id uuid.UUID) error
	// GetGroupInvitationByID returns the group invitation with the given ID, or a NoSuchGroupInvitationError if no such group invitation exists.
	GetGroupInvitationByID(id uuid.UUID) (types.GroupInvitation, error)
	// GetGroupInvitationsByUserID returns all group invitations for the given user.
	GetGroupInvitationsByUserID(userID uuid.UUID) ([]types.GroupInvitation, error)
}
