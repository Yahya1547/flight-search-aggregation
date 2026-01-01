package service

import (
	"flight-search-aggregation/models"
	"sort"
)

func FilterByPrice(flights []models.Flight, max float64) []models.Flight {
	var res []models.Flight
	for _, f := range flights {
		if f.Price <= max {
			res = append(res, f)
		}
	}
	return res
}

func SortByPrice(flights []models.Flight, asc bool) {
	sort.Slice(flights, func(i, j int) bool {
		if asc {
			return flights[i].Price < flights[j].Price
		}
		return flights[i].Price > flights[j].Price
	})
}