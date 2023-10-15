package web

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/shaninalex/homefilestorage/pkg/database"
)

func (web *WebApp) filesUploadHandler(w http.ResponseWriter, r *http.Request) {
	web.State.Error = ""

	err := r.ParseMultipartForm(200000)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	formdata := r.MultipartForm
	files := formdata.File["files"]

	for i, _ := range files {
		file, err := files[i].Open()
		if err != nil {
			web.State.Error = err.Error()
			http.Redirect(w, r, "/", http.StatusMovedPermanently)
			return
		}
		defer file.Close()

		system_path := fmt.Sprintf("%s/%s", web.Config.Storage.Path, files[i].Filename)
		out, err := os.Create(system_path)
		if err != nil {
			web.State.Error = "Unable to create the file for writing. Check your write access privilege"
			http.Redirect(w, r, "/", http.StatusMovedPermanently)
			return
		}
		defer out.Close()

		_, err = io.Copy(out, file)
		if err != nil {
			fmt.Fprintln(w, err)
			web.State.Error = "Unable to create the file for writing"
			http.Redirect(w, r, "/", http.StatusMovedPermanently)
			return
		}
		db_file := &database.File{
			Name:       files[i].Filename,
			MimeType:   files[i].Header.Get("Content-Type"),
			Size:       files[i].Size,
			SystemPath: system_path,
			FolderId:   0,
			CreatedAt:  time.Now(),
		}

		if err := web.Database.SaveFile(db_file); err != nil {
			web.State.Error = err.Error()
		}
	}
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func (web *WebApp) filesGetHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	file_id, err := strconv.ParseInt(id, 10, 8)
	if err != nil {
		http.Redirect(w, r, "/", http.StatusMovedPermanently)
	}
	dbfile, err := web.Database.GetFile(file_id)
	if err != nil {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}

	file, err := os.Open(dbfile.SystemPath)
	if err != nil {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}
	defer file.Close()

	w.Header().Set("Content-Type", dbfile.MimeType)
	w.Header().Set("Content-Disposition", "attachment; filename="+dbfile.Name)
	_, err = io.Copy(w, file)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
