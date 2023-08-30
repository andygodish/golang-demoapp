package coinbase

import (
	"encoding/json"
	"errors"
	"fmt"

	"net/http"
)

const (
	defaultBaseURL = "https://api.coinbase.com/v2"
)

var ErrFailedAPICall = errors.New("bad response from API")

type CoinbasePriceResponse struct {
	Data Price `json:"data"`
}

type Price struct {
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
}

type PricePopulator struct {
	apiURL string
}

func (pp *PricePopulator) GetSellPrice() (Price, error) {
	res, err := http.Get(btcPriceCoinbaseURL(pp.apiURL))
	if err != nil {
		return Price{}, fmt.Errorf("failed to get releases: %v: %w", err, ErrFailedAPICall)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return Price{}, fmt.Errorf("unexpected status code: %d - %s: %w", res.StatusCode, res.Status, ErrFailedAPICall)
	}

	var response CoinbasePriceResponse
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return Price{}, fmt.Errorf("failed to decode the body: %w", err)
	}

	price := response.Data

	return price, nil
}

func NewPricePoplulator() *PricePopulator {
	return &PricePopulator{apiURL: defaultBaseURL}
}

func btcPriceCoinbaseURL(apiURL string) string {
	return fmt.Sprintf("%s/prices/BTC-USD/sell", apiURL)
}
