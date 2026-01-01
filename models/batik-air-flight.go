package models

import "time"
import "flight-search-aggregation/utils"

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

{
      "flightNumber": "ID6514",
      "airlineName": "Batik Air",
      "airlineIATA": "ID",
      "origin": "CGK",
      "destination": "DPS",
      "departureDateTime": "2025-12-15T07:15:00+0700",
      "arrivalDateTime": "2025-12-15T10:00:00+0800",
      "travelTime": "1h 45m",
      "numberOfStops": 0,
      "fare": {
        "basePrice": 980000,
        "taxes": 120000,
        "totalPrice": 1100000,
        "currencyCode": "IDR",
        "class": "Y"
      },
      "seatsAvailable": 32,
      "aircraftModel": "Airbus A320",
      "baggageInfo": "7kg cabin, 20kg checked",
      "onboardServices": [
        "Snack",
        "Beverage"
      ]
    },
    {
      "flightNumber": "ID6520",
      "airlineName": "Batik Air",
      "airlineIATA": "ID",
      "origin": "CGK",
      "destination": "DPS",
      "departureDateTime": "2025-12-15T13:30:00+0700",
      "arrivalDateTime": "2025-12-15T16:20:00+0800",
      "travelTime": "1h 50m",
      "numberOfStops": 0,
      "fare": {
        "basePrice": 1050000,
        "taxes": 130000,
        "totalPrice": 1180000,
        "currencyCode": "IDR",
        "class": "Y"
      },
      "seatsAvailable": 18,
      "aircraftModel": "Boeing 737-800",
      "baggageInfo": "7kg cabin, 20kg checked",
      "onboardServices": [
        "Meal",
        "Beverage",
        "Entertainment"
      ]
    },
    {
      "flightNumber": "ID7042",
      "airlineName": "Batik Air",
      "airlineIATA": "ID",
      "origin": "CGK",
      "destination": "DPS",
      "departureDateTime": "2025-12-15T18:45:00+0700",
      "arrivalDateTime": "2025-12-15T23:50:00+0800",
      "travelTime": "3h 5m",
      "numberOfStops": 1,
      "connections": [
        {
          "stopAirport": "UPG",
          "stopDuration": "55m"
        }
      ],
      "fare": {
        "basePrice": 850000,
        "taxes": 100000,
        "totalPrice": 950000,
        "currencyCode": "IDR",
        "class": "Y"
      },
      "seatsAvailable": 41,
      "aircraftModel": "Airbus A320",
      "baggageInfo": "7kg cabin, 20kg checked",
      "onboardServices": [
        "Snack"
      ]
    }