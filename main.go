package main

import (
	"fmt"
	"log"
)

// create an HTTP endpoint that returns the price of Bitcoin

func main() {
	pp := NewPricePoplulator()

	price, err := pp.GetSellPrice()
	if err != nil {
		log.Fatalf("Failed to get BTC sell price: %v", err)
	}

	// Print the retrieved price
	fmt.Printf("BTC Sell Price: %s %s\n", price.Amount, price.Currency)
}
