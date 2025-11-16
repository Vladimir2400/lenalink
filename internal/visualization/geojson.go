package visualization

import (
	"github.com/lenalink/backend/internal/graph"
	"github.com/lenalink/backend/internal/routing"
)

// GeoJSONFeature represents a GeoJSON feature
type GeoJSONFeature struct {
	Type       string                 `json:"type"`
	Geometry   GeoJSONGeometry        `json:"geometry"`
	Properties map[string]interface{} `json:"properties"`
}

// GeoJSONGeometry represents geometry in GeoJSON
type GeoJSONGeometry struct {
	Type        string      `json:"type"`
	Coordinates interface{} `json:"coordinates"`
}

// GeoJSONFeatureCollection represents a collection of GeoJSON features
type GeoJSONFeatureCollection struct {
	Type     string            `json:"type"`
	Features []GeoJSONFeature  `json:"features"`
}

// GenerateGeoJSON creates GeoJSON from a path
func GenerateGeoJSON(path *routing.Path) *GeoJSONFeatureCollection {
	features := make([]GeoJSONFeature, 0)

	// Add route line
	lineCoords := make([][]float64, 0)
	for _, node := range path.Nodes {
		lineCoords = append(lineCoords, []float64{node.Longitude, node.Latitude})
	}

	routeLine := GeoJSONFeature{
		Type: "Feature",
		Geometry: GeoJSONGeometry{
			Type:        "LineString",
			Coordinates: lineCoords,
		},
		Properties: map[string]interface{}{
			"type":     "route",
			"distance": path.TotalDistance,
			"price":    path.TotalPrice,
		},
	}
	features = append(features, routeLine)

	// Add markers for stops
	for i, node := range path.Nodes {
		marker := GeoJSONFeature{
			Type: "Feature",
			Geometry: GeoJSONGeometry{
				Type:        "Point",
				Coordinates: []float64{node.Longitude, node.Latitude},
			},
			Properties: map[string]interface{}{
				"name":  node.Name,
				"type":  node.Type,
				"order": i,
			},
		}
		features = append(features, marker)
	}

	return &GeoJSONFeatureCollection{
		Type:     "FeatureCollection",
		Features: features,
	}
}

// GetSegmentColor returns color for transport type
func GetSegmentColor(transportType string) string {
	colors := map[string]string{
		"air":   "#FF6B6B",
		"rail":  "#4ECDC4",
		"bus":   "#45B7D1",
		"river": "#96CEB4",
		"walk":  "#FFEAA7",
		"taxi":  "#DFE6E9",
	}
	if color, exists := colors[transportType]; exists {
		return color
	}
	return "#95A5A6"
}
