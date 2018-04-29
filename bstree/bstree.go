package bstree

import (
	"fmt"
	"sync"

	"github.com/danrl/golibby/util"
)

// Item holds the key and value of a node to be returned by an iterator
type Item struct {
	Key string
	Val interface{}
}

type node struct {
	key   string
	val   interface{}
	left  *node
	right *node
}

// BSTree represents a binary search tree
type BSTree struct {
	lock sync.RWMutex
	root *node
}

// ErrorNotFound is returned when a key is not in the binary search tree
var ErrorNotFound = fmt.Errorf("not found")

func (n *node) value(key string) (interface{}, error) {
	if n == nil {
		return nil, ErrorNotFound
	}
	if key < n.key {
		return n.left.value(key)
	} else if key > n.key {
		return n.right.value(key)
	}
	return n.val, nil
}

// Value returns the data associated with a given key
func (b *BSTree) Value(key string) (interface{}, error) {
	b.lock.RLock()
	defer b.lock.RUnlock()
	return b.root.value(key)
}

func (n *node) upsert(key string, val interface{}) {
	if key < n.key {
		if n.left == nil {
			n.left = &node{key: key, val: val}
		} else {
			n.left.upsert(key, val)
		}
	} else if key > n.key {
		if n.right == nil {
			n.right = &node{key: key, val: val}
		} else {
			n.right.upsert(key, val)
		}
	} else {
		n.val = val
	}
}

// Upsert updates or inserts data associated to a given key
func (b *BSTree) Upsert(key string, val interface{}) {
	b.lock.Lock()
	defer b.lock.Unlock()
	// if root node is empty, new node is root now
	if b.root == nil {
		b.root = &node{key: key, val: val}
	} else {
		b.root.upsert(key, val)
	}
}

func (n *node) isLeaf() bool {
	return !n.hasLeft() && !n.hasRight()
}

func (n *node) hasLeft() bool {
	return n.left != nil
}

func (n *node) hasRight() bool {
	return n.right != nil
}

func (n *node) min() *node {
	for ; n.left != nil; n = n.left {
	}
	return n
}

func (n *node) delete(key string) (*node, error) {
	var err error
	if n == nil {
		return nil, ErrorNotFound
	}
	if key < n.key {
		n.left, err = n.left.delete(key)
		return n, err
	}
	if key > n.key {
		n.right, err = n.right.delete(key)
		return n, err
	}
	// case 1: node is leaf node
	if n.isLeaf() {
		return nil, nil
	}
	// case 2a: node has left child only
	if n.hasLeft() && !n.hasRight() {
		return n.left, nil
	}
	// case 2b: node has right child only
	if n.hasRight() && !n.hasLeft() {
		return n.right, nil
	}
	// case 3: node has two children
	min := n.right.min()
	n.key = min.key
	n.val = min.val
	n.right, err = n.right.delete(min.key)
	return n, err
}

// Delete removes a key and associated data from a binsary search tree
func (b *BSTree) Delete(k string) error {
	var err error
	b.lock.RLock()
	b.root, err = b.root.delete(k)
	b.lock.RUnlock()
	return err
}

func (n *node) height() int {
	if n == nil {
		return 0
	}
	return 1 + util.Max(n.left.height(), n.right.height())
}

// Height returns the height of a binary search tree
func (b *BSTree) Height() int {
	return b.root.height()
}

func (n *node) iter(ch chan<- Item) {
	if n == nil {
		return
	}
	n.left.iter(ch)
	ch <- Item{
		Key: n.key,
		Val: n.val,
	}
	n.right.iter(ch)
}

// Iter provides an iterator to walk through the binary search tree
func (b *BSTree) Iter() <-chan Item {
	ch := make(chan Item)
	b.lock.RLock()
	go func() {
		b.root.iter(ch)
		b.lock.RUnlock()
		close(ch)
	}()
	return ch
}
