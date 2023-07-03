package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	ory "github.com/ory/kratos-client-go"
	"github.com/shaninalex/homefilestorage/internal/database"
	"github.com/shaninalex/homefilestorage/internal/filemanager"
)

type Api struct {
	router      *gin.Engine
	filemanager *filemanager.FileManager
	database    *database.Database
	ory         *ory.APIClient
	kratos_path string
}

func CreateApi(filemanager *filemanager.FileManager, database *database.Database, kratos_path string) (*Api, error) {
	var api Api

	api.database = database
	api.filemanager = filemanager
	api.router = gin.Default()
	api.kratos_path = kratos_path

	configuration := ory.NewConfiguration()
	configuration.Servers = []ory.ServerConfiguration{
		{
			URL: "http://kratos:4433",
		},
	}
	api.ory = ory.NewAPIClient(configuration)
	api.initializeRoutes()

	return &api, nil
}

func (api *Api) initializeRoutes() {
	api.router.GET("/health", api.AppHealth)

	api.router.GET("/files/list", api.FilesList)
	api.router.POST("/files/upload", api.FilesUpload)

	api.router.GET("/user/info", api.GetUserInfoBySession)
	api.router.GET("/user/check", api.CheckUserSession)

}

func (api *Api) Run(port int) {
	api.router.Run(fmt.Sprintf(":%d", port))
}
