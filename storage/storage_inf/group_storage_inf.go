package storage_inf

import (
	. "github.com/google/uuid"
	. "split-the-bill-server/domain/model"
)

type IGroupStorage interface {
	// AddGroup adds the given group to the storage. If a group with the same ID or name already exists, a GroupAlreadyExistsError is returned.
	AddGroup(group GroupModel) error

	// GetGroupByID returns the group with the given ID, or a NoSuchGroupError if no such group exists.
	GetGroupByID(id UUID) (GroupModel, error)

	// AddMemberToGroup adds the given member to the group with the given ID. If the group does not exist, a NoSuchGroupError is returned.
	AddMemberToGroup(memberID UUID, groupID UUID) error

	// AddBillToGroup adds the given bill to the group with the given ID. If the group does not exist, a NoSuchGroupError is returned.
	AddBillToGroup(bill *BillModel, groupID UUID) error
}