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
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

<<<<<<< HEAD
	result, err := newUser.Create(app.DB)
=======
	ID, err := newUser.Create(app.DB)
>>>>>>> 8898747274b1392e4411946473428bf6c315fbaf
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

<<<<<<< HEAD
	Publish("register", fmt.Sprintf("New User registered with %s id", result), app.MQChannel, app.MQQueue)
=======
	Publish("register", fmt.Sprintf("New User registered with %s id", ID), app.MQChannel, app.MQQueue)
>>>>>>> 8898747274b1392e4411946473428bf6c315fbaf
	c.JSON(http.StatusCreated, gin.H{"success": true})
}

func (app *App) GetUser(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		log.Println("Empty account id")
	}
	uintID, err := strconv.Atoi(id)

	user, err := models.Get(app.DB, uint(uintID))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "User does not exists"})
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
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	if reflect.ValueOf(payload).IsZero() {
		log.Println(fmt.Errorf("payload %v is empty or contain wrong values. Nothing to udpate", payload))
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Cant handle payload (empty or incorrect)"})
		return
	}

	err := .Update(app.DB, id, payload)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Unable to update account"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true})
}
