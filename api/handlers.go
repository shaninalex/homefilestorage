package api

import (
	"io/ioutil"
	"log"
	"mime"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (api *Api) AppHealth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": true})
}

func (api *Api) FilesList(c *gin.Context) {
	user_id := c.Request.Header.Get("X-User")
	folder_id, _ := strconv.Atoi(c.Query("folder_id"))
	files, err := api.database.GetUserFiles(user_id, int64(folder_id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
	c.JSON(http.StatusOK, gin.H{"files": files})
}

func (api *Api) FilesUpload(c *gin.Context) {
	d, err := ioutil.ReadAll(c.Request.Body)
	filename := handleMediaType(c.Request.Header.Get("Content-Disposition"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error reading input data"})
	}
	fileinfo, err := api.filemanager.SaveFile(filename, d)
	// TODO: Save fileinfo to user
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error reading input data"})
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
