package test

import (
	"split-the-bill-server/types"
)

var chars = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func GenerateUsernames(amount int, chars []rune) []string {
	usernames := make([]string, amount)
	mod := len(chars)
	for i := 0; i < amount; i++ {
		username := make([]rune, i/mod+1)
		for j := 0; j < len(username); j++ {
			username[j] = chars[i%mod]
		}
		usernames[i] = string(username)
	}
	return usernames
}

// GenerateDifferentUsers generates a slice of users with different usernames. The IDs are generated randomly, but
// collisions are highly unlikely.
func GenerateDifferentUsers(amount int) []types.User {
	usernames := GenerateUsernames(amount, chars)
	return GenerateUsersWithUsernames(usernames)
}

func GenerateUsersWithUsernames(usernames []string) []types.User {
	users := make([]types.User, len(usernames))
	for i := 0; i < len(usernames); i++ {
		users[i] = types.NewUser(usernames[i])
	}
	return users
}
