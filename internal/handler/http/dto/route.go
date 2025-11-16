package dto

import "time"

// SearchRouteRequest represents a request to search for routes
type SearchRouteRequest struct {
	From          string `json:"from" validate:"required"`
	To            string `json:"to" validate:"required"`
	DepartureDate string `json:"departure_date" validate:"required"` // YYYY-MM-DD format
	Passengers    int    `json:"passengers" validate:"omitempty,min=1,max=10"`
}

// RouteResponse represents a route in API response
type RouteResponse struct {
	ID               string            `json:"id"`
	Type             string            `json:"type"` // optimal, fastest, cheapest
	Segments         []SegmentResponse `json:"segments"`
	TotalPrice       float64           `json:"total_price"`
	TotalDistance    int               `json:"total_distance"`
	TotalDuration    string            `json:"total_duration"` // e.g., "6h 30m"
	ReliabilityScore float64           `json:"reliability_score,omitempty"`
}

// SegmentResponse represents a transport segment
type SegmentResponse struct {
	ID            string       `json:"id"`
	TransportType string       `json:"transport_type"` // air, rail, bus, river, taxi, walk
	Provider      string       `json:"provider"`
	From          StopResponse `json:"from"`
	To            StopResponse `json:"to"`
	DepartureTime time.Time    `json:"departure_time"`
	ArrivalTime   time.Time    `json:"arrival_time"`
	Duration      string       `json:"duration"` // e.g., "2h 30m"
	Price         float64      `json:"price"`
	Distance      int          `json:"distance"`
	SeatCount     int          `json:"seat_count"`
}

// StopResponse represents a stop/station
type StopResponse struct {
	ID        string  `json:"id"`
	Name      string  `json:"name"`
	City      string  `json:"city"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

// RouteSearchResponse represents search results
type RouteSearchResponse struct {
	Routes         []RouteResponse    `json:"routes"`
	SearchCriteria SearchRouteRequest `json:"search_criteria"`
}

// RouteDetailsResponse represents detailed route information
type RouteDetailsResponse struct {
	Route                RouteResponse           `json:"route"`
	CommissionBreakdown  *CommissionBreakdown    `json:"commission_breakdown,omitempty"`
	InsuranceAvailable   bool                    `json:"insurance_available"`
	InsurancePremium     float64                 `json:"insurance_premium,omitempty"`
	InsuranceBreakdown   *InsuranceBreakdown     `json:"insurance_breakdown,omitempty"`
}

// CommissionBreakdown shows pricing breakdown with commission
type CommissionBreakdown struct {
	BasePrice  float64                    `json:"base_price"`
	Commission float64                    `json:"commission"`
	GrandTotal float64                    `json:"grand_total"`
	Segments   []SegmentCommissionDetails `json:"segments"`
}

// SegmentCommissionDetails shows commission for individual segment
type SegmentCommissionDetails struct {
	SegmentID      string  `json:"segment_id"`
	TransportType  string  `json:"transport_type"`
	BasePrice      float64 `json:"base_price"`
	CommissionRate float64 `json:"commission_rate"` // e.g., 0.07 for 7%
	Commission     float64 `json:"commission"`
	Total          float64 `json:"total"`
}

// InsuranceBreakdown shows insurance calculation details
type InsuranceBreakdown struct {
	BasePremium                float64 `json:"base_premium"`
	TightConnectionSurcharge   float64 `json:"tight_connection_surcharge,omitempty"`
	NightFlightSurcharge       float64 `json:"night_flight_surcharge,omitempty"`
	RiverTransportSurcharge    float64 `json:"river_transport_surcharge,omitempty"`
	MultiSegmentSurcharge      float64 `json:"multi_segment_surcharge,omitempty"`
	Total                      float64 `json:"total"`
}
