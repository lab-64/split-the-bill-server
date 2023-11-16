package eph_storages

import (
	"github.com/google/uuid"
	"split-the-bill-server/domain/model"
	"split-the-bill-server/storage/ephemeral"
	"split-the-bill-server/storage/storage_inf"
)

type InvitationStorage struct {
	e *ephemeral.Ephemeral
}

func NewInvitationStorage(ephemeral *ephemeral.Ephemeral) storage_inf.IInvitationStorage {
	return &InvitationStorage{e: ephemeral}
}

func (i InvitationStorage) GetGroupInvitationByID(id uuid.UUID) (model.GroupInvitationModel, error) {
	//TODO implement me
	panic("implement me")
}

func (i InvitationStorage) GetGroupInvitationsByUserID(userID uuid.UUID) ([]model.GroupInvitationModel, error) {
	//TODO implement me
	panic("implement me")
}

func (i InvitationStorage) AddGroupInvitation(invitation model.GroupInvitationModel) error {
	//TODO implement me
	panic("implement me")
}

func (i InvitationStorage) DeleteGroupInvitation(id uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}
