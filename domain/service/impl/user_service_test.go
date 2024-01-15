package impl

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"split-the-bill-server/domain/model"
	"split-the-bill-server/storage/mocks"
	"testing"
)

// Testdata
var (
	TestUser = model.UserModel{
		ID:    uuid.New(),
		Email: "test@mail.com",
	}
	TestUser2 = model.UserModel{
		ID:    uuid.New(),
		Email: "test2@mail.com",
	}
)

func TestUserService_GetByID(t *testing.T) {

	// mock method
	mocks.MockUserGetByID = func(id uuid.UUID) (model.UserModel, error) {
		return TestUser, nil
	}
	ret, err := userService.GetByID(TestUser.ID)
	assert.NotNilf(t, ret, "Returned data should not be nil")
	assert.Nilf(t, err, "Error should be nil")
	assert.Equalf(t, TestUser.ID, ret.ID, "Returned ID should be equal to the given ID")
	assert.EqualValuesf(t, TestUser.Email, ret.Email, "Returned email should be equal to the given email")
}

func TestUserService_GetAll(t *testing.T) {

	// mock method
	mocks.MockUserGetAll = func() ([]model.UserModel, error) {
		return []model.UserModel{TestUser, TestUser2}, nil
	}

	ret, err := userService.GetAll()
	assert.NotNilf(t, ret, "Returned data should not be nil")
	assert.Nilf(t, err, "Error should be nil")
	assert.Equalf(t, 2, len(ret), "Returned data should have 2 elements")
	assert.EqualValuesf(t, TestUser.Email, ret[0].Email, "Returned email should be equal to the given email")
	assert.EqualValuesf(t, TestUser2.Email, ret[1].Email, "Returned email should be equal to the given email")
}
