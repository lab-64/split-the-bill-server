package integration_tests

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"os"
	"split-the-bill-server/authentication"
	"split-the-bill-server/domain/service/impl"
	"split-the-bill-server/presentation/handler"
	"split-the-bill-server/presentation/router"
	"split-the-bill-server/storage/database"
	"split-the-bill-server/storage/database/db_storages"
	. "split-the-bill-server/storage/database/entity"
	"split-the-bill-server/storage/storage_inf"
	"testing"
)

var (
	db  *database.Database
	app *fiber.App
)

// TestMain initializes the test environment. It is called before the tests are executed.
func TestMain(m *testing.M) {
	setupTestEnv()
	os.Exit(m.Run())
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
	passwordValidator, err := authentication.NewPasswordValidator()
	if err != nil {
		log.Fatal(err)
	}

	// handlers
	userHandler := handler.NewUserHandler(&userService, passwordValidator)
	groupHandler := handler.NewGroupHandler(&groupService, &invitationService)
	billHandler := handler.NewBillHandler(&billService, &groupService)
	invitationHandler := handler.NewInvitationHandler(&invitationService)

	// authenticator
	authenticator := authentication.NewAuthenticator(&cookieStorage)

	// create webserver
	fiberApp := fiber.New()

	// setupTestEnv routing
	router.SetupRoutes(fiberApp, *userHandler, *groupHandler, *billHandler, *invitationHandler, *authenticator)

	app = fiberApp
}

// setupStorage initializes and configures the storage components for the integration tests.
func setupStorage() (storage_inf.IUserStorage, storage_inf.IGroupStorage, storage_inf.ICookieStorage, storage_inf.IBillStorage, storage_inf.IInvitationStorage) {
	return db_storages.NewUserStorage(db), db_storages.NewGroupStorage(db), db_storages.NewCookieStorage(db), db_storages.NewBillStorage(db), db_storages.NewInvitationStorage(db)

}

// initDB initializes the test database connection.
func initDB() {
	db = &database.Database{}

	sqliteDB, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	err = sqliteDB.AutoMigrate(&User{}, &AuthCookie{}, &Credentials{}, &Group{}, &GroupInvitation{}, &Bill{}, &Item{})
	if err != nil {
		panic("failed to migrate database")
	}
	db.Context = sqliteDB

}

// refreshDB deletes all entries from the database.
func refreshDB() {
	db.Context.Unscoped().Where("1 = 1").Delete(&User{})
	return
}
