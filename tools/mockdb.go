package tools

import "time"

type mockDB struct {}

// Mock user login details
var mockLoginDetails = map[string]LoginDetails {
	"alex": {
		AuthToken: "123ABC",
		Username: "alex",
	},
	"jason": {
		AuthToken: "456DEF",
		Username: "jason",
	},
	"giselle": {
		AuthToken: "789GHI",
		Username: "giselle",
	},
}

// Mock user coin details
var mockCoinDetails = map[string]CoinDetails {
	"alex": {
		Coins: 100,
		Username: "alex",
	},
	"jason": {
		Coins: 250,
		Username: "jason",
	},
	"giselle": {
		Coins: 500,
		Username: "giselle",
	},
}

func (db *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	// Simulate db call
	time.Sleep(time.Second * 1)

	var clientData = LoginDetails{}
	clientData, ok := mockLoginDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (db *mockDB) GetUserCoins(username string) *CoinDetails {
	// Simulate db call
	time.Sleep(time.Second * 1)

	var clientData = CoinDetails{}
	clientData, ok := mockCoinDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (db *mockDB) SetupDatabase() error {
	return nil
}