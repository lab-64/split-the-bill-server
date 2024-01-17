package impl

import (
	"github.com/google/uuid"
	"split-the-bill-server/domain/model"
)

// Testdata
var (
	TestBill = model.BillModel{
		ID:      uuid.New(),
		Name:    "Test Bill",
		OwnerID: TestUser.ID,
		Items:   []model.ItemModel{TestItem1, TestItem2},
	}

	TestItem1 = model.ItemModel{
		ID:           uuid.New(),
		Name:         "Test Item 1",
		Price:        10,
		Contributors: []uuid.UUID{TestUser.ID},
	}

	TestItem2 = model.ItemModel{
		ID:           uuid.New(),
		Name:         "Test Item 2",
		Price:        18.5,
		Contributors: []uuid.UUID{TestUser.ID, TestUser2.ID},
	}
)
