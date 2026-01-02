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
	flightChannel := make(chan []models.Flight, len(providersList))
	errorChannel := make(chan error, len(providersList))

	for _, p := range providersList {
		wg.Add(1)
		go func(provider provider.AirlineProvider) {
			defer wg.Done()
			flights, err := provider.GetFlights(ctx, req)
			if err != nil {
                errorChannel <- err
                return
            }
			var validFlights []models.Flight
			for _, flight := range flights {
				if IsFlightValid(flight) {
					validFlights = append(validFlights, flight)
				}
			}
			flightChannel <- validFlights
		}(p)
	}

	wg.Wait()
	close(flightChannel)
	close(errorChannel)

	var results []models.Flight
	for flightsByProvider := range flightChannel {
		results = append(results, flightsByProvider...)
	}

	failedProviders := 0
	for _ = range errorChannel {
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

func IsFlightValid(flight models.Flight) bool {
	if (flight.AvailableSeats <= 0) {
		return false
	}

	if (flight.Price.Amount <= 0) {
		return false
	}

	if (flight.Duration.TotalMinutes <= 0) {
		return false
	}

	if (flight.Departure.Timestamp <= 0 || flight.Arrival.Timestamp <= 0) {
		return false
	}

	if (flight.Arrival.Timestamp <= flight.Departure.Timestamp) {
		return false
	}

	return true
}