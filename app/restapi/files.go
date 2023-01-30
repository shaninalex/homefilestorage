package restapi

import (
	"homestorage/app/filesystem"
	"homestorage/app/utils"
	"io"
	"net/http"
	"strings"

	"github.com/uptrace/bunrouter"
)

func (h *BaseHandler) RouteSaveFile(w http.ResponseWriter, req bunrouter.Request) error {
	// Get user
	token := req.Header.Get("Authorization")
	_, id, err := utils.IdentifyJWT(strings.Replace(token, "Bearer ", "", 1))
	if err != nil {
		return err
	}

	// Get and validate request file
	rFile, fileHeader, err := req.FormFile("file")
	if err != nil {
		return err
	}
	defer rFile.Close()

	robots, err := io.ReadAll(rFile)
	if err != nil {
		return err
	}
	file_type := http.DetectContentType(robots)

	err = filesystem.ValidateFile(fileHeader.Size, file_type)
	if err != nil {
		return err
	}

	systemPath, hash, err := filesystem.SaveFileToStorage(rFile, fileHeader)
	if err != nil {
		return err
	}

	file := filesystem.File{
		Owner:      *id,
		MimeType:   file_type,
		Size:       int(fileHeader.Size),
		SystemPath: systemPath,
		Hash:       hash,
	}

	err = h.db.SaveFileRecord(&file)

	if err != nil {
		return err
	}

	return nil
}
