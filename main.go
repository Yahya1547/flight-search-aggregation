package main

import (
	"log"
	"net/http"

	"flight-search-aggregation/api"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/flights/search", api.SearchFlightsHandler)
	mux.HandleFunc("/garudaindonesia/search", api.GarudaHandler)
	mux.HandleFunc("/lionair/search", api.LionAirHandler)
	mux.HandleFunc("/batikair/search", api.BatikAirHandler)
	mux.HandleFunc("/airasia/search", api.AirAsiaHandler)

	log.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}