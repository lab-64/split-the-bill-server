package db_storages

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	. "split-the-bill-server/domain/model"
	"split-the-bill-server/storage"
	. "split-the-bill-server/storage/database"
	. "split-the-bill-server/storage/database/entity"
)

type InvitationStorage struct {
	DB *gorm.DB
}

func NewInvitationStorage(DB *Database) storage.IInvitationStorage {
	return &InvitationStorage{DB: DB.Context}
}

func (i InvitationStorage) AddGroupInvitation(invitation GroupInvitationModel) (GroupInvitationModel, error) {
	// make group invitation entity
	groupInvitation := CreateGroupInvitationEntity(invitation)

	// store group invitation in db
	res := i.DB.Create(&groupInvitation)
	if res.RowsAffected == 0 {
		return GroupInvitationModel{}, storage.GroupInvitationAlreadyExistsError
	}

	return ConvertToGroupInvitationModel(groupInvitation), res.Error
}

func (i InvitationStorage) AcceptGroupInvitation(invitationID uuid.UUID, userID uuid.UUID) error {
	var groupInvitation GroupInvitation

	//TODO: generalize error messages
	// TODO: test behavior
	// Check if the group invitation exists
	if err := i.DB.First(&groupInvitation, "id = ?", invitationID).Error; err != nil {
		return err
	}

	// add the user to the group members
	group := Group{Base: Base{ID: groupInvitation.GroupID}}
	user := User{Base: Base{ID: userID}}

	res := i.DB.Model(&group).Association("Members").Append(&user)

	return res
}

func (i InvitationStorage) GetGroupInvitationByID(id uuid.UUID) (GroupInvitationModel, error) {
	var groupInvitation GroupInvitation
	tx := i.DB.
		Preload("Group.Owner").
		Limit(1).
		Find(&groupInvitation, "id = ?", id)
	if tx.RowsAffected == 0 {
		return GroupInvitationModel{}, storage.NoSuchGroupInvitationError
	}
	return ConvertToGroupInvitationModel(groupInvitation), tx.Error
}
