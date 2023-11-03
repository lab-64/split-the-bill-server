package impl

import (
	"github.com/google/uuid"
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

func (i InvitationService) CreateGroupInvitation(request dto.GroupInvitationDTO) error {
	// get invites from request
	invites := request.Invitees

	// handle group invitations for all invitees
	for _, invitee := range invites {
		groupInvitation := types.CreateGroupInvitation(request.GroupID, invitee)
		err := i.IInvitationStorage.AddGroupInvitation(groupInvitation)
		if err != nil {
			return err
		}

	}

	return nil
}

func (i InvitationService) GetGroupInvitationByID(id uuid.UUID) (dto.GroupInvitationOutputDTO, error) {
	group, err := i.IInvitationStorage.GetGroupInvitationByID(id)
	if err != nil {
		return dto.GroupInvitationOutputDTO{}, err
	}
	return dto.ToGroupInvitationDTO(group), nil
}

func (i InvitationService) AcceptGroupInvitation(invitationID uuid.UUID, userID uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}

func (i InvitationService) DeclineGroupInvitation(invitation uuid.UUID, userID uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}
