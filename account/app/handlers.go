package app

import (
	"account/app/models"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func (app *App) CreateUser(c *gin.Context) {
	var newUser models.User

	if err := c.ShouldBindJSON(&newUser); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ID, err := newUser.Create(app.DB)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	Publish("register", fmt.Sprintf("New User registered with %v id", ID), app.MQChannel, app.MQQueue)
	c.JSON(http.StatusCreated, gin.H{"success": true})
}

func (app *App) GetUser(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		log.Println("Empty account id")
	}
	intID, err := strconv.Atoi(id)
	if err != nil {
		log.Println("Cant parse id")
		c.JSON(http.StatusNotFound, gin.H{"message": "User does not exists"})
		return
	}

	user, err := models.GetUser(app.DB, int64(intID))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusNotFound, gin.H{"message": "User does not exists"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (app *App) UpdateUser(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		log.Println("Empty account id")
	}

	var payload models.UpdateUser
	if err := c.ShouldBindJSON(&payload); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if reflect.ValueOf(payload).IsZero() {
		log.Println(fmt.Errorf("payload %v is empty or contain wrong values. Nothing to udpate", payload))
		c.JSON(http.StatusBadRequest, gin.H{"message": "Cant handle payload (empty or incorrect)"})
		return
	}

	intID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"message": "Cant parse id"})
		return
	}
	user, err := models.GetUser(app.DB, int64(intID))
	if err != nil {
		log.Println("Cant get user id")
		c.JSON(http.StatusNotFound, gin.H{"message": "User does not exists"})
		return
	}

	// User itself can change only email and username since
	// user active status and password can be changed only via
	// special pipes - change/restore password, activate/deactivate account
	if payload.Email != nil {
		user.Email = *payload.Email
	}

	if payload.Username != nil {
		user.Username = *payload.Username
	}

	err = user.Update(app.DB)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Unable to update account"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true})
}
