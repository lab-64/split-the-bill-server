package converter

import (
	"split-the-bill-server/domain/model"
	"split-the-bill-server/storage/database/entity"
)

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// AUTH COOKIE
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func ToAuthCookieEntity(authCookie model.AuthCookie) entity.AuthCookie {
	return entity.AuthCookie{
		Base:        entity.Base{ID: authCookie.Token},
		UserID:      authCookie.UserID,
		ValidBefore: authCookie.ValidBefore,
	}
}

func ToAuthCookieModel(authCookie *entity.AuthCookie) model.AuthCookie {
	return model.AuthCookie{
		UserID:      authCookie.UserID,
		Token:       authCookie.ID,
		ValidBefore: authCookie.ValidBefore,
	}
}

func ToAuthCookieModels(authCookies []entity.AuthCookie) []model.AuthCookie {
	s := make([]model.AuthCookie, len(authCookies))
	for i, cookie := range authCookies {
		s[i] = ToAuthCookieModel(&cookie)
	}
	return s
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// BILL
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func ToBillEntity(bill model.Bill) entity.Bill {

	// convert items
	var items []entity.Item
	for _, item := range bill.Items {
		items = append(items, ToItemEntity(item))
	}

	return entity.Bill{
		Base:    entity.Base{ID: bill.ID},
		Name:    bill.Name,
		Date:    bill.Date,
		GroupID: bill.GroupID,
		OwnerID: bill.Owner.ID,
		Items:   items,
	}
}

func ToBillModel(bill entity.Bill) model.Bill {
	items := make([]model.Item, len(bill.Items))
	for i, item := range bill.Items {
		items[i] = ToItemModel(item)
	}

	return model.Bill{
		ID:      bill.ID,
		Name:    bill.Name,
		Date:    bill.Date,
		Owner:   ToUserModel(bill.Owner),
		GroupID: bill.GroupID,
		Items:   items,
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// GROUP
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func ToGroupEntity(group model.Group) entity.Group {
	// convert uuids to users
	var members []*entity.User
	for _, member := range group.Members {
		members = append(members, &entity.User{Base: entity.Base{ID: member.ID}})
	}
	return entity.Group{
		Base:            entity.Base{ID: group.ID},
		OwnerUID:        group.Owner.ID,
		Name:            group.Name,
		Members:         members,
		GroupInvitation: entity.GroupInvitation{Base: entity.Base{ID: group.InvitationID}},
	}
}

func ToGroupModel(group entity.Group) model.Group {
	members := make([]model.User, len(group.Members))
	for i, member := range group.Members {
		members[i] = ToUserModel(*member)
	}
	bills := make([]model.Bill, len(group.Bills))
	for i, bill := range group.Bills {
		bills[i] = ToBillModel(bill)
	}

	return model.Group{
		ID:           group.ID,
		Name:         group.Name,
		Owner:        ToUserModel(group.Owner),
		Members:      members,
		Bills:        bills,
		InvitationID: group.GroupInvitation.ID,
	}
}

func ToGroupModels(groups []entity.Group) []model.Group {
	s := make([]model.Group, len(groups))
	for i, group := range groups {
		s[i] = ToGroupModel(group)
	}
	return s
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// ITEM
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func ToItemEntity(item model.Item) entity.Item {
	// create user entities with the given ids for the contributors
	var contributors []*entity.User
	for _, contributor := range item.Contributors {
		contributors = append(contributors, &entity.User{Base: entity.Base{ID: contributor.ID}})
	}
	return entity.Item{
		Base:         entity.Base{ID: item.ID},
		Name:         item.Name,
		Price:        item.Price,
		BillID:       item.BillID,
		Contributors: contributors,
	}
}

func ToItemModel(item entity.Item) model.Item {
	contributors := make([]model.User, len(item.Contributors))
	for i, contributor := range item.Contributors {
		contributors[i] = ToUserModel(*contributor)
	}
	return model.Item{
		ID:           item.ID,
		Name:         item.Name,
		Price:        item.Price,
		BillID:       item.BillID,
		Contributors: contributors,
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// USER
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func ToUserEntity(user model.User) entity.User {
	return entity.User{
		Base:           entity.Base{ID: user.ID},
		Email:          user.Email,
		Username:       user.Username,
		ProfileImgPath: user.ProfileImgPath,
	}
}

func ToUserModel(user entity.User) model.User {
	return model.User{
		ID:             user.ID,
		Email:          user.Email,
		Username:       user.Username,
		ProfileImgPath: user.ProfileImgPath,
	}
}

func ToUserModels(users []entity.User) []model.User {
	s := make([]model.User, len(users))
	for i, user := range users {
		s[i] = ToUserModel(user)
	}
	return s
}
