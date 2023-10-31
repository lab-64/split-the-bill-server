package impl

import (
	. "github.com/google/uuid"
	"split-the-bill-server/domain/model"
	. "split-the-bill-server/domain/service/service_inf"
	. "split-the-bill-server/presentation/dto"
	. "split-the-bill-server/storage/storage_inf"
)

type InvitationService struct {
	invitationStorage IInvitationStorage
	userStorage       IUserStorage
}

func NewInvitationService(invitationStorage *IInvitationStorage, userStorage *IUserStorage) IInvitationService {
	return &InvitationService{invitationStorage: *invitationStorage, userStorage: *userStorage}
}

func (i InvitationService) CreateGroupInvitation(request GroupInputDTO, groupID UUID) error {
	// get invites from request
	invites := request.Invites
	// TODO: change, wrong implementation, look up how to store association in gorm

	// handle group invitations
	for _, invitee := range invites {
		groupInvitation := model.CreateGroupInvitation(groupID, invitee)
		err := i.invitationStorage.AddGroupInvitation(groupInvitation)
		if err != nil {
			return err
		}

	}

	/*
		// add group invitation to all users
		for _, userID := range invites {
			err = i.IUserStorage.AddGroupInvitation(groupInvitation, userID)
			// TODO: error handling, should return to which users the invitation could not be added
			core.LogError(err)
		}

	*/

	return nil
}

func (i InvitationService) AcceptGroupInvitation(invitation UUID, userID UUID) error {
	//TODO implement me
	panic("implement me")
}

func (i InvitationService) DeclineGroupInvitation(invitation UUID, userID UUID) error {
	//TODO implement me
	panic("implement me")
}
