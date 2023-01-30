package filesystem

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"

	"github.com/google/uuid"
)

func SaveFileToStorage(file multipart.File, fileHeader *multipart.FileHeader) (*string, *string, error) {

	ext1 := filepath.Ext(fileHeader.Filename)
	new_generated_name := uuid.New().String() + ext1

	date := time.Now()
	save_path := fmt.Sprintf("./tmp/%d/%d/%d/%s", date.Year(), date.Month(), date.Day(), new_generated_name)

	if err := os.MkdirAll(filepath.Dir(save_path), 0770); err != nil {
		return nil, nil, err
	}

	f, err := os.Create(save_path)
	if err != nil {
		return nil, nil, err
	}

	_, err = io.Copy(f, file)
	if err != nil {
		return nil, nil, err
	}

	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return nil, nil, err
	}
	hash_string := hex.EncodeToString(hash.Sum(nil))

	return &save_path, &hash_string, nil
}
