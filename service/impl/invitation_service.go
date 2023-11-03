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
	storage.IGroupStorage
}

func NewInvitationService(invitationStorage *storage.IInvitationStorage, groupStorage *storage.IGroupStorage) service.IInvitationService {
	return &InvitationService{IInvitationStorage: *invitationStorage, IGroupStorage: *groupStorage}
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

// TODO: Handle consistency. AddMemberToGroup and DeleteGroupInvitation should be performed both or none.
func (i InvitationService) AcceptGroupInvitation(invitationID uuid.UUID, userID uuid.UUID) error {
	// TODO: is this really necessary?
	// get group from invitation
	invitation, err := i.IInvitationStorage.GetGroupInvitationByID(invitationID)
	if err != nil {
		return err
	}
	groupID := invitation.Group.ID
	// add user to group
	err = i.AddMemberToGroup(userID, groupID)
	if err != nil {
		return err
	}
	// delete invitation
	err = i.IInvitationStorage.DeleteGroupInvitation(invitationID)
	return err
}

func (i InvitationService) DeclineGroupInvitation(invitation uuid.UUID, userID uuid.UUID) error {
	err := i.IInvitationStorage.DeleteGroupInvitation(invitation)
	return err
}
