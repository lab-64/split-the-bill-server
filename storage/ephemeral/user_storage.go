package ephemeral

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"log"
	"split-the-bill-server/storage"
	"split-the-bill-server/types"
)

type UserStorage struct {
	e *Ephemeral
}

func NewUserStorage(ephemeral *Ephemeral) storage.IUserStorage {
	return &UserStorage{e: ephemeral}
}

func (u *UserStorage) Delete(id uuid.UUID) error {
	u.e.lock.Lock()
	defer u.e.lock.Unlock()
	user, exists := u.e.users[id]
	if !exists {
		return nil
	}
	delete(u.e.nameIndex, user.Username)
	delete(u.e.users, id)
	delete(u.e.passwords, id)
	return nil
}

func (u *UserStorage) GetAll() ([]types.User, error) {
	u.e.lock.Lock()
	defer u.e.lock.Unlock()
	users := make([]types.User, len(u.e.users))
	i := 0
	for _, user := range u.e.users {
		users[i] = user
		i++
	}
	return users, nil
}

func (u *UserStorage) GetByID(id uuid.UUID) (types.User, error) {
	u.e.lock.Lock()
	defer u.e.lock.Unlock()
	user, ok := u.e.users[id]
	if !ok {
		return user, storage.NoSuchUserError
	}
	return user, nil
}

func (u *UserStorage) GetByUsername(username string) (types.User, error) {
	u.e.lock.Lock()
	defer u.e.lock.Unlock()
	id, ok := u.e.nameIndex[username]
	if !ok {
		return types.User{}, storage.NoSuchUserError
	}
	user, ok := u.e.users[id]
	if !ok {
		log.Printf("FATAL error: user storage inconsistent: username '%s' points to non-existent user", username)
		return user, fmt.Errorf("user storage inconsistent: username '%s' points to non-existent user", username)
	}
	return user, nil
}

func (u *UserStorage) Create(user types.User, hash []byte) error {
	u.e.lock.Lock()
	defer u.e.lock.Unlock()

	if _, ok := u.e.nameIndex[user.Username]; ok {
		return storage.UserAlreadyExistsError
	}

	_, ok := u.e.users[user.ID]
	if ok {
		return storage.UserAlreadyExistsError
	}

	u.e.users[user.ID] = user
	u.e.nameIndex[user.Username] = user.ID

	_, exists := u.e.passwords[user.ID]
	if exists {
		return errors.New("fatal: user already has saved password")
	}
	u.e.passwords[user.ID] = hash
	return nil
}

func (u *UserStorage) GetCredentials(id uuid.UUID) ([]byte, error) {
	u.e.lock.Lock()
	defer u.e.lock.Unlock()
	hash, exists := u.e.passwords[id]
	if !exists {
		return nil, storage.NoCredentialsError
	}
	return hash, nil
}

func (u *UserStorage) AddGroupInvitationToUser(invitation types.GroupInvitation, userID uuid.UUID) error {
	u.e.lock.Lock()
	defer u.e.lock.Unlock()
	user, exists := u.e.users[userID]
	if !exists {
		return storage.NoSuchUserError
	}
	// update invitation list
	invitations := append(user.PendingGroupInvitations, &invitation)
	user.PendingGroupInvitations = invitations
	u.e.users[userID] = user
	return nil
}

func (u *UserStorage) HandleInvitation(invitationType string, userID uuid.UUID, invitationID uuid.UUID, accept bool) error {
	u.e.lock.Lock()
	defer u.e.lock.Unlock()
	// get user
	user, exists := u.e.users[userID]
	if !exists {
		return storage.NoSuchUserError
	}
	// handle group invitation reply
	if invitationType == "group" {
		return u.handleGroupInvitation(user, invitationID, accept)
	}
	// TODO: handle further invitation replies
	return storage.InvitationNotFoundError
}

// handleGroupInvitation handles the reply to a group invitation. If the invitation gets accepted, the user gets added to the group and the invitations gets deleted.
// If the invitation gets declined, the invitation gets deleted.
func (u *UserStorage) handleGroupInvitation(user types.User, invitationID uuid.UUID, accept bool) error {
	// if invitation gets accepted, add user to group
	for _, invitation := range user.PendingGroupInvitations {
		if invitation.ID == invitationID {
			if accept {
				// get group
				group, exists := u.e.groups[invitation.For.ID]
				if !exists {
					return storage.NoSuchGroupError
				}
				// insert user into group members
				group.Members = append(group.Members, user)
				u.e.groups[group.ID] = group
				// add group pointer to user struct
				groupList := append(user.Groups, group)
				user.Groups = groupList
			}
			// remove invitation
			user.PendingGroupInvitations = removeInvitation(user.PendingGroupInvitations, invitationID)
			u.e.users[user.ID] = user
			return nil
		}
	}
	return nil
}

// removeInvitation removes the invitation with the given ID from the given invitation list.
func removeInvitation(invitations []*types.GroupInvitation, id uuid.UUID) []*types.GroupInvitation {
	for i, invitation := range invitations {
		if invitation.ID == id {
			return append(invitations[:i], invitations[i+1:]...)
		}
	}
	return invitations
}
