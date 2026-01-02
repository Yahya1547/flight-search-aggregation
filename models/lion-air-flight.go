package models

import (
	"strings"
	"time"
	"flight-search-aggregation/utils"
)

type LionAirResponse struct {
	Success bool `json:"success"`
	Data LionAirDataResponse `json:"data"`
}

type LionAirDataResponse struct {
	AvailableFlights []LionAirFlight `json:"available_flights"`
}

type LionAirFlight struct {
	Id       string    `json:"id"`
	Carrier  CarrierInfo `json:"carrier"`
	Route    RouteInfo   `json:"route"`
	Schedule ScheduleInfo `json:"schedule"`
	FlightTime int       `json:"flight_time"`
	IsDirect   bool      `json:"is_direct"`
	Pricing   PricingInfo `json:"pricing"`
	SeatsLeft int       `json:"seats_left"`
	PlaneType string    `json:"plane_type"`
	Services  ServiceInfo `json:"services"`
	Layovers  []LayoverInfo `json:"layovers,omitempty"`
	StopCount int       `json:"stop_count"`
}

type CarrierInfo struct {
	Name string `json:"name"`
	Iata string `json:"iata"`
}

type RouteInfo struct {
	From AirportInfo `json:"from"`
	To   AirportInfo `json:"to"`
}

type AirportInfo struct {
	Code string `json:"code"`
	Name string `json:"name"`
	City string `json:"city"`
}

type ScheduleInfo struct {
	Departure          string `json:"departure"`
	DepartureTimezone  string    `json:"departure_timezone"`
	Arrival            string `json:"arrival"`
	ArrivalTimezone    string    `json:"arrival_timezone"`
}

type PricingInfo struct {
	Total    float64 `json:"total"`
	Currency string  `json:"currency"`
	FareType string  `json:"fare_type"`
}

type ServiceInfo struct {
	WifiAvailable   bool           `json:"wifi_available"`
	MealsIncluded  bool           `json:"meals_included"`
	BaggageAllowance BaggageAllowanceInfo `json:"baggage_allowance"`
}

type BaggageAllowanceInfo struct {
	Cabin string `json:"cabin"`
	Hold  string  `json:"hold"`
}

type LayoverInfo struct {
	Airport string `json:"airport"`
	DurationMinutes    int `json:"duration_minutes"`
}

func (lionAir LionAirFlight) ToFlight() Flight {
	var stopCount int
	if (lionAir.IsDirect) {
		stopCount = 0
	} else {
		stopCount = lionAir.StopCount
	}
	
	var amenities []string = []string{}
	if (lionAir.Services.WifiAvailable) {
		amenities = append(amenities, "WiFi")
	}
	if (lionAir.Services.MealsIncluded) {
		amenities = append(amenities, "Meals")
	}

	var layoverTime int
	if (lionAir.IsDirect) {
		layoverTime = 0
	} else {
		for _, layover := range lionAir.Layovers {
			layoverTime += layover.DurationMinutes
		}
	}

	flightTime := lionAir.FlightTime + layoverTime
	
	departureTime, _ := time.Parse("2006-01-02T15:04:05", lionAir.Schedule.Departure)
	arrivalTime, _ := time.Parse("2006-01-02T15:04:05", lionAir.Schedule.Arrival)
	
	return Flight {
		Id:       lionAir.Id + "_LionAir",
		Provider: "Lion Air",
		Airline: AirlineInfo{
			Name: lionAir.Carrier.Name,
			Code: lionAir.Carrier.Iata,
		},
		FlightNumber: lionAir.Id,
		Departure: FlightPointInfo{
			Airport:  lionAir.Route.From.Code,
			City:     lionAir.Route.From.City,
			Datetime: strings.Split(lionAir.Schedule.Departure, "T")[0],
			Timestamp: int64(departureTime.Unix()),
		},
		Arrival: FlightPointInfo{
			Airport:  lionAir.Route.To.Code,
			City:     lionAir.Route.To.City,
			Datetime: strings.Split(lionAir.Schedule.Arrival, "T")[0],
			Timestamp: int64(arrivalTime.Unix()),
		},
		Duration: DurationInfo{
			TotalMinutes: flightTime,
			Formatted: utils.FormatDuration(flightTime),
		},
		Stops: stopCount,
		Aircraft: lionAir.PlaneType,
		Price: PriceInfo{
			Amount:   lionAir.Pricing.Total,
			Currency: lionAir.Pricing.Currency,
		},
		AvailableSeats: lionAir.SeatsLeft,
		CabinClass:     lionAir.Pricing.FareType,
		Baggage: BaggageInfo{
			CarryOn: lionAir.Services.BaggageAllowance.Cabin,
			Checked: lionAir.Services.BaggageAllowance.Hold,
		},
		Amenities: amenities,
	}
}