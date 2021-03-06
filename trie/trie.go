package trie

import (
	"sync"
)

// Trie holds the trie's root node
type Trie struct {
	lock sync.RWMutex
	root *node
}

// Upsert assigns arbitrary data to a node in a trie identified by a path of
// keys
func (t *Trie) Upsert(path []string, value interface{}) {
	t.lock.Lock()
	defer t.lock.Unlock()
	if t.root == nil {
		t.root = &node{}
	}
	t.root.upsert(path, value)
}

// Data retrieves data assigned to a node in a trie identified by a path of keys
func (t *Trie) Data(path []string) (interface{}, error) {
	t.lock.RLock()
	defer t.lock.RUnlock()
	return t.root.data(path)
}

// Delete deletes data from a trie identified by a path of keys
func (t *Trie) Delete(path []string) error {
	t.lock.Lock()
	defer t.lock.Unlock()
	_, err := t.root.delete(path)
	return err
}
