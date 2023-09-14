package server

import (
	"encoding/json"
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
	sellPrice, err := m.PricePopulator.GetSellPrice()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(sellPrice)
}
