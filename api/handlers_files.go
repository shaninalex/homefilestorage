package api

import (
	"io"
	"log"
	"mime"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/shaninalex/homefilestorage/internal/typedefs"
)

func (api *Api) AppHealth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": true})
}

func (api *Api) FilesList(c *gin.Context) {
	user_id := c.MustGet("user_id").(string)
	folder_id, _ := strconv.Atoi(c.Query("folder_id"))
	files, err := api.database.GetUserFiles(user_id, int64(folder_id))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
	c.JSON(http.StatusOK, gin.H{"files": files})
}

func (api *Api) FilesUpload(c *gin.Context) {
	user_id := c.MustGet("user_id").(string)

	d, err := io.ReadAll(c.Request.Body)
	filename := handleMediaType(c.Request.Header.Get("Content-Disposition"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error reading input data"})
		return
	}
	f, err := api.filemanager.SaveFile(filename, d)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "cant save file"})
		return
	}

	file := &typedefs.File{
		Name:       f.Name,
		MimeType:   f.MimeType,
		Size:       uint(f.Size),
		SystemPath: f.SystemPath,
		Hash:       f.Hash,
		Owner:      user_id,
		FolderId:   0,
		Public:     false,
	}
	err = api.database.FileSave(file)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "cant save file to db"})
		return
	}

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "error reading input data"})
		return
	}

	// TODO: write metrics and logs
	c.JSON(http.StatusOK, &file)
}

func (api *Api) FilesItem(c *gin.Context) {
	user_id := c.MustGet("user_id").(string)
	file_id, exists := c.Params.Get("file_id")
	if !exists {
		log.Println(exists)
		c.JSON(http.StatusNotFound, gin.H{"error": "Not found"})
		return
	}

	_file_id, err := strconv.Atoi(file_id)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid file id"})
		return
	}
	file, err := api.database.GetFile(user_id, int64(_file_id))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid file id"})
		return
	}
	log.Println(file.SystemPath)
	c.File(file.SystemPath)
}

func handleMediaType(header_media_type string) string {
	_, params, err := mime.ParseMediaType(header_media_type)
	if err != nil {
		log.Println(err)
	}
	return params["filename"]
}
