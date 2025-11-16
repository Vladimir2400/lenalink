package graph

import "sync"

// Graph represents a transportation network
type Graph struct {
	mu    sync.RWMutex
	Nodes map[string]*Node    // node_id -> Node
	Edges map[string][]*Edge  // from_node_id -> []Edge
}

// NewGraph creates an empty graph
func NewGraph() *Graph {
	return &Graph{
		Nodes: make(map[string]*Node),
		Edges: make(map[string][]*Edge),
	}
}

// AddNode adds a node to the graph
func (g *Graph) AddNode(node *Node) {
	g.mu.Lock()
	defer g.mu.Unlock()
	g.Nodes[node.ID] = node
}

// AddEdge adds an edge to the graph
func (g *Graph) AddEdge(edge *Edge) {
	g.mu.Lock()
	defer g.mu.Unlock()

	// Ensure both nodes exist
	if _, exists := g.Nodes[edge.FromNodeID]; !exists {
		return
	}
	if _, exists := g.Nodes[edge.ToNodeID]; !exists {
		return
	}

	// Add edge
	g.Edges[edge.FromNodeID] = append(g.Edges[edge.FromNodeID], edge)
}

// GetNode retrieves a node by ID
func (g *Graph) GetNode(nodeID string) (*Node, bool) {
	g.mu.RLock()
	defer g.mu.RUnlock()
	node, exists := g.Nodes[nodeID]
	return node, exists
}

// GetNeighbors returns all edges from a given node (returns a copy to avoid race conditions)
func (g *Graph) GetNeighbors(nodeID string) []*Edge {
	g.mu.RLock()
	defer g.mu.RUnlock()

	edges, exists := g.Edges[nodeID]
	if !exists {
		return []*Edge{}
	}

	// Return a copy to prevent race conditions
	edgesCopy := make([]*Edge, len(edges))
	copy(edgesCopy, edges)
	return edgesCopy
}

// NodeCount returns the number of nodes in the graph
func (g *Graph) NodeCount() int {
	g.mu.RLock()
	defer g.mu.RUnlock()
	return len(g.Nodes)
}

// EdgeCount returns the total number of edges in the graph
func (g *Graph) EdgeCount() int {
	g.mu.RLock()
	defer g.mu.RUnlock()
	count := 0
	for _, edges := range g.Edges {
		count += len(edges)
	}
	return count
}
