package eph_storages

import (
	"github.com/google/uuid"
	"split-the-bill-server/domain/model"
	"split-the-bill-server/storage"
	"split-the-bill-server/storage/ephemeral"
)

type InvitationStorage struct {
	e *ephemeral.Ephemeral
}

func (i InvitationStorage) DeleteGroupInvitation(id uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}

func NewInvitationStorage(ephemeral *ephemeral.Ephemeral) storage.IInvitationStorage {
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

func (i InvitationStorage) AcceptGroupInvitation(id uuid.UUID, userId uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}

func (i InvitationStorage) AddGroupInvitation(invitation model.GroupInvitationModel) (model.GroupInvitationModel, error) {
	//TODO implement me
	panic("implement me")
}
