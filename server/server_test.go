package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

type StubBitcoinPriceStore struct {
	price string
}

func (s *StubBitcoinPriceStore) GetSellPrice() string {
	return s.price
}

func TestGETBtcSellPrice(t *testing.T) {
	store := &StubBitcoinPriceStore{price: "1020.25"}
	server := &MyServer{
		Store: store,
	}
	t.Run("Returns the current sell price of bitcoin", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/prices/BTC-USD/sell", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assert.Equal(t, "1020.25", response.Body.String())
	})
}
