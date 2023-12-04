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

func (g *GroupService) Create(groupDTO GroupInputDTO) (GroupOutputDTO, error) {

	// create group with the only member being the owner
	group := ToGroupModel(groupDTO)

	// store group in db
	err := g.groupStorage.AddGroup(group)
	if err != nil {
		return GroupOutputDTO{}, err
	}

	return ToGroupDTO(group), nil
}

func (g *GroupService) GetByID(id UUID) (GroupOutputDTO, error) {
	group, err := g.groupStorage.GetGroupByID(id)
	if err != nil {
		return GroupOutputDTO{}, err
	}
	return ToGroupDTO(group), nil
}

func (g *GroupService) GetAllByUser(userID UUID) ([]GroupOutputDTO, error) {
	groups, err := g.groupStorage.GetGroupsByUserID(userID)
	if err != nil {
		return nil, err
	}

	groupsDTO := make([]GroupOutputDTO, len(groups))
	for i := range groups {
		groupsDTO[i] = ToGroupDTO(groups[i])
	}

	return groupsDTO, nil
}
