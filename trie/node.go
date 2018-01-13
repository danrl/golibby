package trie

import (
	"fmt"
)

// node is a node in a trie data structure
type node struct {
	keys  map[string]*node
	value interface{}
	set   bool
}

var (
	// ErrorNotFound indicates that the requested node does not exist
	ErrorNotFound = fmt.Errorf("not found")
	// ErrorNoData indicates that the requested node does not hold data
	ErrorNoData = fmt.Errorf("no data")
)

func (n *node) node(path []string, create bool) (*node, error) {
	if len(path) == 0 {
		return n, nil
	}
	nd, ok := n.keys[path[0]]
	if !ok {
		if !create {
			return nil, ErrorNotFound
		}
		nd = &node{}
		if n.keys == nil {
			n.keys = make(map[string]*node)
		}
		n.keys[path[0]] = nd
	}
	return nd.node(path[1:], create)
}

func (n *node) upsert(path []string, value interface{}) {
	nd, _ := n.node(path, true)
	nd.value = value
	nd.set = true
}

func (n *node) data(path []string) (interface{}, error) {
	if n == nil {
		return nil, ErrorNotFound
	}
	nd, err := n.node(path, false)
	if err != nil {
		return nil, err
	}
	if !nd.set {
		return nil, ErrorNoData
	}
	return nd.value, nil
}

func (n *node) delete(path []string) (int, error) {
	if n == nil {
		return 0, ErrorNotFound
	}
	// remove data and return if this node is the last node in the path
	if len(path) == 0 {
		n.value = nil
		n.set = false
		return len(n.keys), nil
	}
	// find next node in path
	nd, ok := n.keys[path[0]]
	if !ok {
		return 0, ErrorNotFound
	}
	// recurse down
	nk, err := nd.delete(path[1:])
	if err != nil {
		return 0, err
	}
	// remove empty leaf node
	if nk == 0 {
		delete(n.keys, path[0])
	}
	return len(n.keys), nil
}
