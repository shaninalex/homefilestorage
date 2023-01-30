package restapi

import (
	"encoding/json"
	"homestorage/app/database"
	"homestorage/app/utils"
	"io"
	"net/http"
	"strings"

	"github.com/uptrace/bunrouter"
)

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

	// - validate request
	payload := database.CreateUserPayload{Email: p.Email, HashedPassword: p.Password}
	// - save user
	encodedHash, err := utils.GenerateFromPassword(payload.HashedPassword)
	if err != nil {
		return err
	}
	payload.HashedPassword = encodedHash
	err = h.db.CreateUser(&payload)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&BooleanResponse{Status: true})

	// RabbitMQ
	// - send email to admin ( new user created )
	// - send email to new user ( confirm your account )
	return nil
}

func (h *BaseHandler) RouteLoginUser(w http.ResponseWriter, req bunrouter.Request) error {
	data, err := io.ReadAll(req.Body)
	if err != nil {
		return ErrParse
	}

	p := loginUserRequestPayload{}
	json.Unmarshal(data, &p)

	user, err := h.db.GetUserByEmail(p.Email)
	if err != nil {
		return err
	}

	match, err := utils.CheckPassword(p.Password, user.Hashed_password)
	if err != nil {
		return err
	}

	if !match {
		return ErrPasswordDoesNotMatch
	}

	token, err := utils.GenerateJWT(user.Email, user.Id)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(token)

	return nil
}

func (h *BaseHandler) RouteRefreshToken(w http.ResponseWriter, req bunrouter.Request) error {
	data, err := io.ReadAll(req.Body)
	if err != nil {
		return ErrParse
	}

	p := RefreshToken{}
	json.Unmarshal(data, &p)

	access_credentials, err := utils.RefreshJWT(p.Access)
	if err != nil {
		return ErrCantRafreshToken
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(access_credentials)
	return nil
}

func (h *BaseHandler) RouteGetAccount(w http.ResponseWriter, req bunrouter.Request) error {
	token := req.Header.Get("Authorization")
	_, id, err := utils.IdentifyJWT(strings.Replace(token, "Bearer ", "", 1))
	if err != nil {
		return err
	}

	user, err := h.db.GetUser(id)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(utils.PublicUser(user))
	return nil
}
