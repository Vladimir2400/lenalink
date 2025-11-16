package graph

import (
	"github.com/lenalink/backend/internal/domain"
	"github.com/lenalink/backend/pkg/utils"
)

// Builder constructs a graph from route data
type Builder struct {
	graph *Graph
}

// NewBuilder creates a new graph builder
func NewBuilder() *Builder {
	return &Builder{
		graph: NewGraph(),
	}
}

// BuildFromRoutes constructs a graph from a list of routes
func (b *Builder) BuildFromRoutes(routes []domain.Route) *Graph {
	// Add all unique cities/stations as nodes
	nodeMap := make(map[string]bool)

	for _, route := range routes {
		for _, segment := range route.Segments {
			// Add start stop as node
			if !nodeMap[segment.StartStop.ID] {
				node := NewNode(
					segment.StartStop.ID,
					segment.StartStop.City,
					segment.StartStop.Latitude,
					segment.StartStop.Longitude,
					"city",
				)
				b.graph.AddNode(node)
				nodeMap[segment.StartStop.ID] = true
			}

			// Add end stop as node
			if !nodeMap[segment.EndStop.ID] {
				node := NewNode(
					segment.EndStop.ID,
					segment.EndStop.City,
					segment.EndStop.Latitude,
					segment.EndStop.Longitude,
					"city",
				)
				b.graph.AddNode(node)
				nodeMap[segment.EndStop.ID] = true
			}

			// Add segment as edge
			edge := NewEdge(
				segment.ID,
				segment.StartStop.ID,
				segment.EndStop.ID,
				string(segment.TransportType),
				segment.Provider,
				float64(segment.Distance),
				segment.Duration,
				segment.Price,
			)
			edge.DepartureTime = segment.DepartureTime
			edge.ArrivalTime = segment.ArrivalTime

			b.graph.AddEdge(edge)
		}
	}

	return b.graph
}

// AddCity manually adds a city node
func (b *Builder) AddCity(id, name string, lat, lon float64) {
	node := NewNode(id, name, lat, lon, "city")
	b.graph.AddNode(node)
}

// AddConnection manually adds a connection between two nodes
func (b *Builder) AddConnection(fromID, toID, transportType, provider string, distance float64, durationMinutes int, price float64) {
	edgeID := utils.GenerateID()
	edge := NewEdge(
		edgeID,
		fromID,
		toID,
		transportType,
		provider,
		distance,
		0, // duration will be calculated
		price,
	)
	b.graph.AddEdge(edge)
}

// GetGraph returns the built graph
func (b *Builder) GetGraph() *Graph {
	return b.graph
}
