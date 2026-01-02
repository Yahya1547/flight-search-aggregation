package service

import (
	"flight-search-aggregation/models"
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
	res := []models.Flight{}
	for _, flight := range flights {
		if flight.Duration.TotalMinutes / 60 <= max && flight.Duration.TotalMinutes / 60 >= min {
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

func FilterByAirlines(flights []models.Flight, airlines []string) []models.Flight {
	if len(airlines) == 0 {
		return flights
	}

	var res []models.Flight
	airlineMap := make(map[string]bool)
	for _, airline := range airlines {
		airlineMap[airline] = true
	}

	for _, flight := range flights {
		if airlineMap[flight.Provider] {
			res = append(res, flight)
		}
	}
	return res
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

func SortByDepartureTime(flights []models.Flight, asc bool) {
	sort.Slice(flights, func(i, j int) bool {
		if asc {
			return flights[i].Departure.Timestamp < flights[j].Departure.Timestamp
		}
		return flights[i].Departure.Timestamp > flights[j].Departure.Timestamp
	})
}

func SortByArrivalTime(flights []models.Flight, asc bool) {
	sort.Slice(flights, func(i, j int) bool {
		if asc {
			return flights[i].Arrival.Timestamp < flights[j].Arrival.Timestamp
		}
		return flights[i].Arrival.Timestamp > flights[j].Arrival.Timestamp
	})
}