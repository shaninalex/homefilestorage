package restapi

import (
	"encoding/json"
	"homestorage/app/database"
	"homestorage/app/utils"
	"io"
	"log"
	"net/http"

	"github.com/uptrace/bunrouter"
)

type BaseHandler struct {
	db *database.DatabaseRepository
}

type createUserRequestPayload struct {
	Email           string `json:"email"`
	Password        string `json:"password"`
	PasswordConfirm string `json:"password_confirm"`
}

type ErrorResponse struct {
	Errors []string `json:"errors"`
}

func Handlers(db *database.DatabaseRepository) *BaseHandler {
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

	data, err := io.ReadAll(req.Body)
	if err != nil {
		return ErrParse
	}

	p := createUserRequestPayload{}
	json.Unmarshal(data, &p)

	if p.Password != p.PasswordConfirm {
		return ErrRegistrationPasswordConfirm
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(p)

	// - validate request
	payload := database.CreateUserPayload{Email: p.Email, HashedPassword: p.Password}
	// - save user
	encodedHash, err := utils.GenerateFromPassword(payload.HashedPassword)
	if err != nil {
		log.Fatal(err)
	}
	payload.HashedPassword = encodedHash
	h.db.CreateUser(&payload)

	// RabbitMQ
	// - send email to admin ( new user created )
	// - send email to new user ( confirm your account )
	return nil
}
