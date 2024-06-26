package db_storages

import (
	"github.com/google/uuid"
	"split-the-bill-server/domain/model"
)

var TestGroup = model.Group{
	ID:   uuid.New(),
	Name: "Test Group",
}

// TODO: update test
/*
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
				dbMock.ExpectExec(`DELETE FROM "groups"`).WithArgs(TestGroup.ID).WillReturnResult(sqlmock.NewResult(0, 0))
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
*/
