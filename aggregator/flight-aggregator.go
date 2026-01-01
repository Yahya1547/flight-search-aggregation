package aggregator

import (
	"context"
	"sync"

	"flight-search-aggregation/models"
	"flight-search-aggregation/provider"
)

func Aggregate(
	ctx context.Context,
	req provider.SearchRequest,
	providersList []provider.AirlineProvider,
) ([]models.Flight, error) {

	var wg sync.WaitGroup
	ch := make(chan []models.Flight, len(providersList))

	for _, p := range providersList {
		wg.Add(1)
		go func(provider provider.AirlineProvider) {
			defer wg.Done()
			flights, _ := provider.GetFlights(ctx, req)
			ch <- flights
		}(p)
	}

	wg.Wait()
	close(ch)

	var results []models.Flight
	for f := range ch {
		results = append(results, f...)
	}

	return results, nil
}