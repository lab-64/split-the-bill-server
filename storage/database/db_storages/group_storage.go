package db_storages

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	. "split-the-bill-server/domain/model"
	"split-the-bill-server/storage"
	. "split-the-bill-server/storage/database"
	. "split-the-bill-server/storage/database/entity"
)

type GroupStorage struct {
	DB *gorm.DB
}

func NewGroupStorage(DB *Database) storage.IGroupStorage {
	return &GroupStorage{DB: DB.Context}
}

func (g *GroupStorage) AddGroup(group GroupModel) (GroupModel, error) {
	groupItem := CreateGroupEntity(group)

	// .First(...) in the end enables preload on create (kind of workaround)
	// https://github.com/go-gorm/gen/issues/618
	res := g.DB.
		Preload(clause.Associations).
		Create(&groupItem).
		First(&groupItem)

	// RowsAffected == 0 -> group already exists
	if res.RowsAffected == 0 {
		return GroupModel{}, storage.GroupAlreadyExistsError
	}
	return ConvertToGroupModel(groupItem, true), res.Error
}

func (g *GroupStorage) UpdateGroup(group GroupModel) (GroupModel, error) {
	groupEntity := CreateGroupEntity(group)

	res := g.DB.
		Preload(clause.Associations).
		Model(&groupEntity).
		Updates(&groupEntity).
		First(&groupEntity)

	// TODO: add finer error handling
	if res.Error != nil {
		return GroupModel{}, res.Error
	}

	if res.RowsAffected == 0 {
		return GroupModel{}, storage.NoSuchGroupError
	}

	return ConvertToGroupModel(groupEntity, true), nil
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
	return ConvertToGroupModel(group, true), nil
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
