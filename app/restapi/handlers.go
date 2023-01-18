package restapi

import (
	"homestorage/app/database"
	"net/http"

	"github.com/uptrace/bunrouter"
	"gorm.io/gorm"
)

type BaseHandler struct {
	db *gorm.DB
}

func Handlers(db *gorm.DB) *BaseHandler {
	return &BaseHandler{
		db: db,
	}
}

func (h *BaseHandler) RouteIndex(w http.ResponseWriter, req bunrouter.Request) error {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Read documentation..."))
	return nil
}

func (h *BaseHandler) RouteCreateUser(w http.ResponseWriter, req bunrouter.Request) error {
	// - validate request
	payload := database.CreateUserPayload{Email: "test@test.com", HashedPassword: "secret"}
	// - save user
	database.CreateUser(&payload, h.db)

	// RabbitMQ
	// - send email to admin ( new user created )
	// - send email to new user ( confirm your account )
	return nil
}

// type createUserRequestPayload struct {
// 	email            string
// 	password         string
// 	password_confirm string
// }
