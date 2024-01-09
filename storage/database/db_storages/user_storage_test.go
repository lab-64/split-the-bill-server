package db_storages

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"log"
	"os"
	"split-the-bill-server/storage/database/entity"
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
// ExpectRollback for ROLLBACK if DB query fails

func TestGetByID(t *testing.T) {

	tests := []struct {
		name        string
		userUID     uuid.UUID
		mock        func()
		wantErr     bool
		expectedErr error
		want        entity.User
	}{
		{
			// When everything works as expected
			name:    "Get User by ID Success",
			userUID: User.ID,
			mock: func() {
				// We added one row
				userRows := sqlmock.NewRows([]string{"ID", "Email"}).AddRow(User.ID, "mail@mail.com")
				mock.ExpectQuery(`SELECT (.+) FROM "users"`).WithArgs(User.ID).WillReturnRows(userRows)
				groupInvitationRows := sqlmock.NewRows([]string{"ID", "InviteeID"}).AddRow(uuid.New(), User.ID) // Include field where user is stored
				mock.ExpectQuery(`SELECT (.+) FROM "group_invitations"`).WithArgs(User.ID).WillReturnRows(groupInvitationRows)
				groupMemberRows := sqlmock.NewRows([]string{"ID", "OwnerUID"}).AddRow(uuid.New(), User.ID)
				mock.ExpectQuery(`SELECT (.+) FROM "group_members"`).WithArgs(User.ID).WillReturnRows(groupMemberRows)
			},
			wantErr:     false,
			expectedErr: nil,
			want:        entity.User{Base: entity.Base{ID: User.ID}},
		},
		/*		{
					//When the role tried to access is not found
					name:    "Not Found",
					storage:       userStorage,
					userUID: uuid.New(),
					mock: func() {
						rows := sqlmock.NewRows([]string{"Id", "Email", "CreatedAt"}) //observe that we didnt add any role here
						mock.ExpectPrepare("SELECT (.+) FROM users").ExpectQuery().WithArgs(1).WillReturnRows(rows)
					},
					wantErr: true,
					want:    entity.User{},
				},
				{
					//When invalid statement is provided, ie the SQL syntax is wrong(in this case, we provided a wrong database)
					name:    "Invalid Prepare",
					storage:       userStorage,
					userUID: uuid.New(),
					mock: func() {
						rows := sqlmock.NewRows([]string{"Id", "Title", "Body", "CreatedAt"}).AddRow(1, "title", "body", created_at)
						mock.ExpectPrepare("SELECT (.+) FROM wrong_table").ExpectQuery().WithArgs(1).WillReturnRows(rows)
					},
					wantErr: true,
					want:    entity.User{},
				},*/
	}
	for _, testcase := range tests {
		t.Run(testcase.name, func(t *testing.T) {
			testcase.mock()
			got, err := userStorage.GetByID(testcase.userUID)

			// Test validation
			// test error
			assert.Equalf(t, testcase.wantErr, err != nil, "Get() error = %v, wantErr %v", err, testcase.wantErr)
			assert.Equalf(t, testcase.expectedErr, err, "Get() error = %v, expectedErr %v", err, testcase.expectedErr)
			// test returned data
			assert.Equalf(t, testcase.want.ID, got.ID, "Get() = %v, want %v", got.ID, testcase.want.ID)

			// Ensure all expectations were met
			if err = mock.ExpectationsWereMet(); err != nil {
				t.Errorf("Unfulfilled expectations: %s", err)
			}
		})
	}
}
