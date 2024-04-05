package types_test

import (
	"github.com/google/uuid"
	"split-the-bill-server/domain/model"
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
func GenerateDifferentUsers(amount int) []model.User {
	// TODO: generate emails instead of usernames
	usernames := GenerateUsernames(amount, chars)
	return GenerateUsersWithEmails(usernames)
}

func GenerateUsersWithEmails(emails []string) []model.User {
	users := make([]model.User, len(emails))
	for i := 0; i < len(emails); i++ {
		users[i] = model.User{
			ID:    uuid.New(),
			Email: emails[i]}
	}
	return users
}
