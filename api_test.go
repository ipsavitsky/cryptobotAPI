package cryptobotAPI_test

import (
	"testing"

	crypto "github.com/ipsavitsky/cryptobotAPI"
	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
)

func TestGetMe(t *testing.T) {
	defer gock.Off()
	gock.New("https://pay.crypt.bot").
		Get("/api/getMe").
		Reply(200).
		BodyString(`
		{
			"ok": true,
			"result": {
				"app_id": 1,
				"name": "mock_app",
				"payment_processing_bot_username": "CryptoBot"
			}	
		}`)
	api := crypto.CryptoBotAPI{
		Options: crypto.APIOptions{
			Protocol: "https",
			Host:     "pay.crypt.bot",
		},
		Api_key: "123312",
	}
	resp, err := api.GetMe()
	if assert.NoError(t, err) {
		assert.Equal(t, 1, resp.App_id)
		assert.Equal(t, "mock_app", resp.Name)
		assert.Equal(t, "CryptoBot", resp.Payment_processing_bot_username)
	}
}

func TestGetMeError(t *testing.T) {
	defer gock.Off()
	gock.New("https://pay.crypt.bot").
		Get("/api/getMe").
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
	resp, err := api.GetMe()
	if assert.Error(t, err) {
		assert.Equal(t, "1: mock_error", err.Error())
		assert.Nil(t, resp)
	}
}
