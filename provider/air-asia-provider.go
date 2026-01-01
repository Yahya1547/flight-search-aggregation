package provider

import (
	"context"
    "encoding/json"
    "math/rand"
    "os"
    "errors"

	"flight-search-aggregation/models"
    "flight-search-aggregation/utils"
)

type AirAsiaProvider struct{}

func (airAsia *AirAsiaProvider) GetFlights(ctx context.Context, req SearchRequest) ([]models.Flight, error) {
	utils.RandomDelay(50, 150)

	// 90% success rate
	if rand.Float64() > 0.9 {
        return nil, errors.New("AirAsia service unavailable")
		// return nil, http.Error("AirAsia service unavailable")
	}

	data, _ := os.ReadFile("mock/airasia_search_response.json")

	var airAsiaResponse models.AirAsiaResponse
	json.Unmarshal(data, &airAsiaResponse)

	var results []models.Flight
    for _, data := range airAsiaResponse.Flights {
        results = append(results, data.ToFlight())
    }

	return results, nil
}