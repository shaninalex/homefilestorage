package filesystem

import "mime/multipart"

func SaveFileToStorage(file multipart.File, fileHeader *multipart.FileHeader) (string, string, error) {
	return "filesistem", "hash", nil
}
