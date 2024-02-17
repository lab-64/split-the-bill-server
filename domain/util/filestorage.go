package util

import (
	"errors"
	"github.com/google/uuid"
	"os"
	"path/filepath"
	"split-the-bill-server/presentation/router"
)

func StoreFile(file []byte, userID uuid.UUID) (string, error) {
	// create file name
	fileName := userID.String() + ".png"
	// save file
	filePath := filepath.Join(router.StorePath, fileName)
	err := os.WriteFile(filePath, file, 0644)
	if err != nil {
		return "", ErrStoreFile
	}
	return filePath, nil
}

var ErrStoreFile = errors.New("cannot store file")
