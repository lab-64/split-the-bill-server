package mocks

import (
	"github.com/google/uuid"
	"split-the-bill-server/domain/model"
	"split-the-bill-server/storage"
)

var (
	MockGroupAddGroup     func(model.GroupModel) (model.GroupModel, error)
	MockGroupUpdateGroup  func(model.GroupModel) (model.GroupModel, error)
	MockGroupGetGroupByID func(uuid.UUID) (model.GroupModel, error)
	MockGroupGetGroups    func(uuid.UUID) ([]model.GroupModel, error)
)

func NewGroupStorageMock() storage.IGroupStorage {
	return &GroupStorageMock{}
}

type GroupStorageMock struct {
}

func (g GroupStorageMock) AddGroup(group model.GroupModel) (model.GroupModel, error) {
	return MockGroupAddGroup(group)
}

func (g GroupStorageMock) UpdateGroup(group model.GroupModel) (model.GroupModel, error) {
	return MockGroupUpdateGroup(group)
}

func (g GroupStorageMock) GetGroupByID(id uuid.UUID) (model.GroupModel, error) {
	return MockGroupGetGroupByID(id)
}

func (g GroupStorageMock) GetGroups(userID uuid.UUID, invitationID uuid.UUID) ([]model.GroupModel, error) {
	return MockGroupGetGroups(userID)
}

func (g GroupStorageMock) AcceptGroupInvitation(invitationID uuid.UUID, userID uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}
