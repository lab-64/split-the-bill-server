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

type InvitationStorage struct {
	DB *gorm.DB
}

func NewInvitationStorage(DB *Database) IInvitationStorage {
	return &InvitationStorage{DB: DB.Context}
}

func (i InvitationStorage) AddGroupInvitation(invitation GroupInvitationModel) (GroupInvitationModel, error) {
	// make group invitation entity
	groupInvitation := ToGroupInvitationEntity(invitation)

	// store group invitation in db
	res := i.DB.Where(&groupInvitation).Preload(clause.Associations).FirstOrCreate(&groupInvitation)
	if res.RowsAffected == 0 {
		return GroupInvitationModel{}, storage.GroupInvitationAlreadyExistsError
	}

	return ToGroupInvitationModel(groupInvitation), res.Error
}

func (i InvitationStorage) DeleteGroupInvitation(id uuid.UUID) error {
	tx := i.DB.Delete(&GroupInvitation{}, id)
	return tx.Error
}

func (i InvitationStorage) AcceptGroupInvitation(id uuid.UUID) error {
	var groupInvitation GroupInvitation

	//TODO: generalize error messages

	// Check if the group invitation exists
	if err := i.DB.First(&groupInvitation, "id = ?", id).Error; err != nil {
		return err
	}

	// Update group members within a transaction
	err := i.DB.Transaction(func(tx *gorm.DB) error {
		// Append the invitee to the group members
		group := Group{Base: Base{ID: groupInvitation.GroupID}}
		user := User{Base: Base{ID: groupInvitation.InviteeID}}

		res := tx.Model(&group).Association("Members").Append(&user)

		if res != nil {
			return res
		}

		// Delete invitation
		res = tx.Delete(&groupInvitation).Error
		if res != nil {
			return res
		}

		return nil
	})

	return err
}

func (i InvitationStorage) GetGroupInvitationByID(id uuid.UUID) (GroupInvitationModel, error) {
	var groupInvitation GroupInvitation
	tx := i.DB.
		Preload("For.Owner").
		Preload("For.Members").
		Limit(1).
		Find(&groupInvitation, "id = ?", id)
	if tx.RowsAffected == 0 {
		return GroupInvitationModel{}, storage.NoSuchGroupInvitationError
	}
	return ToGroupInvitationModel(groupInvitation), tx.Error
}

func (i InvitationStorage) GetGroupInvitationsByUserID(userID uuid.UUID) ([]GroupInvitationModel, error) {
	var groupInvitations []GroupInvitation
	tx := i.DB.
		Preload("For.Owner").
		Preload("For.Members").
		Find(&groupInvitations, "Invitee_id = ?", userID)
	if tx.Error != nil {
		return nil, tx.Error
	}
	var result []GroupInvitationModel
	for _, groupInvitation := range groupInvitations {
		result = append(result, ToGroupInvitationModel(groupInvitation))
	}
	return result, nil
}
