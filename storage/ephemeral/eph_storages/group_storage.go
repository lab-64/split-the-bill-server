package eph_storages

import (
	"github.com/google/uuid"
	"split-the-bill-server/domain/model"
	"split-the-bill-server/storage"
	eph "split-the-bill-server/storage/ephemeral"
)

type GroupStorage struct {
	e *eph.Ephemeral
}

func NewGroupStorage(ephemeral *eph.Ephemeral) storage.IGroupStorage {
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

func (g *GroupStorage) UpdateGroup(group model.GroupModel) (model.GroupModel, error) {
	//TODO implement me
	panic("implement me")
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

func (g *GroupStorage) GetGroups(userID uuid.UUID, invitationID uuid.UUID) ([]model.GroupModel, error) {
	//TODO implement me
	panic("implement me")
}

func (g *GroupStorage) AcceptGroupInvitation(invitationID uuid.UUID, userID uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}
