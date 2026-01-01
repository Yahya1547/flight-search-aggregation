package service

import (
	"flight-search-aggregation/models"
	"flight-search-aggregation/provider"
	"sort"
)

func FilterByPrice(flights []models.Flight, max float64, min float64) []models.Flight {
	var res []models.Flight
	for _, flight := range flights {
		if flight.Price.Amount <= max && flight.Price.Amount >= min {
			res = append(res, flight)
		}
	}
	return res
}

func FilterByDuration(flights []models.Flight, max int, min int) []models.Flight {
	var res []models.Flight
	for _, flight := range flights {
		if flight.Duration.TotalMinutes <= max && flight.Duration.TotalMinutes >= min {
			res = append(res, flight)
		}
	}
	return res
}

func FilterByNumberOfStops(flights []models.Flight, numberOfStops []int) []models.Flight {
	if len(numberOfStops) == 0 {
		return flights
	}

	var res []models.Flight
	stopsMap := make(map[int]bool)
	for _, stops := range numberOfStops {
		stopsMap[stops] = true
	}

	for _, flight := range flights {
		if stopsMap[flight.Stops] {
			res = append(res, flight)
		}
	}
	return res
}

func GetAirlineProvidersFilterByAirlines(airlines []string) []provider.AirlineProvider {
	var defaultProviders []provider.AirlineProvider
	defaultProviders = append(defaultProviders, &provider.LionAirProvider{})
	defaultProviders = append(defaultProviders, &provider.AirAsiaProvider{})
	defaultProviders = append(defaultProviders, &provider.BatikAirProvider{})
	defaultProviders = append(defaultProviders, &provider.GarudaIndonesiaProvider{})

	if len(airlines) == 0 {
		return defaultProviders
	}

	var providers []provider.AirlineProvider
	for _, airline := range airlines {
		switch airline {
		case "Lion Air":
			providers = append(providers, &provider.LionAirProvider{})
		case "Air Asia":
			providers = append(providers, &provider.AirAsiaProvider{})
		case "Batik Air":
			providers = append(providers, &provider.BatikAirProvider{})
		case "Garuda Indonesia":
			providers = append(providers, &provider.GarudaIndonesiaProvider{})
		}
	}
	return providers
}

func SortByPrice(flights []models.Flight, asc bool) {
	sort.Slice(flights, func(i, j int) bool {
		if asc {
			return flights[i].Price.Amount < flights[j].Price.Amount
		}
		return flights[i].Price.Amount > flights[j].Price.Amount
	})
}

func SortByDuration(flights []models.Flight, asc bool) {
	sort.Slice(flights, func(i, j int) bool {
		if asc {
			return flights[i].Duration.TotalMinutes < flights[j].Duration.TotalMinutes
		}
		return flights[i].Duration.TotalMinutes > flights[j].Duration.TotalMinutes
	})
}