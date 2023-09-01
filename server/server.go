package server

import (
	"encoding/json"
	"net/http"

	"github.com/andygodish/golang-demoapp/coinbase"
)

type BitcoinPriceStore interface {
	GetSellPrice() coinbase.Price
}

type MyServer struct {
	Store        BitcoinPriceStore
	http.Handler // the server now has all the methods of http.Handler embbded in it
}

func NewServer(store BitcoinPriceStore) *MyServer {
	s := new(MyServer)

	router := http.NewServeMux()

	router.Handle("/prices/BTC-USD/sell", http.HandlerFunc(s.BtcSellPriceHandler))
	s.Handler = router

	return s
}

func (m *MyServer) BtcSellPriceHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(m.getSellPrice())
}

func (m *MyServer) getSellPrice() coinbase.Price {
	// pp := coinbase.NewPricePoplulator()
	// price, err := pp.GetSellPrice()
	// if err != nil {
	// 	log.Fatalf("Failed to get BTC sell price: %v", err)
	// }
	// return price
	return coinbase.Price{
		Amount:   "1020.25",
		Currency: "USD",
	}
}
