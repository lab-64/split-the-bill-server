package converter

import (
	"split-the-bill-server/domain/model"
	"split-the-bill-server/presentation/dto"
)

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// BILL
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func ToBillDetailedDTOs(bills []model.Bill) []dto.BillDetailedOutput {
	billsDTO := make([]dto.BillDetailedOutput, len(bills))

	for i, bill := range bills {
		billsDTO[i] = ToBillDetailedDTO(bill)
	}
	return billsDTO
}

func ToBillDetailedDTO(bill model.Bill) dto.BillDetailedOutput {
	itemsDTO := make([]dto.ItemOutput, len(bill.Items))

	for i, item := range bill.Items {
		itemsDTO[i] = ToItemDTO(item)
	}

	return dto.BillDetailedOutput{
		ID:      bill.ID,
		Name:    bill.Name,
		Date:    bill.Date,
		Items:   itemsDTO,
		Owner:   ToUserCoreDTO(&bill.Owner),
		GroupID: bill.GroupID,
		Balance: bill.Balance,
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// ITEM
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func ToItemDTO(item model.Item) dto.ItemOutput {
	contributors := make([]dto.UserCoreOutput, len(item.Contributors))
	for i, cont := range item.Contributors {
		contributors[i] = ToUserCoreDTO(&cont)
	}

	return dto.ItemOutput{
		ID:           item.ID,
		Name:         item.Name,
		Price:        item.Price,
		BillID:       item.BillID,
		Contributors: contributors,
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// GROUP
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func ToGroupDetailedDTO(g model.Group) dto.GroupDetailedOutput {

	billsDTO := ToBillDetailedDTOs(g.Bills)
	owner := ToUserCoreDTO(&g.Owner)
	members := ToUserCoreDTOs(g.Members)
	return dto.GroupDetailedOutput{
		Owner:        owner,
		ID:           g.ID,
		Name:         g.Name,
		Members:      members,
		Bills:        billsDTO,
		Balance:      g.Balance,
		InvitationID: g.InvitationID,
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// USER
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func ToUserCoreDTO(u *model.User) dto.UserCoreOutput {
	return dto.UserCoreOutput{
		ID:       u.ID,
		Email:    u.Email,
		Username: u.Username,
	}
}

func ToUserCoreDTOs(users []model.User) []dto.UserCoreOutput {
	usersDTO := make([]dto.UserCoreOutput, len(users))

	for i, user := range users {
		usersDTO[i] = ToUserCoreDTO(&user)
	}
	return usersDTO
}
