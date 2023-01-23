package restapi

import (
	"encoding/json"
	"io"
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

type createUserRequestPayload struct {
	Email           string `json:"email"`
	Password        string `json:"password"`
	PasswordConfirm string `json:"password_confirm"`
}

type ErrorResponse struct {
	Errors []string `json:"errors"`
}

func (h *BaseHandler) RouteCreateUser(w http.ResponseWriter, req bunrouter.Request) error {

	data, err := io.ReadAll(req.Body)
	if err != nil {
		error_data := ErrorResponse{Errors: []string{"Cannot parse request"}}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(error_data)
		return nil
	}

	// p := createUserRequestPayload{Email: "email@test.com", Password: "123", Password_confirm: "123"}
	p := createUserRequestPayload{}
	json.Unmarshal(data, &p)

	if p.Password != p.PasswordConfirm {
		error_data := ErrorResponse{Errors: []string{"Password and password confirm are not the same."}}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(error_data)
		return nil
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(p)

	// - validate request
	// payload := database.CreateUserPayload{Email: "test@test.com", HashedPassword: "secret"}
	// - save user
	// database.CreateUser(&payload, h.db)

	// RabbitMQ
	// - send email to admin ( new user created )
	// - send email to new user ( confirm your account )
	return nil
}
