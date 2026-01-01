package provider

import (
	"context"
	"encoding/json"
	"os"

	"flight-search-aggregation/models"
	"flight-search-aggregation/utils"
)

type LionAirProvider struct{}

func (lionAir *LionAirProvider) GetFlights(ctx context.Context, req SearchRequest) ([]models.Flight, error) {
	utils.RandomDelay(100, 200)

	data, _ := os.ReadFile("mock/lion_air_search_response.json")
	
	var lionAirResponse models.LionAirResponse
	json.Unmarshal(data, &lionAirResponse)

	var results []models.Flight
    for _, data := range lionAirResponse.Data.AvailableFlights {
        results = append(results, data.ToFlight())
    }

	return results, nil
}