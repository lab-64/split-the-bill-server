package database

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"split-the-bill-server/storage"
	. "split-the-bill-server/storage/database/entity"
	"split-the-bill-server/types"
)

type InvitationStorage struct {
	DB *gorm.DB
}

func NewInvitationStorage(DB *Database) storage.IInvitationStorage {
	return &InvitationStorage{DB: DB.context}
}

func (i InvitationStorage) AddGroupInvitation(invitation types.GroupInvitation) error {
	// make group invitation entity
	groupInvitationItem := MakeGroupInvitation(invitation)

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

func (i InvitationStorage) GetGroupInvitationByID(id uuid.UUID) (types.GroupInvitation, error) {
	var groupInvitation GroupInvitation
	tx := i.DB.Preload("For.Members").Limit(1).Find(&groupInvitation, "id = ?", id)
	if tx.RowsAffected == 0 {
		return types.GroupInvitation{}, storage.NoSuchGroupInvitationError
	}
	return groupInvitation.ToGroupInvitation(), tx.Error
}
