package storage_inf

import (
	. "github.com/google/uuid"
	. "split-the-bill-server/domain/model"
)

type IInvitationStorage interface {
	// AddGroupInvitation adds the given group invitation to the storage.
	AddGroupInvitation(invitation GroupInvitationModel) error
	// DeleteGroupInvitation deletes the group invitation with the given ID from the storage, if it exists.
	DeleteGroupInvitation(id UUID) error
	// GetGroupInvitationByID returns the group invitation with the given ID, or a NoSuchGroupInvitationError if no such group invitation exists.
	GetGroupInvitationByID(id UUID) (GroupInvitationModel, error)
	// GetGroupInvitationsByUserID returns all group invitations for the given user.
	GetGroupInvitationsByUserID(userID UUID) ([]GroupInvitationModel, error)
}
