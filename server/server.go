package server

import (
	"fmt"
	"net/http"
)

type BitcoinPriceStore interface {
	GetSellPrice() string
}

type MyServer struct {
	Store BitcoinPriceStore
}

func (m *MyServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, m.Store.GetSellPrice())
}
