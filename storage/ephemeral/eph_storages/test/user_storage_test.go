package test

import (
	"github.com/stretchr/testify/require"
	"math/rand"
	"split-the-bill-server/authentication"
	. "split-the-bill-server/domain/model"
	types_test "split-the-bill-server/domain/model/test"
	"split-the-bill-server/storage"
	. "split-the-bill-server/storage/storage_inf"
	"testing"
)

func addUsers(uut IUserStorage, users []UserModel, t *testing.T, finished chan<- struct{}) {
	for _, user := range users {
		pw, err := authentication.HashPassword("ehhh")
		require.NoError(t, err)
		err = uut.Create(user, pw)
		require.NoError(t, err)
	}
	close(finished)
}

func getUsers(uut IUserStorage, users []UserModel, t *testing.T, finished chan<- struct{}) {
	for _, user := range users {
		res, err := uut.GetByID(user.ID)
		require.NoError(t, err)
		require.True(t, user.Equals(res))
		res2, err := uut.GetByEmail(user.Email)
		require.NoError(t, err)
		require.True(t, user.Equals(res2))
	}
	close(finished)
}

func deleteUsersAndAssert(uut IUserStorage, users []UserModel, t *testing.T, finished chan<- struct{}) {
	for _, user := range users {
		err := uut.Delete(user.ID)
		require.NoError(t, err)
		_, err = uut.GetByID(user.ID)
		require.ErrorIs(t, err, storage.NoSuchUserError)
		_, err = uut.GetByEmail(user.Email)
		require.ErrorIs(t, err, storage.NoSuchUserError)
	}
	close(finished)
}

func UserStorageTest(e storage.Connection, uut IUserStorage, t *testing.T) {
	const amount = 10
	const concurrency = 10
	users := types_test.GenerateDifferentUsers(amount)
	err := e.Connect()
	require.NoError(t, err)
	allUsers, err := uut.GetAll()
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
	allUsers, err = uut.GetAll()
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
	allUsers, err = uut.GetAll()
	require.NoError(t, err)
	require.Equal(t, 0, len(allUsers))
}

func UserStorageEdgeCaseTest(e storage.Connection, uut IUserStorage, t *testing.T) {
	err := e.Connect()
	require.NoError(t, err)
	users := types_test.GenerateUsersWithEmails([]string{"a", "a"})
	err = uut.Delete(users[0].ID)
	require.NoError(t, err)
	pw, err := authentication.HashPassword("ehhh")
	require.NoError(t, err)
	err = uut.Create(users[0], pw)
	require.NoError(t, err)
	res, err := uut.GetByEmail("a")
	require.NoError(t, err)
	require.True(t, users[0].Equals(res))
	pw, err = authentication.HashPassword("ehhh")
	require.NoError(t, err)
	err = uut.Create(users[1], pw)
	require.ErrorIs(t, err, storage.UserAlreadyExistsError)
	err = uut.Delete(users[1].ID)
	require.NoError(t, err)
	res, err = uut.GetByEmail("a")
	require.NoError(t, err)
	require.True(t, users[0].Equals(res))
}
