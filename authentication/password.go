package authentication

import (
	"errors"
	"github.com/caitlinelfring/nist-password-validator/password"
	"golang.org/x/crypto/bcrypt"
	"os"
	"path/filepath"
	"runtime"
)

func NewPasswordValidator() (*password.Validator, error) {
	// Setup Password Validation
	var validator = password.NewValidator(false, 8, 64)
	// Generate dynamically the path to the common-password-list.txt
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		// TODO: change error message
		return nil, errors.New("cannot get current file path")
	}
	dir := filepath.Dir(filename)
	path := filepath.Join(dir, "common-password-list.txt")
	// Load common password list
	commonPasswords, err := os.Open(path)
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
		return InvalidCredentials
	}
	return nil
}

var InvalidCredentials = errors.New("invalid credentials")
