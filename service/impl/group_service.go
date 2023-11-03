package impl

import (
	"github.com/google/uuid"
	"split-the-bill-server/common"
	"split-the-bill-server/dto"
	"split-the-bill-server/service"
	"split-the-bill-server/storage"
)

type GroupService struct {
	storage.IGroupStorage
	storage.IUserStorage
}

func NewGroupService(groupStorage *storage.IGroupStorage, userStorage *storage.IUserStorage) service.IGroupService {
	return &GroupService{IGroupStorage: *groupStorage, IUserStorage: *userStorage}
}

func (g *GroupService) Create(groupDTO dto.GroupInputDTO) (dto.GroupOutputDTO, error) {

	// TODO: get user id from authenticated user
	// TODO: delete, just for testing
	user, err := g.IUserStorage.GetByUsername("felix")
	if err != nil {
		return dto.GroupOutputDTO{}, err
	}

	// create group with the only member being the owner
	group := groupDTO.ToGroup(user.ID, []uuid.UUID{user.ID})

	// store group in db
	err = g.IGroupStorage.AddGroup(group)
	if err != nil {
		return dto.GroupOutputDTO{}, err
	}

	return dto.ToGroupDTO(group), err
}

func (g *GroupService) GetByID(id uuid.UUID) (dto.GroupOutputDTO, error) {
	group, err := g.IGroupStorage.GetGroupByID(id)
	common.LogError(err)

	return dto.ToGroupDTO(group), err
}
