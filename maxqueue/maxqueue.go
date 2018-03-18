package maxqueue

import (
	"fmt"
	"sync"

	"github.com/danrl/golibby/queue"
)

// MaxQueue represents a maxqueue
type MaxQueue struct {
	lock   sync.RWMutex
	queue  queue.Queue
	maxlen int
}

var (
	// ErrorEmpty is returned on illegal operations on an empty maxqueue
	ErrorEmpty = fmt.Errorf("empty queue")
	// ErrorFull is returned on illegal operations on a full maxqueue
	ErrorFull = fmt.Errorf("full queue")
	// ErrorIllegalLength is returned on illegal maximum length
	ErrorIllegalLength = fmt.Errorf("illegal legnth")
)

// New creates a new maxqueue
func New(maxlen int) (*MaxQueue, error) {
	if maxlen < 1 {
		return nil, ErrorIllegalLength
	}
	return &MaxQueue{
		maxlen: maxlen,
	}, nil
}

// Len returns the number of items in the maxqueue
func (q *MaxQueue) Len() int {
	q.lock.RLock()
	defer q.lock.RUnlock()
	return q.queue.Len()
}

// Add adds an item at the end of the maxqueue
func (q *MaxQueue) Add(item interface{}) error {
	q.lock.Lock()
	defer q.lock.Unlock()
	if q.queue.Len() >= q.maxlen {
		return ErrorFull
	}
	q.queue.Add(item)
	return nil
}

// Peek returns the first item from the maxqueue without removing it
func (q *MaxQueue) Peek() (interface{}, error) {
	q.lock.RLock()
	defer q.lock.RUnlock()
	item, err := q.queue.Peek()
	if err != nil {
		return nil, ErrorEmpty
	}
	return item, nil
}

// Remove returns the first item from the maxqueue
func (q *MaxQueue) Remove() (interface{}, error) {
	q.lock.Lock()
	defer q.lock.Unlock()
	item, err := q.queue.Remove()
	if err != nil {
		return nil, ErrorEmpty
	}
	return item, nil
}
