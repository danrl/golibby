package limitedqueue

import (
	"fmt"
	"sync"

	"github.com/danrl/golibby/queue"
)

// LimitedQueue represents a limited queue
type LimitedQueue struct {
	queue  queue.Queue
	maxlen int
	lock   sync.RWMutex
}

var (
	// ErrorEmpty is returned on illegal operations on an empty limited queue
	ErrorEmpty = fmt.Errorf("empty queue")
	// ErrorFull is returned on illegal operations on a full limited queue
	ErrorFull = fmt.Errorf("full queue")
	// ErrorIllegalLength is returned on illegal maximum length
	ErrorIllegalLength = fmt.Errorf("illegal legnth")
)

// New creates a new limited queue
func New(maxlen int) (*LimitedQueue, error) {
	if maxlen < 1 {
		return nil, ErrorIllegalLength
	}
	return &LimitedQueue{
		maxlen: maxlen,
	}, nil
}

// Len returns the number of items in the limited queue
func (q *LimitedQueue) Len() int {
	q.lock.RLock()
	defer q.lock.RUnlock()
	return q.queue.Len()
}

// Add adds an item at the end of the limited queue
func (q *LimitedQueue) Add(item interface{}) error {
	q.lock.Lock()
	defer q.lock.Unlock()
	if q.queue.Len() >= q.maxlen {
		return ErrorFull
	}
	q.queue.Add(item)
	return nil
}

// Peek returns the first item from the limited queue without removing it
func (q *LimitedQueue) Peek() (interface{}, error) {
	q.lock.RLock()
	defer q.lock.RUnlock()
	item, err := q.queue.Peek()
	if err != nil {
		return nil, ErrorEmpty
	}
	return item, nil
}

// Remove returns the first item from the limited queue
func (q *LimitedQueue) Remove() (interface{}, error) {
	q.lock.Lock()
	defer q.lock.Unlock()
	item, err := q.queue.Remove()
	if err != nil {
		return nil, ErrorEmpty
	}
	return item, nil
}
