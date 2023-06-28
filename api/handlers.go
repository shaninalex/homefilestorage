package api

import (
	"io/ioutil"
	"log"
	"mime"
	"net/http"
	"os"
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
	// move things below to filemanager package
	fileinfo, err := api.filemanager.SaveFile(filename, d)
	tmpfile, err := os.Create("./" + filename)
	defer tmpfile.Close()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error writing file"})
	}
	tmpfile.Write(d)
	c.JSON(http.StatusOK, gin.H{"success": "file saved"})
}

func handleMediaType(header_media_type string) string {
	_, params, err := mime.ParseMediaType(header_media_type)
	if err != nil {
		log.Println(err)
	}
	return params["filename"]
}
