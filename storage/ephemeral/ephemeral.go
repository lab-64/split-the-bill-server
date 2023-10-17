package ephemeral

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"log"
	"split-the-bill-server/storage"
	"split-the-bill-server/types"
	"sync"
)

type Ephemeral struct {
	lock            sync.Mutex
	userStorage     map[uuid.UUID]types.User
	nameIndex       map[string]uuid.UUID
	passwordStorage map[uuid.UUID][]byte
	cookieStorage   map[uuid.UUID][]types.AuthenticationCookie
	groupStorage    map[uuid.UUID]types.Group
}

func NewEphemeral() *Ephemeral {
	return &Ephemeral{
		userStorage:     make(map[uuid.UUID]types.User),
		nameIndex:       make(map[string]uuid.UUID),
		passwordStorage: make(map[uuid.UUID][]byte),
		cookieStorage:   make(map[uuid.UUID][]types.AuthenticationCookie),
		groupStorage:    make(map[uuid.UUID]types.Group),
	}
}

func (e *Ephemeral) Connect() error {
	return nil
}

// User Section

func (e *Ephemeral) RegisterUser(user types.User, hash []byte) error {
	err := e.AddUser(user)
	if err != nil {
		return err
	}
	e.lock.Lock()
	defer e.lock.Unlock()
	_, exists := e.passwordStorage[user.ID]
	if exists {
		return errors.New("fatal: user already has saved password")
	}
	e.passwordStorage[user.ID] = hash
	return nil
}

func (e *Ephemeral) AddUser(user types.User) error {
	e.lock.Lock()
	defer e.lock.Unlock()
	if _, ok := e.nameIndex[user.Username]; ok {
		return storage.UserAlreadyExistsError
	}
	_, ok := e.userStorage[user.ID]
	if ok {
		return storage.UserAlreadyExistsError
	}
	e.userStorage[user.ID] = user
	e.nameIndex[user.Username] = user.ID
	return nil
}

func (e *Ephemeral) DeleteUser(id uuid.UUID) error {
	e.lock.Lock()
	defer e.lock.Unlock()
	user, exists := e.userStorage[id]
	if !exists {
		return nil
	}
	delete(e.nameIndex, user.Username)
	delete(e.userStorage, id)
	delete(e.passwordStorage, id)
	return nil
}

func (e *Ephemeral) GetAllUsers() ([]types.User, error) {
	e.lock.Lock()
	defer e.lock.Unlock()
	users := make([]types.User, len(e.userStorage))
	i := 0
	for _, user := range e.userStorage {
		users[i] = user
		i++
	}
	return users, nil
}

func (e *Ephemeral) GetUserByID(id uuid.UUID) (types.User, error) {
	e.lock.Lock()
	defer e.lock.Unlock()
	user, ok := e.userStorage[id]
	if !ok {
		return user, storage.NoSuchUserError
	}
	return user, nil
}

func (e *Ephemeral) GetUserByUsername(username string) (types.User, error) {
	e.lock.Lock()
	defer e.lock.Unlock()
	id, ok := e.nameIndex[username]
	if !ok {
		return types.User{}, storage.NoSuchUserError
	}
	user, ok := e.userStorage[id]
	if !ok {
		log.Printf("FATAL error: user storage inconsistent: username '%s' points to non-existent user", username)
		return user, fmt.Errorf("user storage inconsistent: username '%s' points to non-existent user", username)
	}
	return user, nil
}

func (e *Ephemeral) GetCredentials(id uuid.UUID) ([]byte, error) {
	e.lock.Lock()
	defer e.lock.Unlock()
	hash, exists := e.passwordStorage[id]
	if !exists {
		return nil, storage.NoCredentialsError
	}
	return hash, nil
}

func (e *Ephemeral) AddGroupInvitationToUser(invitation types.GroupInvitation, userId uuid.UUID) error {
	e.lock.Lock()
	defer e.lock.Unlock()
	user, exists := e.userStorage[userId]
	if !exists {
		return storage.NoSuchUserError
	}
	// update invitation list
	invitations := append(user.PendingGroupInvitations, invitation)
	user.PendingGroupInvitations = invitations
	e.userStorage[userId] = user
	return nil
}

func (e *Ephemeral) HandleInvitation(invitationType string, userID uuid.UUID, invitationID uuid.UUID, accept bool) error {
	e.lock.Lock()
	defer e.lock.Unlock()
	// get user
	user, exists := e.userStorage[userID]
	if !exists {
		return storage.NoSuchUserError
	}
	// handle group invitation reply
	if invitationType == "group" {
		return e.handleGroupInvitation(user, invitationID, accept)
	}
	// TODO: handle further invitation replies
	return storage.InvitationNotFoundError
}

// handleGroupInvitation handles the reply to a group invitation. If the invitation gets accepted, the user gets added to the group and the invitations gets deleted.
// If the invitation gets declined, the invitation gets deleted.
func (e *Ephemeral) handleGroupInvitation(user types.User, invitationID uuid.UUID, accept bool) error {
	// if invitation gets accepted, add user to group
	for _, invitation := range user.PendingGroupInvitations {
		if invitation.ID == invitationID {
			if accept {
				// get group
				group, exists := e.groupStorage[invitation.For.ID]
				if !exists {
					return storage.NoSuchGroupError
				}
				// insert user into group members
				group.Members = append(group.Members, user.ID)
				e.groupStorage[group.ID] = group
				// add group pointer to user struct
				groupList := append(*user.Groups, group)
				user.Groups = &groupList
			}
			// remove invitation
			user.PendingGroupInvitations = removeInvitation(user.PendingGroupInvitations, invitationID)
			e.userStorage[user.ID] = user
			return nil
		}
	}
	return nil
}

// removeInvitation removes the invitation with the given ID from the given invitation list.
func removeInvitation(invitations []types.GroupInvitation, id uuid.UUID) []types.GroupInvitation {
	for i, invitation := range invitations {
		if invitation.ID == id {
			return append(invitations[:i], invitations[i+1:]...)
		}
	}
	return invitations
}

// Cookie Section

func (e *Ephemeral) AddAuthenticationCookie(cookie types.AuthenticationCookie) {
	e.lock.Lock()
	defer e.lock.Unlock()
	cookies, exists := e.cookieStorage[cookie.UserID]
	if !exists {
		cookies = make([]types.AuthenticationCookie, 0)
	}
	cookies = append(cookies, cookie)
	e.cookieStorage[cookie.UserID] = cookies
}

func (e *Ephemeral) GetCookiesForUser(userID uuid.UUID) []types.AuthenticationCookie {
	e.lock.Lock()
	defer e.lock.Unlock()
	return e.cookieStorage[userID]
}

func (e *Ephemeral) GetCookieFromToken(token uuid.UUID) (types.AuthenticationCookie, error) {
	e.lock.Lock()
	defer e.lock.Unlock()
	for _, cookies := range e.cookieStorage {
		for _, cookie := range cookies {
			if cookie.Token == token {
				return cookie, nil
			}
		}
	}
	return types.AuthenticationCookie{}, storage.NoSuchCookieError
}

// Group Section

func (e *Ephemeral) AddGroup(group types.Group) error {
	e.lock.Lock()
	defer e.lock.Unlock()
	_, exists := e.groupStorage[group.ID]
	if exists {
		return storage.GroupAlreadyExistsError
	}
	e.groupStorage[group.ID] = group
	return nil
}

func (e *Ephemeral) GetGroupByID(id uuid.UUID) (types.Group, error) {
	e.lock.Lock()
	defer e.lock.Unlock()
	group, exists := e.groupStorage[id]
	if !exists {
		return group, storage.NoSuchGroupError
	}
	return group, nil
}

func (e *Ephemeral) AddMemberToGroup(memberID uuid.UUID, groupID uuid.UUID) error {
	e.lock.Lock()
	defer e.lock.Unlock()
	group, exists := e.groupStorage[groupID]
	if !exists {
		return storage.NoSuchGroupError
	}
	group.Members = append(group.Members, memberID)
	e.groupStorage[groupID] = group
	return nil
}
