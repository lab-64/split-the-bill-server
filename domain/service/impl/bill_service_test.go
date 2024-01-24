package impl

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"split-the-bill-server/domain/model"
	"split-the-bill-server/presentation/dto"
	"split-the-bill-server/storage/mocks"
	"testing"
)

// Testdata
var (
	TestItem1 = model.ItemModel{
		ID:           uuid.New(),
		Name:         "Test Item 1",
		Price:        10,
		Contributors: []uuid.UUID{TestUser.ID},
	}

	TestItem1Updated = model.ItemModel{
		ID:           TestItem1.ID,
		Name:         "Test Item 1 Updated",
		Price:        12,
		Contributors: []uuid.UUID{TestUser.ID},
	}

	TestItem2 = model.ItemModel{
		ID:           uuid.New(),
		Name:         "Test Item 2",
		Price:        18.5,
		Contributors: []uuid.UUID{TestUser.ID, TestUser2.ID},
	}

	TestBill = model.BillModel{
		ID:      uuid.New(),
		Name:    "Test Bill",
		OwnerID: TestUser.ID,
		Items:   []model.ItemModel{TestItem1, TestItem2},
	}

	TestBillUpdated = model.BillModel{
		ID:      TestBill.ID,
		Name:    "Test Bill Updated",
		OwnerID: TestUser.ID,
		Items:   []model.ItemModel{TestItem1Updated, TestItem2},
	}
)

func TestBillService_Update(t *testing.T) {

	// mock method
	mocks.MockBillUpdate = func(bill model.BillModel) (model.BillModel, error) {
		return TestBillUpdated, nil
	}
	mocks.MockBillGetByID = func(id uuid.UUID) (model.BillModel, error) {
		return TestBill, nil
	}

	itemUpdated := dto.ItemInputDTO{
		Name:         TestItem1Updated.Name,
		Price:        TestItem1Updated.Price,
		Contributors: TestItem1Updated.Contributors,
	}
	item2 := dto.ItemInputDTO{
		Name:         TestItem2.Name,
		Price:        TestItem2.Price,
		Contributors: TestItem2.Contributors,
	}

	// updated fields
	billUpdated := dto.BillInputDTO{
		Owner: TestBillUpdated.OwnerID,
		Name:  TestBillUpdated.Name,
		Items: []dto.ItemInputDTO{itemUpdated, item2},
	}

	ret, err := billService.Update(TestUser.ID, TestBill.ID, billUpdated)
	assert.NoError(t, err)
	assert.NotNil(t, ret)
	assert.Equal(t, TestBillUpdated.ID, ret.ID)
	assert.Equal(t, len(TestBillUpdated.Items), len(ret.Items))
	for i, item := range ret.Items {
		assert.Equal(t, TestBillUpdated.Items[i].ID, item.ID)
		assert.Equal(t, TestBillUpdated.Items[i].Name, item.Name)
		assert.Equal(t, TestBillUpdated.Items[i].Price, item.Price)
		assert.Equal(t, TestBillUpdated.Items[i].Contributors, item.Contributors)
	}
}
