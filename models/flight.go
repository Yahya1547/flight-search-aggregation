package models

type FlightResponse struct {
	SearchCriteria FlightSearchCriteria `json:"search_criteria"`
	Flights        []Flight            `json:"flights"`
	Metadata       FlightMetadata     `json:"metadata"`
}

type FlightSearchCriteria struct {
	Origin        string `json:"origin"`
	Destination   string `json:"destination"`
	DepartureDate string `json:"departure_date"`
	ReturnDate    string `json:"return_date,omitempty"`
	Passengers    int    `json:"passengers"`
	CabinClass    string `json:"cabin_class"`
}

type FlightMetadata struct {
	TotalResults 	  int  `json:"total_results"`
	ProvidersQueried  int  `json:"providers_queried"`
	ProvidersSucceeded int `json:"providers_succeeded"`
	ProvidersFailed    int `json:"providers_failed"`
	SearchTimeMs       int `json:"search_time_ms"`
	CacheHit           bool `json:"cache_hit"`
}

type FlightAggregatedData struct {
	Flights []Flight `json:"flights"`
	Metadata FlightMetadata `json:"metadata"`
}

type Flight struct {
	Id       string    `json:"id"`
	Provider string    `json:"provider"`
	Airline	AirlineInfo `json:"airline"`
	FlightNumber string    `json:"flight_number"`
	Departure FlightPointInfo `json:"departure"`
	Arrival   FlightPointInfo `json:"arrival"`
	Duration DurationInfo `json:"duration"`
	Stops    int       `json:"stops"`
	Price    PriceInfo `json:"price"`
	AvailableSeats int `json:"available_seats"`
	CabinClass     string `json:"cabin_class"`
	Aircraft	   string    `json:"aircraft,omitempty"`
	Amenities      []string  `json:"amenities,omitempty"`
	Baggage        BaggageInfo `json:"baggage"`
}

type FlightPointInfo struct {
	Airport  string    `json:"airport"`
	City     string    `json:"city"`
	Datetime string `json:"datetime"`
	Timestamp int64    `json:"timestamp"`
}

type AirlineInfo struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

type DurationInfo struct {
	TotalMinutes int    `json:"total_minutes"`
	Formatted    string `json:"formatted"`
}

type PriceInfo struct {
	Amount   float64 `json:"amount"`
	FormattedAmount string `json:"formatted_amount"`
	Currency string  `json:"currency"`
}

type BaggageInfo struct {
	CarryOn string `json:"carry_on"`
	Checked string `json:"checked"`
}