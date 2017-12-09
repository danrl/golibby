package trie

import "fmt"

// Node is a node in a trie data structure
type Node struct {
	keys map[string]*Node
	data interface{}
	set  bool
}

var (
	// ErrorNotFound indicates that the requested node does not exist
	ErrorNotFound = fmt.Errorf("not found")
	// ErrorNoData indicates that the requested node does not hold data
	ErrorNoData = fmt.Errorf("no data")
)

// New creates a new, empty trie
func New() *Node {
	return &Node{
		keys: make(map[string]*Node),
	}
}

func (n *Node) node(path []string, create bool) (*Node, error) {
	if len(path) == 0 {
		return n, nil
	}
	nd, ok := n.keys[path[0]]
	if !ok {
		if !create {
			return nil, ErrorNotFound
		}
		nd = &Node{
			keys: make(map[string]*Node),
		}
		n.keys[path[0]] = nd
	}
	return nd.node(path[1:], create)
}

// Set assigns arbitrary data to a node identified by a path of keys
func (n *Node) Set(path []string, data interface{}) {
	nd, _ := n.node(path, true)
	nd.data = data
	nd.set = true
}

// Data retrieves data  assigned to a node identified by a path of keys
func (n *Node) Data(path []string) (interface{}, error) {
	nd, err := n.node(path, false)
	if err != nil {
		return nil, err
	}
	if !nd.set {
		return nil, ErrorNoData
	}
	return nd.data, nil
}

func (n *Node) remove(path []string) (int, error) {
	// remove data and return if this node is the last node in the path
	if len(path) == 0 {
		n.data = nil
		n.set = false
		return len(n.keys), nil
	}
	// find next node in path
	nd, ok := n.keys[path[0]]
	if !ok {
		return 0, ErrorNotFound
	}
	// recurse down
	nk, err := nd.remove(path[1:])
	if err != nil {
		return 0, err
	}
	// remove empty leaf node
	if nk == 0 {
		delete(n.keys, path[0])
	}
	return len(n.keys), nil
}

// Remove deletes data from a node identified by a path of keys
func (n *Node) Remove(path []string) error {
	_, err := n.remove(path)
	return err
}
