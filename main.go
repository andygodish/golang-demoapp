package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/andygodish/golang-demoapp/coinbase"
	"github.com/andygodish/golang-demoapp/server"
)

func main() {
	fmt.Println("Starting server on port 8080")
	
	pp := coinbase.NewPricePoplulator()
	server := server.NewServer(pp)
	log.Fatal(http.ListenAndServe(":8080", server))
}
