package aggregator

import (
	"context"
	"sync"

	"flight-search-aggregation/models"
	"flight-search-aggregation/provider"

	"github.com/patrickmn/go-cache"
)

func Aggregate(
	ctx context.Context,
	req provider.SearchRequest,
	providersList []provider.AirlineProvider,
) (models.FlightAggregatedData, error) {

	cacheKey := cacheKey(req.Origin, req.Destination, req.DepartureDate)
	if cachedData, found := cacheInstance.Get(cacheKey); found {
		aggregatedData := cachedData.(models.FlightAggregatedData)
		aggregatedData.Metadata.CacheHit = true
		return aggregatedData, nil
	}
	
	var wg sync.WaitGroup
	ch := make(chan []models.Flight, len(providersList))
	errCh := make(chan error, len(providersList))

	for _, p := range providersList {
		wg.Add(1)
		go func(provider provider.AirlineProvider) {
			defer wg.Done()
			flights, err := provider.GetFlights(ctx, req)
			if err != nil {
                errCh <- err
                return
            }
			ch <- flights
		}(p)
	}

	wg.Wait()
	close(ch)
	close(errCh)

	var results []models.Flight
	for f := range ch {
		results = append(results, f...)
	}

	failedProviders := 0
	for _ = range errCh {
		failedProviders++
	}

	aggregatedData := models.FlightAggregatedData{
		Flights:		   results,
		Metadata: models.FlightMetadata{
			TotalResults: 	len(results),
			ProvidersQueried: len(providersList),
			ProvidersSucceeded: len(providersList) - failedProviders,
			ProvidersFailed:	failedProviders,
			SearchTimeMs:		0,
			CacheHit:			false,
		},
	}

	cacheInstance.Set(cacheKey, aggregatedData, cache.DefaultExpiration)

	return aggregatedData, nil
}