package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/andygodish/golang-demoapp/coinbase"
	"github.com/andygodish/golang-demoapp/server"
)

type InMemoryPlayerStore struct {
}

func (i *InMemoryPlayerStore) GetSellPrice() string {
	return "1020.25"
}

// create an HTTP endpoint that returns the price of Bitcoin

func main() {
	pp := coinbase.NewPricePoplulator()

	price, err := pp.GetSellPrice()
	if err != nil {
		log.Fatalf("Failed to get BTC sell price: %v", err)
	}

	// Print the retrieved price
	fmt.Printf("BTC Sell Price: %s %s\n", price.Amount, price.Currency)

	server := &server.MyServer{
		Store: &InMemoryPlayerStore{},
	}
	log.Fatal(http.ListenAndServe(":5000", server))
}
