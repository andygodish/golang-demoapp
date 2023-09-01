package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/andygodish/golang-demoapp/coinbase"
)

type PricePopulator interface {
	GetSellPrice() (coinbase.Price, error)
}

type MyServer struct {
	PricePopulator
	http.Handler // the server now has all the methods of http.Handler embbded in it
}

func NewServer(pp PricePopulator) *MyServer {
	s := &MyServer{
		PricePopulator: pp,
	}

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
	price, err := m.GetSellPrice()
	if err != nil {
		log.Fatalf("Failed to get BTC sell price: %v", err)
	}
	return price
}
