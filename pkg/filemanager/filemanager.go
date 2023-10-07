package filemanager

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/google/uuid"
	"github.com/h2non/filetype"
	"github.com/shaninalex/homefilestorage/internal/typedefs"
)

type FileManager struct {
	ServiceStorage string
}

func Initialize(storage_service_url string) *FileManager {
	fm := &FileManager{}
	fm.ServiceStorage = storage_service_url
	return fm
}

func (filemanager *FileManager) SaveFile(filename string, d []byte) (*typedefs.File, error) {

	// Create path
	u := uuid.New()
	new_generated_name := fmt.Sprint(fmt.Sprintf("%v", u) + filepath.Ext(filename))
	date := time.Now()
	uploads_path := fmt.Sprintf("%s/%d/%d/%d/%s", filemanager.ServiceStorage, date.Year(), date.Month(), date.Day(), new_generated_name)
	if err := os.MkdirAll(filepath.Dir(uploads_path), 0770); err != nil {
		log.Println(err)
		return nil, err
	}

	// create file
	tmpfile, err := os.Create(uploads_path)
	if err != nil {
		return nil, err
	}
	defer tmpfile.Close()
	if err != nil {
		return nil, err
	}
	tmpfile.Write(d)

	hash, err := hashFileSha1(uploads_path)
	if err != nil {
		return nil, err
	}
	mime, err := getMIME(uploads_path)
	if err != nil {
		return nil, err
	}
	file_size, err := getFileSize(uploads_path)
	if err != nil {
		return nil, err
	}

	response := &typedefs.File{
		Hash:       hash,
		MimeType:   mime,
		SystemPath: uploads_path,
		Name:       filename,
		Size:       uint(file_size),
	}

	return response, nil
}

func getMIME(save_path string) (string, error) {
	log.Println(save_path)

	buf, err := ioutil.ReadFile(save_path)
	if err != nil {
		return "", err
	}

	kind, err := filetype.Match(buf)
	if err != nil {
		return "", err
	}

	// TODO: Handle unknow file type error ( if it will be error )
	// if filetype.IsImage(buf) && filetype.IsDocument(buf) && filetype.IsAudio(buf) && filetype.IsArchive(buf) && filetype.IsVideo(buf) {
	// 	os.Remove(save_path)
	// 	return "", errors.New("unknow file type")
	// }
	return kind.MIME.Value, nil
}

func hashFileSha1(filePath string) (string, error) {
	//Initialize variable returnMD5String now in case an error has to be returned
	var returnSHA1String string

	//Open the filepath passed by the argument and check for any error
	file, err := os.Open(filePath)
	if err != nil {
		return returnSHA1String, err
	}

	//Tell the program to call the following function when the current function returns
	defer file.Close()

	//Open a new SHA1 hash interface to write to
	hash := sha1.New()

	//Copy the file in the hash interface and check for any error
	if _, err := io.Copy(hash, file); err != nil {
		return returnSHA1String, err
	}

	//Get the 20 bytes hash
	hashInBytes := hash.Sum(nil)[:20]

	//Convert the bytes to a string
	returnSHA1String = hex.EncodeToString(hashInBytes)

	return returnSHA1String, nil
}

func getFileSize(uploads_path string) (int64, error) {
	fileInfo, err := os.Stat(uploads_path)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	return fileInfo.Size(), nil
}
