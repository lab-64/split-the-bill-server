package db_storages

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"log"
	"split-the-bill-server/domain/model"
	"split-the-bill-server/storage"
	"split-the-bill-server/storage/database"
	. "split-the-bill-server/storage/database/entity"
	"split-the-bill-server/storage/storage_inf"
)

type UserStorage struct {
	DB *gorm.DB
}

func NewUserStorage(DB *database.Database) storage_inf.IUserStorage {
	return &UserStorage{DB: DB.Context}
}

func (u *UserStorage) Create(user model.User) error {
	item := MakeUser(user)
	// FIXME: This is a little bit of TOCTOU:
	// If a user with the same username is created after the check, we DO NOT RETURN AN ERROR.
	// We also do not overwrite the existing user.
	// Checking, if the username already exists is still better, as we will receive an error at least in most cased,
	// where there are no unlikely race condition.
	// This could be fixed, if there was some way to check, whether FirstOrCreate actually created a new user or not.
	_, err := u.GetByUsername(user.Username) // check if user already exists
	if err == nil {
		return storage.UserAlreadyExistsError
	}
	res := u.DB.Where(User{Username: user.Username}).FirstOrCreate(&item) // write new user if not exists
	return res.Error
}

func (u *UserStorage) Delete(id uuid.UUID) error {
	tx := u.DB.Delete(&User{}, "id = ?", id)
	return tx.Error
}

func (u *UserStorage) GetAll() ([]model.User, error) {
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

// TODO: include pending invitations & groups in query
func (u *UserStorage) GetByID(id uuid.UUID) (model.User, error) {
	var user User
	tx := u.DB.Limit(1).Preload("Groups").Preload("GroupInvitations.For").Find(&user, "id = ?", id) // lade noch alle groups mit FK = user &
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return model.User{}, storage.NoSuchUserError
		}
		return model.User{}, tx.Error
	}
	if tx.RowsAffected == 0 {
		return model.User{}, storage.NoSuchUserError
	}
	return user.ToUser(), nil
}

func (u *UserStorage) GetByUsername(username string) (model.User, error) {
	var user User
	// TODO: Verify that this is in fact not injectable
	tx := u.DB.Take(&user, "username = ?", username)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return model.User{}, storage.NoSuchUserError
		}
		return model.User{}, tx.Error
	}
	if tx.RowsAffected == 0 {
		return model.User{}, storage.NoSuchUserError
	}
	return user.ToUser(), nil
}

func (u *UserStorage) Register(user model.User, passwordHash []byte) error {
	item := MakeUser(user)
	// store user
	res := u.DB.Where(User{Username: item.Username}).FirstOrCreate(&item)
	log.Println(res)
	if res.Error != nil {
		return storage.UserAlreadyExistsError
	}
	// TODO: handle error case, user should not be created, if credentials cannot be stored
	// store credentials
	res = u.DB.Where(Credentials{UserID: item.ID}).FirstOrCreate(&Credentials{UserID: item.ID, Hash: passwordHash})
	if res.Error != nil {
		// TODO: create suitable error msg
		return res.Error
	}
	return nil
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

func (u *UserStorage) AddGroupInvitation(invitation model.GroupInvitation, userID uuid.UUID) error {
	// make group invitation entity
	groupInvitationItem := MakeGroupInvitation(invitation)
	// TODO: update user, maybe it gets updated automatically
	res := u.DB.Model(&User{Base: Base{ID: userID}}).Association("PendingGroupInvitation").Append(&groupInvitationItem)
	log.Println("AddGroupInvitation: ", res)
	return res
}

func (u *UserStorage) HandleInvitation(invitationType string, userID uuid.UUID, invitationID uuid.UUID, accept bool) error {
	//TODO implement me
	panic("implement me")
}
