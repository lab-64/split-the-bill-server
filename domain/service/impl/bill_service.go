package impl

import (
	"github.com/google/uuid"
	. "split-the-bill-server/domain/service/service_inf"
	. "split-the-bill-server/presentation/dto"
	. "split-the-bill-server/storage/storage_inf"
)

type BillService struct {
	billStorage  IBillStorage
	groupStorage IGroupStorage
}

func NewBillService(billStorage *IBillStorage, groupStorage *IGroupStorage) IBillService {
	return &BillService{billStorage: *billStorage, groupStorage: *groupStorage}
}

func (b *BillService) Create(billDTO BillInputDTO) (BillDetailedOutputDTO, error) {

	// create bill model including items
	bill := ToBillModel(billDTO)
	// store bill in billStorage
	bill, err := b.billStorage.Create(bill)
	if err != nil {
		return BillDetailedOutputDTO{}, err
	}

	return ToBillDetailedDTO(bill), err
}

func (b *BillService) GetByID(id uuid.UUID) (BillDetailedOutputDTO, error) {
	bill, err := b.billStorage.GetByID(id)
	if err != nil {
		return BillDetailedOutputDTO{}, err
	}

	return ToBillDetailedDTO(bill), err
}

func (b *BillService) AddItem(itemDTO ItemInputDTO) (ItemOutputDTO, error) {
	item := ToItemModel(uuid.Nil, itemDTO)

	item, err := b.billStorage.CreateItem(item)
	if err != nil {
		return ItemOutputDTO{}, err
	}

	return ToItemDTO(item), err
}

func (b *BillService) ChangeItem(itemID uuid.UUID, itemDTO ItemInputDTO) (ItemOutputDTO, error) {
	item := ToItemModel(itemID, itemDTO)

	item, err := b.billStorage.UpdateItem(item)
	if err != nil {
		return ItemOutputDTO{}, err
	}

	return ToItemDTO(item), err
}

func (b *BillService) GetItemByID(id uuid.UUID) (ItemOutputDTO, error) {
	item, err := b.billStorage.GetItemByID(id)
	if err != nil {
		return ItemOutputDTO{}, err
	}

	return ToItemDTO(item), err
}
