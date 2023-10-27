package impl

import (
	"github.com/google/uuid"
	"split-the-bill-server/common"
	"split-the-bill-server/dto"
	"split-the-bill-server/service"
	"split-the-bill-server/storage"
	"split-the-bill-server/types"
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
	group := groupDTO.ToGroup(user, []types.User{user})

	// Create a group invitation for each invited user
	for _, member := range groupDTO.Invites {
		groupInvitation := types.CreateGroupInvitation(&group)
		// store group invitation for user
		err = g.IUserStorage.AddGroupInvitationToUser(groupInvitation, member)
		common.LogError(err)
	}

	err = g.IGroupStorage.AddGroup(group)

	return dto.ToGroupDTO(&group), err
}

func (g *GroupService) GetByID(id uuid.UUID) (dto.GroupOutputDTO, error) {
	group, err := g.IGroupStorage.GetGroupByID(id)
	common.LogError(err)

	return dto.ToGroupDTO(&group), err
}
