package impl

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"split-the-bill-server/domain/model"
	"split-the-bill-server/presentation/dto"
	"split-the-bill-server/storage"
	"split-the-bill-server/storage/mocks"
	"testing"
)

// Testdata
var (
	TestItem1Stored = model.ItemModel{
		ID:           uuid.New(),
		Name:         "Test Item 1",
		Price:        10,
		Contributors: []uuid.UUID{TestUser.ID},
	}

	TestItem2Stored = model.ItemModel{
		ID:           uuid.New(),
		Name:         "Test Item 2",
		Price:        18.5,
		Contributors: []uuid.UUID{TestUser.ID, TestUser2.ID},
	}

	TestBillStored = model.BillModel{
		ID:      uuid.New(),
		Name:    "Test Bill",
		OwnerID: TestUser.ID,
		Items:   []model.ItemModel{TestItem1Stored, TestItem2Stored},
	}

	TestItem1Updated = model.ItemModel{
		ID:           TestItem1Stored.ID,
		Name:         "Test Item 1 Updated",
		Price:        12,
		Contributors: []uuid.UUID{TestUser.ID},
	}

	TestBillUpdated = model.BillModel{
		ID:      TestBillStored.ID,
		Name:    "Test Bill Updated",
		OwnerID: TestUser.ID,
		Items:   []model.ItemModel{TestItem1Updated, TestItem2Stored},
	}
)

// TestBillService_Create tests whether the service creates a new bill model with its items from the given dto. The uuid should be newly set by the service, the id from the dto should be ignored.
func TestBillService_Create(t *testing.T) {

	// mock method
	mocks.MockBillCreate = func(bill model.BillModel) (model.BillModel, error) {
		return bill, nil
	}

	// input data
	// for the test we set each ID to uuid.Nil, service should create a new one
	item1 := dto.Item{
		ID: uuid.Nil,
		BaseItem: dto.BaseItem{
			Name:           TestItem1Stored.Name,
			Price:          TestItem1Stored.Price,
			ContributorIDs: TestItem1Stored.Contributors,
		},
	}
	item2 := dto.Item{
		ID: uuid.Nil,
		BaseItem: dto.BaseItem{
			Name:           TestItem2Stored.Name,
			Price:          TestItem2Stored.Price,
			ContributorIDs: TestItem2Stored.Contributors,
		},
	}
	bill := dto.Bill{
		ID: uuid.Nil,
		BaseBill: dto.BaseBill{
			OwnerID: TestBillStored.OwnerID,
			Name:    TestBillStored.Name,
			Items:   []dto.Item{item1, item2},
		},
	}

	ret, err := billService.Create(bill)
	assert.NoError(t, err)
	assert.NotNil(t, ret)
	assert.NotEqual(t, ret.ID, uuid.Nil)
	assert.Equal(t, bill.Name, ret.Name)
	assert.Equal(t, len(bill.Items), len(ret.Items))
	for i, item := range ret.Items {
		assert.NotEqual(t, uuid.Nil, item.ID)
		assert.Equal(t, bill.Items[i].Name, item.Name)
		assert.Equal(t, bill.Items[i].Price, item.Price)
		assert.Equal(t, bill.Items[i].ContributorIDs, item.ContributorIDs)
	}
}

func TestBillService_Update(t *testing.T) {

	tests := []struct {
		name        string
		mock        func()
		inputBill   dto.Bill
		expectedErr error
		want        model.BillModel
	}{
		{
			name: "Update Success, ID included in items",
			mock: func() {
				mocks.MockBillUpdate = func(bill model.BillModel) (model.BillModel, error) {
					return bill, nil
				}
				mocks.MockBillGetByID = func(id uuid.UUID) (model.BillModel, error) {
					return TestBillStored, nil
				}
			},
			inputBill: dto.Bill{
				ID: TestBillUpdated.ID,
				BaseBill: dto.BaseBill{
					OwnerID: TestBillUpdated.OwnerID,
					Name:    TestBillUpdated.Name,
					Items: []dto.Item{
						{
							ID: TestItem1Stored.ID,
							BaseItem: dto.BaseItem{
								Name:           TestItem1Updated.Name,
								Price:          TestItem1Updated.Price,
								ContributorIDs: TestItem1Updated.Contributors,
							},
						},
						{
							ID: TestItem2Stored.ID,
							BaseItem: dto.BaseItem{
								Name:           TestItem2Stored.Name,
								Price:          TestItem2Stored.Price,
								ContributorIDs: TestItem2Stored.Contributors,
							},
						},
					},
				},
			},
			expectedErr: nil,
			want:        TestBillUpdated,
		},
		{
			name: "Update Unsuccessful, Bad ID / not included in items",
			mock: func() {
				mocks.MockBillUpdate = func(bill model.BillModel) (model.BillModel, error) {
					return model.BillModel{}, storage.NoSuchItemError
				}
				mocks.MockBillGetByID = func(id uuid.UUID) (model.BillModel, error) {
					return TestBillStored, nil
				}
			},
			inputBill: dto.Bill{
				ID: TestBillUpdated.ID,
				BaseBill: dto.BaseBill{
					OwnerID: TestBillUpdated.OwnerID,
					Name:    TestBillUpdated.Name,
					Items: []dto.Item{
						{
							ID: TestItem1Stored.ID,
							BaseItem: dto.BaseItem{
								Name:           TestItem1Updated.Name,
								Price:          TestItem1Updated.Price,
								ContributorIDs: TestItem1Updated.Contributors,
							},
						},
					},
				},
			},
			expectedErr: storage.NoSuchItemError,
			want:        model.BillModel{},
		},
	}

	for _, testcase := range tests {
		t.Run(testcase.name, func(t *testing.T) {
			testcase.mock()

			ret, err := billService.Update(TestUser.ID, TestBillStored.ID, testcase.inputBill)
			assert.Equal(t, testcase.expectedErr, err)
			if err != nil { // -> test return
				assert.Equal(t, testcase.want.ID, ret.ID)
				assert.Equal(t, len(testcase.want.Items), len(ret.Items))
				for i, item := range ret.Items {
					assert.Equal(t, testcase.want.Items[i].ID, item.ID)
					assert.Equal(t, testcase.want.Items[i].Name, item.Name)
					assert.Equal(t, testcase.want.Items[i].Price, item.Price)
					assert.Equal(t, testcase.want.Items[i].Contributors, item.ContributorIDs)
				}
			}
		})
	}
}
