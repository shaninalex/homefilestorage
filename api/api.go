package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shaninalex/homefilestorage/pkg/config"
	"github.com/shaninalex/homefilestorage/pkg/database"
)

type Api struct {
	router   *gin.Engine
	database database.Repository
	config   *config.Config
}

func CreateApi(database database.Repository, config *config.Config) (*Api, error) {
	var api Api

	gin.SetMode(config.GIN.Mode)
	api.database = database
	api.router = gin.Default()
	api.config = config

	api.initializeRoutes()

	return &api, nil
}

func (api *Api) initializeRoutes() {
	api.router.GET("/health", api.AppHealth)

	account := api.router.Group("api/v2/account")
	{
		account.GET("/", nil)
		account.PATCH("/", nil)
		account.POST("/login", nil)
		account.GET("/logout", nil)
	}
}

func (api *Api) Run(port int64) {
	api.router.Run(fmt.Sprintf(":%d", port))
}

func (api *Api) AppHealth(c *gin.Context) {
	c.JSON(http.StatusOK, nil)
}
