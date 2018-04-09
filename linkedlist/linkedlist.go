package linkedlist

import (
	"fmt"
	"sync"
)

type item struct {
	next *item
	val  interface{}
}

// LinkedList represents a single linked list
type LinkedList struct {
	lock sync.RWMutex
	head *item
}

var (
	// ErrorNotFound is returned when an item is not in the single linked list
	ErrorNotFound = fmt.Errorf("not found")
)

// Append adds val to the end of the single linked list
func (s *LinkedList) Append(val interface{}) {
	s.lock.Lock()
	defer s.lock.Unlock()
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

// Remove deletes the first occurrence of val from the single linked list
func (s *LinkedList) Remove(val interface{}) error {
	s.lock.Lock()
	defer s.lock.Unlock()
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
func (s *LinkedList) Iter() <-chan interface{} {
	ch := make(chan interface{})
	s.lock.RLock()
	go func() {
		for cur := s.head; cur != nil; cur = cur.next {
			ch <- cur.val
		}
		s.lock.RUnlock()
		close(ch)
	}()
	return ch
}

// Len returns the number of items in the single linked list
func (s *LinkedList) Len() int {
	s.lock.RLock()
	defer s.lock.RUnlock()
	var i int
	for cur := s.head; cur != nil; cur = cur.next {
		i++
	}
	return i
}
