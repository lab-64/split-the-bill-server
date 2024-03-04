package db_storages

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"split-the-bill-server/domain/model"
	"split-the-bill-server/storage"
	"testing"
)

// Testdata
var (
	TestUser = model.User{
		ID:    uuid.New(),
		Email: "test@mail.com",
	}
	TestUserWithEmptyEmail = model.User{
		ID:    uuid.New(),
		Email: "",
	}
	TestPasswordHash, _ = bcrypt.GenerateFromPassword([]byte("test1337"), 10)
)

func TestGetByID(t *testing.T) {

	tests := []struct {
		name        string
		userUID     uuid.UUID
		mock        func()
		wantErr     bool
		expectedErr error
		want        model.User
	}{
		{
			name:    "Success",
			userUID: TestUser.ID,
			mock: func() {
				userRows := sqlmock.NewRows([]string{"ID", "Email"}).AddRow(TestUser.ID, TestUser.Email)
				dbMock.ExpectQuery(`SELECT (.+) FROM "users"`).WithArgs(TestUser.ID).WillReturnRows(userRows)
			},
			wantErr:     false,
			expectedErr: nil,
			want:        TestUser,
		},
		{
			name:    "Not Found",
			userUID: TestUser.ID,
			mock: func() {
				dbMock.ExpectQuery(`SELECT (.+) FROM "users"`).WithArgs(TestUser.ID).WillReturnError(gorm.ErrRecordNotFound) // if this query returns an error the other queries featuring the preload entities won't be executed
			},
			wantErr:     true,
			expectedErr: storage.NoSuchUserError,
			want:        model.User{},
		},
	}
	for _, testcase := range tests {
		t.Run(testcase.name, func(t *testing.T) {
			testcase.mock()
			got, err := userStorage.GetByID(testcase.userUID)

			// Validate error
			assert.Equalf(t, testcase.wantErr, err != nil, "Get() error = %v, wantErr %v", err, testcase.wantErr)
			assert.Equalf(t, testcase.expectedErr, err, "Get() error = %v, expectedErr %v", err, testcase.expectedErr)
			// Validate returned data if err == nil
			if err == nil {
				assert.Equalf(t, testcase.want.ID, got.ID, "Get() = %v, want %v", got.ID, testcase.want.ID)
			}

			// Ensure all expectations were met
			if err = dbMock.ExpectationsWereMet(); err != nil {
				t.Errorf("Unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestCreate(t *testing.T) {

	tests := []struct {
		name        string
		user        model.User
		mock        func()
		wantErr     bool
		expectedErr error
		want        model.User
	}{
		{
			name: "Success",
			user: TestUser,
			mock: func() {
				dbMock.ExpectBegin()
				dbMock.ExpectExec(`INSERT INTO "users"`).
					WithArgs(TestUser.ID, sqlmock.AnyArg(), sqlmock.AnyArg(), TestUser.Email, "", "").
					WillReturnResult(sqlmock.NewResult(1, 1))
				dbMock.ExpectExec(`INSERT INTO "credentials"`).
					WithArgs(TestUser.ID, TestPasswordHash).
					WillReturnResult(sqlmock.NewResult(1, 1))
				dbMock.ExpectCommit()
			},
			wantErr:     false,
			expectedErr: nil,
			want:        TestUser,
		},
		{
			name: "User Already Exists",
			user: TestUser,
			mock: func() {
				dbMock.ExpectBegin()
				dbMock.ExpectExec(`INSERT INTO "users"`).
					WithArgs(TestUser.ID, sqlmock.AnyArg(), sqlmock.AnyArg(), TestUser.Email, "", "").
					WillReturnError(gorm.ErrDuplicatedKey)
				dbMock.ExpectRollback()
			},
			wantErr:     true,
			expectedErr: storage.UserAlreadyExistsError,
			want:        model.User{},
		},
		{
			name: "Invalid User Input",
			user: TestUserWithEmptyEmail,
			mock: func() {
				dbMock.ExpectBegin()
				dbMock.ExpectExec(`INSERT INTO "users"`).
					WithArgs(TestUserWithEmptyEmail.ID, sqlmock.AnyArg(), sqlmock.AnyArg(), TestUserWithEmptyEmail.Email, "", "").
					WillReturnError(gorm.ErrInvalidData)
				dbMock.ExpectRollback()
			},
			wantErr:     true,
			expectedErr: storage.InvalidUserInputError,
			want:        model.User{},
		},
	}

	for _, testcase := range tests {
		t.Run(testcase.name, func(t *testing.T) {
			testcase.mock()
			ret, err := userStorage.Create(testcase.user, TestPasswordHash)

			// Validate error
			assert.Equalf(t, testcase.wantErr, err != nil, "Get() error = %v, wantErr %v", err, testcase.wantErr)
			assert.Equalf(t, testcase.expectedErr, err, "Get() error = %v, expectedErr %v", err, testcase.expectedErr)

			// Validate returned data if err == nil
			if err == nil {
				assert.Equalf(t, testcase.want.ID, ret.ID, "Get() = %v, want %v", ret.ID, testcase.want.ID)
			}

			// Ensure all expectations were met
			if err = dbMock.ExpectationsWereMet(); err != nil {
				t.Errorf("Unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestUpdate(t *testing.T) {

	tests := []struct {
		name        string
		user        model.User
		mock        func()
		expectedErr error
		want        model.User
	}{
		{
			name: "Success",
			user: TestUser,
			mock: func() {
				dbMock.ExpectBegin()
				dbMock.ExpectExec(`UPDATE "users"`).
					WithArgs(sqlmock.AnyArg(), TestUser.ID).
					WillReturnResult(sqlmock.NewResult(1, 1))
				dbMock.ExpectCommit()
				dbMock.ExpectQuery(`SELECT (.+) FROM "users"`).WithArgs(TestUser.ID).WillReturnRows(sqlmock.NewRows([]string{"ID", "Email"}).AddRow(TestUser.ID, TestUser.Email))
			},
			expectedErr: nil,
			want:        TestUser,
		},
	}

	for _, testcase := range tests {
		t.Run(testcase.name, func(t *testing.T) {
			testcase.mock()
			ret, err := userStorage.Update(testcase.user)

			// Validate error
			assert.Equal(t, testcase.expectedErr, err)

			// Validate returned data if err == nil
			if err == nil {
				assert.Equal(t, testcase.want.ID, ret.ID)
			}

			// Ensure all expectations were met
			if err = dbMock.ExpectationsWereMet(); err != nil {
				t.Errorf("Unfulfilled expectations: %s", err)
			}
		})
	}
}
