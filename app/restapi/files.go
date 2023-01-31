package restapi

import (
	"encoding/json"
	"homestorage/app/utils"
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
	p := req.URL.Query().Get("parent")
	parent_id, _ := strconv.Atoi(p)
	token := req.Header.Get("Authorization")
	_, owner_id, err := utils.IdentifyJWT(strings.Replace(token, "Bearer ", "", 1))
	if err != nil {
		return err
	}

	resp_data, err := h.db.GetScreenListData(parent_id, *owner_id)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp_data)
	return nil
}
