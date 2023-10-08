package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/shaninalex/homefilestorage/pkg/database"
	"github.com/shaninalex/homefilestorage/pkg/filemanager"
)

type Api struct {
	router      *gin.Engine
	filemanager *filemanager.FileManager
	database    *database.Database
}

func CreateApi(filemanager *filemanager.FileManager, database *database.Database) (*Api, error) {
	var api Api

	api.database = database
	api.filemanager = filemanager
	api.router = gin.Default()

	api.initializeRoutes()

	return &api, nil
}

func (api *Api) initializeRoutes() {
	api.router.GET("/health", api.AppHealth)

	files := api.router.Group("/files")
	files.Use(SetUserID())
	{
		files.GET("/:file_id", api.FilesItem)
		files.GET("/list", api.FilesList)
		files.POST("/upload", api.FilesUpload)
	}

	user := api.router.Group("/user")
	user.Use(SetUserID())
	{
		user.GET("/info", api.GetUserInfoBySession)
		user.GET("/check", api.CheckUserSession)
	}
}

func (api *Api) Run(port int) {
	api.router.Run(fmt.Sprintf(":%d", port))
}
