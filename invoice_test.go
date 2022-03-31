package cryptobotAPI_test

import (
	"testing"

	crypto "github.com/ipsavitsky/cryptobotAPI"
	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
)

func TestCreateInvoice(t *testing.T) {
	defer gock.Off()
	gock.New("https://pay.crypt.bot").
		Get("/api/createInvoice").
		MatchHeader("Crypto-Pay-API-Token", "123312").
		MatchParam("amount", "5").
		MatchParam("asset", "TON").
		Reply(200).
		BodyString(`
		{
			"ok": true,
			"result": {
				"invoice_id": 1,
				"status": "active",
				"hash": "123",
				"asset": "TON",
				"amount": "5",
				"pay_url": "https://t.me/CryptoBot?start=123",
				"created_at": "1970-01-01T00:00:01.001Z",
				"allow_comments": true,
				"allow_anonymous": true
			}
		}`)
	api := crypto.CryptoBotAPI{
		Options: crypto.APIOptions{
			Protocol: "https",
			Host:     "pay.crypt.bot",
		},
		Api_key: "123312",
	}
	resp, err := api.CreateInvoice("TON", "5")
	if assert.NoError(t, err) {
		assert.Equal(t, 1, resp.Invoice_id)
		assert.Equal(t, "active", resp.Status)
		assert.Equal(t, "123", resp.Hash)
		assert.Equal(t, "TON", resp.Asset)
		assert.Equal(t, "5", resp.Amount)
		assert.Equal(t, "https://t.me/CryptoBot?start=123", resp.Pay_url)
		assert.Equal(t, "1970-01-01T00:00:01.001Z", resp.Created_at)
		assert.Equal(t, true, resp.Allow_comments)
		assert.Equal(t, true, resp.Allow_anonymous)
	}
}

func TestCreateInvoiceError(t *testing.T) {
	defer gock.Off()
	gock.New("https://pay.crypt.bot").
		Get("/api/createInvoice").
		MatchHeader("Crypto-Pay-API-Token", "123312").
		MatchParam("amount", "5").
		MatchParam("asset", "TON").
		Reply(200).
		BodyString(`
		{
			"ok":false,
			"error":{
				"code":1,
				"name":"mock_error"
			}
		}`)
	api := crypto.CryptoBotAPI{
		Options: crypto.APIOptions{
			Protocol: "https",
			Host:     "pay.crypt.bot",
		},
		Api_key: "123312",
	}
	resp, err := api.CreateInvoice("TON", "5")
	if assert.Error(t, err) {
		assert.Equal(t, "1: mock_error", err.Error())
		assert.Nil(t, resp)
	}
}

func TestGetInvoices(t *testing.T) {
	defer gock.Off()
	gock.New("https://pay.crypt.bot").
		Get("/api/getInvoices").
		MatchHeader("Crypto-Pay-API-Token", "123312").
		Reply(200).
		BodyString(`
		{
			"ok": true,
			"result": {
				"items": [
					{
						"invoice_id": 0,
						"status": "active",
						"hash": "123",
						"asset": "TON",
						"amount": "5",
						"pay_url": "https://t.me/CryptoBot?start=123",
						"created_at": "1970-01-01T00:00:01.001Z",
						"allow_comments": true,
						"allow_anonymous": true
					},
					{
						"invoice_id": 1,
						"status": "active",
						"hash": "321",
						"asset": "BNB",
						"amount": "10",
						"pay_url": "https://t.me/CryptoBot?start=321",
						"created_at": "1970-01-02T00:00:01.001Z",
						"allow_comments": true,
						"allow_anonymous": true
					}
				]
			}
		}`)
	api := crypto.CryptoBotAPI{
		Options: crypto.APIOptions{
			Protocol: "https",
			Host:     "pay.crypt.bot",
		},
		Api_key: "123312",
	}
	resp, err := api.GetInvoices()
	invoice_array := *resp
	if assert.NoError(t, err) {
		assert.Equal(t, 2, len(invoice_array))
		assert.Equal(t, 0, invoice_array[0].Invoice_id)
		assert.Equal(t, "active", invoice_array[0].Status)
		assert.Equal(t, "123", invoice_array[0].Hash)
		assert.Equal(t, "TON", invoice_array[0].Asset)
		assert.Equal(t, "5", invoice_array[0].Amount)
		assert.Equal(t, "https://t.me/CryptoBot?start=123", invoice_array[0].Pay_url)
		assert.Equal(t, "1970-01-01T00:00:01.001Z", invoice_array[0].Created_at)
		assert.True(t, invoice_array[0].Allow_comments)
		assert.True(t, invoice_array[0].Allow_anonymous)
		assert.Equal(t, 1, invoice_array[1].Invoice_id)
		assert.Equal(t, "active", invoice_array[1].Status)
		assert.Equal(t, "321", invoice_array[1].Hash)
		assert.Equal(t, "BNB", invoice_array[1].Asset)
		assert.Equal(t, "10", invoice_array[1].Amount)
		assert.Equal(t, "https://t.me/CryptoBot?start=321", invoice_array[1].Pay_url)
		assert.Equal(t, "1970-01-02T00:00:01.001Z", invoice_array[1].Created_at)
		assert.True(t, invoice_array[1].Allow_comments)
		assert.True(t, invoice_array[1].Allow_anonymous)
	}
}

func TestGetInvoicesError(t *testing.T) {
	defer gock.Off()
	gock.New("https://pay.crypt.bot").
		Get("/api/getInvoices").
		MatchHeader("Crypto-Pay-API-Token", "123312").
		Reply(200).
		BodyString(`
		{
			"ok":false,
			"error":{
				"code":1,
				"name":"mock_error"
			}
		}`)
	api := crypto.CryptoBotAPI{
		Options: crypto.APIOptions{
			Protocol: "https",
			Host:     "pay.crypt.bot",
		},
		Api_key: "123312",
	}
	resp, err := api.GetInvoices()
	if assert.Error(t, err) {
		assert.Equal(t, "1: mock_error", err.Error())
		assert.Nil(t, resp)
	}
}
