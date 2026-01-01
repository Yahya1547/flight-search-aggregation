package models

import "time"
import "flight-search-aggregation/utils"

type BatikAirResponse struct {
	Code int   `json:"code"`
	Message string `json:"message"`
	Results []BatikAirFlight `json:"results"`
}

type BatikAirFlight struct {
	FlightNumber      string    `json:"flightNumber"`
	AirlineName       string    `json:"airlineName"`
	AirlineIATA       string    `json:"airlineIATA"`
	Origin            string    `json:"origin"`
	Destination       string    `json:"destination"`
	DepartureDateTime time.Time `json:"departureDateTime"`
	ArrivalDateTime   time.Time `json:"arrivalDateTime"`
	TravelTime        string    `json:"travelTime"`
	NumberOfStops     int       `json:"numberOfStops"`
	Fare              BatikAirFareInfo `json:"fare"`
	SeatsAvailable    int       `json:"seatsAvailable"`
	AircraftModel     string    `json:"aircraftModel"`
	BaggageInfo       string    `json:"baggageInfo"`
	OnboardServices   []string  `json:"onboardServices"`
	Connections 	[]BatikAirConnectionInfo `json:"connections,omitempty"`
}

type BatikAirFareInfo struct {
	BasePrice    float64 `json:"basePrice"`
	Taxes        float64 `json:"taxes"`
	TotalPrice   float64 `json:"totalPrice"`
	CurrencyCode string  `json:"currencyCode"`
	Class        string  `json:"class"`
}

type BatikAirConnectionInfo struct {
	StopAirport  string `json:"stopAirport"`
	StopDuration string `json:"stopDuration"`
}

func (batikAir BatikAirFlight) ToFlight() Flight {
	durationMinutes := utils.ParseDurationToMinutes(batikAir.TravelTime)

	return Flight{
		Id:       batikAir.FlightNumber + "_BatikAir",
		Provider: "Batik Air",
		Airline: AirlineInfo{
			Name: batikAir.AirlineName,
			Code: batikAir.AirlineIATA,
		},
		FlightNumber: batikAir.FlightNumber,
		Departure: FlightPointInfo{
			Airport:  batikAir.Origin,
			City:     "",
			Datetime: batikAir.DepartureDateTime,
			Timestamp: batikAir.DepartureDateTime.Unix(),
		},
		Arrival: FlightPointInfo{
			Airport:  batikAir.Destination,
			City:     "",
			Datetime: batikAir.ArrivalDateTime,
			Timestamp: batikAir.ArrivalDateTime.Unix(),
		},
		Duration: DurationInfo{
			TotalMinutes: durationMinutes,
			Formatted:    batikAir.TravelTime,
		},
		Stops: batikAir.NumberOfStops,
		Price: PriceInfo{
			Amount:   batikAir.Fare.TotalPrice,
			Currency: batikAir.Fare.CurrencyCode,
		},
		AvailableSeats: batikAir.SeatsAvailable,
		CabinClass:     batikAir.Fare.Class,
		Aircraft:       batikAir.AircraftModel,
		Baggage: BaggageInfo{
			CarryOn: parseBaggageCarryOn(batikAir.BaggageInfo),
			Checked: parseBaggageChecked(batikAir.BaggageInfo),
		},
		Amenities: batikAir.OnboardServices,
	}
}

func parseBaggageCarryOn(baggageInfo string) string {
	// Example baggageInfo: "7kg cabin, 20kg checked"
	var carryOn string
	fmt.Sscanf(baggageInfo, "%s cabin,", &carryOn)
	return carryOn
}

func parseBaggageChecked(baggageInfo string) string {
	// Example baggageInfo: "7kg cabin, 20kg checked"
	var checked string
	fmt.Sscanf(baggageInfo, ", %s checked", &checked)
	return checked
}
