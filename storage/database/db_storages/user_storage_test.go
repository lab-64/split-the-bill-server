package db_storages

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"log"
	"os"
	. "split-the-bill-server/domain/model"
	"split-the-bill-server/storage"
	. "split-the-bill-server/storage/database/test_util"
	"testing"
)

var (
	sqlDB       *sql.DB
	mock        sqlmock.Sqlmock
	gormDB      *gorm.DB
	userStorage UserStorage
)

func TestMain(m *testing.M) {
	// Perform setup
	var err error
	// Initialize mock db
	sqlDB, gormDB, mock, err = InitMockDB()
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

// Reminder:
// ExpectQuery for SELECT Query
// ExpectExec for INSERT, UPDATE, DELETE, ...
// ExpectRollback if DB query fails on INSERT
// ExpectCommit if DB query succeeds on INSERT
// ExpectBegin if TRANSACTION is started

func TestGetByID(t *testing.T) {

	tests := []struct {
		name        string
		userUID     uuid.UUID
		mock        func()
		wantErr     bool
		expectedErr error
		want        UserModel
	}{
		{
			name:    "Success",
			userUID: User.ID,
			mock: func() {
				userRows := sqlmock.NewRows([]string{"ID", "Email"}).AddRow(User.ID, "mail@mail.com")
				mock.ExpectQuery(`SELECT (.+) FROM "users"`).WithArgs(User.ID).WillReturnRows(userRows)
				groupInvitationRows := sqlmock.NewRows([]string{"ID", "InviteeID"}).AddRow(uuid.New(), User.ID) // Include field where user is stored
				mock.ExpectQuery(`SELECT (.+) FROM "group_invitations"`).WithArgs(User.ID).WillReturnRows(groupInvitationRows)
				groupMemberRows := sqlmock.NewRows([]string{"ID", "OwnerUID"}).AddRow(uuid.New(), User.ID)
				mock.ExpectQuery(`SELECT (.+) FROM "group_members"`).WithArgs(User.ID).WillReturnRows(groupMemberRows)
			},
			wantErr:     false,
			expectedErr: nil,
			want:        UserModel{ID: User.ID, Email: User.Email},
		},
		{
			name:    "Not Found",
			userUID: User.ID,
			mock: func() {
				mock.ExpectQuery(`SELECT (.+) FROM "users"`).WithArgs(User.ID).WillReturnError(gorm.ErrRecordNotFound) // if this query returns an error the other queries featuring the preload entities won't be executed
			},
			wantErr:     true,
			expectedErr: storage.NoSuchUserError,
			want:        UserModel{},
		},
	}
	for _, testcase := range tests {
		t.Run(testcase.name, func(t *testing.T) {
			testcase.mock()
			got, err := userStorage.GetByID(testcase.userUID)

			// Validate error
			assert.Equalf(t, testcase.wantErr, err != nil, "Get() error = %v, wantErr %v", err, testcase.wantErr)
			assert.Equalf(t, testcase.expectedErr, err, "Get() error = %v, expectedErr %v", err, testcase.expectedErr)
			// Validate returned data if err != nil
			if err != nil {
				assert.Equalf(t, testcase.want.ID, got.ID, "Get() = %v, want %v", got.ID, testcase.want.ID)
			}

			// Ensure all expectations were met
			if err = mock.ExpectationsWereMet(); err != nil {
				t.Errorf("Unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestCreate(t *testing.T) {

	tests := []struct {
		name        string
		user        UserModel
		mock        func()
		wantErr     bool
		expectedErr error
		want        UserModel
	}{
		{
			name: "Success",
			user: User,
			mock: func() {
				mock.ExpectBegin()
				mock.ExpectExec(`INSERT INTO "users"`).
					WithArgs(User.ID, sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), User.Email).
					WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectExec(`INSERT INTO "credentials"`).
					WithArgs(User.ID, PasswordHash).
					WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()
			},
			wantErr:     false,
			expectedErr: nil,
			want:        UserModel{},
		},
		{
			name: "User Already Exists",
			user: User,
			mock: func() {
				mock.ExpectBegin()
				mock.ExpectExec(`INSERT INTO "users"`).
					WithArgs(User.ID, sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), User.Email).
					WillReturnError(gorm.ErrDuplicatedKey)
				mock.ExpectRollback()
			},
			wantErr:     true,
			expectedErr: storage.UserAlreadyExistsError,
			want:        UserModel{},
		},
		{
			name: "Invalid User Input",
			user: UserWithEmptyEmail,
			mock: func() {
				mock.ExpectBegin()
				mock.ExpectExec(`INSERT INTO "users"`).
					WithArgs(UserWithEmptyEmail.ID, sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), UserWithEmptyEmail.Email).
					WillReturnError(gorm.ErrInvalidData)
				mock.ExpectRollback()
			},
			wantErr:     true,
			expectedErr: storage.InvalidUserInputError,
			want:        UserModel{},
		},
	}

	for _, testcase := range tests {
		t.Run(testcase.name, func(t *testing.T) {
			testcase.mock()
			err := userStorage.Create(testcase.user, PasswordHash)

			// Validate error
			assert.Equalf(t, testcase.wantErr, err != nil, "Get() error = %v, wantErr %v", err, testcase.wantErr)
			assert.Equalf(t, testcase.expectedErr, err, "Get() error = %v, expectedErr %v", err, testcase.expectedErr)

			// Ensure all expectations were met
			if err = mock.ExpectationsWereMet(); err != nil {
				t.Errorf("Unfulfilled expectations: %s", err)
			}
		})
	}
}
