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

	err := r.ParseMultipartForm(200000) // grab the multipart form
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	formdata := r.MultipartForm // ok, no problem so far, read the Form data

	files := formdata.File["files"] // grab the filenames

	for i, _ := range files { // loop through the files one by one
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

		_, err = io.Copy(out, file) // file not files[i] !

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
		// redirect to error page
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

	// Set the appropriate content type for the file
	w.Header().Set("Content-Type", dbfile.MimeType)

	// Set the content-disposition header to specify the filename
	w.Header().Set("Content-Disposition", "attachment; filename="+dbfile.Name)

	// Copy the file content to the response writer
	_, err = io.Copy(w, file)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
