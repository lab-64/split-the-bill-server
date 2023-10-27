package ephemeral

import (
	"github.com/google/uuid"
	"split-the-bill-server/storage"
	"split-the-bill-server/types"
)

type GroupStorage struct {
	e *Ephemeral
}

func NewGroupStorage(ephemeral *Ephemeral) storage.IGroupStorage {
	return &GroupStorage{e: ephemeral}
}

func (g *GroupStorage) AddGroup(group types.Group) error {
	g.e.lock.Lock()
	defer g.e.lock.Unlock()
	_, exists := g.e.groups[group.ID]
	if exists {
		return storage.GroupAlreadyExistsError
	}
	g.e.groups[group.ID] = &group
	return nil
}

func (g *GroupStorage) GetGroupByID(id uuid.UUID) (types.Group, error) {
	g.e.lock.Lock()
	defer g.e.lock.Unlock()
	group, exists := g.e.groups[id]
	if !exists {
		return *group, storage.NoSuchGroupError
	}
	return *group, nil
}

func (g *GroupStorage) AddMemberToGroup(memberID uuid.UUID, groupID uuid.UUID) error {
	g.e.lock.Lock()
	defer g.e.lock.Unlock()
	group, exists := g.e.groups[groupID]
	if !exists {
		return storage.NoSuchGroupError
	}
	user, exists := g.e.users[memberID]
	if exists {
		return storage.NoSuchUserError
	}
	group.Members = append(group.Members, user)
	g.e.groups[groupID] = group
	return nil
}

func (g *GroupStorage) AddBillToGroup(bill *types.Bill, groupID uuid.UUID) error {
	g.e.lock.Lock()
	defer g.e.lock.Unlock()
	group, exists := g.e.groups[groupID]
	if !exists {
		return storage.NoSuchGroupError
	}

	// change group
	group.Bills = append(group.Bills, bill)
	g.e.groups[group.ID] = group
	return nil
}
