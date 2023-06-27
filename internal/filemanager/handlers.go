package filemanager

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/shaninalex/homefilestorage/internal/filemanager"
)

type FileResponse struct {
	Name       string `json:"name"`
	MimeType   string `json:"mime_type"`
	Size       uint   `json:"size"`
	SystemPath string `json:"system_path"`
	Hash       string `json:"hash"`
}

func (app *App) GetSingleFile(file_id int, user_id string) (*File, error) {
	file, err := GetFile(app.DB, user_id, int64(file_id))
	if err != nil {
		return nil, err
	}
	return file, nil
}

func (app *App) GetFilesList(c *gin.Context) {
	user_id := c.Request.Header.Get("X-User")
	folder_id, _ := strconv.Atoi(c.Query("folder_id"))
	files, err := GetUserFiles(app.DB, user_id, int64(folder_id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
	c.JSON(http.StatusOK, gin.H{"files": files})
}

func (app *App) SaveFile(c *gin.Context) {
	// check user existens ( this step require several steps - does it exists in database, active/inactive, Personal store GB limit)
	user_id := c.Request.Header.Get("X-User")
	respAccount, err := http.Get(fmt.Sprintf("%s/account/%s", filemanager.ServiceAccount, user_id))
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

	c.JSON(http.StatusOK, gin.H{"file": file})
}

func (app *App) FileData(c *gin.Context) {
	file_id, _ := c.Params.Get("file_id")
	user_id := c.Request.Header.Get("X-User")
	file_id_int, _ := strconv.Atoi(file_id)

	file, err := database.GetFile(app.DB, user_id, int64(file_id_int))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	log.Printf("%s%s\n", app.ServiceStorage, file.SystemPath)

	var fileBody io.Reader
	req, err := http.NewRequest("GET", fmt.Sprintf("%s%s", app.ServiceStorage, file.SystemPath), fileBody)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()

	fileBites, _ := ioutil.ReadAll(resp.Body)

	c.Data(http.StatusOK, file.MimeType, fileBites)
}