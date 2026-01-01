package models

import "time"
import "flight-search-aggregation/utils"

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
			Formatted:    utils.formatDuration(durationMinutes),
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
		}
		Amenities: []string{},
	}
}

{
      "flight_code": "QZ520",
      "airline": "AirAsia",
      "from_airport": "CGK",
      "to_airport": "DPS",
      "depart_time": "2025-12-15T04:45:00+07:00",
      "arrive_time": "2025-12-15T07:25:00+08:00",
      "duration_hours": 1.67,
      "direct_flight": true,
      "price_idr": 650000,
      "seats": 67,
      "cabin_class": "economy",
      "baggage_note": "Cabin baggage only, checked bags additional fee"
    },
    {
      "flight_code": "QZ524",
      "airline": "AirAsia",
      "from_airport": "CGK",
      "to_airport": "DPS",
      "depart_time": "2025-12-15T10:00:00+07:00",
      "arrive_time": "2025-12-15T12:45:00+08:00",
      "duration_hours": 1.75,
      "direct_flight": true,
      "price_idr": 720000,
      "seats": 54,
      "cabin_class": "economy",
      "baggage_note": "Cabin baggage only, checked bags additional fee"
    },
    {
      "flight_code": "QZ532",
      "airline": "AirAsia",
      "from_airport": "CGK",
      "to_airport": "DPS",
      "depart_time": "2025-12-15T19:30:00+07:00",
      "arrive_time": "2025-12-15T22:10:00+08:00",
      "duration_hours": 1.67,
      "direct_flight": true,
      "price_idr": 595000,
      "seats": 72,
      "cabin_class": "economy",
      "baggage_note": "Cabin baggage only, checked bags additional fee"
    },
    {
      "flight_code": "QZ7250",
      "airline": "AirAsia",
      "from_airport": "CGK",
      "to_airport": "DPS",
      "depart_time": "2025-12-15T15:15:00+07:00",
      "arrive_time": "2025-12-15T20:35:00+08:00",
      "duration_hours": 4.33,
      "direct_flight": false,
      "stops": [
        {
          "airport": "SOC",
          "wait_time_minutes": 95
        }
      ],
      "price_idr": 485000,
      "seats": 88,
      "cabin_class": "economy",
      "baggage_note": "Cabin baggage only, checked bags additional fee"
    }