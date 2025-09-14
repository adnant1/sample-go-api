package tools

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