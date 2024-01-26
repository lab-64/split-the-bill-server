package impl

import (
	"github.com/google/uuid"
	"split-the-bill-server/domain"
	. "split-the-bill-server/domain/service"
	. "split-the-bill-server/presentation/dto"
	"split-the-bill-server/storage"
)

type BillService struct {
	billStorage  storage.IBillStorage
	groupStorage storage.IGroupStorage
}

func NewBillService(billStorage *storage.IBillStorage, groupStorage *storage.IGroupStorage) IBillService {
	return &BillService{billStorage: *billStorage, groupStorage: *groupStorage}
}

func (b *BillService) Create(billDTO BillInputDTO) (BillDetailedOutputDTO, error) {

	// create bill model including items
	bill := CreateBillModel(uuid.New(), billDTO)
	// store bill in billStorage
	bill, err := b.billStorage.Create(bill)
	if err != nil {
		return BillDetailedOutputDTO{}, err
	}

	return ConvertToBillDetailedDTO(bill), err
}

func (b *BillService) Update(userID uuid.UUID, billID uuid.UUID, billDTO BillInputDTO) (BillDetailedOutputDTO, error) {
	bill, err := b.billStorage.GetByID(billID)

	if err != nil {
		return BillDetailedOutputDTO{}, err
	}

	// Authorize
	if userID != bill.Owner.ID {
		return BillDetailedOutputDTO{}, domain.ErrNotAuthorized
	}

	updatedBill := CreateBillModel(billID, billDTO)
	updatedBill.ID = bill.ID

	bill, err = b.billStorage.UpdateBill(updatedBill)
	if err != nil {
		return BillDetailedOutputDTO{}, err
	}

	return ConvertToBillDetailedDTO(bill), err

}
func (b *BillService) GetByID(id uuid.UUID) (BillDetailedOutputDTO, error) {
	bill, err := b.billStorage.GetByID(id)
	if err != nil {
		return BillDetailedOutputDTO{}, err
	}
	return ConvertToBillDetailedDTO(bill), err
}

func (b *BillService) AddItem(itemDTO ItemInputDTO) (ItemOutputDTO, error) {
	item := CreateItemModel(uuid.Nil, itemDTO)

	item, err := b.billStorage.CreateItem(item)
	if err != nil {
		return ItemOutputDTO{}, err
	}

	return ConvertToItemDTO(item), err
}

func (b *BillService) ChangeItem(itemID uuid.UUID, itemDTO ItemInputDTO) (ItemOutputDTO, error) {
	item := CreateItemModel(itemID, itemDTO)

	item, err := b.billStorage.UpdateItem(item)
	if err != nil {
		return ItemOutputDTO{}, err
	}

	return ConvertToItemDTO(item), err
}

func (b *BillService) GetItemByID(id uuid.UUID) (ItemOutputDTO, error) {
	item, err := b.billStorage.GetItemByID(id)
	if err != nil {
		return ItemOutputDTO{}, err
	}
	return ConvertToItemDTO(item), err
}
