package db_storages

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"log"
	"split-the-bill-server/domain/model"
	"split-the-bill-server/storage"
	. "split-the-bill-server/storage/database"
	"split-the-bill-server/storage/database/converter"
	"split-the-bill-server/storage/database/entity"
)

type UserStorage struct {
	DB *gorm.DB
}

func NewUserStorage(DB *Database) storage.IUserStorage {
	return &UserStorage{DB: DB.Context}
}

func (u *UserStorage) Delete(id uuid.UUID) error {
	tx := u.DB.Delete(&entity.User{}, "id = ?", id)
	return tx.Error
}

func (u *UserStorage) GetAll() ([]model.User, error) {
	var users []entity.User
	// find all users in the database
	// TODO: GetAllUsers should not return an error, if no users are found
	tx := u.DB.
		Find(&users)
	if tx.Error != nil {
		return nil, tx.Error
	}

	// return users
	return converter.ToUserModels(users), nil
}

func (u *UserStorage) GetByID(id uuid.UUID) (model.User, error) {
	var user entity.User
	tx := u.DB.Limit(1).
		Find(&user, "id = ?", id)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return model.User{}, storage.NoSuchUserError
		}
		return model.User{}, tx.Error
	}
	if tx.RowsAffected == 0 {
		return model.User{}, storage.NoSuchUserError
	}
	return converter.ToUserModel(user), nil
}

func (u *UserStorage) GetByEmail(email string) (model.User, error) {
	var user entity.User
	// TODO: Verify that this is in fact not injectable
	tx := u.DB.Take(&user, "email = ?", email)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return model.User{}, storage.NoSuchUserError
		}
		return model.User{}, tx.Error
	}
	if tx.RowsAffected == 0 {
		return model.User{}, storage.NoSuchUserError
	}
	return converter.ToUserModel(user), nil
}

func (u *UserStorage) Create(user model.User, passwordHash []byte) (model.User, error) {
	item := converter.ToUserEntity(user)

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
		res = tx.Create(&entity.Credentials{UserID: item.ID, Hash: passwordHash})
		if res.Error != nil {
			return storage.InvalidUserInputError
		}

		return nil
	})

	return converter.ToUserModel(item), err
}

func (u *UserStorage) Update(user model.User) (model.User, error) {
	userEntity := entity.User{}

	res := u.DB.Model(&entity.User{}).Where("id = ?", user.ID).Updates(entity.User{Username: user.Username, ProfileImgPath: user.ProfileImgPath}).First(&userEntity)
	// TODO: error handling
	log.Println(" Storage | Updated User profile image path: ", userEntity.ProfileImgPath)
	return converter.ToUserModel(userEntity), res.Error
}

func (u *UserStorage) GetCredentials(id uuid.UUID) ([]byte, error) {
	var credentials entity.Credentials
	// get credentials from given user
	res := u.DB.Limit(1).First(&credentials, entity.Credentials{UserID: id})
	if res.Error != nil {
		return nil, storage.NoCredentialsError
	}
	return credentials.Hash, nil
}
