package util

import (
	"errors"
	"github.com/google/uuid"
	"os"
	"path/filepath"
)

func StoreFile(file []byte, userID uuid.UUID) (string, error) {
	// create file name
	fileName := userID.String() + ".gif"
	// save file
	filePath := filepath.Join("./uploads/profileImages", fileName)
	err := os.WriteFile(filePath, file, 0644)
	if err != nil {
		return "", ErrStoreFile
	}
	return filePath, nil
}

var ErrStoreFile = errors.New("cannot store file")
