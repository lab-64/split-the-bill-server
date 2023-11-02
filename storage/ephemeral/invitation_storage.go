package ephemeral

import (
	"github.com/google/uuid"
	"split-the-bill-server/storage"
	"split-the-bill-server/types"
)

type InvitationStorage struct {
	e *Ephemeral
}

func (i InvitationStorage) AddGroupInvitation(invitation types.GroupInvitation) error {
	//TODO implement me
	panic("implement me")
}

func (i InvitationStorage) DeleteGroupInvitation(id uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}

func NewInvitationStorage(ephemeral *Ephemeral) storage.IInvitationStorage {
	return &InvitationStorage{e: ephemeral}
}
