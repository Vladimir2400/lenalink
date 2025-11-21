package aviasales

// Coordinates represents geographic coordinates.
type Coordinates struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

// NameTranslations represents name translations in various languages.
type NameTranslations struct {
	En string `json:"en"`
}

// Airport represents airport data from Aviasales Data API.
type Airport struct {
	Code             string           `json:"code"`             // IATA code (e.g., "SVO", "LED")
	Name             string           `json:"name"`             // Airport name
	NameTranslations NameTranslations `json:"name_translations"` // Name translations
	CityCode         string           `json:"city_code"`        // City IATA code
	CountryCode      string           `json:"country_code"`     // Country code (ISO 3166-1)
	Coordinates      Coordinates      `json:"coordinates"`      // GPS coordinates
	TimeZone         string           `json:"time_zone"`        // Timezone (e.g., "Europe/Moscow")
	IataType         string           `json:"iata_type"`        // Type: airport, bus, etc.
	Flightable       bool             `json:"flightable"`       // Whether airport is flightable
}

// City represents city data from Aviasales API.
type City struct {
	Code             string           `json:"code"`              // IATA code
	Name             string           `json:"name"`              // City name
	NameTranslations NameTranslations `json:"name_translations"` // Name translations
	CountryCode      string           `json:"country_code"`      // Country code
	Coordinates      Coordinates      `json:"coordinates"`       // GPS coordinates
	TimeZone         string           `json:"time_zone"`         // Timezone
}

// Flight represents flight schedule data from Aviasales latest prices API.
type Flight struct {
	DepartDate       string  `json:"depart_date"`      // Departure date (YYYY-MM-DD)
	ReturnDate       string  `json:"return_date"`      // Return date for round trips
	Origin           string  `json:"origin"`           // Origin IATA code (city)
	Destination      string  `json:"destination"`      // Destination IATA code (city)
	Gate             string  `json:"gate"`             // Booking gateway/partner
	FoundAt          string  `json:"found_at"`         // When price was found (ISO 8601)
	TripClass        int     `json:"trip_class"`       // Trip class (0=economy, 1=business, etc.)
	Value            float64 `json:"value"`            // Price value
	NumberOfChanges  int     `json:"number_of_changes"` // Number of transfers/changes
	Duration         int     `json:"duration"`         // Duration in minutes
	Distance         int     `json:"distance"`         // Distance in km
	ShowToAffiliates bool    `json:"show_to_affiliates"`
	Actual           bool    `json:"actual"`           // Whether price is still actual
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
