package avltree

import (
	"sync"
)

// AVLTree represents a concurrency-safe implementation of a self-balancing
// binary search tree as described by Adelson-Velskii and Landis
type AVLTree struct {
	lock sync.RWMutex
	root *node
}

// New initiates a new AVL tree
func New() *AVLTree {
	return &AVLTree{}
}

// Upsert inserts or updates a key value pair
func (a *AVLTree) Upsert(key string, value interface{}) {
	a.lock.Lock()
	defer a.lock.Unlock()

	a.root = a.root.upsert(key, value)
	return
}

// Lookup retrieves a previously saved value from the AVL tree
func (a *AVLTree) Lookup(key string) (interface{}, error) {
	a.lock.RLock()
	defer a.lock.RUnlock()

	return a.root.lookup(key)
}
