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
	//TODO implement me
	panic("implement me")
}
