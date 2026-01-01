package main

import (
	"log"
	"net/http"

	"flight-search-aggregation/api"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/flights/search", api.SearchFlightsHandler)

	log.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}