package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
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
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("%s", body)
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
	bts, err := api.getBytes("GET", fmt.Sprintf("%s://%s/api/getMe", api.Options.Protocol, api.Options.Host))
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

type Invoice struct {
	Ok     bool `json:"ok"`
	Result struct {
		Invoice_id      int    `json:"invoice_id"`
		Status          string `json:"status"`
		Hash            string `json:"hash"`
		Asset           string `json:"asset"`
		Amount          string `json:"amount"`
		Pay_url         string `json:"pay_url"`
		Description     string `json:"description"`
		Created_at      string `json:"created_at"`
		Allow_comments  bool   `json:"allow_comments"`
		Allow_anonymous bool   `json:"allow_anonymous"`
		Expiration_date string `json:"expiration_date"`
		Paid_at         string `json:"paid_at"`
		Paid_anonymous  bool   `json:"paid_anonymous"`
		Comment         string `json:"comment"`
		Hidden_message  string `json:"hidden_message"`
		Payload         string `json:"payload"`
		Paid_btn_name   string `json:"paid_btn_name"`
		Paid_btn_url    string `json:"paid_btn_url"`
	} `json:"result"`
}

func (api *CryptoBotAPI) GetInvoices() (*[]Invoice, error) {
	bts, err := api.getBytes("GET", fmt.Sprintf("%s://%s/api/getInvoices", api.Options.Protocol, api.Options.Host))
	if err != nil {
		return nil, err
	}
	var resp struct {
		Ok     bool `json:"ok"`
		Result struct {
			Items []Invoice `json:"items"`
		} `json:"result"`
	}
	err = json.Unmarshal(bts, &resp)
	if err != nil {
		return nil, err
	}
	return &resp.Result.Items, nil
}

type Currency struct {
	Is_blockchain bool   `json:"is_blockchain"`
	Is_stablecoin bool   `json:"is_stablecoin"`
	Is_fiat       bool   `json:"is_fiat"`
	Name          string `json:"name"`
	Code          string `json:"code"`
	Url           string `json:"url"`
	Decimals      int    `json:"decimals"`
}

func (api *CryptoBotAPI) GetCurrencies() (*[]Currency, error) {
	bts, err := api.getBytes("GET", fmt.Sprintf("%s://%s/api/getCurrencies", api.Options.Protocol, api.Options.Host))
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

type Balance struct {
	Currency_code string `json:"currency_code"`
	Available     string `json:"available"`
}

func (api *CryptoBotAPI) GetBalance() (*[]Balance, error) {
	bts, err := api.getBytes("GET", fmt.Sprintf("%s://%s/api/getBalance", api.Options.Protocol, api.Options.Host))
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

type ExchangeRate struct {
	Is_valid bool   `json:"is_valid"`
	Source   string `json:"source"`
	Target   string `json:"target"`
	Rate     string `json:"rate"`
}

func (api *CryptoBotAPI) GetExchangeRates() (*[]ExchangeRate, error) {
	bts, err := api.getBytes("GET", fmt.Sprintf("%s://%s/api/getExchangeRates", api.Options.Protocol, api.Options.Host))
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

type Transfer struct {
	Transfer_id  int    `json:"transfer_id"`
	User_id      string `json:"user_id"`
	Asset        string `json:"asset"`
	Amount       string `json:"amount"`
	Status       string `json:"status"`
	Completed_at string `json:"completed_at"`
	Comment      string `json:"comment"`
}

func (api *CryptoBotAPI) Transfer(user_id int, asset string, amount string, spend_id string) (*Transfer, error) {
	bts, err := api.getBytes("GET", fmt.Sprintf("%s://%s/api/transfer?user_id=%d&asset=%s&amount=%s&spend_id=%s", api.Options.Protocol, api.Options.Host, user_id, asset, amount, spend_id))
	if err != nil {
		return nil, err
	}
	var resp Transfer
	err = json.Unmarshal(bts, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (api *CryptoBotAPI) CreateInvoice(asset string, amount string) (*Invoice, error) {
	bts, err := api.getBytes("GET", fmt.Sprintf("%s://%s/api/createInvoice?asset=%s&amount=%s", api.Options.Protocol, api.Options.Host, asset, amount))
	if err != nil {
		return nil, err
	}
	var resp Invoice
	err = json.Unmarshal(bts, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
