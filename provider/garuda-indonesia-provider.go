package provider

import (
	"context"
	"time"

	"flight-search-aggregation/models"
)

type GarudaIndonesiaProvider struct{
	
}

func (garudaIndonesia *GarudaIndonesiaProvider) BaseUrl() string {
	return "http://localhost:8080/garudaindonesia"
}

func (garudaIndonesia *GarudaIndonesiaProvider) GetFlights(ctx context.Context, req SearchRequest) ([]models.Flight, error) {
	reqHTTP, _ := http.NewRequestWithContext(ctx, "GET", garudaIndonesia.BaseUrl()+"/search", nil)
    response, err := http.DefaultClient.Do(reqHTTP)
    if err != nil {
        return nil, err
    }
    defer response.Body.Close()

	var garudaIndonesiaResponse models.GarudaIndonesiaResponse

    json.NewDecoder(response.Body).Decode(&garudaIndonesiaResponse)

	var results []models.Flight
    for _, garudaFlight := range garudaIndonesiaResponse.Flights {
        results = append(results, garudaFlight.ToFlight())
    }

	return results, nil
}