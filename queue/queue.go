package queue

import "fmt"

// Queue represents a queue
type Queue struct {
	data []interface{}
}

var (
	// ErrorEmpty is returned on illegal operations on an empty queue
	ErrorEmpty = fmt.Errorf("empty queue")
)

// New creates a new queue
func New() *Queue {
	return &Queue{}
}

// Len returns the number of items in the queue
func (q *Queue) Len() int {
	return len(q.data)
}

// Add adds an item at the end of the queue
func (q *Queue) Add(item interface{}) {
	q.data = append(q.data, item)
}

// Peek returns the first item from the queue without removing it
func (q *Queue) Peek() (interface{}, error) {
	if len(q.data) == 0 {
		return nil, ErrorEmpty
	}
	return q.data[0], nil
}

// Remove returns the first item from the queue
func (q *Queue) Remove() (interface{}, error) {
	if len(q.data) == 0 {
		return nil, ErrorEmpty
	}
	item := q.data[0]
	q.data = q.data[1:]
	return item, nil
}
