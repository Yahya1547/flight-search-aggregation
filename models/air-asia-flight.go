package models

import (
	"time"
	"flight-search-aggregation/utils"
)

type AirAsiaResponse struct {
	Status string `json:"status"`
	Flights []AirAsiaFlight `json:"flights"`
}

type AirAsiaFlight struct {
	FlightCode       string    `json:"flight_code"`
	Airline         string    `json:"airline"`
	FromAirport     string    `json:"from_airport"`
	ToAirport       string    `json:"to_airport"`
	DepartTime      time.Time `json:"depart_time"`
	ArriveTime      time.Time `json:"arrive_time"`
	DurationHours   float64   `json:"duration_hours"`
	DirectFlight    bool      `json:"direct_flight"`
	Stops           []AirAsiaFlightStop `json:"stops,omitempty"`
	PriceIDR        int       `json:"price_idr"`
	Seats          int       `json:"seats"`
	CabinClass      string    `json:"cabin_class"`
	BaggageNote     string    `json:"baggage_note"`
}

type AirAsiaFlightStop struct {
	Airport         string `json:"airport"`
	WaitTimeMinutes int    `json:"wait_time_minutes"`
}

func (airAsia AirAsiaFlight) ToFlight() Flight {
	durationMinutes := int(airAsia.DurationHours * 60)
	var stops int
	if airAsia.DirectFlight {
		stops = 0
	} else {
		stops = len(airAsia.Stops)
	}
	return Flight{
		Id:       airAsia.FlightCode + "_AirAsia",
		Provider: "AirAsia",
		Airline: AirlineInfo{
			Name: airAsia.Airline,
			Code: "QZ",
		},
		FlightNumber: airAsia.FlightCode,
		Departure: FlightPointInfo{
			Airport:  airAsia.FromAirport,
			City:     "",
			Datetime: airAsia.DepartTime,
			Timestamp: airAsia.DepartTime.Unix(),
		},
		Arrival: FlightPointInfo{
			Airport:  airAsia.ToAirport,
			City:     "",
			Datetime: airAsia.ArriveTime,
			Timestamp: airAsia.ArriveTime.Unix(),
		},
		Duration: DurationInfo{
			TotalMinutes: durationMinutes,
			Formatted:    utils.FormatDuration(durationMinutes),
		},
		Stops: stops,
		Price: PriceInfo{
			Amount:   float64(airAsia.PriceIDR),
			Currency: "IDR",
		},
		AvailableSeats: airAsia.Seats,
		CabinClass:     airAsia.CabinClass,
		Baggage: BaggageInfo{
			CarryOn: "Cabin baggage only",
			Checked: "checked bags additional fee",
		},
		Amenities: []string{},
	}
}

