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

func (n *node) setHeight(height int, left bool) {
	old := n.height()
	if left {
		n.leftHeight = height
	} else {
		n.rightHeight = height
	}
	if n.parent == nil {
		return
	}
	if old == n.height() {
		return
	}
	// update parent
	if n == n.parent.left {
		n.parent.setHeight(1+height, true)
	} else if n == n.parent.right {
		n.parent.setHeight(1+height, false)
	}
}

func (n *node) setLeftHeight(height int) {
	n.setHeight(height, true)
	return
}

func (n *node) setRightHeight(height int) {
	n.setHeight(height, false)
	return
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
	n.setLeftHeight(1)
}

func (n *node) newRightNode(key string, value interface{}) {
	n.right = newNode(n, key, value)
	n.setRightHeight(1)
}

/*
 *    8 <= n
 *     \
 *      9     --> 9 <= nd
 *               /
 *              8
 * ---
 *   4 <= n
 *  / \
 * 2   9    -->      9 <= nd
 *    / \           / \
 *   8   12        4   12
 *                / \
 *               2   8
 */
func (n *node) leftRotate() *node {
	// old root's right child becomes new root
	nr := n.right
	nr.parent = n.parent
	// new root's left child becomes old root's right child
	n.right = nr.left
	if n.right != nil {
		n.right.parent = n
	}
	// old root becomes new's left child
	nr.left = n
	nr.left.parent = nr

	// update heights
	if n.hasRight() {
		n.setRightHeight(n.right.height() + 1)
	} else {
		n.setRightHeight(0)
	}
	nr.setLeftHeight(nr.left.height() + 1)
	return nr
}

/*
 *    8 <= n
 *   /
 *  6         --> 6 <= nd
 *                 \
 *                  8
 * ---
 *      8 <= n
 *     / \
 *    4   9  -->   4 <= nd
 *   / \          / \
 *  2   5        2   8
 *                  / \
 *                 5   9
 */

func (n *node) rightRotate() *node {
	// old root's left child becomes new root
	nr := n.left
	nr.parent = n.parent
	// new root's right child becomes old root's left child
	n.left = nr.right
	if n.left != nil {
		n.left.parent = n
	}
	// old root becomes new's right child
	nr.right = n
	nr.right.parent = nr

	// update heights
	if n.hasLeft() {
		n.setLeftHeight(n.left.height() + 1)
	} else {
		n.setLeftHeight(0)
	}
	nr.setRightHeight(nr.right.height() + 1)
	return nr
}

/*
 *    2             2             2
 *   / \           / \           / \
 *  1   7    -->  1   7    -->  1   4    -->      4
 *     / \           / \           / \           / \
 *    4   9         4   9         5   7         2   7
 *                   \                 \       / \   \
 *                   *5*                9     1   5   9
 */
func (n *node) balance() *node {
	if n.hasLeftViolation() {
		if n.left.hasRightImbalance() {
			/*
			 *     8          8
			 *    /          /
			 *   4    ->    6
			 *    \        /
			 *     6      4
			 */
			n.left = n.left.leftRotate()
		}
		/*
		 *      8
		 *     /
		 *    6    ->    6
		 *   /          / \
		 *  4          4   8
		 */
		n = n.rightRotate()
	} else if n.hasRightViolation() {
		if n.right.hasLeftImbalance() {
			/*
			 *  4          4
			 *   \          \
			 *    8    ->    6
			 *   /            \
			 *  6              8
			 */
			n.right = n.right.rightRotate()
		}
		/*
		 *  4
		 *   \
		 *    6    ->    6
		 *     \        / \
		 *      8      4   8
		 */
		n = n.leftRotate()
	}
	return n
}

func (n *node) upsert(key string, value interface{}) *node {
	if n == nil {
		return newNode(nil, key, value)
	}
	if key == n.key {
		n.value = value
		return n
	}
	if key < n.key {
		if n.hasLeft() {
			n.left = n.left.upsert(key, value)
		} else {
			n.newLeftNode(key, value)
		}
	} else {
		if n.hasRight() {
			n.right = n.right.upsert(key, value)
		} else {
			n.newRightNode(key, value)
		}
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
