package db_storages

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	. "split-the-bill-server/domain/model"
	"split-the-bill-server/storage"
	. "split-the-bill-server/storage/database"
	. "split-the-bill-server/storage/database/entity"
	. "split-the-bill-server/storage/storage_inf"
)

type UserStorage struct {
	DB *gorm.DB
}

func NewUserStorage(DB *Database) IUserStorage {
	return &UserStorage{DB: DB.Context}
}

func (u *UserStorage) Delete(id uuid.UUID) error {
	tx := u.DB.Delete(&User{}, "id = ?", id)
	return tx.Error
}

func (u *UserStorage) GetAll() ([]UserModel, error) {
	var users []User
	// find all users in the database
	// TODO: GetAllUsers should not return an error, if no users are found
	tx := u.DB.Preload("Groups.Members").Preload("GroupInvitations").Find(&users)
	if tx.Error != nil {
		return nil, tx.Error
	}

	// return users
	return ToUserModelSlice(users), nil
}

func (u *UserStorage) GetByID(id uuid.UUID) (UserModel, error) {
	var user User
	tx := u.DB.Limit(1).Preload("Groups.Members").Preload("GroupInvitations").Find(&user, "id = ?", id)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return UserModel{}, storage.NoSuchUserError
		}
		return UserModel{}, tx.Error
	}
	if tx.RowsAffected == 0 {
		return UserModel{}, storage.NoSuchUserError
	}
	return ToUserModel(user), nil
}

func (u *UserStorage) GetByEmail(email string) (UserModel, error) {
	var user User
	// TODO: Verify that this is in fact not injectable
	tx := u.DB.Take(&user, "email = ?", email)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return UserModel{}, storage.NoSuchUserError
		}
		return UserModel{}, tx.Error
	}
	if tx.RowsAffected == 0 {
		return UserModel{}, storage.NoSuchUserError
	}
	return ToUserModel(user), nil
}

func (u *UserStorage) Create(user UserModel, passwordHash []byte) error {
	item := ToUserEntity(user)

	// run as a transaction to ensure consistency. user shouldn't be created if saving credentials failed and vice versa
	err := u.DB.Transaction(func(tx *gorm.DB) error {
		// store user
		res := tx.Create(&item)

		if res.Error != nil {
			if errors.Is(res.Error, gorm.ErrDuplicatedKey) {
				return storage.UserAlreadyExistsError
			}
			return storage.InvalidUserInputError
		}

		// store credentials
		res = tx.Create(&Credentials{UserID: item.ID, Hash: passwordHash})
		if res.Error != nil {
			return storage.InvalidUserInputError
		}

		return nil
	})

	return err
}

func (u *UserStorage) GetCredentials(id uuid.UUID) ([]byte, error) {
	var credentials Credentials
	// get credentials from given user
	res := u.DB.Limit(1).First(&credentials, Credentials{UserID: id})
	if res.Error != nil {
		return nil, storage.NoCredentialsError
	}
	return credentials.Hash, nil
}
