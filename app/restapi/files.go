package restapi

import (
	"encoding/json"
	"errors"
	"homestorage/app/utils"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/uptrace/bunrouter"
)

func (h *BaseHandler) RouteSaveFile(w http.ResponseWriter, req bunrouter.Request) error {
	token := req.Header.Get("Authorization")
	_, id, err := utils.IdentifyJWT(strings.Replace(token, "Bearer ", "", 1))
	if err != nil {
		return err
	}

	file, handler, err := req.FormFile("file")
	if err != nil {
		return err
	}
	defer file.Close()

	dFile := &utils.File{
		Owner:  *id,
		Size:   int(handler.Size),
		Name:   handler.Filename,
		Public: true,
	}

	dFile, err = h.storage.SaveFileToStorage(file, handler, dFile)
	if err != nil {
		return err
	}

	new_file_id, err := h.db.SaveFileRecord(dFile)

	if err != nil {
		return err
	}

	dFile.Id = new_file_id

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(dFile)

	return nil
}

func (h *BaseHandler) RouteFilesList(w http.ResponseWriter, req bunrouter.Request) error {
	token := req.Header.Get("Authorization")
	_, _, err := utils.IdentifyJWT(strings.Replace(token, "Bearer ", "", 1))
	if err != nil {
		return err
	}

	_parent := req.URL.Query().Get("parent")
	parent, err := strconv.Atoi(_parent)
	if err != nil {
		return err
	}

	_offset := req.URL.Query().Get("offset")
	offset, err := strconv.Atoi(_offset)
	if err != nil {
		return err
	}

	_limit := req.URL.Query().Get("limit")
	limit, err := strconv.Atoi(_limit)
	if err != nil {
		return err
	}

	// TODO: max limit offset pagination values
	// 		 to prevent users to get super large values
	log.Printf("Parent: %d, Limit: %d, Offset: %d\n", parent, limit, offset)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(nil)
	return nil
}

func (h *BaseHandler) RouteFileItem(w http.ResponseWriter, req bunrouter.Request) error {
	token := req.Header.Get("Authorization")
	_, userId, err := utils.IdentifyJWT(strings.Replace(token, "Bearer ", "", 1))
	if err != nil {
		return err
	}

	_fileId, exist := req.Params().Get("id")
	fileId, err := strconv.Atoi(_fileId)
	if err != nil {
		return err
	}
	if !exist {
		return errors.New("o file id provided")
	}

	file, err := h.db.GetFile(fileId, *userId)
	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(file)

	return nil
}

func (h *BaseHandler) RouteFileItemContent(w http.ResponseWriter, req bunrouter.Request) error {
	token := req.Header.Get("Authorization")
	_, userId, err := utils.IdentifyJWT(strings.Replace(token, "Bearer ", "", 1))
	if err != nil {
		return err
	}

	_fileId, exist := req.Params().Get("id")
	fileId, err := strconv.Atoi(_fileId)
	if err != nil {
		return err
	}
	if !exist {
		return errors.New("o file id provided")
	}

	filePath, err := h.db.GetFilePath(fileId, *userId)
	if err != nil {
		return err
	}

	http.ServeFile(w, req.Request, *filePath)

	return nil
}
