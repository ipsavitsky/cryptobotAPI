package cryptobotAPI_test

import (
	"testing"

	crypto "github.com/ipsavitsky/cryptobotAPI"
	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
)

func TestGetCurrency(t *testing.T) {
	defer gock.Off()
	gock.New("https://pay.crypt.bot").
		Get("/api/getCurrencies").
		Reply(200).
		BodyString(`
		{
			"ok": true,
			"result": [
				{
					"is_blockchain": true,
					"is_stablecoin": false,
					"is_fiat": false,
					"name": "Toncoin",
					"code": "TON",
					"url": "https://ton.org/",
					"decimals": 9
				},
				{
					"is_blockchain": false,
					"is_stablecoin": true,
					"is_fiat": false,
					"name": "Tether",
					"code": "USDT",
					"url": "https://tether.to/",
					"decimals": 18
				},
				{
					"is_blockchain": false,
					"is_stablecoin": false,
					"is_fiat": true,
					"name": "Russian ruble",
					"code": "RUB",
					"decimals": 8
				},
				{
					"is_blockchain": false,
					"is_stablecoin": false,
					"is_fiat": true,
					"name": "United States dollar",
					"code": "USD",
					"decimals": 8
				}
			]
		}`)
	api := crypto.CryptoBotAPI{
		Options: crypto.APIOptions{
			Protocol: "https",
			Host:     "pay.crypt.bot",
		},
		Api_key: "123312",
	}
	resp, err := api.GetCurrencies()
	currency_array := *resp
	if assert.NoError(t, err) {
		assert.True(t, currency_array[0].Is_blockchain)
		assert.False(t, currency_array[0].Is_stablecoin)
		assert.False(t, currency_array[0].Is_fiat)
		assert.Equal(t, "Toncoin", currency_array[0].Name)
		assert.Equal(t, "TON", currency_array[0].Code)
		assert.Equal(t, "https://ton.org/", currency_array[0].Url)
		assert.Equal(t, 9, currency_array[0].Decimals)
		assert.False(t, currency_array[1].Is_blockchain)
		assert.True(t, currency_array[1].Is_stablecoin)
		assert.False(t, currency_array[1].Is_fiat)
		assert.Equal(t, "Tether", currency_array[1].Name)
		assert.Equal(t, "USDT", currency_array[1].Code)
		assert.Equal(t, "https://tether.to/", currency_array[1].Url)
		assert.Equal(t, 18, currency_array[1].Decimals)
		assert.False(t, currency_array[2].Is_blockchain)
		assert.False(t, currency_array[2].Is_stablecoin)
		assert.True(t, currency_array[2].Is_fiat)
		assert.Equal(t, "Russian ruble", currency_array[2].Name)
		assert.Equal(t, "RUB", currency_array[2].Code)
		assert.Equal(t, 8, currency_array[2].Decimals)
		assert.False(t, currency_array[3].Is_blockchain)
		assert.False(t, currency_array[3].Is_stablecoin)
		assert.True(t, currency_array[3].Is_fiat)
		assert.Equal(t, "United States dollar", currency_array[3].Name)
	}
}

func TestGetCurrencyError(t *testing.T) {
	defer gock.Off()
	gock.New("https://pay.crypt.bot").
		Get("/api/getCurrencies").
		Reply(200).
		BodyString(`
		{
			"ok":false,
			"error":{
				"code":1,
				"name":"mock error"
			}
		}`)
	api := crypto.CryptoBotAPI{
		Options: crypto.APIOptions{
			Protocol: "https",
			Host:     "pay.crypt.bot",
		},
		Api_key: "123312",
	}
	resp, err := api.GetCurrencies()
	if assert.Error(t, err) {
		assert.Equal(t, "1: mock error", err.Error())
		assert.Nil(t, resp)
	}
}

func TestGetBalance(t *testing.T) {
	defer gock.Off()
	gock.New("https://pay.crypt.bot").
		Get("/api/getBalance").
		Reply(200).
		BodyString(`
		{
			"ok": true,
			"result": [
				{
					"currency_code": "TON",
					"available": "15"
				},
				{
					"currency_code": "USDT",
					"available": "20"
				}
			]
		}`)
	api := crypto.CryptoBotAPI{
		Options: crypto.APIOptions{
			Protocol: "https",
			Host:     "pay.crypt.bot",
		},
		Api_key: "123312",
	}
	resp, err := api.GetBalance()
	currency_array := *resp
	if assert.NoError(t, err) {
		assert.Equal(t, "TON", currency_array[0].Currency_code)
		assert.Equal(t, "15", currency_array[0].Available)
		assert.Equal(t, "USDT", currency_array[1].Currency_code)
		assert.Equal(t, "20", currency_array[1].Available)
	}
}

func TestGetBalanceError(t *testing.T) {
	defer gock.Off()
	gock.New("https://pay.crypt.bot").
		Get("/api/getBalance").
		Reply(200).
		BodyString(`
		{
			"ok":false,
			"error":{
				"code":1,
				"name":"mock error"
			}
		}`)
	api := crypto.CryptoBotAPI{
		Options: crypto.APIOptions{
			Protocol: "https",
			Host:     "pay.crypt.bot",
		},
		Api_key: "123312",
	}
	resp, err := api.GetBalance()
	if assert.Error(t, err) {
		assert.Equal(t, "1: mock error", err.Error())
		assert.Nil(t, resp)
	}
}

func TestGetExchangeRates(t *testing.T) {
	defer gock.Off()
	gock.New("https://pay.crypt.bot").
		Get("/api/getExchangeRates").
		Reply(200).
		BodyString(`
		{
			"ok": true,
			"result": [
				{
					"is_valid": true,
					"source": "BTC",
					"target": "RUB",
					"rate": "3"
				},
				{
					"is_valid": true,
					"source": "BTC",
					"target": "USD",
					"rate": "5"
				}
			]
		}`)
	api := crypto.CryptoBotAPI{
		Options: crypto.APIOptions{
			Protocol: "https",
			Host:     "pay.crypt.bot",
		},
		Api_key: "123312",
	}
	resp, err := api.GetExchangeRates()
	currency_array := *resp
	if assert.NoError(t, err) {
		assert.Equal(t, "BTC", currency_array[0].Source)
		assert.Equal(t, "RUB", currency_array[0].Target)
		assert.Equal(t, "3", currency_array[0].Rate)
		assert.Equal(t, "BTC", currency_array[1].Source)
		assert.Equal(t, "USD", currency_array[1].Target)
		assert.Equal(t, "5", currency_array[1].Rate)
	}
}

func TestExchangeRatesError(t *testing.T) {
	defer gock.Off()
	gock.New("https://pay.crypt.bot").
		Get("/api/getExchangeRates").
		Reply(200).
		BodyString(`
		{
			"ok":false,
			"error":{
				"code":1,
				"name":"mock error"
			}
		}`)
	api := crypto.CryptoBotAPI{
		Options: crypto.APIOptions{
			Protocol: "https",
			Host:     "pay.crypt.bot",
		},
		Api_key: "123312",
	}
	resp, err := api.GetExchangeRates()
	if assert.Error(t, err) {
		assert.Equal(t, "1: mock error", err.Error())
		assert.Nil(t, resp)
	}
}
