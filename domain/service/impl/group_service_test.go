package impl

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"split-the-bill-server/domain"
	"split-the-bill-server/domain/model"
	"split-the-bill-server/storage"
	"split-the-bill-server/storage/mocks"
	"testing"
)

// Testdata
var (
	TestGroup = model.Group{
		ID:           uuid.New(),
		Name:         "Test Group",
		Owner:        TestUser,
		Members:      []model.User{TestUser, TestUser2},
		Bills:        []model.Bill{TestBill},
		InvitationID: uuid.New(),
	}

	TestGroup2 = model.Group{
		ID:           uuid.New(),
		Name:         "Test Group 2",
		Owner:        TestUser2,
		Members:      []model.User{TestUser, TestUser2},
		InvitationID: uuid.New(),
	}
)

func TestGroupService_GetByID(t *testing.T) {

	tests := []struct {
		name             string
		mock             func()
		requesterID      uuid.UUID
		groupID          uuid.UUID
		expectedGroup    model.Group
		expectedBillSize int
		expectedError    error
	}{
		{
			name: "Success",
			mock: func() {
				mocks.MockGroupGetGroupByID = func(id uuid.UUID) (model.Group, error) {
					return TestGroup, nil
				}
			},
			requesterID:      TestUser.ID,
			groupID:          TestGroup.ID,
			expectedGroup:    TestGroup,
			expectedBillSize: 2,
			expectedError:    nil,
		},
		{
			name: "Not authorized",
			mock: func() {
				mocks.MockGroupGetGroupByID = func(id uuid.UUID) (model.Group, error) {
					return TestGroup, nil
				}
			},
			requesterID:   uuid.New(),
			groupID:       TestGroup.ID,
			expectedError: domain.ErrNotAuthorized,
		},
		{
			name: "No such group",
			mock: func() {
				mocks.MockGroupGetGroupByID = func(id uuid.UUID) (model.Group, error) {
					return model.Group{}, storage.NoSuchGroupError
				}
			},
			requesterID:   TestUser.ID,
			groupID:       uuid.New(),
			expectedError: storage.NoSuchGroupError,
		},
	}

	for _, testcase := range tests {
		t.Run(testcase.name, func(t *testing.T) {
			testcase.mock()
			ret, err := groupService.GetByID(testcase.requesterID, testcase.groupID)
			assert.Equalf(t, testcase.expectedError, err, "Returned error should be equal to the expected error")
			if err == nil {
				assert.Equalf(t, testcase.expectedGroup.ID, ret.ID, "Returned ID should be equal to the given ID")
				assert.Equalf(t, testcase.expectedGroup.Name, ret.Name, "Returned name should be equal to the given name")
				assert.Equalf(t, testcase.expectedGroup.Owner.ID, ret.Owner.ID, "Returned owner ID should be equal to the given owner ID")
				assert.Equalf(t, len(testcase.expectedGroup.Members), len(ret.Members), "Returned members should have the same length as the expected members")
				assert.Equalf(t, len(testcase.expectedGroup.Bills), len(ret.Bills), "Returned bills should have the same length as the expected bills")
				assert.NotNilf(t, ret.Balance, "Returned balance should not be nil")
				assert.Equalf(t, testcase.expectedBillSize, len(ret.Balance), "Returned balance should have the same length as the expected balance")
			}
		})
	}
}

func TestGroupService_GetAll(t *testing.T) {

	tests := []struct {
		name           string
		mock           func()
		requesterID    uuid.UUID
		userID         uuid.UUID
		invitationID   uuid.UUID
		expectedError  error
		expectedGroups []model.Group
	}{
		{
			name: "Success get from userID",
			mock: func() {
				mocks.MockGroupGetGroups = func(userID uuid.UUID, invitationID uuid.UUID) ([]model.Group, error) {
					return []model.Group{TestGroup, TestGroup2}, nil
				}
			},
			requesterID:    TestUser.ID,
			userID:         TestUser.ID,
			invitationID:   uuid.Nil,
			expectedError:  nil,
			expectedGroups: []model.Group{TestGroup, TestGroup2},
		},
		{
			name: "Success get from invitationID",
			mock: func() {
				mocks.MockGroupGetGroups = func(userID uuid.UUID, invitationID uuid.UUID) ([]model.Group, error) {
					return []model.Group{TestGroup}, nil
				}
			},
			requesterID:    TestUser.ID,
			userID:         uuid.Nil,
			invitationID:   TestGroup.InvitationID,
			expectedError:  nil,
			expectedGroups: []model.Group{TestGroup},
		},
		{
			name: "Not authorized",
			mock: func() {
				mocks.MockGroupGetGroups = func(userID uuid.UUID, invitationID uuid.UUID) ([]model.Group, error) {
					return []model.Group{TestGroup, TestGroup2}, nil
				}
			},
			requesterID:   TestUser.ID,
			userID:        TestUser2.ID,
			invitationID:  uuid.Nil,
			expectedError: domain.ErrNotAuthorized,
		},
	}

	for _, testcase := range tests {
		t.Run(testcase.name, func(t *testing.T) {
			testcase.mock()
			ret, err := groupService.GetAll(testcase.requesterID, testcase.userID, testcase.invitationID)
			assert.Equalf(t, testcase.expectedError, err, "Returned error should be equal to the expected error")
			if err == nil {
				assert.Equalf(t, len(testcase.expectedGroups), len(ret), "Returned data should have the same length as the expected data")
				for i, group := range ret {
					assert.Equalf(t, testcase.expectedGroups[i].ID, group.ID, "Returned ID should be equal to the given ID")
					assert.Equalf(t, testcase.expectedGroups[i].Name, group.Name, "Returned name should be equal to the given name")
					assert.NotNilf(t, group.Balance, "Returned balance should not be nil")
				}
			}
		})
	}
}
