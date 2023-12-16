package db_storages

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	. "split-the-bill-server/storage/database/test_util"
	"testing"
)

func TestUserStorage_Create_Success(t *testing.T) {
	sqlDB, gormDB, mock := InitMockDB(t)
	defer sqlDB.Close()

	// Create an instance of UserStorage with the mocked DB
	userStorage := UserStorage{DB: gormDB}

	currentUser := User

	// Expectations for the transaction with mocked behavior
	mock.ExpectBegin()
	mock.ExpectExec(`INSERT INTO "users"`).
		WithArgs(currentUser.ID, sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), currentUser.Email).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec(`INSERT INTO "credentials"`).
		WithArgs(currentUser.ID, PasswordHash).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err := userStorage.Create(currentUser, PasswordHash)

	// Assertions
	assert.NoError(t, err) // Check if the Create method returns no error

	// Ensure all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
}
