package app

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/shaninalex/homefilestorage/filemanger/app/database"
)

type FileResponse struct {
	Name       string `json:"name"`
	MimeType   string `json:"mime_type"`
	Size       uint   `json:"size"`
	SystemPath string `json:"system_path"`
	Hash       string `json:"hash"`
}

func Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

// This api handler get
// - user_id from JWT ( unpacked by krakend )
// - file_id - the id of file he want to get
// If user hase right permissions and file exist function return
// the file itseld
func (app *App) GetFile(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

// To save the file user should provide:
// - user_id
// - file it self
// Optional:
// Folder_id he want to save the file
// Krakend will unpack user JWT token, get his sub and add into URL
func (app *App) SaveFile(c *gin.Context) {
	// check user existens ( this step require several steps - does it exists in database, active/inactive, Personal store GB limit)
	user_id := c.Params.ByName("user_id")
	respAccount, err := http.Get(fmt.Sprintf("%s/account/%s", app.ServiceAccount, user_id))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	if respAccount.StatusCode != 200 {
		// if user not found or something else happening...
		c.Status(respAccount.StatusCode)
		c.Writer.WriteHeaderNow()
		return
	}

	// resend form data into storage service
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/save", app.ServiceStorage), c.Request.Body)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	req.Header.Set("Content-Type", c.Request.Header.Get("Content-Type"))
	req.Header.Set("Content-Length", c.Request.Header.Get("Content-Length"))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	defer resp.Body.Close()

	// Return response from another backend
	rbody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	var responseStorage FileResponse
	if err := json.Unmarshal(rbody, &responseStorage); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	var file database.File
	userIdUint, _ := strconv.Atoi(user_id)

	file.Name = responseStorage.Name
	file.MimeType = responseStorage.MimeType
	file.Size = responseStorage.Size
	file.SystemPath = responseStorage.SystemPath
	file.Hash = responseStorage.Hash
	file.Owner = uint(userIdUint)
	file.Public = true

	err = file.FileSave(app.DB)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	Publish(app.MQQueue.Name, fmt.Sprintf("New file saved from user: %s", user_id), app.MQChannel, app.MQQueue)

	c.JSON(http.StatusOK, file)
}
