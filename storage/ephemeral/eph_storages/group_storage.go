package eph_storages

import (
	"github.com/google/uuid"
	"split-the-bill-server/domain/model"
	"split-the-bill-server/storage"
	"split-the-bill-server/storage/ephemeral"
	"split-the-bill-server/storage/storage_inf"
)

type GroupStorage struct {
	e *ephemeral.Ephemeral
}

func NewGroupStorage(ephemeral *ephemeral.Ephemeral) storage_inf.IGroupStorage {
	return &GroupStorage{e: ephemeral}
}

func (g *GroupStorage) AddGroup(group model.GroupModel) error {
	g.e.Lock.Lock()
	defer g.e.Lock.Unlock()
	_, exists := g.e.Groups[group.ID]
	if exists {
		return storage.GroupAlreadyExistsError
	}
	g.e.Groups[group.ID] = &group
	return nil
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

func (g *GroupStorage) AddBillToGroup(bill *model.BillModel, groupID uuid.UUID) error {
	g.e.Lock.Lock()
	defer g.e.Lock.Unlock()
	group, exists := g.e.Groups[groupID]
	if !exists {
		return storage.NoSuchGroupError
	}

	// change group
	group.Bills = append(group.Bills, *bill)
	g.e.Groups[group.ID] = group
	return nil
}
