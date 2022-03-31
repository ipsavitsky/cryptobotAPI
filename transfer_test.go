package cryptobotAPI_test

import (
	"testing"

	crypto "github.com/ipsavitsky/cryptobotAPI"
	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
)

func TestTransfer(t *testing.T) {
	defer gock.Off()
	gock.New("https://pay.crypt.bot").
		Get("/api/transfer").
		MatchParam("user_id", "1").
		MatchParam("asset", "BTC").
		MatchParam("amount", "0.1").
		MatchParam("spend_id", "123").
		Reply(200).
		BodyString(`
		{
			"ok": true,
			"result": {
				"transfer_id": 1,
				"user_id": "1",
				"asset": "BTC",
				"amount": "0.1",
				"status": "pending",
				"completed_at": "",
				"comment": ""
			}
		}`)
	api := crypto.CryptoBotAPI{
		Options: crypto.APIOptions{
			Protocol: "https",
			Host:     "pay.crypt.bot",
		},
	}
	resp, err := api.Transfer(1, "BTC", "0.1", "123")
	if assert.NoError(t, err) {
		assert.Equal(t, 1, resp.Transfer_id)
		assert.Equal(t, "1", resp.User_id)
		assert.Equal(t, "BTC", resp.Asset)
		assert.Equal(t, "0.1", resp.Amount)
		assert.Equal(t, "pending", resp.Status)
		assert.Equal(t, "", resp.Completed_at)
		assert.Equal(t, "", resp.Comment)
	}
}

func TestTransferError(t *testing.T) {
	defer gock.Off()
	gock.New("https://pay.crypt.bot").
		Get("/api/transfer").
		MatchParam("user_id", "1").
		MatchParam("asset", "BTC").
		MatchParam("amount", "0.1").
		MatchParam("spend_id", "123").
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
	}
	reqr, err := api.Transfer(1, "BTC", "0.1", "123")
	if assert.Error(t, err) {
		assert.Equal(t, "1: mock_error", err.Error())
		assert.Nil(t, reqr)
	}
}
