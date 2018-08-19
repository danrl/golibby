// Package directedgraph implements an directed graph with nodes (vertices),
// edges, and supporting methods.
package directedgraph

import (
	"bytes"
	"fmt"
	"sync"
)

var (
	// ErrorNodeNotFound is returned when trying to access a non-existent node
	ErrorNodeNotFound = fmt.Errorf("node not found")
	// ErrorNodeAlreadyExists is returned when trying to create duplicate nodes
	ErrorNodeAlreadyExists = fmt.Errorf("node already exists")
)

// DirectedGraph holds a graph data structure
type DirectedGraph struct {
	lock  sync.RWMutex
	nodes map[string]interface{}
	edges map[string]map[string]bool
}

// New initializes a new graph
func New() *DirectedGraph {
	return &DirectedGraph{
		nodes: make(map[string]interface{}),
		edges: make(map[string]map[string]bool),
	}
}

// NewNode adds a new node to the graph
func (g *DirectedGraph) NewNode(key string, value interface{}) error {
	g.lock.Lock()
	defer g.lock.Unlock()

	if _, ok := g.nodes[key]; ok {
		return ErrorNodeAlreadyExists
	}
	g.nodes[key] = value
	g.edges[key] = make(map[string]bool)

	return nil
}

// Value retrieves the value assigned to the node identified by key
func (g *DirectedGraph) Value(key string) (interface{}, error) {
	g.lock.RLock()
	defer g.lock.RUnlock()

	if _, ok := g.nodes[key]; !ok {
		return nil, ErrorNodeNotFound
	}
	value := g.nodes[key]
	return value, nil
}

// UpdateValue sets the value of the node identified by key
func (g *DirectedGraph) UpdateValue(key string, value interface{}) error {
	g.lock.Lock()
	defer g.lock.Unlock()

	if _, ok := g.nodes[key]; !ok {
		return ErrorNodeNotFound
	}
	g.nodes[key] = value
	return nil
}

// NewEdge adds an edge between to nodes in the graph
func (g *DirectedGraph) NewEdge(from, to string) error {
	g.lock.Lock()
	defer g.lock.Unlock()

	if _, ok := g.nodes[from]; !ok {
		return ErrorNodeNotFound
	}
	if _, ok := g.nodes[to]; !ok {
		return ErrorNodeNotFound
	}

	g.edges[from][to] = true
	return nil
}

// Edges returns the keys of nodes that are directly connected to the node
func (g *DirectedGraph) Edges(from string) ([]string, error) {
	var edges []string

	g.lock.RLock()
	if _, ok := g.nodes[from]; !ok {
		return edges, ErrorNodeNotFound
	}
	for to := range g.edges[from] {
		if g.edges[from][to] {
			edges = append(edges, to)
		}
	}
	g.lock.RUnlock()

	return edges, nil
}

// Nodes returns a list of all nodes in the graph
func (g *DirectedGraph) Nodes() []string {
	g.lock.RLock()
	defer g.lock.RUnlock()

	i := 0
	nodes := make([]string, len(g.nodes))
	for key := range g.nodes {
		nodes[i] = key
		i++
	}
	return nodes
}

// String returns a human readable multi-line string describing the graph
func (g *DirectedGraph) String() string {
	var out bytes.Buffer

	g.lock.RLock()
	for key, value := range g.nodes {
		out.WriteString(fmt.Sprintf("⦿ `%v` (%v)\n", key, value))
		for to, active := range g.edges[key] {
			if active {
				out.WriteString(fmt.Sprintf("⤷ `%v`\n", to))
			}
		}
	}
	g.lock.RUnlock()

	return out.String()
}
