package storage

import (
	. "github.com/google/uuid"
	"split-the-bill-server/domain/model"
)

type IGroupStorage interface {
	// AddGroup adds the given group to the storage. If a group with the same ID or name already exists, a GroupAlreadyExistsError is returned.
	AddGroup(group model.Group) (model.Group, error)

	// UpdateGroup updates the group
	UpdateGroup(group model.Group) (model.Group, error)

	// GetGroupByID returns the group with the given ID, or a NoSuchGroupError if no such group exists.
	GetGroupByID(id UUID) (model.Group, error)

	// GetGroups returns all groups for the user with the given user ID and/or invitation ID.
	GetGroups(userID UUID, invitationID UUID) ([]model.Group, error)

	// DeleteGroup deletes the group with the given ID, or a NoSuchGroupError if no such group exists.
	DeleteGroup(id UUID) error

	// AcceptGroupInvitation adds the associated user to a group
	AcceptGroupInvitation(invitationID UUID, userID UUID) error
}
