package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shaninalex/homefilestorage/pkg/database"
)

type Api struct {
	router   *gin.Engine
	database database.Repository
}

func CreateApi(database database.Repository) (*Api, error) {
	var api Api

	api.database = database
	api.router = gin.Default()

	api.initializeRoutes()

	return &api, nil
}

func (api *Api) initializeRoutes() {
	api.router.GET("/health", api.AppHealth)
}

func (api *Api) Run(port int) {
	api.router.Run(fmt.Sprintf(":%d", port))
}

func (api *Api) AppHealth(c *gin.Context) {
	c.JSON(http.StatusOK, nil)
}
