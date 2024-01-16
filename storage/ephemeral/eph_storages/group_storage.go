package eph_storages

import (
	"github.com/google/uuid"
	"split-the-bill-server/domain/model"
	"split-the-bill-server/storage"
	"split-the-bill-server/storage/ephemeral"
)

type GroupStorage struct {
	e *ephemeral.Ephemeral
}

func NewGroupStorage(ephemeral *ephemeral.Ephemeral) storage.IGroupStorage {
	return &GroupStorage{e: ephemeral}
}

func (g *GroupStorage) AddGroup(group model.GroupModel) (model.GroupModel, error) {
	g.e.Lock.Lock()
	defer g.e.Lock.Unlock()
	_, exists := g.e.Groups[group.ID]
	if exists {
		return model.GroupModel{}, storage.GroupAlreadyExistsError
	}
	g.e.Groups[group.ID] = &group
	return group, nil
}

func (g *GroupStorage) UpdateGroup(group model.GroupModel) (model.GroupModel, error) {
	//TODO implement me
	panic("implement me")
}

func (g *GroupStorage) GetGroupByID(id uuid.UUID) (model.GroupModel, error) {
	g.e.Lock.Lock()
	defer g.e.Lock.Unlock()
	group, exists := g.e.Groups[id]
	if !exists {
		return *group, storage.NoSuchGroupError
	}
	return *group, nil
}

func (g *GroupStorage) AddMemberToGroup(memberID uuid.UUID, groupID uuid.UUID) error {
	g.e.Lock.Lock()
	defer g.e.Lock.Unlock()
	group, exists := g.e.Groups[groupID]
	if !exists {
		return storage.NoSuchGroupError
	}
	user, exists := g.e.Users[memberID]
	if exists {
		return storage.NoSuchUserError
	}
	group.Members = append(group.Members, user)
	g.e.Groups[groupID] = group
	return nil
}

func (g *GroupStorage) GetGroupsByUserID(userID uuid.UUID) ([]model.GroupModel, error) {
	//TODO implement me
	panic("implement me")
}
