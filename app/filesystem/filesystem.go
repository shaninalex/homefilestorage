package filesystem

import (
	"errors"
	"fmt"
	"homestorage/app/utils"
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"

	"github.com/h2non/filetype"
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

func (s *FileStorage) SaveFileToStorage(file multipart.File, fileHeader *multipart.FileHeader, dFile *utils.File) (*utils.File, error) {

	new_generated_name := fmt.Sprint(time.Now().Unix()) + filepath.Ext(fileHeader.Filename)

	date := time.Now()
	save_path := fmt.Sprintf("%s/%d/%d/%d/%s", s.storage, date.Year(), date.Month(), date.Day(), new_generated_name)

	f, err := os.OpenFile(save_path, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	_, err = io.Copy(f, file) // save content into new file
	if err != nil {
		return nil, err
	}

	hash_string, err := HashFileSha1(save_path)
	if err != nil {
		return nil, err
	}

	buf, _ := ioutil.ReadFile(save_path)
	kind, _ := filetype.Match(buf)

	if filetype.IsImage(buf) && filetype.IsDocument(buf) && filetype.IsAudio(buf) && filetype.IsArchive(buf) && filetype.IsVideo(buf) {
		os.Remove(save_path)
		return nil, errors.New("unknow file type")
	}

	dFile.SystemPath = save_path
	dFile.Hash = hash_string
	dFile.MimeType = kind.MIME.Value

	return dFile, nil
}
