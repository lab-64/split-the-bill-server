package service

import (
	"github.com/google/uuid"
	"split-the-bill-server/presentation/dto"
)

type IBillService interface {
	// Create creates a new bill and returns a detailed description of the created bill.
	// *Authorization required: requester in group.Members
	Create(requesterID uuid.UUID, bill dto.BillInput) (dto.BillDetailedOutput, error)

	// Update updates the bill with the given id with the new bill data.
	// *Authorization required: requester == group.Owner
	Update(requesterID uuid.UUID, billID uuid.UUID, billDTO dto.BillInput) (dto.BillDetailedOutput, error)

	// GetByID returns the bill with the given id.
	// *Authorization required: requester in group.Members
	GetByID(requesterID uuid.UUID, id uuid.UUID) (dto.BillDetailedOutput, error)

	// AddItem adds a new item to the bill.
	// *Authorization required: requester == bill.Owner
	AddItem(requesterID uuid.UUID, item dto.ItemInput) (dto.ItemOutput, error)

	// ChangeItem updates the item with the given id with the new item data.
	// *Authorization required: requester == bill.Owner
	ChangeItem(requesterID uuid.UUID, itemID uuid.UUID, item dto.ItemInput) (dto.ItemOutput, error)

	// TODO: add ChangeItemContribution or allow every group member to update item
	//ChangeItemContribution

	// GetItemByID returns the item with the given id.
	// *Authorization required: requester in group.Members
	GetItemByID(requesterID uuid.UUID, id uuid.UUID) (dto.ItemOutput, error)

	// DeleteItem deletes the item with the given id.
	// *Authorization required: requester == bill.Owner
	DeleteItem(requesterID uuid.UUID, id uuid.UUID) error
}
