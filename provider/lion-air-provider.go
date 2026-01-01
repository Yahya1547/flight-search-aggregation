package provider

import (
	"context"
	"time"

	"flight-search-aggregation/models"
)

type LionAirProvider struct{}

func (lionAir *LionAirProvider) BaseUrl() string {
	return "http://localhost:8080/lionair"
}

func (lionAir *LionAirProvider) GetFlights(ctx context.Context, req SearchRequest) ([]models.Flight, error) {
	reqHTTP, _ := http.NewRequestWithContext(ctx, "GET", lionAir.BaseUrl()+"/search", nil)
    response, err := http.DefaultClient.Do(reqHTTP)
    if err != nil {
        return nil, err
    }
    defer response.Body.Close()

	var lionAirResponse models.LionAirResponse

    json.NewDecoder(response.Body).Decode(&lionAirResponse)

	var results []models.Flight
    for _, data := range lionAirResponse.Data.AvailableFlights {
        results = append(results, data.ToFlight())
    }

	return results, nil
}