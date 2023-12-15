package db_storages

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

func (g *GroupStorage) AddGroup(group GroupModel) (GroupModel, error) {
	groupItem := ToGroupEntity(group)

	// try to store new group in storage
	res := g.DB.
		Where(Group{Base: Base{ID: groupItem.ID}}).
		Preload("Owner").
		FirstOrCreate(&groupItem)
	// RowsAffected == 0 -> group already exists
	if res.RowsAffected == 0 {
		return GroupModel{}, storage.GroupAlreadyExistsError
	}
	return ToGroupModel(&groupItem), res.Error
}

func (g *GroupStorage) GetGroupByID(id uuid.UUID) (GroupModel, error) {
	var group Group

	// load group with related user and members from db
	tx := g.DB.
		Preload(clause.Associations).
		Preload("Bills.Items.Contributors").
		Limit(1).Find(&group, "id = ?", id)

	if tx.Error != nil {
		return GroupModel{}, tx.Error
	}
	if tx.RowsAffected == 0 {
		return GroupModel{}, storage.NoSuchGroupError
	}
	return ToGroupModel(&group), nil
}

func (g *GroupStorage) GetGroupsByUserID(userID uuid.UUID) ([]GroupModel, error) {
	var groups []Group

	tx := g.DB.
		Preload(clause.Associations).
		Preload("Bills.Items.Contributors").
		Where("id IN (SELECT group_id FROM group_members WHERE user_id = ?)", userID).
		Find(&groups)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return ToGroupModelSlice(groups), nil
}
