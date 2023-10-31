package db_storages

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"split-the-bill-server/domain/model"
	"split-the-bill-server/storage"
	"split-the-bill-server/storage/database"
	. "split-the-bill-server/storage/database/entity"
	"split-the-bill-server/storage/storage_inf"
)

type InvitationStorage struct {
	DB *gorm.DB
}

func NewInvitationStorage(DB *database.Database) storage_inf.IInvitationStorage {
	return &InvitationStorage{DB: DB.Context}
}

func (i InvitationStorage) AddGroupInvitation(invitation model.GroupInvitation) error {
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
	//TODO implement me
	panic("implement me")
}
