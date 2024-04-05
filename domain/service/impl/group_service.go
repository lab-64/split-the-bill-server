package impl

import (
	"github.com/google/uuid"
	. "split-the-bill-server/domain"
	"split-the-bill-server/domain/converter"
	"split-the-bill-server/domain/model"
	. "split-the-bill-server/domain/service"
	"split-the-bill-server/domain/util"
	"split-the-bill-server/presentation/dto"
	"split-the-bill-server/storage"
)

type GroupService struct {
	groupStorage storage.IGroupStorage
}

func NewGroupService(groupStorage *storage.IGroupStorage) IGroupService {
	return &GroupService{groupStorage: *groupStorage}
}

func (g *GroupService) Create(requesterID uuid.UUID, groupDTO dto.GroupInput) (dto.GroupDetailedOutput, error) {
	// Authorization
	if requesterID != groupDTO.OwnerID {
		return dto.GroupDetailedOutput{}, ErrNotAuthorized
	}

	// create group with the only member being the owner
	group := model.CreateGroup(uuid.New(), groupDTO, []uuid.UUID{groupDTO.OwnerID})

	// store group in db
	group, err := g.groupStorage.AddGroup(group)
	if err != nil {
		return dto.GroupDetailedOutput{}, err
	}

	return converter.ToGroupDetailedDTO(group), nil
}

func (g *GroupService) Update(requesterID uuid.UUID, groupID uuid.UUID, groupDTO dto.GroupInput) (dto.GroupDetailedOutput, error) {
	group, err := g.groupStorage.GetGroupByID(groupID)

	if err != nil {
		return dto.GroupDetailedOutput{}, err
	}

	// Authorization
	if requesterID != group.Owner.ID {
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

func (g *GroupService) GetByID(requesterID uuid.UUID, id uuid.UUID) (dto.GroupDetailedOutput, error) {
	group, err := g.groupStorage.GetGroupByID(id)
	if err != nil {
		return dto.GroupDetailedOutput{}, err
	}
	// Authorization
	if !group.IsMember(requesterID) {
		return dto.GroupDetailedOutput{}, ErrNotAuthorized
	}

	balance := group.CalculateBalance()
	group.Balance = balance
	return converter.ToGroupDetailedDTO(group), nil
}

func (g *GroupService) GetAll(requesterID uuid.UUID, userID uuid.UUID, invitationID uuid.UUID) ([]dto.GroupDetailedOutput, error) {
	// Authorization
	if userID != uuid.Nil && requesterID != userID {
		return nil, ErrNotAuthorized
	}
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

func (g *GroupService) Delete(requesterID uuid.UUID, id uuid.UUID) (dto.GroupDeletionOutput, error) {
	// get group
	group, err := g.groupStorage.GetGroupByID(id)
	if err != nil {
		return dto.GroupDeletionOutput{}, err
	}
	// Authorization
	if requesterID != group.Owner.ID {
		return dto.GroupDeletionOutput{}, ErrNotAuthorized
	}
	// calculate transactions to clear balance
	transactions := util.ProduceTransactionsFromBalance(group.CalculateBalance())
	// delete group
	return dto.GroupDeletionOutput{Transactions: transactions}, g.groupStorage.DeleteGroup(id)
}

func (g *GroupService) AcceptGroupInvitation(invitationID uuid.UUID, userID uuid.UUID) error {
	err := g.groupStorage.AcceptGroupInvitation(invitationID, userID)
	return err
}
