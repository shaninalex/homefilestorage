package restapi

import (
	"errors"
	"fmt"
	"homestorage/app/database"
	"log"
	"net/http"

	"github.com/uptrace/bunrouter"
)

type BaseHandler struct {
	db *database.DatabaseRepository
}

func Handlers(db *database.DatabaseRepository) *BaseHandler {
	return &BaseHandler{
		db: db,
	}
}

// TODO i18n errors texts
var (
	ErrParse                       = errors.New("cannot parse request")
	ErrRegistrationPasswordConfirm = errors.New("password and password confirm are not the same")
	ErrPasswordDoesNotMatch        = errors.New("password does not match")
	ErrCantRafreshToken            = errors.New("cant refresh token")
)

func Server(db *database.DatabaseRepository, port int) {

	router := bunrouter.New(bunrouter.Use(ErrorHandler))

	h := Handlers(db)

	router.GET("/", h.RouteIndex)
	router.GET("/api/v1/account/", h.RouteGetAccount) // require token
	router.POST("/api/v1/account/create/", h.RouteCreateUser)
	router.POST("/api/v1/account/login/", h.RouteLoginUser)
	router.POST("/api/v1/account/refresh/", h.RouteRefreshToken) // require token

	router.POST("/api/v1/files/upload/", h.RouteSaveFile) // require token
	router.GET("/api/v1/files/", h.RouteFilesList)

	log.Printf("Start server under :%d port...", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), router)
}
