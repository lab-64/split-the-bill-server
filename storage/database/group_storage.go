package database

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"split-the-bill-server/storage"
	. "split-the-bill-server/storage/database/entity"
	"split-the-bill-server/types"
)

type GroupStorage struct {
	DB *gorm.DB
}

func NewGroupStorage(DB *Database) storage.IGroupStorage {
	return &GroupStorage{DB: DB.context}
}

func (g *GroupStorage) AddGroup(group types.Group) error {
	groupItem := MakeGroup(group)

	// check if group already exists
	_, err := g.GetGroupByID(groupItem.ID)
	if err == nil {
		return storage.GroupAlreadyExistsError
	}
	// write new group in storage
	res := g.DB.Where(Group{Base: Base{ID: groupItem.ID}}).FirstOrCreate(&groupItem)
	return res.Error
}

func (g *GroupStorage) GetGroupByID(id uuid.UUID) (types.Group, error) {
	var group Group

	// load group with related user and members from db
	tx := g.DB.Preload("User").Preload("Members").Limit(1).Find(&group, "id = ?", id)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return types.Group{}, storage.NoSuchGroupError
		}
		return types.Group{}, tx.Error
	}
	if tx.RowsAffected == 0 {
		return types.Group{}, storage.NoSuchGroupError
	}
	return group.ToGroup(), nil
}

func (g *GroupStorage) AddMemberToGroup(memberID uuid.UUID, groupID uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}

func (g *GroupStorage) AddBillToGroup(bill *types.Bill, groupID uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}
