package api

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"time"
	"strings"

	"flight-search-aggregation/aggregator"
	"flight-search-aggregation/provider"
	"flight-search-aggregation/service"
	"flight-search-aggregation/models"
)

func SearchFlightsHandler(response http.ResponseWriter, request *http.Request) {
	origin := request.URL.Query().Get("origin")
	destination := request.URL.Query().Get("destination")
	date := request.URL.Query().Get("date")
	passengers := request.URL.Query().Get("passengers")
	cabinClass := request.URL.Query().Get("cabin_class")
	airlines := request.URL.Query().Get("airlines")
	numberOfStops := request.URL.Query().Get("number_of_stops")
	maxPriceStr := request.URL.Query().Get("max_price")
	minPriceStr := request.URL.Query().Get("min_price")
	maxDurationStr := request.URL.Query().Get("max_duration")
	minDurationStr := request.URL.Query().Get("min_duration")
	sortByStr := request.URL.Query().Get("sort_by")
	sortDirectionStr := request.URL.Query().Get("sort_direction")

	numberOfStopFilter := []int{}
	if (numberOfStops != "") { 
		for _, stopStr := range strings.Split(numberOfStops, ",") {
			stopInt, err := strconv.Atoi(stopStr)
			if err == nil {
				numberOfStopFilter = append(numberOfStopFilter, stopInt)
			}
		}
	}

	passengersInt, _ := strconv.Atoi(passengers)
	req := provider.SearchRequest{
		Origin:      origin,
		Destination: destination,
		DepartureDate: date,
		Passengers:  passengersInt,
		CabinClass: cabinClass,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	airlineFilterList := []string{}
	if airlines != "" {
		airlineFilterList = strings.Split(airlines, ",")
	}
	providersList := service.GetAirlineProvidersFilterByAirlines(airlineFilterList)
	aggregatedFlight, _ := aggregator.Aggregate(ctx, req, providersList)
	flights := aggregatedFlight.Flights

	if maxPriceStr != "" && minPriceStr != "" {
		maxPrice, _ := strconv.ParseFloat(maxPriceStr, 64)
		minPrice, _ := strconv.ParseFloat(minPriceStr, 64)
		flights = service.FilterByPrice(flights, maxPrice, minPrice)
	}


	if maxDurationStr != "" && minDurationStr != "" {
		maxDuration, _ := strconv.Atoi(maxDurationStr)
		minDuration, _ := strconv.Atoi(minDurationStr)
		flights = service.FilterByDuration(flights, maxDuration, minDuration)
	}

	if len(numberOfStopFilter) > 0 {
		flights = service.FilterByNumberOfStops(flights, numberOfStopFilter)
	}

	sortBy := strings.ToLower(sortByStr)
	if sortBy == "" {
		sortBy = "price"
	}
	sortDirection := strings.ToLower(sortDirectionStr)
	if sortDirection == "" {
		sortDirection = "asc"
	}

	switch sortBy {
		case "price":
			service.SortByPrice(flights, sortDirection == "asc")
		case "duration":
			service.SortByDuration(flights, sortDirection == "asc")
	}

	flightsResponse := models.FlightResponse {
		SearchCriteria: models.FlightSearchCriteria {
			Origin: origin,
			Destination: destination,
			DepartureDate: date,
			Passengers: passengersInt,
			CabinClass: cabinClass,
		},
		Flights: flights,
		Metadata: aggregatedFlight.Metadata,
	}

	flightsResponse.Metadata.TotalResults = len(flights)

	response.Header().Set("Content-Type", "application/json")
	json.NewEncoder(response).Encode(flightsResponse)
}