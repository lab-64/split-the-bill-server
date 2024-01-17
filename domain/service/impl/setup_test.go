package impl

import (
	"os"
	"split-the-bill-server/domain/service"
	"split-the-bill-server/storage/mocks"
	"testing"
)

var (
	userService  service.IUserService
	groupService service.IGroupService
)

func TestMain(m *testing.M) {
	// setup
	userStorage := mocks.NewUserStorageMock()
	cookieStorage := mocks.NewCookieStorageMock()
	userService = NewUserService(&userStorage, &cookieStorage)
	groupStorage := mocks.NewGroupStorageMock()
	groupService = NewGroupService(&groupStorage, &userStorage)

	// Run tests
	exitCode := m.Run()

	// Exit with the same code as the test run
	os.Exit(exitCode)
}
