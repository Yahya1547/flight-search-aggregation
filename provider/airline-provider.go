package provider

import (
	"time"
	"context"
	"flight-search-aggregation/models"
)

type SearchRequest struct {
	Origin      string `json:"origin"`
	Destination string `json:"destination"`
	DepartureDate time.Time `json:"departureDate"`
	Passengers  int `json:"passengers"`
	CabinClass string `json:"cabinClass"`
}

type AirlineProvider interface {
	GetFlights(ctx context.Context, req SearchRequest) ([]models.Flight, error)
}