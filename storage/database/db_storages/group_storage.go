package db_storages

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	. "split-the-bill-server/domain/model"
	"split-the-bill-server/storage"
	. "split-the-bill-server/storage/database"
	. "split-the-bill-server/storage/database/entity"
	. "split-the-bill-server/storage/storage_inf"
)

type GroupStorage struct {
	DB *gorm.DB
}

func NewGroupStorage(DB *Database) IGroupStorage {
	return &GroupStorage{DB: DB.Context}
}

func (g *GroupStorage) AddGroup(group GroupModel) error {
	groupItem := ToGroupEntity(group)

	// try to store new group in storage
	res := g.DB.Where(Group{Base: Base{ID: groupItem.ID}}).FirstOrCreate(&groupItem)
	// RowsAffected == 0 -> group already exists
	if res.RowsAffected == 0 {
		return storage.GroupAlreadyExistsError
	}
	return res.Error
}

func (g *GroupStorage) GetGroupByID(id uuid.UUID) (GroupModel, error) {
	var group Group

	// load group with related user and members from db
	tx := g.DB.Preload("UserModel").Preload("Members").Limit(1).Find(&group, "id = ?", id)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return GroupModel{}, storage.NoSuchGroupError
		}
		return GroupModel{}, tx.Error
	}
	if tx.RowsAffected == 0 {
		return GroupModel{}, storage.NoSuchGroupError
	}
	return ToGroupModel(&group), nil
}

func (g *GroupStorage) AddMemberToGroup(memberID uuid.UUID, groupID uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}

func (g *GroupStorage) AddBillToGroup(bill *BillModel, groupID uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}