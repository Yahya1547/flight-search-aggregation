package models

import "time"
import "flight-search-aggregation/utils"

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
	Price          PriceInfo `json:"price"`
	AvailableSeats int       `json:"available_seats"`
	FareClass      string    `json:"fare_class"`
	Baggage        GarudaBaggageInfo `json:"baggage"`
	Amenities      []string  `json:"amenities"`
	Segments 	[]GarudaFlightSegment `json:"segments,omitempty"`
}

type GarudaFlightPointInfo struct {
	Airport  string    `json:"airport"`
	City     string    `json:"city"`
	Time     time.Time `json:"time"`
	Terminal string    `json:"terminal"`
}

type PriceInfo struct {
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
			Datetime: garudaIndonesia.Departure.Time,
			Timestamp: int64(garudaIndonesia.Departure.Time.Unix()),
		},
		Arrival: FlightPointInfo{
			Airport:  garudaIndonesia.Arrival.Airport,
			City:     garudaIndonesia.Arrival.City,
			Datetime: garudaIndonesia.Arrival.Time,
			Timestamp: int64(garudaIndonesia.Arrival.Time.Unix()),
		},
		Duration: DurationInfo{
			TotalMinutes: garudaIndonesia.DurationMin,
			Formatted: utils.formatDuration(garudaIndonesia.DurationMin),
		},
		Stops:    garudaIndonesia.Stops,
		Aircraft: garudaIndonesia.Aircraft,
		Baggage: BaggageInfo{
			CarryOn: garudaIndonesia.Baggage.CarryOn,
			Checked: garudaIndonesia.Baggage.Checked,
		},
        Price: PriceInfo{
            Amount:   garudaIndonesia.Price.Amount,
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

