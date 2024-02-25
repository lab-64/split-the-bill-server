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

func (b *BillService) Create(requesterID uuid.UUID, billDTO dto.BillInput) (dto.BillDetailedOutput, error) {
	// Validate groupID
	group, err := b.groupStorage.GetGroupByID(billDTO.GroupID)
	if err != nil {
		return dto.BillDetailedOutput{}, err
	}
	// Authorization
	if !group.IsMember(requesterID) {
		return dto.BillDetailedOutput{}, domain.ErrNotAuthorized
	}

	// add all group members to unseen list except the owner
	groupMembers := group.Members
	unseenFrom := make([]uuid.UUID, 0)
	for _, member := range groupMembers {
		if member.ID != group.Owner.ID {
			unseenFrom = append(unseenFrom, member.ID)
		}
	}
	// create bill model including items
	bill := model.CreateBill(uuid.New(), billDTO, time.Now(), unseenFrom)
	// store bill in billStorage
	bill, err = b.billStorage.Create(bill)
	if err != nil {
		return dto.BillDetailedOutput{}, err
	}

	return converter.ToBillDetailedDTO(bill), err
}

func (b *BillService) Update(requesterID uuid.UUID, billID uuid.UUID, billDTO dto.BillInput) (dto.BillDetailedOutput, error) {
	// Get bill
	bill, err := b.billStorage.GetByID(billID)
	if err != nil {
		return dto.BillDetailedOutput{}, err
	}
	// Authorization
	if requesterID != bill.Owner.ID {
		return dto.BillDetailedOutput{}, domain.ErrNotAuthorized
	}
	// TODO: change unseen list if bill gets updated?
	updatedBill := model.CreateBill(billID, billDTO, billDTO.Date, bill.UnseenFromUserID)
	updatedBill.ID = bill.ID
	bill, err = b.billStorage.UpdateBill(updatedBill)
	if err != nil {
		return dto.BillDetailedOutput{}, err
	}

	return converter.ToBillDetailedDTO(bill), err

}
func (b *BillService) GetByID(requesterID uuid.UUID, id uuid.UUID) (dto.BillDetailedOutput, error) {
	bill, err := b.billStorage.GetByID(id)
	if err != nil {
		return dto.BillDetailedOutput{}, err
	}

	// Get group
	group, err := b.groupStorage.GetGroupByID(bill.GroupID)
	if err != nil {
		return dto.BillDetailedOutput{}, err
	}
	// Authorization
	if !group.IsMember(requesterID) {
		return dto.BillDetailedOutput{}, domain.ErrNotAuthorized
	}

	balance := bill.CalculateBalance()
	bill.Balance = balance

	return converter.ToBillDetailedDTO(bill), err
}

func (b *BillService) GetAllByUserID(requesterID uuid.UUID, userID uuid.UUID, isUnseen bool, isOwner bool) ([]dto.BillDetailedOutput, error) {
	// Authorization
	if userID != requesterID {
		return nil, domain.ErrNotAuthorized
	}
	// get groups from user
	groups, err := b.groupStorage.GetGroups(userID, uuid.Nil)
	if err != nil {
		return nil, err
	}
	// run through all bills and return only the ones that match the filter
	var billDTOs []dto.BillDetailedOutput
	for _, group := range groups {
		for _, bill := range group.Bills {
			// apply isUnseen and isOwner filter
			if (isUnseen && bill.IsUnseen(userID)) || (isOwner && bill.Owner.ID == userID) {
				// set balance
				billPointer := &bill
				billPointer.Balance = billPointer.CalculateBalance()
				// store bill in return array
				billDTOs = append(billDTOs, converter.ToBillDetailedDTO(bill))
			}
			// if no filter is set, return all bills
			if !isUnseen && !isOwner {
				// set balance
				billPointer := &bill
				billPointer.Balance = billPointer.CalculateBalance()
				// store bill in return array
				billDTOs = append(billDTOs, converter.ToBillDetailedDTO(bill))
			}
		}
	}
	return billDTOs, err
}

func (b *BillService) AddItem(requesterID uuid.UUID, itemDTO dto.ItemInput) (dto.ItemOutput, error) {
	// Validate billID
	bill, err := b.billStorage.GetByID(itemDTO.BillID)
	if err != nil {
		return dto.ItemOutput{}, err
	}
	// Authorization
	if requesterID != bill.Owner.ID {
		return dto.ItemOutput{}, domain.ErrNotAuthorized
	}

	// TODO: check if contributors are members of the group

	item := model.CreateItem(uuid.New(), itemDTO)

	item, err = b.billStorage.CreateItem(item)
	if err != nil {
		return dto.ItemOutput{}, err
	}

	return converter.ToItemDTO(item), err
}

func (b *BillService) ChangeItem(requesterID uuid.UUID, itemID uuid.UUID, itemDTO dto.ItemInput) (dto.ItemOutput, error) {
	// Validate billID
	bill, err := b.billStorage.GetByID(itemDTO.BillID)
	if err != nil {
		return dto.ItemOutput{}, err
	}
	// TODO: allow all group members to change the contribution lst
	// Authorization
	if requesterID != bill.Owner.ID {
		return dto.ItemOutput{}, domain.ErrNotAuthorized
	}

	item := model.CreateItem(itemID, itemDTO)

	item, err = b.billStorage.UpdateItem(item)
	if err != nil {
		return dto.ItemOutput{}, err
	}

	return converter.ToItemDTO(item), err
}

func (b *BillService) GetItemByID(requesterID uuid.UUID, id uuid.UUID) (dto.ItemOutput, error) {
	// Get item
	item, err := b.billStorage.GetItemByID(id)
	if err != nil {
		return dto.ItemOutput{}, err
	}
	// Get bill
	bill, err := b.billStorage.GetByID(item.BillID)
	if err != nil {
		return dto.ItemOutput{}, err
	}
	// Get group
	group, err := b.groupStorage.GetGroupByID(bill.GroupID)
	if err != nil {
		return dto.ItemOutput{}, err
	}
	// Authorization
	if !group.IsMember(requesterID) {
		return dto.ItemOutput{}, domain.ErrNotAuthorized
	}

	return converter.ToItemDTO(item), err
}

func (b *BillService) DeleteItem(requesterID uuid.UUID, itemID uuid.UUID) error {
	// Get item
	item, err := b.billStorage.GetItemByID(itemID)
	if err != nil {
		return err
	}
	// Get bill
	bill, err := b.billStorage.GetByID(item.BillID)
	if err != nil {
		return err
	}
	// Authorization
	if requesterID != bill.Owner.ID {
		return domain.ErrNotAuthorized
	}
	// Delete item
	return b.billStorage.DeleteItem(itemID)
}
