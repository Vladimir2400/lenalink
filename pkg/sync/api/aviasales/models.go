package aviasales

// Airport represents airport data from Aviasales Data API.
type Airport struct {
	Code         string   `json:"code"`           // IATA code (e.g., "SVO", "LED")
	Name         string   `json:"name"`           // Airport name
	NameEn       string   `json:"name_translations,omitempty"` // English name
	CityCode     string   `json:"city_code"`      // City IATA code
	CityName     string   `json:"city_name"`      // City name
	CountryCode  string   `json:"country_code"`   // Country code (ISO 3166-1)
	CountryName  string   `json:"country_name"`   // Country name
	Latitude     float64  `json:"coordinates.lat"` // GPS latitude
	Longitude    float64  `json:"coordinates.lon"` // GPS longitude
	TimeZone     string   `json:"time_zone"`      // Timezone (e.g., "Europe/Moscow")
	IsActive     bool     `json:"is_active"`      // Whether airport is active
}

// City represents city data from Aviasales API.
type City struct {
	Code        string  `json:"code"`         // IATA code
	Name        string  `json:"name"`         // City name
	CountryCode string  `json:"country_code"` // Country code
	Latitude    float64 `json:"coordinates.lat"`
	Longitude   float64 `json:"coordinates.lon"`
	TimeZone    string  `json:"time_zone"`
}

// Flight represents flight schedule data.
type Flight struct {
	ID                string  `json:"id"`                  // Unique flight identifier
	DepartureAt       string  `json:"departure_at"`        // ISO 8601 datetime
	ReturnAt          string  `json:"return_at,omitempty"` // For round trips
	Origin            string  `json:"origin"`              // Origin IATA code
	Destination       string  `json:"destination"`         // Destination IATA code
	OriginAirport     string  `json:"origin_airport"`
	DestinationAirport string `json:"destination_airport"`
	Airline           string  `json:"airline"`             // Airline IATA code
	FlightNumber      string  `json:"flight_number"`
	Price             float64 `json:"price"`               // Price in RUB
	Currency          string  `json:"currency"`            // Currency code
	Transfers         int     `json:"transfers"`           // Number of transfers
	Duration          int     `json:"duration"`            // Duration in minutes
	AvailableSeats    int     `json:"available_seats"`     // Available seats count
}

// PriceResponse represents the response from prices endpoint.
type PriceResponse struct {
	Success bool     `json:"success"`
	Data    []Flight `json:"data"`
	Currency string  `json:"currency"`
}

// AirportsResponse represents the response from airports endpoint.
type AirportsResponse struct {
	Success bool      `json:"success"`
	Data    []Airport `json:"data"`
}

// CitiesResponse represents the response from cities endpoint.
type CitiesResponse struct {
	Success bool   `json:"success"`
	Data    []City `json:"data"`
}
