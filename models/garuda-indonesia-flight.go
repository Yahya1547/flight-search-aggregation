package models

import (
	"time"
	"fmt"
	"strconv"
	"strings"

	"flight-search-aggregation/utils"
)

type GarudaIndonesiaResponse struct {
	Status string `json:"status"`
	Flights []GarudaIndonesiaFlight `json:"flights"`
}

type GarudaIndonesiaFlight struct {
	FlightId       string    `json:"flight_id"`
	Airline        string    `json:"airline"`
	AirlineCode    string    `json:"airline_code"`
	Departure   GarudaFlightPointInfo `json:"departure"`
	Arrival   GarudaFlightPointInfo `json:"arrival"`
	DurationMin	int       `json:"duration_minutes"`
	Stops          int       `json:"stops"`
	Aircraft       string    `json:"aircraft"`
	Price          GarudaPriceInfo `json:"price"`
	AvailableSeats int       `json:"available_seats"`
	FareClass      string    `json:"fare_class"`
	Baggage        GarudaBaggageInfo `json:"baggage"`
	Amenities      []string  `json:"amenities"`
	Segments 	[]GarudaFlightSegment `json:"segments,omitempty"`
}

type GarudaFlightPointInfo struct {
	Airport  string    `json:"airport"`
	City     string    `json:"city"`
	Time     string `json:"time"`
	Terminal string    `json:"terminal"`
}

type GarudaPriceInfo struct {
	Amount   float64 `json:"amount"`
	Currency string  `json:"currency"`
}

type GarudaBaggageInfo struct {
	CarryOn int `json:"carry_on"`
	Checked int `json:"checked"`
}

type GarudaFlightSegment struct {
	FlightNumber   string                 `json:"flight_number"`
	Departure      GarudaSegmentPointInfo `json:"departure"`
	Arrival        GarudaSegmentPointInfo `json:"arrival"`
	DurationMin    int                    `json:"duration_minutes"`
	LayoverMinutes int                    `json:"layover_minutes,omitempty"`
}

type GarudaSegmentPointInfo struct {
	Airport string    `json:"airport"`
	Time    time.Time `json:"time"`
}

func (garudaIndonesia GarudaIndonesiaFlight) ToFlight() Flight {
	departureTime, _ := time.Parse(time.RFC3339, garudaIndonesia.Departure.Time)
	arrivalTime, _ := time.Parse(time.RFC3339, garudaIndonesia.Arrival.Time)
	return Flight {
		Id:       garudaIndonesia.FlightId + "_GarudaIndonesia",
		Provider: "Garuda Indonesia",
		Airline: AirlineInfo{
			Name: garudaIndonesia.Airline,
			Code: garudaIndonesia.AirlineCode,
		},
		FlightNumber: garudaIndonesia.FlightId,
		Departure: FlightPointInfo{
			Airport:  garudaIndonesia.Departure.Airport,
			City:     garudaIndonesia.Departure.City,
			Datetime: strings.Split(garudaIndonesia.Departure.Time, "T")[0],
			Timestamp: int64(departureTime.Unix()),
		},
		Arrival: FlightPointInfo{
			Airport:  garudaIndonesia.Arrival.Airport,
			City:     garudaIndonesia.Arrival.City,
			Datetime: strings.Split(garudaIndonesia.Arrival.Time, "T")[0],
			Timestamp: int64(arrivalTime.Unix()),
		},
		Duration: DurationInfo{
			TotalMinutes: garudaIndonesia.DurationMin,
			Formatted: utils.FormatDuration(garudaIndonesia.DurationMin),
		},
		Stops:    garudaIndonesia.Stops,
		Aircraft: garudaIndonesia.Aircraft,
		Baggage: BaggageInfo{
			CarryOn: strconv.Itoa(garudaIndonesia.Baggage.CarryOn),
			Checked: strconv.Itoa(garudaIndonesia.Baggage.Checked),
		},
        Price: PriceInfo{
            Amount:   garudaIndonesia.Price.Amount,
			FormattedAmount: utils.FormatIDR(garudaIndonesia.Price.Amount),
            Currency: "IDR",
        },
        AvailableSeats:garudaIndonesia.AvailableSeats,
        CabinClass:garudaIndonesia.FareClass,
		Amenities: garudaIndonesia.Amenities,
    }
}

func formatDuration(totalMinutes int) string {
	hours := totalMinutes / 60
	minutes := totalMinutes % 60
	return  fmt.Sprintf("%dh %dm", hours, minutes)
}

