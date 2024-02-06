package eph_storages

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"log"
	"split-the-bill-server/domain/model"
	"split-the-bill-server/storage"
	eph "split-the-bill-server/storage/ephemeral"
)

type UserStorage struct {
	e *eph.Ephemeral
}

func NewUserStorage(ephemeral *eph.Ephemeral) storage.IUserStorage {
	return &UserStorage{e: ephemeral}
}

func (u *UserStorage) Delete(id uuid.UUID) error {
	r := u.e.Locker.Lock(eph.RUsers, eph.RNameIndex, eph.RPasswords)
	defer u.e.Locker.Unlock(r)
	user, exists := u.e.Users[id]
	if !exists {
		return nil
	}
	delete(u.e.NameIndex, user.Email)
	delete(u.e.Users, id)
	delete(u.e.Passwords, id)
	return nil
}

func (u *UserStorage) GetAll() ([]model.UserModel, error) {
	r := u.e.Locker.Lock(eph.RUsers)
	defer u.e.Locker.Unlock(r)
	users := make([]model.UserModel, len(u.e.Users))
	i := 0
	for _, user := range u.e.Users {
		users[i] = user
		i++
	}
	return users, nil
}

func (u *UserStorage) GetByID(id uuid.UUID) (model.UserModel, error) {
	r := u.e.Locker.Lock(eph.RUsers)
	defer u.e.Locker.Unlock(r)
	user, ok := u.e.Users[id]
	if !ok {
		return user, storage.NoSuchUserError
	}
	return user, nil
}

func (u *UserStorage) GetByEmail(email string) (model.UserModel, error) {
	r := u.e.Locker.Lock(eph.RUsers, eph.RNameIndex)
	defer u.e.Locker.Unlock(r)
	id, ok := u.e.NameIndex[email]
	if !ok {
		return model.UserModel{}, storage.NoSuchUserError
	}
	user, ok := u.e.Users[id]
	if !ok {
		log.Printf("FATAL error: user storage inconsistent: email '%s' points to non-existent user", email)
		return user, fmt.Errorf("user storage inconsistent: email '%s' points to non-existent user", email)
	}
	return user, nil
}

func (u *UserStorage) Create(user model.UserModel, hash []byte) (model.UserModel, error) {
	r := u.e.Locker.Lock(eph.RUsers, eph.RNameIndex, eph.RPasswords)
	defer u.e.Locker.Unlock(r)

	if _, ok := u.e.NameIndex[user.Email]; ok {
		return model.UserModel{}, storage.UserAlreadyExistsError
	}

	_, ok := u.e.Users[user.ID]
	if ok {
		return model.UserModel{}, storage.UserAlreadyExistsError
	}

	u.e.Users[user.ID] = user
	u.e.NameIndex[user.Email] = user.ID

	_, exists := u.e.Passwords[user.ID]
	if exists {
		return model.UserModel{}, errors.New("fatal: user already has saved password")
	}
	u.e.Passwords[user.ID] = hash
	return user, nil
}

func (u *UserStorage) GetCredentials(id uuid.UUID) ([]byte, error) {
	r := u.e.Locker.Lock(eph.RPasswords)
	defer u.e.Locker.Unlock(r)
	hash, exists := u.e.Passwords[id]
	if !exists {
		return nil, storage.NoCredentialsError
	}
	return hash, nil
}

func (u *UserStorage) Update(user model.UserModel) (model.UserModel, error) {
	//TODO implement me
	panic("implement me")
}
