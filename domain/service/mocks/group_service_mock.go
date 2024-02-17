package mocks

import (
	"github.com/google/uuid"
	"split-the-bill-server/domain/service"
	"split-the-bill-server/presentation/dto"
)

var (
	MockGroupCreate                func(requesterID uuid.UUID, groupDTO dto.GroupInput) (dto.GroupDetailedOutput, error)
	MockGroupUpdate                func(requesterID uuid.UUID, groupID uuid.UUID, group dto.GroupInput) (dto.GroupDetailedOutput, error)
	MockGroupGetByID               func(requesterID uuid.UUID, id uuid.UUID) (dto.GroupDetailedOutput, error)
	MockGroupGetAll                func(requesterID uuid.UUID, userID uuid.UUID, invitationID uuid.UUID) ([]dto.GroupDetailedOutput, error)
	MockGroupAcceptGroupInvitation func(invitationID uuid.UUID, userID uuid.UUID) error
	MockGroupDelete                func(requesterID uuid.UUID, id uuid.UUID) error
)

func NewGroupServiceMock() service.IGroupService {
	return &GroupServiceMock{}
}

type GroupServiceMock struct {
}

func (g GroupServiceMock) Create(requesterID uuid.UUID, groupDTO dto.GroupInput) (dto.GroupDetailedOutput, error) {
	return MockGroupCreate(requesterID, groupDTO)
}

func (g GroupServiceMock) Update(requesterID uuid.UUID, groupID uuid.UUID, group dto.GroupInput) (dto.GroupDetailedOutput, error) {
	return MockGroupUpdate(requesterID, groupID, group)
}

func (g GroupServiceMock) GetByID(requesterID uuid.UUID, id uuid.UUID) (dto.GroupDetailedOutput, error) {
	return MockGroupGetByID(requesterID, id)
}

func (g GroupServiceMock) GetAll(requesterID uuid.UUID, userID uuid.UUID, invitationID uuid.UUID) ([]dto.GroupDetailedOutput, error) {
	return MockGroupGetAll(requesterID, userID, invitationID)
}

func (g GroupServiceMock) Delete(requesterID uuid.UUID, id uuid.UUID) error {
	return MockGroupDelete(requesterID, id)
}

func (g GroupServiceMock) AcceptGroupInvitation(invitationID uuid.UUID, userID uuid.UUID) error {
	return MockGroupAcceptGroupInvitation(invitationID, userID)
}
