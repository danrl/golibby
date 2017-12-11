package linkedlist

import "fmt"

type item struct {
	next *item
	val  interface{}
}

// Single represents a single linked list
type Single struct {
	head *item
}

var (
	// ErrorNotFound is returned when an item is not in the single linked list
	ErrorNotFound = fmt.Errorf("not found")
)

// NewSingle creates a new single linked list
func NewSingle() *Single {
	return &Single{}
}

// Append adds val to the end of the single linked list
func (s *Single) Append(val interface{}) {
	new := &item{
		val: val,
	}
	if s.head == nil {
		s.head = new
		return
	}
	var cur *item
	for cur = s.head; cur.next != nil; cur = cur.next {
	}
	cur.next = new
}

// Remove deletes the first occurence of val from the single linked list
func (s *Single) Remove(val interface{}) error {
	var prev *item
	for cur := s.head; cur != nil; cur = cur.next {
		if cur.val == val {
			if prev == nil {
				s.head = cur.next
			} else {
				prev.next = cur.next
			}
			return nil
		}
		prev = cur
	}
	return ErrorNotFound
}

// Iter provides an iterator to walk through the single linked list
func (s *Single) Iter() <-chan interface{} {
	ch := make(chan interface{})
	go func() {
		for cur := s.head; cur != nil; cur = cur.next {
			ch <- cur.val
		}
		close(ch)
	}()
	return ch
}

// Len returns the number of items in the single linked list
func (s *Single) Len() int {
	var i int
	for cur := s.head; cur != nil; cur = cur.next {
		i++
	}
	return i
}
