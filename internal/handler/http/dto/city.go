package dto

// CityResponse represents a city in autocomplete
type CityResponse struct {
	Name      string  `json:"name"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

// CitiesResponse represents a list of cities
type CitiesResponse struct {
	Cities []CityResponse `json:"cities"`
}
