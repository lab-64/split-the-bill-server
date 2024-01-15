package impl

import (
	"os"
	"split-the-bill-server/domain/service/service_inf"
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
