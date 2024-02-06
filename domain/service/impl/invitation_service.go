package impl

import (
	"github.com/google/uuid"
	. "split-the-bill-server/domain/service"
	"split-the-bill-server/presentation/dto"
	"split-the-bill-server/storage"
)

type InvitationService struct {
	invitationStorage storage.IInvitationStorage
	groupStorage      storage.IGroupStorage
}

func NewInvitationService(invitationStorage *storage.IInvitationStorage, groupStorage *storage.IGroupStorage) IInvitationService {
	return &InvitationService{invitationStorage: *invitationStorage, groupStorage: *groupStorage}
}

func (i InvitationService) CreateGroupInvitations(groupID uuid.UUID) (dto.GroupInvitationOutputDTO, error) {
	invitation := dto.CreateGroupInvitationModel(uuid.New(), groupID)
	// store invitation
	invitation, err := i.invitationStorage.AddGroupInvitation(invitation)
	if err != nil {
		return dto.GroupInvitationOutputDTO{}, err
	}
	return dto.ConvertToGroupInvitationDTO(invitation), nil
}

func (i InvitationService) GetGroupInvitationByID(invitationID uuid.UUID) (dto.GroupInvitationOutputDTO, error) {
	group, err := i.invitationStorage.GetGroupInvitationByID(invitationID)
	if err != nil {
		return dto.GroupInvitationOutputDTO{}, err
	}
	return dto.ConvertToGroupInvitationDTO(group), nil
}

func (i InvitationService) AcceptGroupInvitation(invitationID uuid.UUID, userID uuid.UUID) error {
	err := i.invitationStorage.AcceptGroupInvitation(invitationID, userID)
	return err
}
