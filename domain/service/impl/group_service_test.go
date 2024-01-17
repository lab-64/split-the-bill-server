package impl

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"split-the-bill-server/domain/model"
	"split-the-bill-server/storage"
	"split-the-bill-server/storage/mocks"
	"testing"
)

// Testdata
var (
	TestGroup = model.GroupModel{
		ID:      uuid.New(),
		Name:    "Test Group",
		Owner:   TestUser,
		Members: []model.UserModel{TestUser, TestUser2},
		Bills:   []model.BillModel{TestBill},
	}

	TestGroup2 = model.GroupModel{
		ID:      uuid.New(),
		Name:    "Test Group 2",
		Owner:   TestUser2,
		Members: []model.UserModel{TestUser, TestUser2},
	}
)

func TestBalanceCalculation(t *testing.T) {

	balance := calculateBalance(TestGroup)
	assert.Equalf(t, 2, len(balance), "Balance should inclue 2 group members")
	assert.Equalf(t, 9.25, balance[TestUser.ID], "Balance for TestUser should be 9.25")
	assert.Equalf(t, -9.25, balance[TestUser2.ID], "Balance for TestUser2 should be -9.25")

}

func TestGroupService_GetByID(t *testing.T) {

	// mock method
	mocks.MockGroupGetGroupByID = func(id uuid.UUID) (model.GroupModel, error) {
		return TestGroup, nil
	}
	ret, err := groupService.GetByID(TestGroup.ID)
	assert.NotNilf(t, ret, "Returned data should not be nil")
	assert.Nilf(t, err, "Error should be nil")
	assert.Equalf(t, TestGroup.ID, ret.ID, "Returned ID should be equal to the given ID")
	assert.EqualValuesf(t, TestGroup.Name, ret.Name, "Returned name should be equal to the given name")
	assert.EqualValuesf(t, TestGroup.Owner.ID, ret.Owner.ID, "Returned owner should be equal to the given owner")
	assert.Equalf(t, 2, len(ret.Members), "Returned members should have 2 elements")
	assert.Equalf(t, 1, len(ret.Bills), "Returned bills should have 1 element")
	assert.NotNilf(t, ret.Balance, "Returned balance should not be nil")
	assert.Equalf(t, 2, len(ret.Balance), "Returned balance should have 2 elements")

	// mock method with error
	mocks.MockGroupGetGroupByID = func(id uuid.UUID) (model.GroupModel, error) {
		return model.GroupModel{}, storage.NoSuchGroupError
	}
	ret, err = groupService.GetByID(TestGroup.ID)
	assert.NotNilf(t, err, "Error should not be nil")
	assert.Errorf(t, err, storage.NoSuchGroupError.Error(), "Error should be NoSuchGroupError")
}

func TestGroupService_GetAllByUser(t *testing.T) {

	// mock method
	mocks.MockGroupGetGroupsByUserID = func(userID uuid.UUID) ([]model.GroupModel, error) {
		return []model.GroupModel{TestGroup, TestGroup2}, nil
	}
	ret, err := groupService.GetAllByUser(TestUser.ID)
	assert.NotNilf(t, ret, "Returned data should not be nil")
	assert.Nilf(t, err, "Error should be nil")
	assert.Equalf(t, 2, len(ret), "Returned data should have 2 elements")
	assert.EqualValuesf(t, TestGroup.Name, ret[0].Name, "Returned name should be equal to the given name")
	assert.EqualValuesf(t, TestGroup2.Name, ret[1].Name, "Returned name should be equal to the given name")
	assert.NotNilf(t, ret[0].Balance, "Returned balance should not be nil")
	assert.NotNilf(t, ret[1].Balance, "Returned balance should not be nil")

	// mock method with error
	mocks.MockGroupGetGroupsByUserID = func(userID uuid.UUID) ([]model.GroupModel, error) {
		return nil, storage.NoSuchUserError
	}
	ret, err = groupService.GetAllByUser(TestUser.ID)
	assert.NotNilf(t, err, "Error should not be nil")
}
