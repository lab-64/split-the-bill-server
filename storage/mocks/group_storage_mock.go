package mocks

import (
	"github.com/google/uuid"
	"split-the-bill-server/domain/model"
	"split-the-bill-server/storage"
)

var (
	MockGroupAddGroup              func(model.Group) (model.Group, error)
	MockGroupUpdateGroup           func(model.Group) (model.Group, error)
	MockGroupGetGroupByID          func(uuid.UUID) (model.Group, error)
	MockGroupGetGroups             func(uuid.UUID, uuid.UUID) ([]model.Group, error)
	MockGroupDeleteGroup           func(uuid.UUID) error
	MockGroupAcceptGroupInvitation func(uuid.UUID, uuid.UUID) error
	MockGroupCreateTransaction     func(transaction model.GroupTransaction) (model.GroupTransaction, error)
	MockGroupGetTransactions       func(uuid.UUID) ([]model.GroupTransaction, error)
)

func NewGroupStorageMock() storage.IGroupStorage {
	return &GroupStorageMock{}
}

type GroupStorageMock struct {
}

func (g GroupStorageMock) AddGroup(group model.Group) (model.Group, error) {
	return MockGroupAddGroup(group)
}

func (g GroupStorageMock) UpdateGroup(group model.Group) (model.Group, error) {
	return MockGroupUpdateGroup(group)
}

func (g GroupStorageMock) GetGroupByID(id uuid.UUID) (model.Group, error) {
	return MockGroupGetGroupByID(id)
}

func (g GroupStorageMock) GetGroups(userID uuid.UUID, invitationID uuid.UUID) ([]model.Group, error) {
	return MockGroupGetGroups(userID, invitationID)
}

func (g GroupStorageMock) DeleteGroup(id uuid.UUID) error {
	return MockGroupDeleteGroup(id)
}

func (g GroupStorageMock) AcceptGroupInvitation(invitationID uuid.UUID, userID uuid.UUID) error {
	return MockGroupAcceptGroupInvitation(invitationID, userID)
}

func (g GroupStorageMock) CreateGroupTransaction(transaction model.GroupTransaction) (model.GroupTransaction, error) {
	return MockGroupCreateTransaction(transaction)
}

func (g GroupStorageMock) GetAllGroupTransactions(userID uuid.UUID) ([]model.GroupTransaction, error) {
	return MockGroupGetTransactions(userID)
}
