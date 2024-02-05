package impl

import (
	"github.com/google/uuid"
	. "split-the-bill-server/domain/service"
	. "split-the-bill-server/presentation/dto"
	"split-the-bill-server/storage"
)

type InvitationService struct {
	invitationStorage storage.IInvitationStorage
	groupStorage      storage.IGroupStorage
}

func NewInvitationService(invitationStorage *storage.IInvitationStorage, groupStorage *storage.IGroupStorage) IInvitationService {
	return &InvitationService{invitationStorage: *invitationStorage, groupStorage: *groupStorage}
}

func (i InvitationService) CreateGroupInvitations(request GroupInvitationInputDTO) ([]GroupInvitationOutputDTO, error) {
	// get invites from request
	invites := request.Invitees
	var result []GroupInvitationOutputDTO

	// handle group invitations for all invitees
	for _, invitee := range invites {
		groupInvitation := CreateGroupInvitationModel(uuid.New(), request.GroupID, invitee)
		groupInvitation, err := i.invitationStorage.AddGroupInvitation(groupInvitation)

		if err != nil {
			return nil, err
		}

		result = append(result, ConvertToGroupInvitationDTO(groupInvitation))
	}

	return result, nil
}

func (i InvitationService) GetGroupInvitationByID(invitationID uuid.UUID) (GroupInvitationOutputDTO, error) {
	group, err := i.invitationStorage.GetGroupInvitationByID(invitationID)
	if err != nil {
		return GroupInvitationOutputDTO{}, err
	}
	return ConvertToGroupInvitationDTO(group), nil
}

func (i InvitationService) GetGroupInvitationsByUser(userID uuid.UUID) ([]GroupInvitationOutputDTO, error) {
	groupInvitations, err := i.invitationStorage.GetGroupInvitationsByUserID(userID)
	if err != nil {
		return nil, err
	}
	// convert group invitations to data transfer objects
	var result []GroupInvitationOutputDTO
	for _, groupInvitation := range groupInvitations {
		result = append(result, ConvertToGroupInvitationDTO(groupInvitation))
	}
	return result, nil
}

func (i InvitationService) HandleGroupInvitation(invitationID uuid.UUID, isAccept bool) error {
	//TODO: authorization

	if !isAccept {
		if err := i.invitationStorage.DeleteGroupInvitation(invitationID); err != nil {
			return err
		}
	}

	if err := i.invitationStorage.AcceptGroupInvitation(invitationID); err != nil {
		return err
	}

	return nil
}
