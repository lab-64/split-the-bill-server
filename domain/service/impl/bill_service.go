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

func (b *BillService) Create(billDTO BillInputDTO) (BillOutputDTO, error) {

	// create bill model
	bill, err := ToBillModel(billDTO)
	if err != nil {
		return BillOutputDTO{}, err
	}

	// store bill in billStorage
	err = b.billStorage.Create(bill)
	if err != nil {
		return BillOutputDTO{}, err
	}

	return ToBillDTO(bill), err
}

func (b *BillService) GetByID(id uuid.UUID) (BillOutputDTO, error) {
	bill, err := b.billStorage.GetByID(id)
	if err != nil {
		return BillOutputDTO{}, err
	}

	return ToBillDTO(bill), err
}

func (b *BillService) AddItem(itemDTO ItemInputDTO) (ItemOutputDTO, error) {
	item := ToItemModel(itemDTO)

	item, err := b.billStorage.CreateItem(item)
	if err != nil {
		return ItemOutputDTO{}, err
	}

	return ToItemDTO(item), err
}

func (b *BillService) AddItemContributor(itemContributorDTO ItemContributorInputDTO) (ItemOutputDTO, error) {
	// get item
	item, err := b.billStorage.GetItemByID(itemContributorDTO.ItemID)
	if err != nil {
		return ItemOutputDTO{}, err
	}

	// TODO: update item
	item.Contributors = append(item.Contributors, itemContributorDTO.Contributors...)
	// store item
	item, err = b.billStorage.UpdateItem(item)
	if err != nil {
		return ItemOutputDTO{}, err
	}

	return ToItemDTO(item), err
}
