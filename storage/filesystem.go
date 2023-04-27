package main

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"

	"github.com/h2non/filetype"
)

type File struct {
	Id         int       `json:"id"`
	Name       string    `json:"name"`
	MimeType   string    `json:"mime_type"`
	Size       int       `json:"size"`
	SystemPath string    `json:"system_path"`
	Owner      int       `json:"owner"`
	Hash       string    `json:"hash"`
	Public     bool      `json:"public"`
	Folder     *int      `json:"folder"`
	Created_at time.Time `json:"created_at"`
}

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

func (s *FileStorage) SaveFileToStorage(file multipart.File, fileHeader *multipart.FileHeader, dFile *File) (*File, error) {

	new_generated_name := fmt.Sprint(time.Now().Unix()) + filepath.Ext(fileHeader.Filename)

	date := time.Now()
	save_path := fmt.Sprintf("%s/%d/%d/%d/%s", s.storage, date.Year(), date.Month(), date.Day(), new_generated_name)

	// Create file path before creating
	if err := os.MkdirAll(filepath.Dir(save_path), 0770); err != nil {
		return nil, err
	}

	// Opening file
	f, err := os.OpenFile(save_path, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer f.Close()

	// save content into file
	_, err = io.Copy(f, file)
	if err != nil {
		return nil, err
	}

	hash_string, err := HashFileSha1(save_path)
	if err != nil {
		return nil, err
	}

	buf, err := ioutil.ReadFile(save_path)
	if err != nil {
		return nil, err
	}

	kind, err := filetype.Match(buf)
	if err != nil {
		return nil, err
	}

	if filetype.IsImage(buf) && filetype.IsDocument(buf) && filetype.IsAudio(buf) && filetype.IsArchive(buf) && filetype.IsVideo(buf) {
		os.Remove(save_path)
		return nil, errors.New("unknow file type")
	}

	dFile.SystemPath = save_path
	dFile.Hash = hash_string
	dFile.MimeType = kind.MIME.Value

	return dFile, nil
}
