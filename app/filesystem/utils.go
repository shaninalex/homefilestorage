package filesystem

import (
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"io"
	"os"
)

var (
	MAX_FILE_SIZE = 314572800 // 300Mb ( NOTE: This constant can be in configuration files )

	ALLOWED_FILE_TYPES = [19]string{

		// DOCUMENTS and PRESENTATIONS
		"application/msword", // .doc
		"application/vnd.openxmlformats-officedocument.wordprocessingml.document",   // .docx
		"application/vnd.ms-powerpoint",                                             // .ppt
		"application/vnd.openxmlformats-officedocument.presentationml.presentation", // .pptx
		"application/vnd.oasis.opendocument.presentation",                           // .odp
		"application/vnd.oasis.opendocument.spreadsheet",                            // .ods
		"application/vnd.oasis.opendocument.text",                                   // .odt
		"application/rtf",
		"application/pdf", // .pdf
		"text/plain",

		// IMAGES
		"image/jpeg", //.jpeg .jpg
		"image/png",  // png

		// AUDIO
		"audio/mpeg", // mp3
		"audio/3gpp", // 3gp
		"audio/ogg",  // oga

		// VIDEO
		"video/mpeg",      // .mpeg
		"video/x-msvideo", // .avi
		"video/ogg",       // .ogv
		"video/mp4",       // .mp4
	}
)

var (
	ERROR_FILE_TOO_BIG          = "File is too big"
	ERROR_EMPTY_FILE            = "Empty file error"
	ERROR_NOT_ALLOWED_FILE_TYPE = "Not allowed file type"
)

func isValidMaxSize(file_size int64) bool {
	if file_size > int64(MAX_FILE_SIZE) {
		return false
	}
	return true
}

func isAllowedFIleType(str string) bool {
	for _, v := range ALLOWED_FILE_TYPES {
		if v == str {
			return true
		}
	}
	return false
}

func ValidateFile(file_size int64, mime_type string) error {

	if isValidMaxSize(file_size) != true {
		return errors.New(ERROR_FILE_TOO_BIG)
	}

	if file_size == 0 {
		return errors.New(ERROR_EMPTY_FILE)
	}

	if isAllowedFIleType(mime_type) != true {
		return errors.New(ERROR_NOT_ALLOWED_FILE_TYPE)
	}

	return nil
}

func HashFileSha1(filePath string) (string, error) {
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
