package util

import (
	"errors"
	"github.com/google/uuid"
	"os"
	"path/filepath"
)

func StoreFile(file []byte, userID uuid.UUID) (string, error) {
	// TODO: we can create the folder by hand and leave this out here
	storagePath := "./uploads/profileImages"
	// create storage directory for id
	if err := os.MkdirAll(storagePath, os.ModePerm); err != nil {
		return "", ErrStoreFile
	}
	// TODO: is the image type important?
	fileName := userID.String() + ".png"
	// save file
	filePath := filepath.Join(storagePath, fileName)
	err := os.WriteFile(filePath, file, 0644)
	if err != nil {
		return "", ErrStoreFile
	}
	return filePath, nil
}

var ErrStoreFile = errors.New("cannot store file")
