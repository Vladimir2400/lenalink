package rzd

import "time"

// Station represents a railway station.
type Station struct {
	Code        string  `json:"code"`         // Station code (e.g., "2000000")
	Name        string  `json:"name"`         // Station name
	City        string  `json:"city"`         // City name
	Region      string  `json:"region"`       // Region name
	Country     string  `json:"country"`      // Country code
	Latitude    float64 `json:"latitude"`     // GPS latitude
	Longitude   float64 `json:"longitude"`    // GPS longitude
	TimeZone    string  `json:"time_zone"`    // Timezone
}

// Train represents train schedule data.
type Train struct {
	TrainNumber    string    `json:"train_number"`    // Train number (e.g., "002А")
	TrainName      string    `json:"train_name"`      // Train name
	OriginStation  string    `json:"origin_station"`  // Origin station code
	DestStation    string    `json:"dest_station"`    // Destination station code
	DepartureTime  time.Time `json:"departure_time"`  // Departure datetime
	ArrivalTime    time.Time `json:"arrival_time"`    // Arrival datetime
	Duration       int       `json:"duration"`        // Duration in minutes
	Carrier        string    `json:"carrier"`         // Carrier name (e.g., "РЖД")
	TrainType      string    `json:"train_type"`      // Train type (e.g., "Скоростной")
	Distance       int       `json:"distance"`        // Distance in km
}

// Ticket represents ticket pricing information.
type Ticket struct {
	TrainNumber    string  `json:"train_number"`
	CarType        string  `json:"car_type"`        // Car class (Плацкарт, Купе, СВ)
	Price          float64 `json:"price"`           // Price in RUB
	AvailableSeats int     `json:"available_seats"` // Available seats
	ServiceClass   string  `json:"service_class"`   // Service class
}
