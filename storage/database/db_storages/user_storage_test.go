package db_storages

import (
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	. "split-the-bill-server/storage"
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
	assert.NoError(t, err) // Check if the Create method returns no error

	// Ensure all expectations were met
	if err = mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
}

func TestUserStorage_Create_Bad_Inputs(t *testing.T) {
	// Initialize mock DB and UserStorage
	sqlDB, gormDB, mock := InitMockDB(t)
	defer sqlDB.Close()
	userStorage := UserStorage{DB: gormDB}

	// Test case: Create user with empty email
	userWithEmptyEmail := UserWithEmptyEmail
	mock.ExpectBegin()
	mock.ExpectExec(`INSERT INTO "users"`).
		WithArgs(userWithEmptyEmail.ID, sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), userWithEmptyEmail.Email).
		WillReturnError(InvalidUserInputError)
	mock.ExpectRollback()

	err := userStorage.Create(userWithEmptyEmail, PasswordHash)
	assert.Error(t, err)
	assert.EqualError(t, err, InvalidUserInputError.Error())

	// Ensure all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
}

func TestUserStorage_Create_Already_Exist(t *testing.T) {
	// Initialize mock DB and UserStorage
	sqlDB, gormDB, mock := InitMockDB(t)
	defer sqlDB.Close()
	userStorage := UserStorage{DB: gormDB}

	// Test case: Successful creation
	user := User
	mock.ExpectBegin()
	mock.ExpectExec(`INSERT INTO "users"`).
		WithArgs(user.ID, sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), user.Email).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec(`INSERT INTO "credentials"`).
		WithArgs(user.ID, PasswordHash).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err := userStorage.Create(user, PasswordHash)
	assert.NoError(t, err)

	// Test case: Create user with same ID
	userWithSameID := UserWithSameID
	mock.ExpectBegin()
	mock.ExpectExec(`INSERT INTO "users"`).
		WithArgs(userWithSameID.ID, sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), userWithSameID.Email).
		WillReturnError(errors.New("duplicate key value violates unique constraint \"users_pkey\""))
	mock.ExpectRollback()

	err = userStorage.Create(userWithSameID, PasswordHash)
	assert.Error(t, err)
	assert.EqualError(t, err, UserAlreadyExistsError.Error())

	// Test case: Create user with same email
	userWithSameEmail := UserWithSameEmail
	mock.ExpectBegin()
	mock.ExpectExec(`INSERT INTO "users"`).
		WithArgs(userWithSameEmail.ID, sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), userWithSameEmail.Email).
		WillReturnError(errors.New("duplicate key value violates unique constraint \"users_email_key\""))
	mock.ExpectRollback()

	err = userStorage.Create(userWithSameEmail, PasswordHash)
	assert.Error(t, err)
	assert.EqualError(t, err, UserAlreadyExistsError.Error())

	// Ensure all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
}
