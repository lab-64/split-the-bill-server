package mocks

import (
	"github.com/google/uuid"
	"split-the-bill-server/domain/model"
	"split-the-bill-server/storage"
)

var (
	MockGroupAddGroup          func(model.GroupModel) (model.GroupModel, error)
	MockGroupUpdateGroup       func(model.GroupModel) (model.GroupModel, error)
	MockGroupGetGroupByID      func(uuid.UUID) (model.GroupModel, error)
	MockGroupGetGroupsByUserID func(uuid.UUID) ([]model.GroupModel, error)
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

func (g GroupStorageMock) GetGroupsByUserID(userID uuid.UUID) ([]model.GroupModel, error) {
	return MockGroupGetGroupsByUserID(userID)
}
