package provider

import (
	"context"
	"time"

	"flight-search-aggregation/models"
)

type BatikAirProvider struct{}

func (batikAir *BatikAirProvider) BaseUrl() string {
	return "http://localhost:8080/batikair"
}

func (batikAir *BatikAirProvider) GetFlights(ctx context.Context, req SearchRequest) ([]models.Flight, error) {
	reqHTTP, _ := http.NewRequestWithContext(ctx, "GET", batikAir.BaseUrl()+"/search", nil)
    response, err := http.DefaultClient.Do(reqHTTP)
    if err != nil {
        return nil, err
    }
    defer response.Body.Close()

	var batikAirResponse models.BatikAirResponse

    json.NewDecoder(response.Body).Decode(&batikAirResponse)

	var results []models.Flight
    for _, data := range batikAirResponse.Results {
        results = append(results, data.ToFlight())
    }

	return results, nil
}