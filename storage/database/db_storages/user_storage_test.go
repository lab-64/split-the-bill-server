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
	TestUser = model.UserModel{
		ID:    uuid.New(),
		Email: "test@mail.com",
	}
	TestUserWithEmptyEmail = model.UserModel{
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
		want        model.UserModel
	}{
		{
			// TODO: Use seed user instead of test user
			name:    "Success",
			userUID: TestUser.ID,
			mock: func() {
				userRows := sqlmock.NewRows([]string{"ID", "Email"}).AddRow(TestUser.ID, TestUser.Email)
				dbMock.ExpectQuery(`SELECT (.+) FROM "users"`).WithArgs(TestUser.ID).WillReturnRows(userRows)
				groupInvitationRows := sqlmock.NewRows([]string{"ID", "InviteeID"}).AddRow(uuid.New(), TestUser.ID) // Include field where user is stored
				dbMock.ExpectQuery(`SELECT (.+) FROM "group_invitations"`).WithArgs(TestUser.ID).WillReturnRows(groupInvitationRows)
				groupMemberRows := sqlmock.NewRows([]string{"ID", "OwnerUID"}).AddRow(uuid.New(), TestUser.ID)
				dbMock.ExpectQuery(`SELECT (.+) FROM "group_members"`).WithArgs(TestUser.ID).WillReturnRows(groupMemberRows)
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
			want:        model.UserModel{},
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
		user        model.UserModel
		mock        func()
		wantErr     bool
		expectedErr error
		want        model.UserModel
	}{
		{
			name: "Success",
			user: TestUser,
			mock: func() {
				dbMock.ExpectBegin()
				dbMock.ExpectExec(`INSERT INTO "users"`).
					WithArgs(TestUser.ID, sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), TestUser.Email).
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
					WithArgs(TestUser.ID, sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), TestUser.Email).
					WillReturnError(gorm.ErrDuplicatedKey)
				dbMock.ExpectRollback()
			},
			wantErr:     true,
			expectedErr: storage.UserAlreadyExistsError,
			want:        model.UserModel{},
		},
		{
			name: "Invalid User Input",
			user: TestUserWithEmptyEmail,
			mock: func() {
				dbMock.ExpectBegin()
				dbMock.ExpectExec(`INSERT INTO "users"`).
					WithArgs(TestUserWithEmptyEmail.ID, sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), TestUserWithEmptyEmail.Email).
					WillReturnError(gorm.ErrInvalidData)
				dbMock.ExpectRollback()
			},
			wantErr:     true,
			expectedErr: storage.InvalidUserInputError,
			want:        model.UserModel{},
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
