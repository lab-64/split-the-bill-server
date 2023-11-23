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

	// TODO: get user id from authenticated user
	// TODO: delete, just for testing
	user, err := g.userStorage.GetByUsername("felix")
	if err != nil {
		return GroupOutputDTO{}, err
	}
	groupDTO.Owner = user.ID

	// create group with the only member being the owner
	group := ToGroupModel(groupDTO, []UUID{user.ID})

	// store group in db
	err = g.groupStorage.AddGroup(group)
	if err != nil {
		return GroupOutputDTO{}, err
	}

	return ToGroupDTO(group), err
}

func (g *GroupService) GetByID(id UUID) (GroupOutputDTO, error) {
	group, err := g.groupStorage.GetGroupByID(id)
	if err != nil {
		return GroupOutputDTO{}, err
	}
	return ToGroupDTO(group), err
}
