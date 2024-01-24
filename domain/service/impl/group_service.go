package impl

import (
	. "github.com/google/uuid"
	. "split-the-bill-server/domain"
	. "split-the-bill-server/domain/service"
	. "split-the-bill-server/presentation/dto"
	"split-the-bill-server/storage"
)

type GroupService struct {
	groupStorage storage.IGroupStorage
	userStorage  storage.IUserStorage
}

func NewGroupService(groupStorage *storage.IGroupStorage, userStorage *storage.IUserStorage) IGroupService {
	return &GroupService{groupStorage: *groupStorage, userStorage: *userStorage}
}

func (g *GroupService) Create(groupDTO GroupInputDTO) (GroupDetailedOutputDTO, error) {

	// create group with the only member being the owner
	group := ToGroupModel(groupDTO)

	// store group in db
	group, err := g.groupStorage.AddGroup(group)
	if err != nil {
		return GroupDetailedOutputDTO{}, err
	}

	return ToGroupDetailedDTO(group), nil
}

func (g *GroupService) Update(userID UUID, groupID UUID, groupDTO GroupInputDTO) (GroupDetailedOutputDTO, error) {
	group, err := g.groupStorage.GetGroupByID(groupID)

	if err != nil {
		return GroupDetailedOutputDTO{}, err
	}

	// Authorize
	if userID != group.Owner.ID {
		return GroupDetailedOutputDTO{}, ErrNotAuthorized
	}

	// Update fields
	group.Name = groupDTO.Name
	group.Owner.ID = groupDTO.OwnerID

	group, err = g.groupStorage.UpdateGroup(group)
	if err != nil {
		return GroupDetailedOutputDTO{}, err
	}

	return ToGroupDetailedDTO(group), err
}

func (g *GroupService) GetByID(id UUID) (GroupDetailedOutputDTO, error) {
	group, err := g.groupStorage.GetGroupByID(id)
	if err != nil {
		return GroupDetailedOutputDTO{}, err
	}

	balance := group.CalculateBalance()
	group.Balance = balance
	return ToGroupDetailedDTO(group), nil
}

func (g *GroupService) GetAllByUser(userID UUID) ([]GroupDetailedOutputDTO, error) {
	groups, err := g.groupStorage.GetGroupsByUserID(userID)
	if err != nil {
		return nil, err
	}

	var groupsDTO []GroupDetailedOutputDTO
	for _, group := range groups {
		balance := group.CalculateBalance()
		group.Balance = balance
		groupsDTO = append(groupsDTO, ToGroupDetailedDTO(group))
	}

	return groupsDTO, nil
}
