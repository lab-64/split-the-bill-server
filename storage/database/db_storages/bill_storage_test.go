package db_storages

import (
	"github.com/google/uuid"
	"split-the-bill-server/domain/model"
)

// Testdata
var (
	TestBill = model.Bill{
		ID:      uuid.New(),
		Name:    "Test Bill",
		GroupID: TestGroup.ID,
		Owner:   TestUser,
	}
)

/*
// TODO: Add expected SQL mock updating the items

func TestBillStorage_UpdateBill(t *testing.T) {
	tests := []struct {
		name           string
		bill           model.Bill
		mock           func()
		expectedErr    error
		expectedReturn model.Bill
	}{
		{
			name: "Success",
			bill: TestBill,
			mock: func() {
				dbMock.ExpectBegin()
				dbMock.ExpectBegin()
				dbMock.ExpectExec(`UPDATE "bills"`).WithArgs(sqlmock.AnyArg(), TestBill.Owner.ID, TestBill.Name, TestBill.GroupID, TestBill.ID).WillReturnResult(sqlmock.NewResult(1, 1))
				dbMock.ExpectCommit()
				dbMock.ExpectExec(`UPDATE "bills"`).WithArgs(sqlmock.AnyArg(), TestBill.ID).WillReturnResult(sqlmock.NewResult(1, 1))
				dbMock.ExpectExec(`DELETE FROM "unseen_bills"`).WithArgs(TestBill.ID).WillReturnResult(sqlmock.NewResult(1, 1))
				dbMock.ExpectCommit()
			},
			expectedErr:    nil,
			expectedReturn: TestBill,
		},
	}

	for _, testcase := range tests {
		t.Run(testcase.name, func(t *testing.T) {
			testcase.mock()
			bill, err := billStorage.UpdateBill(testcase.bill)
			assert.NoErrorf(t, err, "Wrong err")
			assert.Equalf(t, testcase.expectedReturn.Name, bill.Name, "Wrong bill name")

			// Ensure all expectations were met
			if err = dbMock.ExpectationsWereMet(); err != nil {
				t.Errorf("Unfulfilled expectations: %s", err)
			}
		})
	}
}

*/
