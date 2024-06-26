package impl

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"split-the-bill-server/domain"
	"split-the-bill-server/domain/model"
	"split-the-bill-server/presentation/dto"
	"split-the-bill-server/storage/mocks"
	"testing"
)

// Testdata
var (
	TestItem1 = model.Item{
		ID:           uuid.New(),
		Name:         "Test Item 1",
		Price:        10,
		Contributors: []model.User{TestUser},
	}

	TestItem2 = model.Item{
		ID:           uuid.New(),
		Name:         "Test Item 2",
		Price:        18.5,
		Contributors: []model.User{TestUser, TestUser2},
	}

	TestBill = model.Bill{
		ID:    uuid.New(),
		Name:  "Test Bill",
		Owner: TestUser,
		Items: []model.Item{TestItem1, TestItem2},
	}

	TestBillUpdated = model.Bill{
		ID:    TestBill.ID,
		Name:  "Test Bill Updated",
		Owner: TestUser,
		Items: TestBill.Items,
	}
)

func TestBillService_Update(t *testing.T) {

	tests := []struct {
		name          string
		mock          func()
		requesterID   uuid.UUID
		billID        uuid.UUID
		billUpdated   dto.BillUpdate
		expectedError error
		expectedBill  model.Bill
	}{
		{
			name: "Success",
			mock: func() {
				mocks.MockBillGetByID = func(id uuid.UUID) (model.Bill, error) {
					return TestBill, nil
				}
				mocks.MockGroupGetGroupByID = func(id uuid.UUID) (model.Group, error) {
					return TestGroup, nil
				}
				mocks.MockBillUpdate = func(bill model.Bill) (model.Bill, error) {
					return TestBillUpdated, nil
				}

			},
			requesterID: TestUser.ID,
			billID:      TestBill.ID,
			billUpdated: dto.BillUpdate{
				Name: TestBillUpdated.Name,
				Date: TestBillUpdated.Date,
			},
			expectedError: nil,
			expectedBill:  TestBillUpdated,
		},
		{
			name: "Not authorized",
			mock: func() {
				mocks.MockBillGetByID = func(id uuid.UUID) (model.Bill, error) {
					return TestBill, nil
				}
				mocks.MockGroupGetGroupByID = func(id uuid.UUID) (model.Group, error) {
					return TestGroup, nil
				}
			},
			requesterID:   uuid.New(),
			billID:        TestBill.ID,
			billUpdated:   dto.BillUpdate{},
			expectedError: domain.ErrNotAuthorized,
		},
	}

	for _, testcase := range tests {
		t.Run(testcase.name, func(t *testing.T) {
			testcase.mock()
			ret, err := billService.Update(testcase.requesterID, testcase.billID, testcase.billUpdated)
			assert.Equalf(t, testcase.expectedError, err, "Wrong error")
			if err == nil {
				assert.Equalf(t, testcase.expectedBill.ID, ret.ID, "Wrong BillID") // expect no changes in ID
				assert.Equalf(t, testcase.expectedBill.Name, ret.Name, "Wrong Bill Name")
				assert.Equalf(t, len(testcase.expectedBill.Items), len(ret.Items), "Wrong number of items") // expect no changes in items
				for i, item := range ret.Items {
					assert.Equalf(t, testcase.expectedBill.Items[i].ID, item.ID, "Wrong ItemID")
					assert.Equalf(t, testcase.expectedBill.Items[i].Name, item.Name, "Wrong Item Name")
					assert.Equalf(t, testcase.expectedBill.Items[i].Price, item.Price, "Wrong Item Price")
					for j, contributor := range item.Contributors {
						assert.Equalf(t, testcase.expectedBill.Items[i].Contributors[j].ID, contributor.ID, "Wrong ContributorID")
					}
				}
			}
		})
	}
}

func TestBillService_Create(t *testing.T) {
	tests := []struct {
		name         string
		mock         func()
		requesterID  uuid.UUID
		billDTO      dto.BillCreate
		expectedErr  error
		expectedBill model.Bill
	}{
		{
			name: "Success",
			mock: func() {
				mocks.MockBillCreate = func(bill model.Bill) (model.Bill, error) {
					return bill, nil
				}
				mocks.MockGroupGetGroupByID = func(id uuid.UUID) (model.Group, error) {
					return TestGroup, nil
				}
				mocks.MockUserGetByID = func(id uuid.UUID) (model.User, error) {
					return TestUser, nil
				}
			},
			requesterID: TestUser.ID,
			billDTO: dto.BillCreate{
				OwnerID: TestBill.Owner.ID,
				Name:    TestBill.Name,
				Items: []dto.ItemInput{
					{
						Name:         TestItem1.Name,
						Price:        TestItem1.Price,
						Contributors: []uuid.UUID{TestUser.ID},
					},
					{
						Name:         TestItem2.Name,
						Price:        TestItem2.Price,
						Contributors: []uuid.UUID{TestUser.ID, TestUser2.ID},
					},
				},
			},
			expectedErr:  nil,
			expectedBill: TestBill,
		},
		{
			name: "Unauthorized",
			mock: func() {
				mocks.MockGroupGetGroupByID = func(id uuid.UUID) (model.Group, error) {
					return TestGroup, nil
				}
				mocks.MockUserGetByID = func(id uuid.UUID) (model.User, error) {
					return TestUser, nil
				}
			},
			requesterID: uuid.New(),
			billDTO:     dto.BillCreate{},
			expectedErr: domain.ErrNotAuthorized,
		},
	}

	for _, testcase := range tests {
		t.Run(testcase.name, func(t *testing.T) {
			testcase.mock()
			ret, err := billService.Create(testcase.requesterID, testcase.billDTO)
			assert.Equalf(t, testcase.expectedErr, err, "Wrong error")
			if err == nil {
				assert.Equalf(t, testcase.expectedBill.Name, ret.Name, "Wrong Bill Name")
				assert.Equalf(t, len(testcase.expectedBill.Items), len(ret.Items), "Wrong number of items")
				for i, item := range ret.Items {
					assert.Equalf(t, testcase.expectedBill.Items[i].Name, item.Name, "Wrong Item Name")
					assert.Equalf(t, testcase.expectedBill.Items[i].Price, item.Price, "Wrong Item Price")
					for j, contributor := range item.Contributors {
						assert.Equalf(t, testcase.expectedBill.Items[i].Contributors[j].ID, contributor.ID, "Wrong ContributorID")
					}
				}
			}
		})
	}
}
