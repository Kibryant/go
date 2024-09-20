package middleware

import (
	"errors"
	"net/http"

	"second_project/api"
	"second_project/internal/tools"

	log "github.com/sirupsen/logrus"
)

var UnauthorizedError = errors.New("Unauthorized")

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var username string = r.URL.Query().Get("username")
		var token string = r.Header.Get("Authorization")
		var err error

		if username == "" || token == "" {
			log.Error(UnauthorizedError)

			api.RequestErrorHandler(w, UnauthorizedError)

			return
		}

		var database *tools.Database

		database, err = tools.GetDatabase()

		if err != nil {
			log.Error(err)
			api.InternalServerErrorHandler(w)
		}

		var loginDetails *tools.LoginDetails

		loginDetails, err = (*database).GetUserLoginDetails(username)

		if (loginDetails == nil) || (token != loginDetails.Token) {
			log.Error(UnauthorizedError)
			api.RequestErrorHandler(w, UnauthorizedError)
			return
		}

		next.ServeHTTP(w, r)
	})
}
