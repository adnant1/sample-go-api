package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/adnant1/sample-go-api/api"
	"github.com/adnant1/sample-go-api/internal/tools"
	"github.com/gorilla/schema"
	log "github.com/sirupsen/logrus"
)

// GetCoinBalance handles the GET /account/coins endpoint
func GetCoinBalance(w http.ResponseWriter, r *http.Request) {
	var params = api.CoinBalanceParams{}
	var decoder *schema.Decoder = schema.NewDecoder()
	var err error

	// Parse query parameters
	err = decoder.Decode(&params, r.URL.Query())
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	// Instantiate database connection
	var database *tools.DatabaseInterface
	database, err = tools.NewDatabase()
	if err != nil {
		api.InternalErrorHandler(w)
		return
	}

	// Query the database for user coin details
	var coinDetails *tools.coinDetails
	coinDetails = (*database).GetUserCoins(params.Username)
	if coinDetails == nil {
		api.InternalErrorHandler(w)
		return
	}

	// Create response
	var res = api.CoinBalanceReponse{
		Code: http.StatusOK,
		Balance: (*coinDetails).Coins,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}
}