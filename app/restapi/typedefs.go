package restapi

type createUserRequestPayload struct {
	Email           string `json:"email"`
	Password        string `json:"password"`
	PasswordConfirm string `json:"password_confirm"`
}

type loginUserRequestPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RefreshToken struct {
	Access string `json:"access"`
}

type ErrorResponse struct {
	Errors []string `json:"errors"`
}

type BooleanResponse struct {
	Status bool `json:"status"`
}
