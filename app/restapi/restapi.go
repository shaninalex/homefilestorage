package restapi

import (
	"errors"
	"fmt"
	"homestorage/app/database"
	"homestorage/app/filesystem"
	"log"
	"net/http"

	"github.com/uptrace/bunrouter"
)

type BaseHandler struct {
	db      *database.DatabaseRepository
	storage *filesystem.FileStorage
}

func Handlers(db *database.DatabaseRepository, path string) *BaseHandler {
	storage, err := filesystem.CreateFileStorage(path)
	if err != nil {
		log.Println(err)
	}
	return &BaseHandler{
		db:      db,
		storage: storage,
	}
}

// TODO i18n errors texts
var (
	ErrParse                       = errors.New("cannot parse request")
	ErrRegistrationPasswordConfirm = errors.New("password and password confirm are not the same")
	ErrPasswordDoesNotMatch        = errors.New("password does not match")
	ErrCantRafreshToken            = errors.New("cant refresh token")
)

func Server(db *database.DatabaseRepository, storage_path string, port int) {

	router := bunrouter.New(bunrouter.Use(ErrorHandler))

	h := Handlers(db, storage_path)

	router.GET("/", h.RouteIndex)
	router.GET("/api/v1/account/", h.RouteGetAccount)
	router.POST("/api/v1/account/create/", h.RouteCreateUser) // public
	router.POST("/api/v1/account/login/", h.RouteLoginUser)   // public
	router.POST("/api/v1/account/refresh/", h.RouteRefreshToken)

	router.POST("/api/v1/files/upload/", h.RouteSaveFile)
	router.GET("/api/v1/files/", h.RouteFilesList)
	router.GET("/api/v1/file/:id/", h.RouteFileItem)

	// public ( shared links )

	log.Printf("Start server under :%d port...", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), router)
}
