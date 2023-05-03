package main

import (
	"encoding/json"
	"net/http"
)

type JsonResponsePayload struct {
	Status string      `json:"status"` // success, error
	Data   interface{} `json:"data"`
}

type JsonFrontRequestPayload struct {
	Action string      `json:"action"`
	Data   interface{} `json:"data"`
}

func WriteJSON(w http.ResponseWriter, status int, data interface{}, headers ...http.Header) error {
	out, err := json.Marshal(data)
	if err != nil {
		return err
	}

	if len(headers) > 0 {
		for key, value := range headers[0] {
			w.Header()[key] = value
		}
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	_, err = w.Write(out)
	if err != nil {
		return err
	}

	return nil
}

func ReadTypedJSON(w http.ResponseWriter, r *http.Request, data JsonFrontRequestPayload) error {
	err := ReadJSON(w, r, data)
	if err != nil {
		return err
	}

	return nil
}

func ReadJSON(w http.ResponseWriter, r *http.Request, data any) error {
	maxBytes := 1048576 // Max payload size 1 Mb

	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	dec := json.NewDecoder(r.Body)
	err := dec.Decode(data)
	if err != nil {
		return err
	}

	return nil
}

func ErrorJSON(w http.ResponseWriter, err error, status ...int) error {
	statusCode := http.StatusBadRequest

	// if a custom response code is specified, use that instead of bad request
	if len(status) > 0 {
		statusCode = status[0]
	}

	payload := JsonResponsePayload{
		Status: "error",
		Data:   err.Error(),
	}

	return WriteJSON(w, statusCode, payload)
}
