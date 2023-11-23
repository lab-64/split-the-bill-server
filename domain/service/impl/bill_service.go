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
	// TODO: delete if authentication is used
	billDTO.Owner = uuid.MustParse("7f1b2ed5-1201-4443-b997-56877fe31991")

	// create types_test.bill
	bill, err := ToBillModel(billDTO)
	if err != nil {
		return BillOutputDTO{}, err
	}

	// store bill in billStorage
	err = b.billStorage.Create(bill)
	if err != nil {
		return BillOutputDTO{}, err
	}

	// add bill to group
	err = b.groupStorage.AddBillToGroup(&bill, billDTO.Group)
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
