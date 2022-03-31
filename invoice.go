package cryptobotAPI

import (
	"encoding/json"
	"fmt"
	"net/url"
)

// An invoice structure. Contains invoice id, amount, asset, status, etc.
type Invoice struct {
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
}

// Gets all your current invoices
func (api *CryptoBotAPI) GetInvoices() (*[]Invoice, error) {
	url := &url.URL{
		Scheme: api.Options.Protocol,
		Host:   api.Options.Host,
		Path:   "/api/getInvoices",
	}
	bts, err := api.getCryptoBotResponse("GET", url.String())
	if err != nil {
		return nil, err
	}
	var resp struct {
		Ok     bool `json:"ok"`
		Result struct {
			Items []Invoice `json:"items"`
		} `json:"result,omitempty"`
		Error struct {
			Code int    `json:"code"`
			Name string `json:"name"`
		} `json:"error,omitempty"`
	}
	err = json.Unmarshal(bts, &resp)
	if err != nil {
		return nil, err
	}
	if !resp.Ok {
		return nil, fmt.Errorf("%d: %s", resp.Error.Code, resp.Error.Name)
	}
	return &resp.Result.Items, nil
}

// Create a new invoice
func (api *CryptoBotAPI) CreateInvoice(asset string, amount string) (*Invoice, error) {
	url := &url.URL{
		Scheme: api.Options.Protocol,
		Host:   api.Options.Host,
		Path:   "/api/createInvoice",
	}
	q := url.Query()
	q.Set("asset", asset)
	q.Set("amount", amount)
	url.RawQuery = q.Encode()
	bts, err := api.getCryptoBotResponse("GET", url.String())
	if err != nil {
		return nil, err
	}
	var resp struct {
		Ok     bool    `json:"ok"`
		Result Invoice `json:"result,omitempty"`
		Error  struct {
			Code int    `json:"code"`
			Name string `json:"name"`
		} `json:"error,omitempty"`
	}
	err = json.Unmarshal(bts, &resp)
	if err != nil {
		return nil, err
	}
	if !resp.Ok {
		return nil, fmt.Errorf("%d: %s", resp.Error.Code, resp.Error.Name)
	}
	return &resp.Result, nil
}
