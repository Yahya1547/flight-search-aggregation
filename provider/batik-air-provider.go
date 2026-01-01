package provider

import (
	"context"
    "encoding/json"
    "os"

	"flight-search-aggregation/models"
    "flight-search-aggregation/utils"
)

type BatikAirProvider struct{}

func (batikAir *BatikAirProvider) GetFlights(ctx context.Context, req SearchRequest) ([]models.Flight, error) {
	utils.RandomDelay(200, 400)

	data, _ := os.ReadFile("mock/batik_air_search_response.json")

	var batikAirResponse models.BatikAirResponse
	json.Unmarshal(data, &batikAirResponse)

	var results []models.Flight
    for _, data := range batikAirResponse.Results {
        results = append(results, data.ToFlight())
    }

	return results, nil
}