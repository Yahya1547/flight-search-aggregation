package provider

import (
	"context"
	"encoding/json"
	"os"

	"flight-search-aggregation/models"
	"flight-search-aggregation/utils"
)

type GarudaIndonesiaProvider struct{
	
}

func (garudaIndonesia *GarudaIndonesiaProvider) GetFlights(ctx context.Context, req SearchRequest) ([]models.Flight, error) {
	utils.RandomDelay(50, 100)

	data, _ := os.ReadFile("mock/garuda_indonesia_search_response.json")

	var garudaIndonesiaResponse models.GarudaIndonesiaResponse
	json.Unmarshal(data, &garudaIndonesiaResponse)


	var results []models.Flight
    for _, garudaFlight := range garudaIndonesiaResponse.Flights {
        results = append(results, garudaFlight.ToFlight())
    }

	return results, nil
}