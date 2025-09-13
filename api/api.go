package api

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
