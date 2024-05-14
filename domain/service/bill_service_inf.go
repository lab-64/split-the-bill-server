package service

import (
	"github.com/google/uuid"
	"split-the-bill-server/presentation/dto"
)

type IBillService interface {
	// Create creates a new bill and returns a detailed description of the created bill.
	// *Authorization required: requester in group.Members
	Create(requesterID uuid.UUID, bill dto.BillCreate) (dto.BillDetailedOutput, error)

	// Update updates the bill with the given id with the new bill data.
	// *Authorization required: requester in group.Member
	Update(requesterID uuid.UUID, billID uuid.UUID, billDTO dto.BillUpdate) (dto.BillDetailedOutput, error)

	// GetByID returns the bill with the given id.
	// *Authorization required: requester in group.Members
	GetByID(requesterID uuid.UUID, id uuid.UUID) (dto.BillDetailedOutput, error)

	// Delete deletes the bill with the given id.
	// *Authorization required: requester == bill.Owner
	Delete(requesterID uuid.UUID, id uuid.UUID) error

	// GetAllByUserID returns all the bills of the given user according to the filter.
	// If no filter is provided, all bills from the groups in which the user is a member are returned.
	// *Authorization required: requester == userID
	GetAllByUserID(requesterID uuid.UUID, userID uuid.UUID, isUnseen *bool, isOwner *bool) ([]dto.BillDetailedOutput, error)
}
