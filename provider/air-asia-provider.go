package provider

import (
	"context"
	"time"

	"flight-search-aggregation/models"
)

type AirAsiaProvider struct{}

func (airAsia *AirAsiaProvider) BaseUrl() string {
	return "http://localhost:8080/airasia"
}

func (airAsia *AirAsiaProvider) GetFlights(ctx context.Context, req SearchRequest) ([]models.Flight, error) {
	reqHTTP, _ := http.NewRequestWithContext(ctx, "GET", airAsia.BaseUrl()+"/search", nil)
    response, err := http.DefaultClient.Do(reqHTTP)
    if err != nil {
        return nil, err
    }
    defer response.Body.Close()

	var airAsiaResponse models.AirAsiaResponse

    json.NewDecoder(response.Body).Decode(&airAsiaResponse)

	var results []models.Flight
    for _, data := range airAsiaResponse.Flights {
        results = append(results, data.ToFlight())
    }

	return results, nil
}