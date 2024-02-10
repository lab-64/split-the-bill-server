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
	return ConvertToGroupModel(groupItem), res.Error
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

	return ConvertToGroupModel(groupEntity), nil
}

func (g *GroupStorage) GetGroupByID(id uuid.UUID) (GroupModel, error) {
	var group Group

	// load group with related user and members from db
	tx := g.DB.
		Preload(clause.Associations).
		Preload("Bills.Items.Contributors").
		Preload("Bills.Owner").
		Limit(1).Find(&group, "id = ?", id)

	if tx.Error != nil {
		return GroupModel{}, tx.Error
	}
	if tx.RowsAffected == 0 {
		return GroupModel{}, storage.NoSuchGroupError
	}
	return ConvertToGroupModel(group), nil
}

func (g *GroupStorage) GetGroups(userID uuid.UUID, invitationID uuid.UUID) ([]GroupModel, error) {
	var groups []Group

	tx := g.DB.
		Preload(clause.Associations).
		Preload("Bills.Items.Contributors").
		Preload("Bills.Owner")

	if userID != uuid.Nil {
		tx = tx.Where("id IN (SELECT group_id FROM group_members WHERE user_id = ?)", userID)
	}

	if invitationID != uuid.Nil {
		tx = tx.Where("id IN (SELECT group_id FROM group_invitations WHERE id = ?)", invitationID)
	}

	tx.Find(&groups)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return ToGroupModelSlice(groups), nil
}

func (g *GroupStorage) AcceptGroupInvitation(invitationID uuid.UUID, userID uuid.UUID) error {
	var groupInvitation GroupInvitation

	//TODO: generalize error messages
	// TODO: test behavior
	// Check if the group invitation exists
	if err := g.DB.First(&groupInvitation, "id = ?", invitationID).Error; err != nil {
		return err
	}

	// add the user to the group members
	group := Group{Base: Base{ID: groupInvitation.GroupID}}
	user := User{Base: Base{ID: userID}}

	res := g.DB.Model(&group).Association("Members").Append(&user)

	return res
}
