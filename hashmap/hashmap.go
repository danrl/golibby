package hashmap

import (
	"fmt"
	"sync"

	"github.com/danrl/golibby/hash"
)

// HashMap holds a concurrency-safe hashmap implementation
type HashMap struct {
	lock sync.RWMutex
	data [1 << 16][]item
}

type item struct {
	key   string
	value interface{}
}

// ErrorNotFound indicates that the requested item does not exist
var ErrorNotFound = fmt.Errorf("not found")

// Upsert inserts or updates the value for a given key
func (h *HashMap) Upsert(key string, value interface{}) {
	h.lock.Lock()
	defer h.lock.Unlock()
	offset := hash.Pearson16([]byte(key))
	updated := false
	for i := range h.data[offset] {
		if h.data[offset][i].key == key {
			h.data[offset][i].value = value
			updated = true
			break
		}
	}
	if !updated {
		h.data[offset] = append(h.data[offset], item{
			key:   key,
			value: value,
		})
	}
}

// Delete removes the value for a given key from the hash map
func (h *HashMap) Delete(key string) error {
	h.lock.Lock()
	defer h.lock.Unlock()
	offset := hash.Pearson16([]byte(key))
	for i := range h.data[offset] {
		if h.data[offset][i].key == key {
			h.data[offset] = append(h.data[offset][:i], h.data[offset][i+1:]...)
			return nil
		}
	}
	return ErrorNotFound
}

// Value returns the value for a given key in the hash map
func (h *HashMap) Value(key string) (interface{}, error) {
	h.lock.RLock()
	defer h.lock.RUnlock()
	offset := hash.Pearson16([]byte(key))
	for i := range h.data[offset] {
		if h.data[offset][i].key == key {
			return h.data[offset][i].value, nil
		}
	}
	return nil, ErrorNotFound
}
