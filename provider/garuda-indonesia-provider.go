package provider

import (
	"context"
	"time"

	"flight-search-aggregation/models"
)

type GarudaIndonesiaProvider struct{
	
}

func (garudaIndonesia *GarudaIndonesiaProvider) Name() string {
	return "Garuda Indonesia Provider"
}

func (garudaIndonesia *GarudaIndonesiaProvider) BaseUrl() string {
	return "https://api.garudaindonesia.com"
}

func (garudaIndonesia *GarudaIndonesiaProvider) GetFlights(ctx context.Context, req SearchRequest) ([]models.Flight, error) {
	reqHTTP, _ := http.NewRequestWithContext(ctx, "GET", garudaIndonesia.BaseUrl()+"/search", nil)
    response, err := http.DefaultClient.Do(reqHTTP)
    if err != nil {
        return nil, err
    }
    defer response.Body.Close()

	var garudaIndonesiaResponse struct {
        Status string `json:"status"`
        Flights []models.GarudaIndonesiaFlight `json:"flights"`
    }

    json.NewDecoder(response.Body).Decode(&garudaIndonesiaResponse)

	var results []models.Flight
    for _, garudaFlight := range garudaIndonesiaResponse.Flights {
        results = append(results, garudaFlight.ToFlight())
    }

	return results, nil
}