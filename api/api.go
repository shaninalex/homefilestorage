package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/shaninalex/homefilestorage/internal/database"
	"github.com/shaninalex/homefilestorage/internal/filemanager"
)

type Api struct {
	router      *gin.Engine
	filemanager *filemanager.FileManager
	database    *database.Database
}

func CreateApplication(
	filemanager *filemanager.FileManager,
	database *database.Database,
) (*Api, error) {
	var api Api

	api.database = database
	api.filemanager = filemanager
	api.router = gin.Default()
	return &api, nil
}

func (api *Api) initializeRoutes() {
	api.router.GET("/health", api.AppHealth)
	api.router.GET("/api/v2/files/list", api.FilesList)
	api.router.POST("/api/v2/files/upload", api.FilesUpload)
	// TODO: Upload File
}

func (api *Api) Run(port int) {
	api.router.Run(fmt.Sprintf(":%d", port))
}
