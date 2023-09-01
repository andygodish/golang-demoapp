package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/andygodish/golang-demoapp/coinbase"
	"github.com/andygodish/golang-demoapp/server"
)

// This store really isn't necessary other than that it allows me to stub 
// a store in my test in order to mock returns values from the coinbase API
type InMemoryStore struct {}
func (i *InMemoryStore) GetSellPrice() coinbase.Price {
	return coinbase.Price{}
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

	server := server.NewServer(pp)

	log.Fatal(http.ListenAndServe(":5000", server))
}
