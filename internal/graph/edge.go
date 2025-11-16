package graph

import "time"

// Edge represents a connection between two nodes (a transport segment)
type Edge struct {
	ID            string        // Unique identifier
	FromNodeID    string        // Source node
	ToNodeID      string        // Destination node
	TransportType string        // air, rail, bus, river, walk, taxi
	Provider      string        // Operator name (S7, RZD, etc.)
	Distance      float64       // Distance in kilometers
	Duration      time.Duration // Travel time
	Price         float64       // Price in currency
	DepartureTime time.Time     // When it departs (optional, for scheduled transport)
	ArrivalTime   time.Time     // When it arrives (optional)
}

// NewEdge creates a new graph edge
func NewEdge(id, fromNode, toNode, transportType, provider string, distance float64, duration time.Duration, price float64) *Edge {
	return &Edge{
		ID:            id,
		FromNodeID:    fromNode,
		ToNodeID:      toNode,
		TransportType: transportType,
		Provider:      provider,
		Distance:      distance,
		Duration:      duration,
		Price:         price,
	}
}

// Weight calculates the weight for pathfinding algorithms
// By default, uses duration (in minutes)
func (e *Edge) Weight() float64 {
	return e.Duration.Minutes()
}
