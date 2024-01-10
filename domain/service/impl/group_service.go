package impl

import (
	. "github.com/google/uuid"
	. "split-the-bill-server/domain/service/service_inf"
	. "split-the-bill-server/presentation/dto"
	. "split-the-bill-server/storage/storage_inf"
)

type GroupService struct {
	groupStorage IGroupStorage
	userStorage  IUserStorage
}

func NewGroupService(groupStorage *IGroupStorage, userStorage *IUserStorage) IGroupService {
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

	// Validate
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
	return ToGroupDetailedDTO(group), nil
}

func (g *GroupService) GetAllByUser(userID UUID) ([]GroupDetailedOutputDTO, error) {
	groups, err := g.groupStorage.GetGroupsByUserID(userID)
	if err != nil {
		return nil, err
	}

	groupsDTO := make([]GroupDetailedOutputDTO, len(groups))
	for i := range groups {
		groupsDTO[i] = ToGroupDetailedDTO(groups[i])
	}

	return groupsDTO, nil
}
