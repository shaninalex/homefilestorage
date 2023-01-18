package restapi

import (
	"net/http"

	"github.com/uptrace/bunrouter"
)

func RouteIndex(w http.ResponseWriter, req bunrouter.Request) error {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Read documentation..."))
	return nil
}

type createUserRequestPayload struct {
	email            string
	password         string
	password_confirm string
}

func RouteCreateUser(w http.ResponseWriter, req bunrouter.Request) error {
	// - validate request
	// - save user

	// RabbitMQ
	// - send email to admin ( new user created )
	// - send email to new user ( confirm your account )
	return nil
}
