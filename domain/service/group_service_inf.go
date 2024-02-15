package service

import (
	"github.com/google/uuid"
	"split-the-bill-server/presentation/dto"
)

type IGroupService interface {
	// Create creates a new group with the given data.
	// *Authorization required: requesterID == OwnerID
	Create(requesterID uuid.UUID, groupDTO dto.GroupInput) (dto.GroupDetailedOutput, error)

	// Update updates the group with the given id with the new group data.
	// *Authorization required: requester == group.Owner
	Update(requesterID uuid.UUID, groupID uuid.UUID, group dto.GroupInput) (dto.GroupDetailedOutput, error)

	// GetByID returns the group with the given id.
	// *Authorization required: requester in group.Members
	GetByID(requesterID uuid.UUID, id uuid.UUID) (dto.GroupDetailedOutput, error)

	// GetAll returns all groups in which the user is a member or for which the invitation applies.
	// *Authorization required: requesterID == id (for param: userID)
	GetAll(requesterID uuid.UUID, userID uuid.UUID, invitationID uuid.UUID) ([]dto.GroupDetailedOutput, error)

	// Delete deletes the group with the given id.
	// *Authorization required: requester == group.Owner
	Delete(requesterID uuid.UUID, id uuid.UUID) error

	// AcceptGroupInvitation accepts the invitation and adds user to the group.
	AcceptGroupInvitation(invitationID uuid.UUID, userID uuid.UUID) error
}
