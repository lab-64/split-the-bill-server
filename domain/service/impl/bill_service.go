package impl

import (
	"github.com/google/uuid"
	"sort"
	"split-the-bill-server/domain"
	"split-the-bill-server/domain/converter"
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

func (b *BillService) Create(requesterID uuid.UUID, billDTO dto.BillCreate) (dto.BillDetailedOutput, error) {
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
		if member.ID != billDTO.OwnerID {
			unseenFrom = append(unseenFrom, member.ID)
		}
	}
	// create new bill model
	bill := model.CreateBill(uuid.New(), billDTO.OwnerID, billDTO.Name, billDTO.Date, billDTO.GroupID, nil, unseenFrom)
	// create new items
	var items []model.Item
	for _, item := range billDTO.Items {
		items = append(items, model.CreateItem(uuid.New(), bill.ID, item))
	}
	// add items to bill
	bill.Items = items
	// store bill in billStorage
	bill, err = b.billStorage.Create(bill)
	if err != nil {
		return dto.BillDetailedOutput{}, err
	}

	return converter.ToBillDetailedDTO(bill), err
}

func (b *BillService) Update(requesterID uuid.UUID, billID uuid.UUID, billDTO dto.BillUpdate) (dto.BillDetailedOutput, error) {
	// Get bill
	// TODO: just update bill and not create new one
	bill, err := b.billStorage.GetByID(billID)
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

	// update bill base fields
	bill.UpdateBill(billDTO)
	// delete user from unseen list if bill is viewed
	if billDTO.Viewed != nil {
		if *billDTO.Viewed {
			bill.UnseenFromUserID = removeEntryFromSlice(bill.UnseenFromUserID, requesterID)
		}
	}
	// store updated bill
	bill, err = b.billStorage.UpdateBill(bill)
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

func (b *BillService) Delete(requesterID uuid.UUID, id uuid.UUID) error {
	// Get bill
	bill, err := b.billStorage.GetByID(id)
	if err != nil {
		return err
	}
	// Authorization
	if requesterID != bill.Owner.ID {
		return domain.ErrNotAuthorized
	}
	// Delete bill
	return b.billStorage.DeleteBill(id)
}

func (b *BillService) GetAllByUserID(requesterID uuid.UUID, userID uuid.UUID, isUnseen *bool, isOwner *bool) ([]dto.BillDetailedOutput, error) {
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
			// apply isUnseen filter
			if isUnseen != nil && *isUnseen != bill.IsUnseen(userID) {
				continue
			}
			// apply isOwner filter
			if isOwner != nil && *isOwner != (bill.Owner.ID == userID) {
				continue
			}

			// set balance
			billPointer := &bill
			billPointer.Balance = billPointer.CalculateBalance()
			// store bill in return array
			billDTOs = append(billDTOs, converter.ToBillDetailedDTO(bill))
		}
	}
	// sort bills by date descending
	orderedBills := orderBillsByDate(billDTOs)
	return orderedBills, err
}

// removeEntryFromSlice removes the first occurrence of entry from slice
func removeEntryFromSlice(slice []uuid.UUID, entry uuid.UUID) []uuid.UUID {
	for i, e := range slice {
		if e == entry {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}

// orderBillsByDate orders bills by date descending
func orderBillsByDate(bills []dto.BillDetailedOutput) []dto.BillDetailedOutput {
	sort.Slice(bills, func(i, j int) bool {
		return bills[i].Date.After(bills[j].Date)
	})
	return bills
}
