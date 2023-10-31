package impl

import (
	"github.com/google/uuid"
	"split-the-bill-server/core"
	"split-the-bill-server/domain/model"
	"split-the-bill-server/domain/service/service_inf"
	"split-the-bill-server/presentation/dto"
	"split-the-bill-server/storage/storage_inf"
)

type GroupService struct {
	groupStorage storage_inf.IGroupStorage
	userStorage  storage_inf.IUserStorage
}

func NewGroupService(groupStorage *storage_inf.IGroupStorage, userStorage *storage_inf.IUserStorage) service_inf.IGroupService {
	return &GroupService{groupStorage: *groupStorage, userStorage: *userStorage}
}

func (g *GroupService) Create(groupDTO dto.GroupInputDTO) (dto.GroupOutputDTO, error) {

	// TODO: get user id from authenticated user
	// TODO: delete, just for testing
	user, err := g.userStorage.GetByUsername("felix")
	if err != nil {
		return dto.GroupOutputDTO{}, err
	}

	// create group with the only member being the owner
	group := groupDTO.ToGroup(user, []model.User{user})

	// store group in db
	err = g.groupStorage.AddGroup(group)
	if err != nil {
		return dto.GroupOutputDTO{}, err
	}

	return dto.ToGroupDTO(&group), err
}

func (g *GroupService) GetByID(id uuid.UUID) (dto.GroupOutputDTO, error) {
	group, err := g.groupStorage.GetGroupByID(id)
	core.LogError(err)

	return dto.ToGroupDTO(&group), err
}
