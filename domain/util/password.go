package util

import (
	"github.com/caitlinelfring/nist-password-validator/password"
	"golang.org/x/crypto/bcrypt"
	"os"
	"split-the-bill-server/domain"
)

func NewPasswordValidator() (*password.Validator, error) {
	// Setup Password Validation
	var validator = password.NewValidator(false, 8, 64)
	// Load common password list
	var commonPasswords, err = os.Open("common-password-list.txt")
	if err != nil {
		return nil, err
	}
	// Validate Passwords
	err = validator.AddCommonPasswords(commonPasswords)
	if err != nil {
		return nil, err
	}
	return validator, nil
}

func HashPassword(pwd string) ([]byte, error) {
	// Hash Password
	return bcrypt.GenerateFromPassword([]byte(pwd), 10)
}

func ComparePassword(hash []byte, pwd string) error {
	// Compare password with hash
	res := bcrypt.CompareHashAndPassword(hash, []byte(pwd))
	if res != nil {
		return domain.InvalidCredentials
	}
	return nil
}
