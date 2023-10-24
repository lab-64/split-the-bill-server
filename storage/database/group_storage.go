package database

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"split-the-bill-server/storage"
	"split-the-bill-server/types"
)

type GroupStorage struct {
	DB *gorm.DB
}

func NewGroupStorage(DB *Database) storage.IGroupStorage {
	return &GroupStorage{DB: DB.context}
}

func (g *GroupStorage) AddGroup(group types.Group) error {
	//TODO implement me
	panic("implement me")
}

func (g *GroupStorage) GetGroupByID(id uuid.UUID) (types.Group, error) {
	//TODO implement me
	panic("implement me")
}

func (g *GroupStorage) AddMemberToGroup(memberID uuid.UUID, groupID uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}

func (g *GroupStorage) AddBillToGroup(bill *types.Bill, groupID uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}
