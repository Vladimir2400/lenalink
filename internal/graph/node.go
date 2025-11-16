package graph

// Node represents a location in the transportation network
type Node struct {
	ID        string  // Unique identifier
	Name      string  // City or station name
	Latitude  float64 // Geographic latitude
	Longitude float64 // Geographic longitude
	Type      string  // airport, train_station, bus_terminal, port, city_center
}

// NewNode creates a new graph node
func NewNode(id, name string, lat, lon float64, nodeType string) *Node {
	return &Node{
		ID:        id,
		Name:      name,
		Latitude:  lat,
		Longitude: lon,
		Type:      nodeType,
	}
}
