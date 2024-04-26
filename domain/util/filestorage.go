package util

import (
	"errors"
	"github.com/google/uuid"
	"os"
	"path/filepath"
	"strings"
)

func StoreFile(file []byte, userID uuid.UUID) (string, error) {
	// create file name with random string appended
	randomString := uuid.New().String()
	fileName := userID.String() + "_" + randomString + ".gif"
	// save file
	filePath := filepath.Join("./uploads/profileImages", fileName)

	// Delete existing files with the same userID prefix
	if err := deleteFilesWithPrefix("./uploads/profileImages", userID.String()); err != nil {
		return "", err
	}

	err := os.WriteFile(filePath, file, 0644)
	if err != nil {
		return "", ErrStoreFile
	}
	return filePath, nil
}

func deleteFilesWithPrefix(dirPath string, prefix string) error {
	return filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasPrefix(info.Name(), prefix) {
			return os.Remove(path)
		}
		return nil
	})
}

var ErrStoreFile = errors.New("cannot store file")
