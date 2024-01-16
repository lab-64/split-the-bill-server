package impl

import (
	"os"
	"split-the-bill-server/domain/service"
	"split-the-bill-server/storage"
	"split-the-bill-server/storage/mocks"
	"testing"
)

var (
	userStorage   storage.IUserStorage
	cookieStorage storage.ICookieStorage
	userService   service.IUserService
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
