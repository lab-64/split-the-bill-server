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
	groupStorage      IGroupStorage
}

func NewInvitationService(invitationStorage *IInvitationStorage, groupStorage *IGroupStorage) IInvitationService {
	return &InvitationService{invitationStorage: *invitationStorage, groupStorage: *groupStorage}
}

func (i InvitationService) CreateGroupInvitation(request GroupInvitationDTO) error {
	// get invites from request
	invites := request.Invites

	// handle group invitations for all invitees
	for _, invitee := range invites {
		groupInvitation := model.CreateGroupInvitation(request.GroupID, invitee)
		err := i.invitationStorage.AddGroupInvitation(groupInvitation)
		if err != nil {
			return err
		}

	}

	return nil
}

func (i InvitationService) GetGroupInvitationByID(id UUID) (GroupInvitationOutputDTO, error) {
	group, err := i.invitationStorage.GetGroupInvitationByID(id)
	if err != nil {
		return GroupInvitationOutputDTO{}, err
	}
	return ToGroupInvitationDTO(group), nil
}

func (i InvitationService) GetGroupInvitationsFromUser(id UUID) ([]GroupInvitationOutputDTO, error) {
	groupInvitations, err := i.invitationStorage.GetGroupInvitationsByUserID(id)
	if err != nil {
		return nil, err
	}
	// convert group invitations to data transfer objects
	var result []GroupInvitationOutputDTO
	for _, groupInvitation := range groupInvitations {
		result = append(result, ToGroupInvitationDTO(groupInvitation))
	}
	return result, nil
}

// TODO: Handle consistency. AddMemberToGroup and DeleteGroupInvitation should be performed both or none.
func (i InvitationService) AcceptGroupInvitation(invitationID UUID, userID UUID) error {
	// TODO: is this really necessary?
	// get group from invitation
	invitation, err := i.invitationStorage.GetGroupInvitationByID(invitationID)
	if err != nil {
		return err
	}
	groupID := invitation.Group.ID
	// add user to group
	err = i.groupStorage.AddMemberToGroup(userID, groupID)
	if err != nil {
		return err
	}
	// delete invitation
	err = i.invitationStorage.DeleteGroupInvitation(invitationID)
	return err
}

func (i InvitationService) DeclineGroupInvitation(invitation UUID, userID UUID) error {
	err := i.invitationStorage.DeleteGroupInvitation(invitation)
	return err
}
