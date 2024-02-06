package storage

import (
	. "github.com/google/uuid"
	. "split-the-bill-server/domain/model"
)

type IInvitationStorage interface {
	// AddGroupInvitation adds the given group invitation to the storage.
	AddGroupInvitation(invitation GroupInvitationModel) (GroupInvitationModel, error)
	// AcceptGroupInvitation adds the associated user to a group and deletes the invitation.
	AcceptGroupInvitation(invitationID UUID, userID UUID) error
	// GetGroupInvitationByID returns the group invitation with the given ID, or a NoSuchGroupInvitationError if no such group invitation exists.
	GetGroupInvitationByID(id UUID) (GroupInvitationModel, error)
}
