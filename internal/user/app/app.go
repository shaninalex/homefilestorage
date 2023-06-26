package app

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	ory "github.com/ory/kratos-client-go"
)

type App struct {
	router      *gin.Engine
	ory         *ory.APIClient
	kratos_path string
}

func (app *App) Initialize(kratos_path string) error {
	app.kratos_path = kratos_path

	configuration := ory.NewConfiguration()
	configuration.Servers = []ory.ServerConfiguration{
		{
			URL: "http://kratos:4433",
		},
	}
	app.ory = ory.NewAPIClient(configuration)
	app.router = gin.Default()
	app.initializeRoutes()

	return nil
}

func (app *App) initializeRoutes() {
	app.router.GET("/user/info", app.GetUserInfoBySession)
	app.router.GET("/user/check", app.CheckUserSession)
}

func (app *App) Run(port string) {
	portInt, err := strconv.Atoi(port)
	if err != nil {
		panic(err)
	}
	app.router.Run(fmt.Sprintf(":%d", portInt))
}

func (app *App) validateSession(r *http.Request) (*ory.Session, error) {
	cookie, err := r.Cookie("ory_kratos_session")
	if err != nil {
		return nil, err
	}
	if cookie == nil {
		return nil, errors.New("no session found in cookie")
	}
	resp, _, err := app.ory.FrontendApi.ToSession(context.Background()).Cookie(cookie.String()).Execute()
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (app *App) GetUserInfoBySession(c *gin.Context) {

	session, err := app.validateSession(c.Request)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	if !*session.Active {
		log.Println(err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Session not active"})
		return
	}

	c.JSON(http.StatusOK, session.Identity.Traits)

}

func (app *App) CheckUserSession(c *gin.Context) {

	session, err := app.validateSession(c.Request)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	if !*session.Active {
		log.Println(err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Session not active"})
		return
	}

	c.JSON(http.StatusOK, nil)
}
