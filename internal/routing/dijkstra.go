package routing

import (
	"container/heap"
	"github.com/lenalink/backend/internal/graph"
)

// Path represents a route found by pathfinding
type Path struct {
	Nodes    []*graph.Node
	Edges    []*graph.Edge
	TotalCost float64
	TotalDistance float64
	TotalPrice    float64
}

// DijkstraPathfinder implements Dijkstra's shortest path algorithm
type DijkstraPathfinder struct {
	graph *graph.Graph
}

// NewDijkstraPathfinder creates a new Dijkstra pathfinder
func NewDijkstraPathfinder(g *graph.Graph) *DijkstraPathfinder {
	return &DijkstraPathfinder{graph: g}
}

// FindShortestPath finds the shortest path from start to end node
func (d *DijkstraPathfinder) FindShortestPath(startNodeID, endNodeID string) (*Path, error) {
	// Priority queue for unvisited nodes
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)

	// Distance map: node_id -> cost
	dist := make(map[string]float64)
	// Previous node map: node_id -> previous_node_id
	prev := make(map[string]string)
	// Previous edge map: node_id -> edge_used_to_reach_it
	prevEdge := make(map[string]*graph.Edge)

	// Initialize distances
	for nodeID := range d.graph.Nodes {
		if nodeID == startNodeID {
			dist[nodeID] = 0
			heap.Push(&pq, &Item{
				nodeID:   nodeID,
				priority: 0,
				index:    0,
			})
		} else {
			dist[nodeID] = 1e9 // Infinity
		}
	}

	// Main loop
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		currentNodeID := item.nodeID
		currentDist := dist[currentNodeID]

		// If we reached the end node, we're done
		if currentNodeID == endNodeID {
			break
		}

		// Explore neighbors
		neighbors := d.graph.GetNeighbors(currentNodeID)
		for _, edge := range neighbors {
			alt := currentDist + edge.Weight()
			if alt < dist[edge.ToNodeID] {
				dist[edge.ToNodeID] = alt
				prev[edge.ToNodeID] = currentNodeID
				prevEdge[edge.ToNodeID] = edge

				heap.Push(&pq, &Item{
					nodeID:   edge.ToNodeID,
					priority: alt,
				})
			}
		}
	}

	// Reconstruct path
	return d.reconstructPath(startNodeID, endNodeID, prev, prevEdge), nil
}

// reconstructPath builds the path from prev map
func (d *DijkstraPathfinder) reconstructPath(startID, endID string, prev map[string]string, prevEdge map[string]*graph.Edge) *Path {
	path := &Path{
		Nodes: make([]*graph.Node, 0),
		Edges: make([]*graph.Edge, 0),
	}

	// Build path backwards
	nodeIDs := make([]string, 0)
	currentID := endID
	for currentID != startID {
		nodeIDs = append([]string{currentID}, nodeIDs...)
		currentID = prev[currentID]
		if currentID == "" {
			// No path found
			return path
		}
	}
	nodeIDs = append([]string{startID}, nodeIDs...)

	// Fill nodes and edges
	for _, nodeID := range nodeIDs {
		if node, exists := d.graph.GetNode(nodeID); exists {
			path.Nodes = append(path.Nodes, node)
		}
	}

	for i := 1; i < len(nodeIDs); i++ {
		if edge, exists := prevEdge[nodeIDs[i]]; exists {
			path.Edges = append(path.Edges, edge)
			path.TotalCost += edge.Weight()
			path.TotalDistance += edge.Distance
			path.TotalPrice += edge.Price
		}
	}

	return path
}

// Priority Queue implementation for Dijkstra

type Item struct {
	nodeID   string
	priority float64
	index    int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}
