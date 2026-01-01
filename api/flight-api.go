package api

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"time"
	"math/rand"
	"os"

	"flight-search-aggregation/aggregator"
	"flight-search-aggregation/provider"
	"flight-search-aggregation/service"
	"flight-search-aggregation/utils"
)

func SearchFlightsHandler(response http.ResponseWriter, request *http.Request) {
	origin := request.URL.Query().Get("origin")
	destination := request.URL.Query().Get("destination")
	date := request.URL.Query().Get("date")
	maxPriceStr := request.URL.Query().Get("max_price")

	req := provider.SearchRequest{
		Origin:      origin,
		Destination: destination,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	providersList := []provider.AirlineProvider{
		&provider.LionAirProvider{},
		&provider.GarudaIndonesiaProvider{},
		&provider.AirAsiaProvider{},
		&provider.BatikAirProvider{},
	}

	flights, _ := aggregator.Aggregate(ctx, req, providersList)

	if maxPriceStr != "" {
		maxPrice, _ := strconv.ParseFloat(maxPriceStr, 64)
		flights = service.FilterByPrice(flights, maxPrice)
	}

	service.SortByPrice(flights, true)

	response.Header().Set("Content-Type", "application/json")
	json.NewEncoder(response).Encode(flights)
}

func GarudaHandler(response http.ResponseWriter, request *http.Request) {
	utils.RandomDelay(50, 100)

	rdata, _ := os.ReadFile("../mock/garuda_indonesia_search_response.json")

	var garudaIndonesiaResponse models.GarudaIndonesiaResponse
	json.Unmarshal(rdata, &garudaIndonesiaResponse)

	json.NewEncoder(response).Encode(garudaIndonesiaResponse)
}

func LionAirHandler(response http.ResponseWriter, request *http.Request) {
	utils.RandomDelay(100, 200)

	data, _ := os.ReadFile("../mock/lion_air_search_response.json")

	var lionAirResponse models.LionAirResponse
	json.Unmarshal(data, &lionAirResponse)

	json.NewEncoder(response).Encode(lionAirResponse)
}

func BatikAirHandler(response http.ResponseWriter, request *http.Request) {
	utils.RandomDelay(200, 400)

	data, _ := os.ReadFile("../mock/batik_air_search_response.json")

	var batikAirResponse models.BatikAirResponse
	json.Unmarshal(data, &batikAirResponse)

	json.NewEncoder(response).Encode(batikAirResponse)
}

func AirAsiaHandler(response http.ResponseWriter, request *http.Request) {
	utils.RandomDelay(50, 150)

	// 90% success rate
	if rand.Float64() > 0.9 {
		http.Error(response, "AirAsia service unavailable", http.StatusServiceUnavailable)
		return
	}

	data, _ := os.ReadFile("../mock/airasia_search_response.json")

	var airAsiaResponse models.AirAsiaResponse
	json.Unmarshal(data, &airAsiaResponse)

	json.NewEncoder(response).Encode(airAsiaResponse)
}