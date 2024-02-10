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
	TestGroup = model.Group{
		ID:      uuid.New(),
		Name:    "Test Group",
		Owner:   TestUser,
		Members: []model.User{TestUser, TestUser2},
		Bills:   []model.Bill{TestBill},
	}

	TestGroup2 = model.Group{
		ID:      uuid.New(),
		Name:    "Test Group 2",
		Owner:   TestUser2,
		Members: []model.User{TestUser, TestUser2},
	}
)

func TestGroupService_GetByID(t *testing.T) {

	// mock method
	mocks.MockGroupGetGroupByID = func(id uuid.UUID) (model.Group, error) {
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
	mocks.MockGroupGetGroupByID = func(id uuid.UUID) (model.Group, error) {
		return model.Group{}, storage.NoSuchGroupError
	}
	ret, err = groupService.GetByID(TestGroup.ID)
	assert.NotNilf(t, err, "Error should not be nil")
	assert.Errorf(t, err, storage.NoSuchGroupError.Error(), "Error should be NoSuchGroupError")
}

func TestGroupService_GetAll(t *testing.T) {

	// mock method
	mocks.MockGroupGetGroups = func(userID uuid.UUID) ([]model.Group, error) {
		return []model.Group{TestGroup, TestGroup2}, nil
	}
	ret, err := groupService.GetAll(TestUser.ID, uuid.Nil)
	assert.NotNilf(t, ret, "Returned data should not be nil")
	assert.Nilf(t, err, "Error should be nil")
	assert.Equalf(t, 2, len(ret), "Returned data should have 2 elements")
	assert.EqualValuesf(t, TestGroup.Name, ret[0].Name, "Returned name should be equal to the given name")
	assert.EqualValuesf(t, TestGroup2.Name, ret[1].Name, "Returned name should be equal to the given name")
	assert.NotNilf(t, ret[0].Balance, "Returned balance should not be nil")
	assert.NotNilf(t, ret[1].Balance, "Returned balance should not be nil")

	// mock method with error
	mocks.MockGroupGetGroups = func(userID uuid.UUID) ([]model.Group, error) {
		return nil, storage.NoSuchUserError
	}
	ret, err = groupService.GetAll(TestUser.ID, uuid.Nil)
	assert.NotNilf(t, err, "Error should not be nil")
	assert.Equalf(t, len(ret), 0, "Returned data should have 0 elements")
}
