package eph_storages

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"log"
	"split-the-bill-server/domain/model"
	"split-the-bill-server/storage"
	"split-the-bill-server/storage/ephemeral"
	"split-the-bill-server/storage/storage_inf"
)

type UserStorage struct {
	e *ephemeral.Ephemeral
}

func NewUserStorage(ephemeral *ephemeral.Ephemeral) storage_inf.IUserStorage {
	return &UserStorage{e: ephemeral}
}

func (u *UserStorage) Delete(id uuid.UUID) error {
	u.e.Lock.Lock()
	defer u.e.Lock.Unlock()
	user, exists := u.e.Users[id]
	if !exists {
		return nil
	}
	delete(u.e.NameIndex, user.Username)
	delete(u.e.Users, id)
	delete(u.e.Passwords, id)
	return nil
}

func (u *UserStorage) GetAll() ([]model.UserModel, error) {
	u.e.Lock.Lock()
	defer u.e.Lock.Unlock()
	users := make([]model.UserModel, len(u.e.Users))
	i := 0
	for _, user := range u.e.Users {
		users[i] = user
		i++
	}
	return users, nil
}

func (u *UserStorage) GetByID(id uuid.UUID) (model.UserModel, error) {
	u.e.Lock.Lock()
	defer u.e.Lock.Unlock()
	user, ok := u.e.Users[id]
	if !ok {
		return user, storage.NoSuchUserError
	}
	return user, nil
}

func (u *UserStorage) GetByUsername(username string) (model.UserModel, error) {
	u.e.Lock.Lock()
	defer u.e.Lock.Unlock()
	id, ok := u.e.NameIndex[username]
	if !ok {
		return model.UserModel{}, storage.NoSuchUserError
	}
	user, ok := u.e.Users[id]
	if !ok {
		log.Printf("FATAL error: user storage inconsistent: username '%s' points to non-existent user", username)
		return user, fmt.Errorf("user storage inconsistent: username '%s' points to non-existent user", username)
	}
	return user, nil
}

func (u *UserStorage) Create(user model.UserModel, hash []byte) error {
	u.e.Lock.Lock()
	defer u.e.Lock.Unlock()

	if _, ok := u.e.NameIndex[user.Username]; ok {
		return storage.UserAlreadyExistsError
	}

	_, ok := u.e.Users[user.ID]
	if ok {
		return storage.UserAlreadyExistsError
	}

	u.e.Users[user.ID] = user
	u.e.NameIndex[user.Username] = user.ID

	_, exists := u.e.Passwords[user.ID]
	if exists {
		return errors.New("fatal: user already has saved password")
	}
	u.e.Passwords[user.ID] = hash
	return nil
}

func (u *UserStorage) GetCredentials(id uuid.UUID) ([]byte, error) {
	u.e.Lock.Lock()
	defer u.e.Lock.Unlock()
	hash, exists := u.e.Passwords[id]
	if !exists {
		return nil, storage.NoCredentialsError
	}
	return hash, nil
}

// TODO: move to invitation storage
func (u *UserStorage) HandleInvitation(invitationType string, userID uuid.UUID, invitationID uuid.UUID, accept bool) error {
	u.e.Lock.Lock()
	defer u.e.Lock.Unlock()
	// get user
	user, exists := u.e.Users[userID]
	if !exists {
		return storage.NoSuchUserError
	}
	// handle group invitation reply
	if invitationType == "group" {
		return u.handleGroupInvitation(user, invitationID, accept)
	}
	// TODO: handle further invitation replies
	return storage.NoSuchGroupInvitationError
}

// handleGroupInvitation handles the reply to a group invitation. If the invitation gets accepted, the user gets added to the group and the invitations gets deleted.
// If the invitation gets declined, the invitation gets deleted.
func (u *UserStorage) handleGroupInvitation(user model.UserModel, invitationID uuid.UUID, accept bool) error {
	// if invitation gets accepted, add user to group
	for _, invitation := range user.PendingGroupInvitations {
		if invitation.ID == invitationID {
			if accept {
				// get group
				group, exists := u.e.Groups[invitation.Group.ID]
				if !exists {
					return storage.NoSuchGroupError
				}
				// insert user into group members
				group.Members = append(group.Members, user.ID)
				u.e.Groups[group.ID] = group
				// add group pointer to user struct
				groupList := append(user.Groups, *group)
				user.Groups = groupList
			}
			// remove invitation
			/*
				user.PendingGroupInvitations = removeInvitation(user.PendingGroupInvitations, invitationID)
			*/
			u.e.Users[user.ID] = user
			return nil
		}
	}
	return nil
}
