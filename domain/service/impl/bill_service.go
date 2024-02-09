package impl

import (
	"github.com/google/uuid"
	"split-the-bill-server/domain"
	"split-the-bill-server/domain/converter"
	"split-the-bill-server/domain/model"
	. "split-the-bill-server/domain/service"
	"split-the-bill-server/presentation/dto"
	"split-the-bill-server/storage"
	"time"
)

type BillService struct {
	billStorage  storage.IBillStorage
	groupStorage storage.IGroupStorage
}

func NewBillService(billStorage *storage.IBillStorage, groupStorage *storage.IGroupStorage) IBillService {
	return &BillService{billStorage: *billStorage, groupStorage: *groupStorage}
}

func (b *BillService) Create(billDTO dto.BillInput) (dto.BillDetailedOutput, error) {

	// create bill model including items
	bill := model.CreateBill(uuid.New(), billDTO, time.Now())
	// store bill in billStorage
	bill, err := b.billStorage.Create(bill)
	if err != nil {
		return dto.BillDetailedOutput{}, err
	}

	return converter.ToBillDetailedDTO(bill), err
}

func (b *BillService) Update(userID uuid.UUID, billID uuid.UUID, billDTO dto.BillInput) (dto.BillDetailedOutput, error) {
	bill, err := b.billStorage.GetByID(billID)

	if err != nil {
		return dto.BillDetailedOutput{}, err
	}

	// Authorize
	if userID != bill.Owner.ID {
		return dto.BillDetailedOutput{}, domain.ErrNotAuthorized
	}

	updatedBill := model.CreateBill(billID, billDTO, billDTO.Date)
	updatedBill.ID = bill.ID

	bill, err = b.billStorage.UpdateBill(updatedBill)
	if err != nil {
		return dto.BillDetailedOutput{}, err
	}

	return converter.ToBillDetailedDTO(bill), err

}
func (b *BillService) GetByID(id uuid.UUID) (dto.BillDetailedOutput, error) {
	bill, err := b.billStorage.GetByID(id)
	if err != nil {
		return dto.BillDetailedOutput{}, err
	}

	balance := bill.CalculateBalance()
	bill.Balance = balance

	return converter.ToBillDetailedDTO(bill), err
}

func (b *BillService) AddItem(itemDTO dto.ItemInput) (dto.ItemOutput, error) {
	item := model.CreateItem(uuid.Nil, itemDTO)

	item, err := b.billStorage.CreateItem(item)
	if err != nil {
		return dto.ItemOutput{}, err
	}

	return converter.ToItemDTO(item), err
}

func (b *BillService) ChangeItem(itemID uuid.UUID, itemDTO dto.ItemInput) (dto.ItemOutput, error) {
	item := model.CreateItem(itemID, itemDTO)

	item, err := b.billStorage.UpdateItem(item)
	if err != nil {
		return dto.ItemOutput{}, err
	}

	return converter.ToItemDTO(item), err
}

func (b *BillService) GetItemByID(id uuid.UUID) (dto.ItemOutput, error) {
	item, err := b.billStorage.GetItemByID(id)
	if err != nil {
		return dto.ItemOutput{}, err
	}
	return converter.ToItemDTO(item), err
}
