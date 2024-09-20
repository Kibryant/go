package api

import (
	"encoding/json"
	"net/http"
)

type CoinBalanceParams struct {
	Username string
}

type CoinBalanceResponse struct {
	StatusCode int
	Balance    int64
}

type Error struct {
	StatusCode int
	Message    string
}

func writeError(w http.ResponseWriter, StatusCode int, message string) {
	error := Error{
		StatusCode: StatusCode,
		Message:    message,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(StatusCode)

	json.NewEncoder(w).Encode(error)
}

var (
	RequestErrorHandler = func(w http.ResponseWriter, err error) {
		writeError(w, http.StatusNotFound, err.Error())
	}

	InternalServerErrorHandler = func(w http.ResponseWriter) {
		writeError(w, http.StatusInternalServerError, "Internal Server Error")
	}
)
