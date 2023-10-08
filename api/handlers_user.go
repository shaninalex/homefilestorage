package api

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (api *Api) validateSession(r *http.Request) (interface{}, error) {
	cookie, err := r.Cookie("ory_kratos_session")
	return nil, nil
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
