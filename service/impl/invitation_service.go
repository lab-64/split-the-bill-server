package impl

import (
	"github.com/google/uuid"
	"split-the-bill-server/common"
	"split-the-bill-server/dto"
	"split-the-bill-server/service"
	"split-the-bill-server/storage"
	"split-the-bill-server/types"
)

type InvitationService struct {
	storage.IInvitationStorage
	storage.IUserStorage
}

func NewInvitationService(invitationStorage *storage.IInvitationStorage, userStorage *storage.IUserStorage) service.IInvitationService {
	return &InvitationService{IInvitationStorage: *invitationStorage, IUserStorage: *userStorage}
}

func (i InvitationService) CreateGroupInvitation(request dto.GroupInputDTO, groupID uuid.UUID) error {
	// get invites from request
	invites := request.Invites
	// TODO: change, wrong implementation, look up how to store association in gorm
	// create group invitation
	groupInvitation := types.CreateGroupInvitation(groupID)
	// store group invitation
	err := i.IInvitationStorage.AddGroupInvitation(groupInvitation)
	if err != nil {
		return err
	}

	// add group invitation to all users
	for _, userID := range invites {
		err = i.IUserStorage.AddGroupInvitation(groupInvitation, userID)
		// TODO: error handling, should return to which users the invitation could not be added
		common.LogError(err)
	}

	return err
}

func (i InvitationService) AcceptGroupInvitation(invitation uuid.UUID, userID uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}

func (i InvitationService) DeclineGroupInvitation(invitation uuid.UUID, userID uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}
