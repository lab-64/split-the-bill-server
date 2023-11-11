package db_storages

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"log"
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

func (i InvitationStorage) AddGroupInvitation(invitation GroupInvitationModel) error {
	// make group invitation entity
	groupInvitationItem := ToGroupInvitationEntity(invitation)

	// store group invitation in db
	res := i.DB.Where(GroupInvitation{Base: Base{ID: groupInvitationItem.ID}}).FirstOrCreate(&groupInvitationItem)
	if res.RowsAffected == 0 {
		return storage.GroupInvitationAlreadyExistsError
	}
	return res.Error
}

func (i InvitationStorage) DeleteGroupInvitation(id uuid.UUID) error {
	tx := i.DB.Delete(&GroupInvitation{}, id)
	return tx.Error
}

func (i InvitationStorage) GetGroupInvitationByID(id uuid.UUID) (GroupInvitationModel, error) {
	var groupInvitation GroupInvitation
	tx := i.DB.Preload("For.Members").Limit(1).Find(&groupInvitation, "id = ?", id)
	if tx.RowsAffected == 0 {
		return GroupInvitationModel{}, storage.NoSuchGroupInvitationError
	}
	return ToGroupInvitationModel(groupInvitation), tx.Error
}

func (i InvitationStorage) GetGroupInvitationsByUserID(userID uuid.UUID) ([]GroupInvitationModel, error) {
	var groupInvitations []GroupInvitation
	tx := i.DB.Preload("For.Members").Find(&groupInvitations, "Invitee_id = ?", userID)
	if tx.Error != nil {
		return nil, tx.Error
	}
	log.Println("Storage: GetGroupInvitationsByUserID: ", groupInvitations)
	var result []GroupInvitationModel
	for _, groupInvitation := range groupInvitations {
		result = append(result, ToGroupInvitationModel(groupInvitation))
	}
	return result, nil
}
