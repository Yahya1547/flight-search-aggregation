package api

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"flight-search-aggregation/aggregator"
	"flight-search-aggregation/provider"
	"flight-search-aggregation/service"
)

func SearchFlightsHandler(response http.ResponseWriter, request *http.Request) {
	origin := request.URL.Query().Get("origin")
	destination := request.URL.Query().Get("destination")
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