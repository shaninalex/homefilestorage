package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/shaninalex/homefilestorage/internal/filemanager"
)

type Api struct {
	router      *gin.Engine
	filemanager *filemanager.FileManager
}

func CreateApplication(filemanager *filemanager.FileManager) (*Api, error) {
	var api Api

	api.filemanager = filemanager
	api.router = gin.Default()
	return &api, nil
}

func (api *Api) initializeRoutes() {
	api.router.GET("/health", api.AppHealth)

	// TODO: Get files list
	// TODO: Upload File
}

func (api *Api) Run(port int) {
	api.router.Run(fmt.Sprintf(":%d", port))
}

func (api *Api) AppHealth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": true})
}
