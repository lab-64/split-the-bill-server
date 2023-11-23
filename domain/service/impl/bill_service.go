package impl

import (
	"github.com/google/uuid"
	"split-the-bill-server/core"
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

	// TODO: delete or move to ephemeral bill storage to the create function
	// add bill to group
	err = b.groupStorage.AddBillToGroup(&bill, billDTO.Group)
	if err != nil {
		return BillOutputDTO{}, err
	}

	return ToBillDTO(bill), err
}

func (b *BillService) GetByID(id uuid.UUID) (BillOutputDTO, error) {
	bill, err := b.billStorage.GetByID(id)
	core.LogError(err)

	return ToBillDTO(bill), err
}
