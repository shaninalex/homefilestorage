package filemanager

import (
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"

	"github.com/google/uuid"
	"github.com/shaninalex/homefilestorage/pkg/database"
)

var (
	ErrPermission = errors.New("unable to create the file for writing. Check your write access privilege")
	ErrWritting   = errors.New("unable to create the file for writing")
)

type IFileManager interface {
	Save(f *multipart.FileHeader) (*database.File, error)
	Delete(path string) error
}

type FileManager struct {
	Path string
}

func InitFileManager(path string) (*FileManager, error) {
	err := os.MkdirAll(filepath.Dir(path), 0755)
	if err != nil {
		return nil, err
	}
	return &FileManager{
		Path: path,
	}, nil
}

func (fm *FileManager) SaveFile(f *multipart.FileHeader) (*database.File, error) {
	file, err := f.Open()
	if err != nil {
		return nil, err
	}
	defer file.Close()

	system_path := fmt.Sprintf("%s/%s", fm.Path, f.Filename)

	// Create the directory path if it doesn't exist
	if err := os.MkdirAll(filepath.Dir(system_path), 0755); err != nil {
		return nil, err
	}

	// Check if the file already exists, and if so, generate a unique filename
	if _, err := os.Stat(system_path); err == nil {
		system_path = fmt.Sprintf("%s/%s_%s", fm.Path, uuid.New().String(), f.Filename)
	}

	out, err := os.Create(system_path)
	if err != nil {
		return nil, ErrPermission
	}
	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		return nil, ErrWritting
	}

	fileObject := &database.File{
		Name:       f.Filename,
		MimeType:   f.Header.Get("Content-Type"),
		Size:       f.Size,
		SystemPath: system_path,
		FolderId:   0,
		CreatedAt:  time.Now(),
	}
	return fileObject, nil
}

func (fm *FileManager) Delete(path string) error {
	err := os.Remove(path)
	if err != nil {
		return err
	}
	return nil
}
