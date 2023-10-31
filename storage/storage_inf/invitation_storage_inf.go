package storage_inf

import (
	"github.com/google/uuid"
	"split-the-bill-server/domain/model"
)

type IInvitationStorage interface {
	// AddGroupInvitation adds the given group invitation to the storage.
	AddGroupInvitation(invitation model.GroupInvitation) error
	// DeleteGroupInvitation deletes the group invitation with the given ID from the storage, if it exists.
	DeleteGroupInvitation(id uuid.UUID) error
}
