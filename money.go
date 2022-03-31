package cryptobotAPI

import (
	"encoding/json"
	"net/url"
)

// currency strucure used by GetCurrencies()
type Currency struct {
	Is_blockchain bool   `json:"is_blockchain"`
	Is_stablecoin bool   `json:"is_stablecoin"`
	Is_fiat       bool   `json:"is_fiat"`
	Name          string `json:"name"`
	Code          string `json:"code"`
	Url           string `json:"url"`
	Decimals      int    `json:"decimals"`
}

// Get all currencies accessible by your app
func (api *CryptoBotAPI) GetCurrencies() (*[]Currency, error) {
	url := &url.URL{
		Scheme: api.Options.Protocol,
		Host:   api.Options.Host,
		Path:   "/api/getCurrencies",
	}
	bts, err := api.getBytes("GET", url.String())
	if err != nil {
		return nil, err
	}
	var resp struct {
		Ok     bool       `json:"ok"`
		Result []Currency `json:"result"`
	}
	err = json.Unmarshal(bts, &resp)
	if err != nil {
		return nil, err
	}
	return &resp.Result, nil
}

// Balance structure. Contains currency ticker and how much you have available
type Balance struct {
	Currency_code string `json:"currency_code"`
	Available     string `json:"available"`
}

// Get balances available to your app
func (api *CryptoBotAPI) GetBalance() (*[]Balance, error) {
	url := &url.URL{
		Scheme: api.Options.Protocol,
		Host:   api.Options.Host,
		Path:   "/api/getBalance",
	}
	bts, err := api.getBytes("GET", url.String())
	if err != nil {
		return nil, err
	}
	var resp struct {
		Ok     bool      `json:"ok"`
		Result []Balance `json:"result"`
	}
	err = json.Unmarshal(bts, &resp)
	if err != nil {
		return nil, err
	}
	return &resp.Result, nil
}

// ExchangeRate structure
type ExchangeRate struct {
	Is_valid bool   `json:"is_valid"`
	Source   string `json:"source"`
	Target   string `json:"target"`
	Rate     string `json:"rate"`
}

// Get all current exchange rates
func (api *CryptoBotAPI) GetExchangeRates() (*[]ExchangeRate, error) {
	url := &url.URL{
		Scheme: api.Options.Protocol,
		Host:   api.Options.Host,
		Path:   "/api/getExchangeRates",
	}
	bts, err := api.getBytes("GET", url.String())
	if err != nil {
		return nil, err
	}
	var resp struct {
		Ok     bool           `json:"ok"`
		Result []ExchangeRate `json:"result"`
	}
	err = json.Unmarshal(bts, &resp)
	if err != nil {
		return nil, err
	}
	return &resp.Result, nil
}
