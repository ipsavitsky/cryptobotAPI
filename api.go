package cryptobotAPI

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

type APIOptions struct {
	Protocol string
	Host     string
}

type CryptoBotAPI struct {
	Options APIOptions
	Api_key string
}

func doRequest(req *http.Request) ([]byte, error) {
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, err
	}
	return body, nil
}

func (api *CryptoBotAPI) getBytes(method string, url string) ([]byte, error) {
	req, err := http.NewRequest(method, url, nil)
	req.Header.Set("Crypto-Pay-API-Token", api.Api_key)
	if err != nil {
		return nil, err
	}
	return doRequest(req)
}

type GetMeResponse struct {
	App_id                          int    `json:"app_id"`
	Name                            string `json:"name"`
	Payment_processing_bot_username string `json:"payment_processing_bot_username"`
}

func (api *CryptoBotAPI) GetMe() (*GetMeResponse, error) {
	url := &url.URL{
		Scheme: api.Options.Protocol,
		Host:   api.Options.Host,
		Path:   "/api/getMe",
	}
	bts, err := api.getBytes("GET", url.String())
	if err != nil {
		return nil, err
	}
	var resp struct {
		Ok     bool          `json:"ok"`
		Result GetMeResponse `json:"result"`
	}
	err = json.Unmarshal(bts, &resp)
	if err != nil {
		return nil, err
	}
	return &resp.Result, nil
}
