package restapi

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"homestorage/app/database"
	"io"
	"net/http"

	"github.com/uptrace/bunrouter"
)

type BaseHandler struct {
	db *sql.DB
}

func Handlers(db *sql.DB) *BaseHandler {
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

	fmt.Println("this is test message")

	// - validate request
	payload := database.CreateUserPayload{Email: p.Email, HashedPassword: p.Password}
	// - save user
	database.CreateUser(&payload, h.db)

	// RabbitMQ
	// - send email to admin ( new user created )
	// - send email to new user ( confirm your account )
	return nil
}
