package cryptobotAPI

import (
	"encoding/json"
	"net/url"
	"strconv"
)

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
	url := &url.URL{
		Scheme: api.Options.Protocol,
		Host:   api.Options.Host,
		Path:   "/api/transfer",
	}
	q := url.Query()
	q.Set("user_id", strconv.Itoa(user_id))
	q.Set("asset", asset)
	q.Set("amount", amount)
	q.Set("spend_id", spend_id)
	url.RawQuery = q.Encode()
	bts, err := api.getBytes("GET", url.String())
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
