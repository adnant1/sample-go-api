package middleware

import (
	"errors"
	"net/http"

	"github.com/adnant1/sample-go-api/api"
	"github.com/adnant1/sample-go-api/internal/tools"
	log "github.com/sirupsen/logrus"
)

// Global unauthorized error
var UnAuthorizedError = errors.New("Invalid username or token.")

// Authorization middleware
func Authorization(next http.Handler) http.Handler {
	// Return a handler function
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		
		var username string = r.URL.Query().Get("username")
		var token = r.Header.Get("Authorization")
		var err error
		
		// Check if username and token are valid
		if username == "" || token == "" {
			log.Error(UnAuthorizedError)
			api.RequestErrorHandler(w, UnAuthorizedError)
			return
		}

		// Query the database for user login details
		var database *tools.DatabaseInterface
		database, err = tools.NewDatabase()
		if err != nil {
			api.InternalErrorHandler(w)
			return
		}

		var loginDetails *tools.LoginDetails
		loginDetails = (*database).GetUserLoginDetails(username)

		if (loginDetails == nil || (token != (*loginDetails).AuthToken)) {
			log.Error(UnAuthorizedError)
			api.RequestErrorHandler(w, UnAuthorizedError)
			return
		}

		// Call the next handler if authorized
		next.ServeHTTP(w, r)
	})
}