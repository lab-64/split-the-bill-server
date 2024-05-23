package util

import (
	"cloud.google.com/go/storage"
	"context"
	"github.com/google/uuid"
	"io"
	"mime/multipart"
	"os"
)

func StoreFile(file multipart.File, userID uuid.UUID) (string, error) {
	// create file name with random string appended
	randomString := uuid.New().String()
	fileName := userID.String() + "_" + randomString + ".jpg"

	// TODO: Delete existing files with the same userID prefix

	// create new Google Cloud Storage client
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return "", err
	}
	defer client.Close()

	bucketName := os.Getenv("STORAGE_BUCKET_NAME")
	filePath := "https://storage.googleapis.com/" + bucketName + "/" + fileName

	// store file in Google Cloud Storage
	wc := client.Bucket(bucketName).Object(fileName).NewWriter(ctx)
	if _, err = io.Copy(wc, file); err != nil {
		return "", err
	}
	if err = wc.Close(); err != nil {
		return "", err
	}
	return filePath, nil

}
