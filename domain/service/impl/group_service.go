package impl

import (
	"github.com/google/uuid"
	. "split-the-bill-server/domain"
	"split-the-bill-server/domain/converter"
	"split-the-bill-server/domain/model"
	. "split-the-bill-server/domain/service"
	"split-the-bill-server/presentation/dto"
	"split-the-bill-server/storage"
)

type GroupService struct {
	groupStorage storage.IGroupStorage
}

func NewGroupService(groupStorage *storage.IGroupStorage) IGroupService {
	return &GroupService{groupStorage: *groupStorage}
}

func (g *GroupService) Create(groupDTO dto.GroupInput) (dto.GroupDetailedOutput, error) {

	// create group with the only member being the owner
	group := model.CreateGroup(uuid.New(), groupDTO, []uuid.UUID{groupDTO.OwnerID})

	// store group in db
	group, err := g.groupStorage.AddGroup(group)
	if err != nil {
		return dto.GroupDetailedOutput{}, err
	}

	return converter.ToGroupDetailedDTO(group), nil
}

func (g *GroupService) Update(userID uuid.UUID, groupID uuid.UUID, groupDTO dto.GroupInput) (dto.GroupDetailedOutput, error) {
	group, err := g.groupStorage.GetGroupByID(groupID)

	if err != nil {
		return dto.GroupDetailedOutput{}, err
	}

	// Authorize
	if userID != group.Owner.ID {
		return dto.GroupDetailedOutput{}, ErrNotAuthorized
	}

	// Update fields
	group.Name = groupDTO.Name
	group.Owner.ID = groupDTO.OwnerID

	group, err = g.groupStorage.UpdateGroup(group)
	if err != nil {
		return dto.GroupDetailedOutput{}, err
	}

	return converter.ToGroupDetailedDTO(group), err
}

func (g *GroupService) GetByID(id uuid.UUID) (dto.GroupDetailedOutput, error) {
	group, err := g.groupStorage.GetGroupByID(id)
	if err != nil {
		return dto.GroupDetailedOutput{}, err
	}

	balance := group.CalculateBalance()
	group.Balance = balance
	return converter.ToGroupDetailedDTO(group), nil
}

func (g *GroupService) GetAll(userID uuid.UUID, invitationID uuid.UUID) ([]dto.GroupDetailedOutput, error) {
	groups, err := g.groupStorage.GetGroups(userID, invitationID)
	if err != nil {
		return nil, err
	}

	var groupsDTO []dto.GroupDetailedOutput
	for _, group := range groups {
		balance := group.CalculateBalance()
		group.Balance = balance
		groupsDTO = append(groupsDTO, converter.ToGroupDetailedDTO(group))
	}

	return groupsDTO, nil
}

func (g *GroupService) AcceptGroupInvitation(invitationID uuid.UUID, userID uuid.UUID) error {
	err := g.groupStorage.AcceptGroupInvitation(invitationID, userID)
	return err
}
