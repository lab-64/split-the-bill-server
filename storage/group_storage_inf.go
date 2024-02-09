package storage

import (
	. "github.com/google/uuid"
	. "split-the-bill-server/domain/model"
)

type IGroupStorage interface {
	// AddGroup adds the given group to the storage. If a group with the same ID or name already exists, a GroupAlreadyExistsError is returned.
	AddGroup(group GroupModel) (GroupModel, error)

	// UpdateGroup updates the group
	UpdateGroup(group GroupModel) (GroupModel, error)

	// GetGroupByID returns the group with the given ID, or a NoSuchGroupError if no such group exists.
	GetGroupByID(id UUID) (GroupModel, error)

	// GetGroups returns all groups for the user with the given user ID and/or invitation ID.
	GetGroups(userID UUID, invitationID UUID) ([]GroupModel, error)

	// AcceptGroupInvitation adds the associated user to a group
	AcceptGroupInvitation(invitationID UUID, userID UUID) error
}
