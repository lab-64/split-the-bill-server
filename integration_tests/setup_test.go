package integration_tests

import (
	"bytes"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"net/http/httptest"
	"os"
	"split-the-bill-server/domain/service/impl"
	"split-the-bill-server/domain/util"
	"split-the-bill-server/presentation/handler"
	"split-the-bill-server/presentation/middleware"
	"split-the-bill-server/presentation/router"
	"split-the-bill-server/storage"
	"split-the-bill-server/storage/database"
	"split-the-bill-server/storage/database/db_storages"
	. "split-the-bill-server/storage/database/entity"
	"testing"
)

var (
	db  *database.Database
	app *fiber.App
)

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Helper functions
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// login logs in a user with given credentials and returns the generated session cookie used to prove authentication.
func login(email string, password string) (string, error) {
	// login
	req := httptest.NewRequest("POST", "/api/user/login", bytes.NewBufferString(fmt.Sprintf(`{"email": "%s", "password": "%s"}`, email, password)))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req, -1)
	if resp != nil {
		for _, cookie := range resp.Cookies() {
			if cookie.Name == "session_cookie" {
				return cookie.Value, nil
			}
		}
		defer resp.Body.Close()
	}
	return "No cookie found!", err
}

func getStoredUserEntity(id uuid.UUID) (User, error) {
	var user User
	res := db.Context.Limit(1).
		Preload("Groups.Owner").
		Preload("Groups.Members").
		Preload("GroupInvitations").
		Find(&user, "id = ?", id)
	if res.Error != nil {
		return User{}, res.Error
	}
	return user, nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Test environment setup functions
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// TestMain initializes the test environment. It is called before the tests are executed.
func TestMain(m *testing.M) {
	// remove old sqlite
	os.Remove("gorm.db")

	setupTestEnv()

	// seed db
	for _, s := range All() {
		if err := s.Run(db.Context); err != nil {
			log.Fatalf("Running seed '%s', failed with error: %s", s.Name, err)
		}
	}

	// Run tests
	exitCode := m.Run()

	// Exit with the same code as the test run
	os.Exit(exitCode)
}

// setupTestEnv initializes and configures the storage components and the webserver routes for the integration tests.
func setupTestEnv() {
	initDB()

	// setupTestEnv storage
	userStorage, groupStorage, cookieStorage, billStorage, invitationStorage := setupStorage()

	// services
	userService := impl.NewUserService(&userStorage, &cookieStorage)
	groupService := impl.NewGroupService(&groupStorage, &userStorage)
	billService := impl.NewBillService(&billStorage, &groupStorage)
	invitationService := impl.NewInvitationService(&invitationStorage, &groupStorage)

	// password validator
	passwordValidator, err := util.NewPasswordValidator()
	if err != nil {
		log.Fatal(err)
	}

	// handlers
	userHandler := handler.NewUserHandler(&userService, passwordValidator)
	groupHandler := handler.NewGroupHandler(&groupService, &invitationService)
	billHandler := handler.NewBillHandler(&billService, &groupService)
	invitationHandler := handler.NewInvitationHandler(&invitationService)

	// authenticator
	authenticator := middleware.NewAuthenticator(&cookieStorage)

	// create webserver
	fiberApp := fiber.New()

	// setupTestEnv routing
	router.SetupRoutes(fiberApp, *userHandler, *groupHandler, *billHandler, *invitationHandler, *authenticator)

	app = fiberApp
}

// setupStorage initializes and configures the storage components for the integration tests.
func setupStorage() (storage.IUserStorage, storage.IGroupStorage, storage.ICookieStorage, storage.IBillStorage, storage.IInvitationStorage) {
	return db_storages.NewUserStorage(db), db_storages.NewGroupStorage(db), db_storages.NewCookieStorage(db), db_storages.NewBillStorage(db), db_storages.NewInvitationStorage(db)

}

// initDB initializes the test database connection.
func initDB() {
	db = &database.Database{}

	sqliteDB, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Error while connecting to the database: " + err.Error())
	}

	err = sqliteDB.AutoMigrate(&User{}, &AuthCookie{}, &Credentials{}, &Group{}, &GroupInvitation{}, &Bill{}, &Item{})
	if err != nil {
		log.Fatal("Error while migrating the database: " + err.Error())
	}
	db.Context = sqliteDB

}
