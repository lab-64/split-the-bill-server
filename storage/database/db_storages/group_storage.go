package db_storages

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"split-the-bill-server/domain/model"
	"split-the-bill-server/storage"
	. "split-the-bill-server/storage/database"
	"split-the-bill-server/storage/database/converter"
	"split-the-bill-server/storage/database/entity"
)

type GroupStorage struct {
	DB *gorm.DB
}

func NewGroupStorage(DB *Database) storage.IGroupStorage {
	return &GroupStorage{DB: DB.Context}
}

func (g *GroupStorage) AddGroup(group model.Group) (model.Group, error) {
	groupItem := converter.ToGroupEntity(group)

	// .First(...) in the end enables preload on create (kind of workaround)
	// https://github.com/go-gorm/gen/issues/618
	res := g.DB.
		Preload(clause.Associations).
		Create(&groupItem).
		First(&groupItem)

	// RowsAffected == 0 -> group already exists
	if res.RowsAffected == 0 {
		return model.Group{}, storage.GroupAlreadyExistsError
	}
	return converter.ToGroupModel(groupItem), res.Error
}

func (g *GroupStorage) UpdateGroup(group model.Group) (model.Group, error) {
	groupEntity := converter.ToGroupEntity(group)

	res := g.DB.
		Preload(clause.Associations).
		Model(&groupEntity).
		Updates(&groupEntity).
		First(&groupEntity)

	// TODO: add finer error handling
	if res.Error != nil {
		return model.Group{}, res.Error
	}

	if res.RowsAffected == 0 {
		return model.Group{}, storage.NoSuchGroupError
	}

	return converter.ToGroupModel(groupEntity), nil
}

func (g *GroupStorage) GetGroupByID(id uuid.UUID) (model.Group, error) {
	var group entity.Group

	// load group with related user and members from db
	tx := g.DB.
		Preload(clause.Associations).
		Preload("Bills.Items.Contributors").
		Preload("Bills.Owner").
		Limit(1).Find(&group, "id = ?", id)

	if tx.Error != nil {
		return model.Group{}, tx.Error
	}
	if tx.RowsAffected == 0 {
		return model.Group{}, storage.NoSuchGroupError
	}
	return converter.ToGroupModel(group), nil
}

func (g *GroupStorage) GetGroups(userID uuid.UUID, invitationID uuid.UUID) ([]model.Group, error) {
	var groups []entity.Group

	tx := g.DB.
		Preload(clause.Associations).
		Preload("Bills.Items.Contributors").
		Preload("Bills.Owner").
		Preload("Bills.UnseenFrom")

	if userID != uuid.Nil {
		tx = tx.Where("id IN (SELECT group_id FROM group_members WHERE user_id = ?)", userID)
	}

	if invitationID != uuid.Nil {
		tx = tx.Where("id IN (SELECT group_id FROM group_invitations WHERE id = ?)", invitationID)
	}

	tx.Find(&groups)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return converter.ToGroupModels(groups), nil
}

func (g *GroupStorage) DeleteGroup(id uuid.UUID) error {
	// Delete entries from the unseen_bills table
	if err := g.DB.Exec("DELETE FROM unseen_bills WHERE bill_id IN (SELECT id FROM bills WHERE group_id = ?)", id).Error; err != nil {
		return err
	}

	// Delete the bills associated with the group
	if err := g.DB.Where("group_id = ?", id).Delete(&entity.Bill{}).Error; err != nil {
		return err
	}

	// Delete the group
	if err := g.DB.Delete(&entity.Group{}, id).Error; err != nil {
		return storage.NoSuchGroupError
	}
	return nil
}

func (g *GroupStorage) AcceptGroupInvitation(invitationID uuid.UUID, userID uuid.UUID) error {
	var groupInvitation entity.GroupInvitation

	//TODO: generalize error messages
	// TODO: test behavior
	// Check if the group invitation exists
	if err := g.DB.First(&groupInvitation, "id = ?", invitationID).Error; err != nil {
		return err
	}

	// add the user to the group members
	group := entity.Group{
		Base: entity.Base{ID: groupInvitation.GroupID},
	}

	user := entity.User{
		Base: entity.Base{ID: userID},
	}

	res := g.DB.Model(&group).Association("Members").Append(&user)

	return res
}

func (g *GroupStorage) CreateGroupTransaction(transaction model.GroupTransaction) (model.GroupTransaction, error) {
	groupTransactionEntity := converter.ToGroupTransactionEntity(transaction)

	err := g.DB.Transaction(func(tx *gorm.DB) error {
		// delete all bills and items associated with the group
		if err := tx.Exec("DELETE FROM items WHERE bill_id IN (SELECT id FROM bills WHERE group_id = ?)", transaction.GroupID).Error; err != nil {
			return err
		}

		// add new group transaction
		res := tx.
			Preload(clause.Associations).
			Preload("Transactions.Debtor").
			Preload("Transactions.Creditor").
			Create(&groupTransactionEntity).
			First(&groupTransactionEntity)

		return res.Error
	})

	return converter.ToGroupTransactionModel(groupTransactionEntity), err
}

func (g *GroupStorage) GetAllGroupTransactions(userID uuid.UUID) ([]model.GroupTransaction, error) {
	var groupTransactions []entity.GroupTransaction

	tx := g.DB.
		Preload(clause.Associations).
		Preload("Transactions.Debtor").
		Preload("Transactions.Creditor").
		Where("group_id IN (SELECT group_id FROM group_members WHERE user_id = ?)", userID).
		Find(&groupTransactions)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return converter.ToGroupTransactionModels(groupTransactions), nil
}
