package provider

import (
	"context"
	"time"

	"flight-search-aggregation/models"
)

type AirAsiaProvider struct{}

func (airAsia *AirAsiaProvider) Name() string {
	return "AirAsiaProvider"
}

func (airAsia *AirAsiaProvider) BaseUrl() string {
	return "https://api.AirAsia.com"
}

func (airAsia *AirAsiaProvider) GetFlights(ctx context.Context, req SearchRequest) ([]models.Flight, error) {
	reqHTTP, _ := http.NewRequestWithContext(ctx, "GET", airAsia.BaseUrl()+"/search", nil)
    response, err := http.DefaultClient.Do(reqHTTP)
    if err != nil {
        return nil, err
    }
    defer response.Body.Close()

	var airAsiaResponse struct {
        Status string `json:"status"`
        Flights []models.AirAsiaFlight `json:"flights"`
    }

    json.NewDecoder(response.Body).Decode(&airAsiaResponse)

	var results []models.Flight
    for _, data := range airAsiaResponse.Flights {
        results = append(results, data.ToFlight())
    }

	return results, nil
}