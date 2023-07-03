package api

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	ory "github.com/ory/kratos-client-go"
)

func (api *Api) validateSession(r *http.Request) (*ory.Session, error) {
	cookie, err := r.Cookie("ory_kratos_session")
	if err != nil {
		return nil, err
	}
	if cookie == nil {
		return nil, errors.New("no session found in cookie")
	}
	resp, _, err := api.ory.FrontendApi.ToSession(context.Background()).Cookie(cookie.String()).Execute()
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (api *Api) GetUserInfoBySession(c *gin.Context) {

	session, err := api.validateSession(c.Request)
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

func (api *Api) CheckUserSession(c *gin.Context) {

	session, err := api.validateSession(c.Request)
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
