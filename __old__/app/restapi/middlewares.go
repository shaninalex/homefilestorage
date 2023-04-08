package restapi

import (
	"database/sql"
	"io"
	"net/http"

	"github.com/uptrace/bunrouter"
)

type HTTPError struct {
	StatusCode int `json:"status_code"`

	Code    string `json:"code"`
	Message string `json:"message"`
}

func (e HTTPError) Error() string {
	return e.Message
}

func NewHTTPError(err error) HTTPError {
	switch err {
	case io.EOF:
		return HTTPError{
			StatusCode: http.StatusBadRequest,

			Code:    "eof",
			Message: "EOF reading HTTP request body",
		}
	case sql.ErrNoRows:
		return HTTPError{
			StatusCode: http.StatusNotFound,

			Code:    "not_found",
			Message: "Not Found",
		}
	default:
		return HTTPError{
			StatusCode: http.StatusBadRequest,

			Code:    "bad_request",
			Message: err.Error(),
		}
	}

	// TODO: internal server error in case when we have not handler
	//		 for this error
	// return HTTPError{
	// 	StatusCode: http.StatusInternalServerError,

	// 	Code:    "internal",
	// 	Message: "Internal server error",
	// }
}

func ErrorHandler(next bunrouter.HandlerFunc) bunrouter.HandlerFunc {
	return func(w http.ResponseWriter, req bunrouter.Request) error {
		// Call the next handler on the chain to get the error.

		err := next(w, req)

		switch err := err.(type) {
		case nil:
			w.Header().Set("Content-Type", "application/json")
			// no error
		case HTTPError: // already a HTTPError
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(err.StatusCode)
			_ = bunrouter.JSON(w, err)
		default:
			w.Header().Set("Content-Type", "application/json")
			httpErr := NewHTTPError(err)
			w.WriteHeader(httpErr.StatusCode)
			_ = bunrouter.JSON(w, httpErr)
		}

		return err // return the err in case there other middlewares
	}
}
