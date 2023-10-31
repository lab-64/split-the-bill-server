package db_storages

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"split-the-bill-server/domain/model"
	"split-the-bill-server/storage"
	"split-the-bill-server/storage/database"
	. "split-the-bill-server/storage/database/entity"
	"split-the-bill-server/storage/storage_inf"
)

type GroupStorage struct {
	DB *gorm.DB
}

func NewGroupStorage(DB *database.Database) storage_inf.IGroupStorage {
	return &GroupStorage{DB: DB.Context}
}

func (g *GroupStorage) AddGroup(group model.Group) error {
	groupItem := MakeGroup(group)

	// try to store new group in storage
	res := g.DB.Where(Group{Base: Base{ID: groupItem.ID}}).FirstOrCreate(&groupItem)
	// RowsAffected == 0 -> group already exists
	if res.RowsAffected == 0 {
		return storage.GroupAlreadyExistsError
	}
	return res.Error
}

func (g *GroupStorage) GetGroupByID(id uuid.UUID) (model.Group, error) {
	var group Group

	// load group with related user and members from db
	tx := g.DB.Preload("User").Preload("Members").Limit(1).Find(&group, "id = ?", id)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return model.Group{}, storage.NoSuchGroupError
		}
		return model.Group{}, tx.Error
	}
	if tx.RowsAffected == 0 {
		return model.Group{}, storage.NoSuchGroupError
	}
	return group.ToGroup(), nil
}

func (g *GroupStorage) AddMemberToGroup(memberID uuid.UUID, groupID uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}

func (g *GroupStorage) AddBillToGroup(bill *model.Bill, groupID uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}
