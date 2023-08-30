package coinbase

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetBtcSellPrice(t *testing.T) {
	testBtcPrice := struct {
		cbStatusCode int
		cbBody       string
	}{
		cbStatusCode: http.StatusOK,
		cbBody: `{
  			"data": {
    			"amount": "1020.25",
    			"currency": "USD"
  			}
		}`,
	}

	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(testBtcPrice.cbStatusCode)
		_, err := w.Write([]byte(testBtcPrice.cbBody))
		require.NoError(t, err)
	}))
	defer testServer.Close()

	pp := NewPricePoplulator()
	pp.apiURL = testServer.URL

	gotPrice, _ := pp.GetSellPrice()

	assert.Equal(t, "1020.25", gotPrice.Amount)
	assert.Equal(t, "USD", gotPrice.Currency)
}


