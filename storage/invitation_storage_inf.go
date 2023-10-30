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
}
