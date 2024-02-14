package db_storages

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"split-the-bill-server/domain/model"
	"testing"
)

var TestGroup = model.Group{
	ID:   uuid.New(),
	Name: "Test Group",
}

func TestGroupStorage_DeleteGroup(t *testing.T) {

	tests := []struct {
		name        string
		groupID     uuid.UUID
		mock        func()
		expectedErr error
	}{
		{
			name:    "Success",
			groupID: TestGroup.ID,
			mock: func() {
				dbMock.ExpectBegin()
				dbMock.ExpectBegin()
				dbMock.ExpectExec(`UPDATE "groups" SET "deleted_at"`).WithArgs(sqlmock.AnyArg(), TestGroup.ID).WillReturnResult(sqlmock.NewResult(0, 0))
				dbMock.ExpectCommit()
				dbMock.ExpectBegin()
				dbMock.ExpectExec(`DELETE FROM "group_members"`).WithArgs(TestGroup.ID).WillReturnResult(sqlmock.NewResult(0, 0))
				dbMock.ExpectExec(`UPDATE "bills" SET "deleted_at"`).WithArgs(sqlmock.AnyArg(), TestGroup.ID).WillReturnResult(sqlmock.NewResult(0, 0))
				dbMock.ExpectExec(`UPDATE "group_invitations" SET "deleted_at"`).WithArgs(sqlmock.AnyArg(), TestGroup.ID).WillReturnResult(sqlmock.NewResult(0, 0))
				dbMock.ExpectExec(`UPDATE "groups" SET "deleted_at"`).WithArgs(sqlmock.AnyArg(), TestGroup.ID).WillReturnResult(sqlmock.NewResult(0, 0))
				dbMock.ExpectCommit()
				dbMock.ExpectCommit()
			},
			expectedErr: nil,
		},
	}

	for _, testcase := range tests {
		t.Run(testcase.name, func(t *testing.T) {
			testcase.mock()
			err := groupStorage.DeleteGroup(testcase.groupID)
			assert.Equalf(t, testcase.expectedErr, err, "Wrong error")

			// Ensure all expectations were met
			if err = dbMock.ExpectationsWereMet(); err != nil {
				t.Errorf("Unfulfilled expectations: %s", err)
			}
		})
	}
}
