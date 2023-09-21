package test

import (
	"math/rand"
	"split-the-bill-server/storage"
	"split-the-bill-server/types"
	"split-the-bill-server/types/test"
	"testing"

	"github.com/stretchr/testify/require"
)

func addUsers(uut storage.UserStorage, users []types.User, t *testing.T, finished chan<- struct{}) {
	for _, user := range users {
		_, err := uut.AddUser(user)
		require.NoError(t, err)
	}
	close(finished)
}

func getUsers(uut storage.UserStorage, users []types.User, t *testing.T, finished chan<- struct{}) {
	for _, user := range users {
		res, err := uut.GetUserByID(user.ID)
		require.NoError(t, err)
		require.True(t, user.Equals(res))
		res2, err := uut.GetUserByUsername(user.Username)
		require.NoError(t, err)
		require.True(t, user.Equals(res2))
	}
	close(finished)
}

func deleteUsersAndAssert(uut storage.UserStorage, users []types.User, t *testing.T, finished chan<- struct{}) {
	for _, user := range users {
		err := uut.DeleteUser(user.ID)
		require.NoError(t, err)
		_, err = uut.GetUserByID(user.ID)
		require.ErrorIs(t, err, storage.NoSuchUserError)
		_, err = uut.GetUserByUsername(user.Username)
		require.ErrorIs(t, err, storage.NoSuchUserError)
	}
	close(finished)
}

func UserStorageTest(uut storage.UserStorage, t *testing.T) {
	const amount = 10000
	const concurrency = 10
	users := test.GenerateDifferentUsers(amount)
	err := uut.Connect()
	require.NoError(t, err)
	allUsers, err := uut.GetAllUsers()
	require.NoError(t, err)
	require.Equal(t, 0, len(allUsers))
	finished := make([]chan struct{}, concurrency)
	for i := 0; i < concurrency; i++ {
		finished[i] = make(chan struct{})
		go addUsers(uut, users[i*amount/concurrency:(i+1)*amount/concurrency], t, finished[i])
	}
	for i := 0; i < concurrency; i++ {
		<-finished[i]
	}
	allUsers, err = uut.GetAllUsers()
	require.NoError(t, err)
	require.Equal(t, amount, len(allUsers))
	rand.Shuffle(len(users), func(i, j int) { users[i], users[j] = users[j], users[i] })
	for i := 0; i < concurrency; i++ {
		finished[i] = make(chan struct{})
		go getUsers(uut, users[i*amount/concurrency:(i+1)*amount/concurrency], t, finished[i])
	}
	for i := 0; i < concurrency; i++ {
		<-finished[i]
	}
	rand.Shuffle(len(users), func(i, j int) { users[i], users[j] = users[j], users[i] })
	for i := 0; i < concurrency; i++ {
		finished[i] = make(chan struct{})
		go deleteUsersAndAssert(uut, users[i*amount/concurrency:(i+1)*amount/concurrency], t, finished[i])
	}
	for i := 0; i < concurrency; i++ {
		<-finished[i]
	}
	allUsers, err = uut.GetAllUsers()
	require.NoError(t, err)
	require.Equal(t, 0, len(allUsers))
}

func UserStorageEdgeCaseTest(uut storage.UserStorage, t *testing.T) {
	err := uut.Connect()
	require.NoError(t, err)
	users := test.GenerateUsersWithUsernames([]string{"a", "a"})
	err = uut.DeleteUser(users[0].ID)
	require.NoError(t, err)
	_, err = uut.AddUser(users[0])
	require.NoError(t, err)
	res, err := uut.GetUserByUsername("a")
	require.NoError(t, err)
	require.True(t, users[0].Equals(res))
	_, err = uut.AddUser(users[1])
	require.ErrorIs(t, err, storage.UserAlreadyExistsError)
	err = uut.DeleteUser(users[1].ID)
	require.NoError(t, err)
	res, err = uut.GetUserByUsername("a")
	require.NoError(t, err)
	require.True(t, users[0].Equals(res))
}
