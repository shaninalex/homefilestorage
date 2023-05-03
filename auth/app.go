package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type AccessToken struct {
	Aud   string   `json:"aud"`
	Iss   string   `json:"iss"`
	Sub   int64    `json:"sub"`
	Jti   string   `json:"jti"`
	Roles []string `json:"roles"`
	Exp   int64    `json:"exp"`
}

type AuthPayload struct {
	AccessToken AccessToken `json:"access_token"`
	Exp         int64       `json:"exp"`
}

type LoginPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

var (
	ErrorDefault            = errors.New("something went wrong")
	ErrorPasswordMismatched = errors.New("password mismatched")
	ErrorNotFound           = errors.New("not found")
	ErrorInvalidCredentials = errors.New("invalid credentials")
)

type App struct {
	router  *gin.Engine
	AuthAud string
	AuthIss string
	DB      *sql.DB
}

func (app *App) initialize(dbURL, brokerURL, auth_aud, auth_iss string) {

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Println(err)
		panic("failed to connect database")
	}
	app.DB = db
	app.router = gin.Default()
	app.AuthAud = auth_aud
	app.AuthIss = auth_iss
	app.initializeRoutes()
}

func (app *App) Run(port string) {
	portInt, err := strconv.Atoi(port)
	if err != nil {
		panic(err)
	}
	app.router.Run(fmt.Sprintf(":%d", portInt))
}

func (app *App) ObtainToken(c *gin.Context) {
	var auth_payload LoginPayload

	err := c.ShouldBindJSON(&auth_payload)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var password string
	var id int64
	sql := "SELECT id, password FROM users WHERE email = $1"
	row := app.DB.QueryRow(sql, auth_payload.Email)
	err = row.Scan(&id, &password)

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	match, err := comparePasswordAndHash(auth_payload.Password, password)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !match {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	exp_timestamp := time.Now().Add(time.Hour * 24 * 7) // 3 days

	payload := AuthPayload{
		AccessToken: AccessToken{
			Aud:   app.AuthAud,
			Iss:   app.AuthIss,
			Sub:   id,
			Jti:   randomString(38),
			Roles: []string{"user"},
			Exp:   exp_timestamp.Unix(),
		},
		Exp: exp_timestamp.Unix(),
	}

	c.JSON(http.StatusCreated, payload)
}

func (app *App) initializeRoutes() {
	app.router.POST("/obtain", app.ObtainToken)
}

// HELPERS
func randomString(length int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, length+2)
	rand.Read(b)
	return fmt.Sprintf("%x", b)[2 : length+2]
}
