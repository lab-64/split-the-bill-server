package impl

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"os"
	"split-the-bill-server/domain/model"
	"split-the-bill-server/domain/service/service_inf"
	. "split-the-bill-server/storage/database/test_util"
	"split-the-bill-server/storage/mocks"
	"split-the-bill-server/storage/storage_inf"
	"testing"
)

var (
	userStorage   storage_inf.IUserStorage
	cookieStorage storage_inf.ICookieStorage
	userService   service_inf.IUserService
)

func TestMain(m *testing.M) {
	// setup
	userStorage = mocks.NewUserStorageMock()
	cookieStorage = mocks.NewCookieStorageMock()
	userService = NewUserService(&userStorage, &cookieStorage)

	// Run tests
	exitCode := m.Run()

	// Exit with the same code as the test run
	os.Exit(exitCode)
}

func TestUserService_GetByID(t *testing.T) {

	// mock method
	mocks.MockUserGetByID = func(id uuid.UUID) (model.UserModel, error) {
		return User, nil
	}
	ret, err := userService.GetByID(User.ID)
	assert.NotNilf(t, ret, "Returned data should not be nil")
	assert.Nilf(t, err, "Error should be nil")
	assert.Equalf(t, User.ID, ret.ID, "Returned ID should be equal to the given ID")
	assert.EqualValuesf(t, User.Email, ret.Email, "Returned email should be equal to the given email")
}

func TestUserService_GetAll(t *testing.T) {

	// mock method
	mocks.MockUserGetAll = func() ([]model.UserModel, error) {
		return []model.UserModel{User, User2}, nil
	}

	ret, err := userService.GetAll()
	assert.NotNilf(t, ret, "Returned data should not be nil")
	assert.Nilf(t, err, "Error should be nil")
	assert.Equalf(t, 2, len(ret), "Returned data should have 2 elements")
	assert.EqualValuesf(t, User.Email, ret[0].Email, "Returned email should be equal to the given email")
	assert.EqualValuesf(t, User2.Email, ret[1].Email, "Returned email should be equal to the given email")
}
