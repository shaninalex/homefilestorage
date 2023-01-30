package filesystem

import (
	"errors"
)

type File struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	MimeType   string `json:"mime_type"`
	Size       int    `json:"size"`
	SystemPath string `json:"system_path"`
	Owner      int    `json:"owner"`
	Hash       string `json:"hash"`
	Public     bool   `json:"public"`
	Created_at string `json:"created_at"`
}

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
