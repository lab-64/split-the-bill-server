package eph_storages

import (
	"github.com/google/uuid"
	"split-the-bill-server/domain/model"
	"split-the-bill-server/storage"
	eph "split-the-bill-server/storage/ephemeral"
	"split-the-bill-server/storage/storage_inf"
)

type GroupStorage struct {
	e *eph.Ephemeral
}

func NewGroupStorage(ephemeral *eph.Ephemeral) storage_inf.IGroupStorage {
	return &GroupStorage{e: ephemeral}
}

func (g *GroupStorage) AddGroup(group model.GroupModel) (model.GroupModel, error) {
	r := g.e.Locker.Lock(eph.RGroups)
	defer g.e.Locker.Unlock(r)
	_, exists := g.e.Groups[group.ID]
	if exists {
		return model.GroupModel{}, storage.GroupAlreadyExistsError
	}
	g.e.Groups[group.ID] = &group
	return group, nil
}

func (g *GroupStorage) GetGroupByID(id uuid.UUID) (model.GroupModel, error) {
	r := g.e.Locker.Lock(eph.RGroups)
	defer g.e.Locker.Unlock(r)
	group, exists := g.e.Groups[id]
	if !exists {
		return *group, storage.NoSuchGroupError
	}
	return *group, nil
}

func (g *GroupStorage) AddMemberToGroup(memberID uuid.UUID, groupID uuid.UUID) error {
	r := g.e.Locker.Lock(eph.RUsers, eph.RGroups)
	defer g.e.Locker.Unlock(r)
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
