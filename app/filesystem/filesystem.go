package filesystem

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"

	"github.com/google/uuid"
)

type FileStorage struct {
	storage string
}

func CreateFileStorage(path string) (*FileStorage, error) {
	_, err := os.Stat(path)
	if err != nil {
		return nil, err
	}
	if os.IsNotExist(err) {
		return nil, err
	}
	// TODO: Create if not exist?...
	return &FileStorage{
		storage: path,
	}, nil
}

func (s *FileStorage) SaveFileToStorage(file multipart.File, fileHeader *multipart.FileHeader) (*string, *string, error) {

	ext1 := filepath.Ext(fileHeader.Filename)
	new_generated_name := uuid.New().String() + ext1

	date := time.Now()
	save_path := fmt.Sprintf("%s/%d/%d/%d/%s", s.storage, date.Year(), date.Month(), date.Day(), new_generated_name)

	f, err := os.OpenFile(save_path, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return nil, nil, err
	}
	defer f.Close()
	io.Copy(f, file) // save content into new file

	// hash_string, err := "HashFileSha1(save_path)", nil
	// if err != nil {
	// 	return nil, nil, err
	// }
	hash_string := "&hash_string"
	return &save_path, &hash_string, nil
}
