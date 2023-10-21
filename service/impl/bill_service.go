package impl

import (
	"github.com/google/uuid"
	"split-the-bill-server/common"
	"split-the-bill-server/dto"
	"split-the-bill-server/service"
	"split-the-bill-server/storage"
)

type BillService struct {
	storage.IBillStorage
	storage.IGroupStorage
}

func NewBillService(billStorage *storage.IBillStorage, groupStorage *storage.IGroupStorage) service.IBillService {
	return &BillService{IBillStorage: *billStorage, IGroupStorage: *groupStorage}
}

func (b *BillService) Create(billDTO dto.BillCreateDTO) (dto.BillDTO, error) {

	// TODO: delete if authentication is used
	userID := uuid.MustParse("7f1b2ed5-1201-4443-b997-56877fe31991")

	// create types_test.bill
	bill, err := billDTO.ToBill(userID)
	common.LogError(err)

	// store bill in billStorage
	err = b.IBillStorage.Create(bill)
	common.LogError(err)

	// add bill to group
	err = b.IGroupStorage.AddBillToGroup(&bill, billDTO.Group)
	common.LogError(err)

	return dto.ToBillDTO(&bill), err
}

func (b *BillService) GetByID(id uuid.UUID) (dto.BillDTO, error) {
	bill, err := b.IBillStorage.GetByID(id)
	common.LogError(err)

	return dto.ToBillDTO(&bill), err
}
