package domain

import "time"

// TransportType defines the type of transport
type TransportType string

const (
	TransportAir    TransportType = "air"
	TransportRail   TransportType = "rail"
	TransportBus    TransportType = "bus"
	TransportRiver  TransportType = "river"
)

// Stop represents a location on a route
type Stop struct {
	ID          string     `json:"id"`
	Name        string     `json:"name"`
	City        string     `json:"city"`
	Latitude    float64    `json:"latitude"`
	Longitude   float64    `json:"longitude"`
	ArrivalAt   *time.Time `json:"arrival_at,omitempty"`
	DepartureAt *time.Time `json:"departure_at,omitempty"`
}

// Segment represents a single transport leg
type Segment struct {
	ID              string        `json:"id"`
	TransportType   TransportType `json:"transport_type"`
	Provider        string        `json:"provider"`
	StartStop       Stop          `json:"start_stop"`
	EndStop         Stop          `json:"end_stop"`
	DepartureTime   time.Time     `json:"departure_time"`
	ArrivalTime     time.Time     `json:"arrival_time"`
	Price           float64       `json:"price"`
	Duration        time.Duration `json:"duration"`
	SeatCount       int           `json:"seat_count"`
	ReliabilityRate float64       `json:"reliability_rate"`
	Distance        int           `json:"distance"`
}

// Connection represents a transfer between segments
type Connection struct {
	From              *Segment      `json:"from,omitempty"`
	To                *Segment      `json:"to,omitempty"`
	TransferDuration  time.Duration `json:"transfer_duration"`
	TransferDistance  int           `json:"transfer_distance"`
	RequiresTransport bool          `json:"requires_transport"`
	IsValid           bool          `json:"is_valid"`
	Gap               time.Duration `json:"gap"`
}

// Route represents a complete journey with multiple segments
type Route struct {
	ID                string           `json:"id"`
	FromCity          string           `json:"from_city"`
	ToCity            string           `json:"to_city"`
	DepartureTime     time.Time        `json:"departure_time"`
	ArrivalTime       time.Time        `json:"arrival_time"`
	TotalDuration     time.Duration    `json:"total_duration"`
	Segments          []Segment        `json:"segments"`
	Connections       []Connection     `json:"connections,omitempty"`
	TotalPrice        float64          `json:"total_price"`
	ReliabilityScore  float64          `json:"reliability_score"`
	InsurancePremium  float64          `json:"insurance_premium"`
	InsuranceIncluded bool             `json:"insurance_included"`
	TransportTypes    []TransportType  `json:"transport_types"`
	SavedAt           time.Time        `json:"saved_at"`
}

// RouteSearchCriteria represents search parameters
type RouteSearchCriteria struct {
	FromCity         string
	ToCity           string
	DepartureDate    time.Time
	PassengerCount   int
	PreferredTransport []TransportType
	MaxConnections   int
	MaxTransferTime  int // minutes
	BudgetMax        float64
	BudgetMin        float64
}

// RouteSearchResult contains 3 optimized routes
type RouteSearchResult struct {
	RequestID     string
	FromCity      string
	ToCity        string
	DepartureDate time.Time
	PassengerCount int
	OptimalRoute  *Route
	FastestRoute  *Route
	CheapestRoute *Route
	SearchedAt    time.Time
}
