package util

import (
	"cloud.google.com/go/storage"
	"context"
	"errors"
	"github.com/google/uuid"
	"io"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
)

func StoreFileInGoogleCloudStorage(file multipart.File, fileName string) (string, error) {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return "", err
	}
	defer client.Close()

	bucketName := "stb-profile-imgs"

	wc := client.Bucket(bucketName).Object(fileName).NewWriter(ctx)
	if _, err = io.Copy(wc, file); err != nil {
		return "", err
	}
	if err := wc.Close(); err != nil {
		return "", err
	}
	log.Println("File uploaded successfully")
	return fileName, nil
}

func StoreFile(file multipart.File, userID uuid.UUID) (string, error) {
	// create file name with random string appended
	randomString := uuid.New().String()
	fileName := userID.String() + "_" + randomString + ".gif"

	// TODO: check
	// Delete existing files with the same userID prefix
	/*if err := deleteFilesWithPrefix("./uploads/profileImages", userID.String()); err != nil {
		return "", err
	}
	*/

	// create new Google Cloud Storage client
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return "", err
	}
	defer client.Close()

	bucketName := "stb-profile-imgs"
	filePath := "https://storage.googleapis.com/" + bucketName + "/" + fileName

	// store file in Google Cloud Storage
	wc := client.Bucket(bucketName).Object(fileName).NewWriter(ctx)
	if _, err = io.Copy(wc, file); err != nil {
		return "", err
	}
	if err = wc.Close(); err != nil {
		return "", err
	}
	log.Println("File uploaded successfully")
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
