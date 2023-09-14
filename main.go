package main

import (
	"log"
	"net/http"

	"github.com/andygodish/golang-demoapp/coinbase"
	"github.com/andygodish/golang-demoapp/server"
)

func main() {
	pp := coinbase.NewPricePoplulator()
	server := server.NewServer(pp)
	log.Fatal(http.ListenAndServe(":8080", server))
}
