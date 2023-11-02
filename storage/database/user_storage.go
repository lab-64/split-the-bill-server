package database

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"split-the-bill-server/storage"
	. "split-the-bill-server/storage/database/entity"
	"split-the-bill-server/types"
)

type UserStorage struct {
	DB *gorm.DB
}

func NewUserStorage(DB *Database) storage.IUserStorage {
	return &UserStorage{DB: DB.context}
}

func (u *UserStorage) Delete(id uuid.UUID) error {
	tx := u.DB.Delete(&User{}, "id = ?", id)
	return tx.Error
}

func (u *UserStorage) GetAll() ([]types.User, error) {
	var users []User
	// find all users in the database
	// TODO: GetAllUsers should not return an error, if no users are found
	tx := u.DB.Find(&users)
	if tx.Error != nil {
		return nil, tx.Error
	}

	// return users
	return ToUserSlice(users), nil
}

func (u *UserStorage) GetByID(id uuid.UUID) (types.User, error) {
	var user User
	tx := u.DB.Limit(1).Find(&user, "id = ?", id)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return types.User{}, storage.NoSuchUserError
		}
		return types.User{}, tx.Error
	}
	if tx.RowsAffected == 0 {
		return types.User{}, storage.NoSuchUserError
	}
	return user.ToUser(), nil
}

func (u *UserStorage) GetByUsername(username string) (types.User, error) {
	var user User
	// TODO: Verify that this is in fact not injectable
	tx := u.DB.Take(&user, "username = ?", username)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return types.User{}, storage.NoSuchUserError
		}
		return types.User{}, tx.Error
	}
	if tx.RowsAffected == 0 {
		return types.User{}, storage.NoSuchUserError
	}
	return user.ToUser(), nil
}

func (u *UserStorage) Create(user types.User, passwordHash []byte) error {
	item := MakeUser(user)

	// run as a transaction to ensure consistency. user shouldn't be created if saving credentials failed and vice versa
	err := u.DB.Transaction(func(tx *gorm.DB) error {
		// store user
		res := tx.Create(&item)

		if res.Error != nil {
			return storage.UserAlreadyExistsError
		}

		// store credentials
		res = tx.Create(&Credentials{UserID: item.ID, Hash: passwordHash})
		if res.Error != nil {
			// TODO: create suitable error msg
			return res.Error
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

func (u *UserStorage) AddGroupInvitationToUser(invitation types.GroupInvitation, userID uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}

func (u *UserStorage) HandleInvitation(invitationType string, userID uuid.UUID, invitationID uuid.UUID, accept bool) error {
	//TODO implement me
	panic("implement me")
}
