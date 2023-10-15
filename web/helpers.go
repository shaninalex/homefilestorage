package web

import (
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

func (web *WebApp) helperFileDelete(w http.ResponseWriter, r *http.Request) {
	web.State.Error = ""
	vars := mux.Vars(r)
	id := vars["id"]
	file_id, err := strconv.ParseInt(id, 10, 8)
	if err != nil {
		http.Error(w, "Unable to parse file id", http.StatusBadRequest)
		return
	}
	fileObject, err := web.Database.GetFile(file_id)
	if err != nil {
		http.Error(w, "File object not found", http.StatusNotFound)
		return
	}
	err = web.FileManager.Delete(fileObject.SystemPath)
	if err != nil {
		http.Error(w, "Unable to delete file form storage", http.StatusBadRequest)
		return
	}
	err = web.Database.DeleteFile(fileObject.ID)
	if err != nil {
		http.Error(w, "Unable to delete file from database", http.StatusBadRequest)
		return
	}
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func (web *WebApp) helperFileGet(w http.ResponseWriter, r *http.Request) {
	web.State.Error = ""
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

func (web *WebApp) helperFileSave(w http.ResponseWriter, r *http.Request) {
	web.State.Error = ""
	err := r.ParseMultipartForm(200000)
	if err != nil {
		web.State.Error = err.Error()
		http.Redirect(w, r, "/", http.StatusMovedPermanently)
		return
	}

	formdata := r.MultipartForm
	files := formdata.File["files"]

	for _, file_item := range files {
		file, err := web.FileManager.SaveFile(file_item)
		if err != nil {
			web.State.Error = err.Error()
			http.Redirect(w, r, "/", http.StatusMovedPermanently)
			return
		}
		if err := web.Database.SaveFile(file); err != nil {
			web.State.Error = err.Error()
		}
	}
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func (web *WebApp) helperFilePreview(w http.ResponseWriter, r *http.Request) {
	web.State.Error = ""
	vars := mux.Vars(r)
	id := vars["id"]
	file_id, err := strconv.ParseInt(id, 10, 8)
	if err != nil {
		log.Println(err)
	}
	dbfile, err := web.Database.GetFile(file_id)
	if err != nil {
		log.Println(err)
		return
	}

	file, err := os.Open(dbfile.SystemPath)
	if err != nil {
		log.Println(err)
		return
	}
	defer file.Close()

	w.Header().Set("Content-Type", dbfile.MimeType)
	_, err = io.Copy(w, file)
	if err != nil {
		log.Println(err)
		return
	}
}
