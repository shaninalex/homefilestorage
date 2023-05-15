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
	Sub   string   `json:"sub"`
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

type DBQueryRequestPayload struct {
	ID       int64
	Password string
}

var (
	ErrorDefault            = errors.New("something went wrong")
	ErrorPasswordMismatched = errors.New("password mismatched")
	ErrorNotFound           = errors.New("not found")
	ErrorInvalidCredentials = errors.New("invalid credentials")
)

type App struct {
	Router  *gin.Engine
	DB      *sql.DB
	AuthAud string
	AuthIss string
}

func (a *App) initialize(dbURL, brokerURL, auth_aud, auth_iss string) {
	var err error
	a.DB, err = sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Connected to Postgres")
	}

	a.Router = gin.Default()
	a.AuthAud = auth_aud
	a.AuthIss = auth_iss

	a.initializeRoutes()
}

func (a *App) Run(port string) {
	log.Printf("Listen port: %s\n", port)
	port_int, err := strconv.Atoi(port)
	if err != nil {
		log.Fatal(err)
	}
	addr := fmt.Sprintf(":%d", port_int)
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

func (a *App) obtainToken(c *gin.Context) {
	var auth_payload LoginPayload
	err := c.ShouldBindJSON(&auth_payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var password string
	var id int64
	sql := "SELECT id, password FROM users WHERE email = $1"
	row := a.DB.QueryRow(sql, auth_payload.Email)
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

	current := time.Now().Add(time.Minute * 15)

	payload := AuthPayload{
		AccessToken: AccessToken{
			Aud:   a.AuthAud,
			Iss:   a.AuthIss,
			Sub:   strconv.Itoa(int(id)),
			Jti:   randomString(38),
			Roles: []string{"user", "admin"},
			Exp:   current.Unix(),
		},
		Exp: current.Unix(),
	}

	c.JSON(http.StatusOK, payload)
	if err != nil {
		log.Println(err)
		return
	}
}

func (a *App) initializeRoutes() {
	a.Router.POST("/obtain", a.obtainToken)
}

// HELPERS
func randomString(length int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, length+2)
	rand.Read(b)
	return fmt.Sprintf("%x", b)[2 : length+2]
}
