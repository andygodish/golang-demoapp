package main

import (
	"log"
	"net/http"

	"github.com/andygodish/golang-demoapp/coinbase"
	"github.com/andygodish/golang-demoapp/server"
)

// This store really isn't necessary other than that it allows me to stub
// a store in my test in order to mock returns values from the coinbase API
type InMemoryStore struct{}

func (i *InMemoryStore) GetSellPrice() coinbase.Price {
	return coinbase.Price{}
}

// create an HTTP endpoint that returns the price of Bitcoin

func main() {
	pp := coinbase.NewPricePoplulator()
	server := server.NewServer(pp)
	log.Fatal(http.ListenAndServe(":8080", server))
}
