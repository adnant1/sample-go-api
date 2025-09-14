package api

import (
	"encoding/json"
	"net/http"
)

// Coin Balance params
type CoinBalanceParams struct {
	Username string
}

// Coin Balance response
type CoinBalanceResponse struct {
	Code int // Success code, usually 200
	Balance int64 // Account balance
}

// Error reponse
type Error struct {
	Code int // Error code
	Message string // Error message
}

func writeError(w http.ResponseWriter, message string, code int) {
	res := Error {
		Code: code,
		Message: message,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	
	json.NewEncoder(w).Encode(res)
}