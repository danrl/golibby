package avltree

import (
	"fmt"

	"github.com/danrl/golibby/utils"
)

type node struct {
	key         string
	value       interface{}
	parent      *node
	left        *node
	right       *node
	leftHeight  int
	rightHeight int
}

// ErrorNotFound is returned when a key was not found in the AVL tree
var ErrorNotFound = fmt.Errorf("not found")

func (n *node) hasParent() bool {
	return n.parent != nil
}

func (n *node) hasLeft() bool {
	return n.left != nil
}

func (n *node) hasRight() bool {
	return n.right != nil
}

func (n *node) hasLeftViolation() bool {
	return n.leftHeight > (n.rightHeight + 1)
}

func (n *node) hasRightViolation() bool {
	return n.rightHeight > (n.leftHeight + 1)
}

func (n *node) hasLeftImbalance() bool {
	return n.leftHeight > n.rightHeight
}

func (n *node) hasRightImbalance() bool {
	return n.rightHeight > n.leftHeight
}

func (n *node) height() int {
	return utils.Max(n.leftHeight, n.rightHeight)
}

func (n *node) updateHeights() {
	if n == nil {
		return
	}
	if n.hasLeft() {
		n.leftHeight = 1 + n.left.height()
	} else {
		n.leftHeight = 0
	}
	if n.hasRight() {
		n.rightHeight = 1 + n.right.height()
	} else {
		n.rightHeight = 0
	}
}

func newNode(parent *node, key string, value interface{}) *node {
	return &node{
		key:    key,
		value:  value,
		parent: parent,
	}
}

func (n *node) newLeftNode(key string, value interface{}) {
	n.left = newNode(n, key, value)
	n.leftHeight = 1
}

func (n *node) newRightNode(key string, value interface{}) {
	n.right = newNode(n, key, value)
	n.rightHeight = 1
}

//
//   n
//  / \
// x1  nr   -->      nr
//    / \           / \
//   x2 x3         n  x3
//                / \
//               x1 x2
//

func (n *node) leftRotate() *node {
	if n.right == nil {
		panic("node not right rotatable")
	}
	// define nr
	nr := n.right
	nr.parent = n.parent
	// cut out nr
	n.right = n.right.left
	if n.hasRight() {
		n.right.parent = n
	}
	// move n down
	nr.left = n
	nr.left.parent = n
	// update heights
	n.updateHeights()
	nr.updateHeights()
	return nr
}

//
//      n
//     / \
//    nr x3  -->   nr
//   / \          / \
//  x1 x2       x1   n
//                  / \
//                 x2 x3
//
func (n *node) rightRotate() *node {
	if n.left == nil {
		panic("node not right rotatable")
	}
	// define nr
	nr := n.left
	nr.parent = n.parent
	// cut out nr
	n.left = n.left.right
	if n.hasLeft() {
		n.left.parent = n
	}
	// move n down
	nr.right = n
	nr.right.parent = n
	// update heights
	n.updateHeights()
	nr.updateHeights()
	return nr
}

//
//    2             2             2
//   / \           / \           / \
//  1   7    -->  1   7    -->  1   4    -->      4
//     / \           / \           / \           / \
//    4   9         4   9         5   7         2   7
//                   \                 \       / \   \
//                   *5*                9     1   5   9
//
func (n *node) balance() *node {
	if n.hasLeftViolation() {
		if n.left.hasRightImbalance() {
			//
			//     8          8
			//    /          /
			//   4    ->    6
			//    \        /
			//     6      4
			//
			n.left = n.left.leftRotate()
			n.updateHeights()
		}
		//
		//      8
		//     /
		//    6    ->    6
		//   /          / \
		//  4          4   8
		//
		n = n.rightRotate()
	} else if n.hasRightViolation() {
		if n.right.hasLeftImbalance() {
			//
			//  4          4
			//   \          \
			//    8    ->    6
			//   /            \
			//  6              8
			//
			n.right = n.right.rightRotate()
			n.updateHeights()
		}
		//
		//  4
		//   \
		//    6    ->    6
		//     \        / \
		//      8      4   8
		//
		n = n.leftRotate()
	}
	n.updateHeights()
	return n
}

func (n *node) upsert(key string, value interface{}) *node {
	if n == nil {
		return newNode(nil, key, value)
	} else if key < n.key {
		if n.hasLeft() {
			n.left = n.left.upsert(key, value)
		} else {
			n.newLeftNode(key, value)
		}
		n.updateHeights()
	} else if key > n.key {
		if n.hasRight() {
			n.right = n.right.upsert(key, value)
		} else {
			n.newRightNode(key, value)
		}
		n.updateHeights()
	} else {
		n.value = value
	}
	return n.balance()
}

func (n *node) lookup(key string) (interface{}, error) {
	if n == nil {
		return nil, ErrorNotFound
	}
	if key < n.key {
		return n.left.lookup(key)
	}
	if key > n.key {
		return n.right.lookup(key)
	}
	return n.value, nil
}

//
//   02        02
//  /  \  ->  /
// 01  03    01
//
func (n *node) delete(key string) (*node, error) {
	if n == nil {
		return n, ErrorNotFound
	}
	var err error
	if key < n.key {
		n.left, err = n.left.delete(key)
	} else if key > n.key {
		n.right, err = n.right.delete(key)
	} else {
		// delete node
		if !n.hasLeft() && !n.hasRight() {
			// case: leaf node
			return nil, nil
		} else if n.hasLeft() && !n.hasRight() {
			// case: left child only
			return n.left, nil
		} else if !n.hasLeft() && n.hasRight() {
			// case: right child only
			return n.right, nil
		}
		// case: two children
		// find leftmost node of right subtree
		nd := n.right
		for ; nd.hasLeft(); nd = nd.left {
		}
		// replace to-be-deleted node's key value pair with leftmost node's
		// key value pair
		n.key = nd.key
		n.value = nd.value
		// delete leftmost node
		n.right, _ = n.right.delete(nd.key)
	}
	n.updateHeights()
	return n.balance(), err
}
