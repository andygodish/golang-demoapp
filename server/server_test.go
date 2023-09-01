package server

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/andygodish/golang-demoapp/coinbase"
	"github.com/stretchr/testify/assert"
)

// type StubBitcoinPriceStore struct {
// 	price coinbase.Price
// }

// func (s *StubBitcoinPriceStore) GetSellPrice() coinbase.Price {
// 	return s.price
// }

type MockPricePopulator struct {
	price coinbase.Price
}

func (m *MockPricePopulator) GetSellPrice() (coinbase.Price, error) {
	return m.price, nil
}

func TestGETBtcSellPrice(t *testing.T) {
	wantedPrice := coinbase.Price{
		Amount:   "1020.25",
		Currency: "USD",
	}

	mock := &MockPricePopulator{wantedPrice}
	server := NewServer(mock)

	t.Run("Returns the current sell price of bitcoin", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/prices/BTC-USD/sell", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		var got coinbase.Price

		err := json.NewDecoder(response.Body).Decode(&got)

		if err != nil {
			t.Fatalf("Unable to parse response from server %q into slice of Price, '%v'", response.Body, err)
		}

		assert.Equal(t, http.StatusOK, response.Code)
		assert.Equal(t, wantedPrice, got)
		assert.Equal(t, "application/json", response.Header().Get("content-type"))
	})
}
