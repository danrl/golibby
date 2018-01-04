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

// Item holds the key and value of a node to be returned by an iterator
type Item struct {
	Key string
	Val interface{}
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

// Delete removes a key value pair from the AVL tree
func (a *AVLTree) Delete(key string) error {
	a.lock.Lock()
	defer a.lock.Unlock()

	var err error
	a.root, err = a.root.delete(key)
	return err
}

// Iter provides an iterator to walk through the AVL tree
func (a *AVLTree) Iter() <-chan Item {
	ch := make(chan Item)
	a.lock.RLock()
	go func() {
		a.root.iter(ch)
		a.lock.RUnlock()
		close(ch)
	}()
	return ch
}
