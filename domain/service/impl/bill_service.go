package impl

import (
	"github.com/google/uuid"
	"split-the-bill-server/domain"
	"split-the-bill-server/domain/model"
	. "split-the-bill-server/domain/service"
	"split-the-bill-server/presentation/dto"
	"split-the-bill-server/storage"
)

type BillService struct {
	billStorage  storage.IBillStorage
	groupStorage storage.IGroupStorage
}

func NewBillService(billStorage *storage.IBillStorage, groupStorage *storage.IGroupStorage) IBillService {
	return &BillService{billStorage: *billStorage, groupStorage: *groupStorage}
}

func (b *BillService) Create(billDTO dto.Bill) (dto.Bill, error) {

	// create new items
	items := make([]model.ItemModel, len(billDTO.Items))
	for i, item := range billDTO.Items {
		items[i] = dto.ToItemModel(uuid.New(), item.BaseItem)
	}
	// create new bill
	bill := dto.ToBillModel(uuid.New(), billDTO.BaseBill, items)

	// store bill in billStorage
	bill, err := b.billStorage.Create(bill)
	if err != nil {
		return dto.Bill{}, err
	}

	return dto.ToBillDTO(bill), err
}

func (b *BillService) Update(userID uuid.UUID, billID uuid.UUID, billDTO dto.Bill) (dto.Bill, error) {
	bill, err := b.billStorage.GetByID(billID)

	if err != nil {
		return dto.Bill{}, err
	}

	// Authorize
	if userID != bill.OwnerID {
		return dto.Bill{}, domain.ErrNotAuthorized
	}

	// update items
	items := make([]model.ItemModel, len(billDTO.Items))
	for i, item := range billDTO.Items {
		items[i] = dto.ToItemModel(item.ID, item.BaseItem) // use id from dto
	}
	// update bill
	updatedBill := dto.ToBillModel(billID, billDTO.BaseBill, items)

	bill, err = b.billStorage.UpdateBill(updatedBill)
	if err != nil {
		return dto.Bill{}, err
	}

	return dto.ToBillDTO(bill), err

}
func (b *BillService) GetByID(id uuid.UUID) (dto.BillDetailedOutputDTO, error) {
	bill, err := b.billStorage.GetByID(id)
	if err != nil {
		return dto.BillDetailedOutputDTO{}, err
	}

	return dto.ToBillDetailedDTO(bill), err
}

func (b *BillService) AddItem(itemDTO dto.Item) (dto.Item, error) {
	item := dto.ToItemModel(uuid.New(), itemDTO.BaseItem)

	item, err := b.billStorage.CreateItem(item)
	if err != nil {
		return dto.Item{}, err
	}

	return dto.ToItemDTO(item), err
}

func (b *BillService) ChangeItem(itemID uuid.UUID, itemDTO dto.Item) (dto.Item, error) {
	item := dto.ToItemModel(itemID, itemDTO.BaseItem)

	item, err := b.billStorage.UpdateItem(item)
	if err != nil {
		return dto.Item{}, err
	}

	return dto.ToItemDTO(item), err
}

func (b *BillService) GetItemByID(id uuid.UUID) (dto.Item, error) {
	item, err := b.billStorage.GetItemByID(id)
	if err != nil {
		return dto.Item{}, err
	}

	return dto.ToItemDTO(item), err
}
