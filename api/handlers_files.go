package api

import (
	"io"
	"log"
	"mime"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	db "github.com/shaninalex/homefilestorage/internal/database"
	fm "github.com/shaninalex/homefilestorage/internal/filemanager"
)

func (api *Api) AppHealth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": true})
}

func (api *Api) FilesList(c *gin.Context) {
	user_id := c.Request.Header.Get("X-User")
	folder_id, _ := strconv.Atoi(c.Query("folder_id"))
	files, err := api.database.GetUserFiles(user_id, int64(folder_id))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
	c.JSON(http.StatusOK, gin.H{"files": files})
}

func (api *Api) FilesUpload(c *gin.Context) {
	user_id := c.Request.Header.Get("X-User")

	d, err := io.ReadAll(c.Request.Body)
	filename := handleMediaType(c.Request.Header.Get("Content-Disposition"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error reading input data"})
		return
	}
	fileinfo, err := api.filemanager.SaveFile(filename, d)
	go func(f *fm.FileResponse) {
		file := &db.File{
			Name:       f.Name,
			MimeType:   f.MimeType,
			Size:       uint(f.Size),
			SystemPath: f.SystemPath,
			Hash:       f.Hash,
			Owner:      user_id,
			FolderId:   0,
			Public:     false,
		}
		err := api.database.FileSave(file)
		if err != nil {
			log.Println(err)
		}
	}(fileinfo)

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "error reading input data"})
		return
	}
	c.JSON(http.StatusOK, fileinfo)
}

func handleMediaType(header_media_type string) string {
	_, params, err := mime.ParseMediaType(header_media_type)
	if err != nil {
		log.Println(err)
	}
	return params["filename"]
}
