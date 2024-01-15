package db_storages

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"testing"
)

var (
	sqlDB       *sql.DB
	dbMock      sqlmock.Sqlmock
	gormDB      *gorm.DB
	userStorage UserStorage
)

func TestMain(m *testing.M) {
	// Perform setup
	var err error
	// Initialize mock db
	err = initMockDB()
	if err != nil {
		log.Fatal(err)
	}
	// Create an instance of UserStorage with the mocked DB
	userStorage = UserStorage{DB: gormDB}

	// Run tests
	exitCode := m.Run()

	// Perform cleanup or resource teardown
	defer sqlDB.Close()

	// Exit with the same code as the test run
	os.Exit(exitCode)
}

// initMockDB initializes a mock database and returns the sql.DB, gorm.DB and sqlmock.Sqlmock instances.
func initMockDB() error {
	var err error
	sqlDB, dbMock, err = sqlmock.New()
	if err != nil {
		return err
	}
	gormDB, err = gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	return err

}

// Reminder:
// ExpectQuery for SELECT Query
// ExpectExec for INSERT, UPDATE, DELETE, ...
// ExpectRollback if DB query fails on INSERT
// ExpectCommit if DB query succeeds on INSERT
// ExpectBegin if TRANSACTION is started
