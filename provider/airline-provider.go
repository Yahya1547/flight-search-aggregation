package provider

import (
	"context"
	"flight-search-aggregation/models"
)

type SearchRequest struct {
	Origin      string
	Destination string
}

type AirlineProvider interface {
	BaseUrl() string
	Name() string
	GetFlights(ctx context.Context, req SearchRequest) ([]models.Flight, error)
}