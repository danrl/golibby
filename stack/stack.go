package stack

import "fmt"

// Stack represents a stack
type Stack struct {
	data []interface{}
}

// ErrorEmptyStack is returned on illegal operations on an empty stack
var ErrorEmptyStack = fmt.Errorf("empty stack")

// New creates a new stack
func New() *Stack {
	var s Stack
	return &s
}

// Len returns the number of items in a stack
func (s *Stack) Len() int {
	return len(s.data)
}

// Push adds an item to the stack
func (s *Stack) Push(item interface{}) {
	s.data = append(s.data, item)
}

// Peek returns the top item from the stack without removing it
func (s *Stack) Peek() (interface{}, error) {
	if len(s.data) == 0 {
		return nil, ErrorEmptyStack
	}
	return s.data[len(s.data)-1], nil
}

// Pop returns the top item from the stack
func (s *Stack) Pop() (interface{}, error) {
	if len(s.data) == 0 {
		return nil, ErrorEmptyStack
	}
	item := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return item, nil
}
