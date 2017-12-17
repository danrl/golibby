package avltree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewOrphanNode(t *testing.T) {
	nd := newNode(nil, "foo", 1337)
	assert.Equal(t, (*node)(nil), nd.parent)
	assert.Equal(t, "foo", nd.key)
	assert.Equal(t, 1337, nd.value)
}

func TestNodeNewLeftNode(t *testing.T) {
	nd := newNode(nil, "", 0)
	nd.newLeftNode("foo", 1337)
	assert.Equal(t, 1, nd.height())
	assert.Equal(t, 1, nd.leftHeight)
	assert.Equal(t, 0, nd.rightHeight)
	assert.Equal(t, nd, nd.left.parent)
	assert.Equal(t, "foo", nd.left.key)
	assert.Equal(t, 1337, nd.left.value)
	assert.Equal(t, 0, nd.left.height())
	assert.Equal(t, 0, nd.left.leftHeight)
	assert.Equal(t, 0, nd.left.rightHeight)
}

func TestNodeNewRightNode(t *testing.T) {
	nd := newNode(nil, "", 0)
	nd.newRightNode("foo", 1337)
	assert.Equal(t, 1, nd.height())
	assert.Equal(t, 0, nd.leftHeight)
	assert.Equal(t, 1, nd.rightHeight)
	assert.Equal(t, nd, nd.right.parent)
	assert.Equal(t, "foo", nd.right.key)
	assert.Equal(t, 1337, nd.right.value)
	assert.Equal(t, 0, nd.right.height())
	assert.Equal(t, 0, nd.right.leftHeight)
	assert.Equal(t, 0, nd.right.rightHeight)
}

func TestNodeHasLeft(t *testing.T) {
	nd := newNode(nil, "", nil)
	assert.Equal(t, false, nd.hasLeft())

	nd.newLeftNode("", nil)
	assert.Equal(t, true, nd.hasLeft())
}

func TestNodeHasRight(t *testing.T) {
	nd := newNode(nil, "", nil)
	assert.Equal(t, false, nd.hasRight())

	nd.newRightNode("", nil)
	assert.Equal(t, true, nd.hasRight())
}

func TestNodeHeight(t *testing.T) {
	nd := newNode(nil, "", nil)
	assert.Equal(t, 0, nd.height())

	nd.newLeftNode("", nil)
	assert.Equal(t, 1, nd.height())

	nd.newRightNode("", nil)
	assert.Equal(t, 1, nd.height())

	nd.right.newRightNode("", nil)
	assert.Equal(t, 2, nd.height())
}

func TestNodeSetHeight(t *testing.T) {
	{
		nd := newNode(nil, "", nil)
		nd.setHeight(1, true)
		assert.Equal(t, 1, nd.leftHeight)
		assert.Equal(t, 0, nd.rightHeight)
	}
	{
		nd := newNode(nil, "", nil)
		nd.setHeight(1, false)
		assert.Equal(t, 0, nd.leftHeight)
		assert.Equal(t, 1, nd.rightHeight)
	}
}

func TestNodeSetLeftHeight(t *testing.T) {
	nd := newNode(nil, "", nil)
	nd.setLeftHeight(1)
	assert.Equal(t, 1, nd.leftHeight)
	assert.Equal(t, 0, nd.rightHeight)
}

func TestNodeSetRightHeight(t *testing.T) {
	nd := newNode(nil, "", nil)
	nd.setRightHeight(1)
	assert.Equal(t, 0, nd.leftHeight)
	assert.Equal(t, 1, nd.rightHeight)
}

func TestNodeLeftRotate(t *testing.T) {
	{
		/*
		 *    8 <= n
		 *     \
		 *      9     --> 9 <= nd
		 *               /
		 *              8
		 */
		nd := newNode(nil, "8", nil)
		nd.newRightNode("9", nil)

		nd = nd.leftRotate()

		assert.Equal(t, "9", nd.key)
		assert.Equal(t, (*node)(nil), nd.parent)
		assert.Equal(t, 1, nd.height())
		assert.Equal(t, 1, nd.leftHeight)
		assert.Equal(t, 0, nd.rightHeight)
		assert.Equal(t, true, nd.hasLeft())
		assert.Equal(t, false, nd.hasRight())

		assert.Equal(t, "8", nd.left.key)
		assert.Equal(t, 0, nd.left.height())
		assert.Equal(t, 0, nd.left.leftHeight)
		assert.Equal(t, 0, nd.left.rightHeight)
		assert.Equal(t, false, nd.left.hasLeft())
		assert.Equal(t, false, nd.left.hasRight())
	}
	{
		/*
		 *   4 <= n
		 *  / \
		 * 2   9    -->      9 <= nd
		 *    / \           / \
		 *   8   12        4   12
		 *                / \
		 *               2   8
		 */
		nd := newNode(nil, "4", nil)
		nd.newLeftNode("2", nil)
		nd.newRightNode("9", nil)
		nd.right.newLeftNode("8", nil)
		nd.right.newRightNode("12", nil)

		nd = nd.leftRotate()

		assert.Equal(t, "9", nd.key)
		assert.Equal(t, (*node)(nil), nd.parent)
		assert.Equal(t, 2, nd.height())
		assert.Equal(t, 2, nd.leftHeight)
		assert.Equal(t, 1, nd.rightHeight)
		assert.Equal(t, true, nd.hasLeft())
		assert.Equal(t, true, nd.hasRight())

		assert.Equal(t, "4", nd.left.key)
		assert.Equal(t, 1, nd.left.height())
		assert.Equal(t, 1, nd.left.leftHeight)
		assert.Equal(t, 1, nd.left.rightHeight)
		assert.Equal(t, true, nd.left.hasLeft())
		assert.Equal(t, true, nd.left.hasRight())

		assert.Equal(t, "2", nd.left.left.key)
		assert.Equal(t, 0, nd.left.left.height())
		assert.Equal(t, 0, nd.left.left.leftHeight)
		assert.Equal(t, 0, nd.left.left.rightHeight)
		assert.Equal(t, false, nd.left.left.hasLeft())
		assert.Equal(t, false, nd.left.left.hasRight())

		assert.Equal(t, "8", nd.left.right.key)
		assert.Equal(t, 0, nd.left.right.height())
		assert.Equal(t, 0, nd.left.right.leftHeight)
		assert.Equal(t, 0, nd.left.right.rightHeight)
		assert.Equal(t, false, nd.left.right.hasLeft())
		assert.Equal(t, false, nd.left.right.hasRight())

		assert.Equal(t, "12", nd.right.key)
		assert.Equal(t, 0, nd.right.height())
		assert.Equal(t, 0, nd.right.leftHeight)
		assert.Equal(t, 0, nd.right.rightHeight)
		assert.Equal(t, false, nd.right.hasLeft())
		assert.Equal(t, false, nd.right.hasRight())
	}
}

func TestNodeRightRotate(t *testing.T) {
	{
		/*
		 *     8 <= n
		 *    /
		 *   6         -->  6 <= nd
		 *                   \
		 *                    8
		 */
		nd := newNode(nil, "8", nil)
		nd.newLeftNode("6", nil)

		nd = nd.rightRotate()

		assert.Equal(t, "6", nd.key)
		assert.Equal(t, (*node)(nil), nd.parent)
		assert.Equal(t, 1, nd.height())
		assert.Equal(t, 0, nd.leftHeight)
		assert.Equal(t, 1, nd.rightHeight)
		assert.Equal(t, false, nd.hasLeft())
		assert.Equal(t, true, nd.hasRight())

		assert.Equal(t, "8", nd.right.key)
		assert.Equal(t, 0, nd.right.height())
		assert.Equal(t, 0, nd.right.leftHeight)
		assert.Equal(t, 0, nd.right.rightHeight)
		assert.Equal(t, false, nd.right.hasLeft())
		assert.Equal(t, false, nd.right.hasRight())
	}
	{
		/*
		 *     8 <= n
		 *    / \
		 *   4   9  -->   4 <= nd
		 *  / \          / \
		 * 2   5        2   8
		 *                 / \
		 *                5   9
		 */
		nd := newNode(nil, "8", nil)
		nd.newLeftNode("4", nil)
		nd.left.newLeftNode("2", nil)
		nd.left.newRightNode("5", nil)
		nd.newRightNode("9", nil)

		nd = nd.rightRotate()

		assert.Equal(t, "4", nd.key)
		assert.Equal(t, (*node)(nil), nd.parent)
		assert.Equal(t, 2, nd.height())
		assert.Equal(t, 1, nd.leftHeight)
		assert.Equal(t, 2, nd.rightHeight)
		assert.Equal(t, true, nd.hasLeft())
		assert.Equal(t, true, nd.hasRight())

		assert.Equal(t, "2", nd.left.key)
		assert.Equal(t, 0, nd.left.height())
		assert.Equal(t, 0, nd.left.leftHeight)
		assert.Equal(t, 0, nd.left.rightHeight)
		assert.Equal(t, false, nd.left.hasLeft())
		assert.Equal(t, false, nd.left.hasRight())

		assert.Equal(t, "8", nd.right.key)
		assert.Equal(t, 1, nd.right.height())
		assert.Equal(t, 1, nd.right.leftHeight)
		assert.Equal(t, 1, nd.right.rightHeight)
		assert.Equal(t, true, nd.right.hasLeft())
		assert.Equal(t, true, nd.right.hasRight())

		assert.Equal(t, "5", nd.right.left.key)
		assert.Equal(t, 0, nd.right.left.height())
		assert.Equal(t, 0, nd.right.left.leftHeight)
		assert.Equal(t, 0, nd.right.left.rightHeight)
		assert.Equal(t, false, nd.right.left.hasLeft())
		assert.Equal(t, false, nd.right.left.hasRight())

		assert.Equal(t, "9", nd.right.right.key)
		assert.Equal(t, 0, nd.right.right.height())
		assert.Equal(t, 0, nd.right.right.leftHeight)
		assert.Equal(t, 0, nd.right.right.rightHeight)
		assert.Equal(t, false, nd.right.right.hasLeft())
		assert.Equal(t, false, nd.right.right.hasRight())
	}
}

func TestNodeHasLeftViolation(t *testing.T) {
	nd := newNode(nil, "", nil)
	assert.Equal(t, false, nd.hasLeftViolation())

	nd.leftHeight++
	assert.Equal(t, false, nd.hasLeftViolation())

	nd.leftHeight++
	assert.Equal(t, true, nd.hasLeftViolation())
}

func TestNodeHasRightViolation(t *testing.T) {
	nd := newNode(nil, "", nil)
	assert.Equal(t, false, nd.hasRightViolation())

	nd.rightHeight++
	assert.Equal(t, false, nd.hasRightViolation())

	nd.rightHeight++
	assert.Equal(t, true, nd.hasRightViolation())
}

func TestNodeHasLeftImbalance(t *testing.T) {
	nd := newNode(nil, "", nil)
	assert.Equal(t, false, nd.hasLeftImbalance())

	nd.leftHeight++
	assert.Equal(t, true, nd.hasLeftImbalance())
}

func TestNodeHasRightImbalance(t *testing.T) {
	nd := newNode(nil, "", nil)
	assert.Equal(t, false, nd.hasRightImbalance())

	nd.rightHeight++
	assert.Equal(t, true, nd.hasRightImbalance())
}

func TestNodeBalance(t *testing.T) {
	// balanced
	{
		nd := newNode(nil, "2", nil)
		nd.newLeftNode("1", nil)
		nd.newRightNode("3", nil)

		nd = nd.balance()
		assert.Equal(t, "2", nd.key)
		assert.Equal(t, "1", nd.left.key)
		assert.Equal(t, "3", nd.right.key)
	}
	// left imbalance
	{
		/*
		 *      8
		 *     /
		 *    6    ->    6
		 *   /          / \
		 *  4          4   8
		 */
		nd := newNode(nil, "8", nil)
		nd.newLeftNode("6", nil)
		nd.left.newLeftNode("4", nil)

		nd = nd.balance()
		assert.Equal(t, "6", nd.key)
		assert.Equal(t, "4", nd.left.key)
		assert.Equal(t, "8", nd.right.key)
	}
	// right imbalance
	{
		/*
		 *  4
		 *   \
		 *    6    ->    6
		 *     \        / \
		 *      8      4   8
		 */
		nd := newNode(nil, "4", nil)
		nd.newRightNode("6", nil)
		nd.right.newRightNode("8", nil)

		nd = nd.balance()
		assert.Equal(t, "6", nd.key)
		assert.Equal(t, "4", nd.left.key)
		assert.Equal(t, "8", nd.right.key)
	}
	// left-right imbalance
	{
		/*
		 *  4
		 *   \
		 *    8    ->     6
		 *   /           / \
		 *  6           4   8
		 */
		nd := newNode(nil, "4", nil)
		nd.newRightNode("8", nil)
		nd.right.newLeftNode("6", nil)

		nd = nd.balance()
		assert.Equal(t, "6", nd.key)
		assert.Equal(t, "4", nd.left.key)
		assert.Equal(t, "8", nd.right.key)
	}
	// right-left imbalance
	{
		/*
		 *     8
		 *    /
		 *   4    ->    6
		 *    \        / \
		 *     6      4   8
		 */
		nd := newNode(nil, "8", nil)
		nd.newLeftNode("4", nil)
		nd.left.newRightNode("6", nil)

		nd = nd.balance()
		assert.Equal(t, "6", nd.key)
		assert.Equal(t, "4", nd.left.key)
		assert.Equal(t, "8", nd.right.key)
	}
	// complex imbalance
	{
		/*
		 *    2
		 *   / \
		 *  1   7    -->      4
		 *     / \           / \
		 *    4   9         2   7
		 *     \           /   / \
		 *      5         1   5   9
		 */
		nd := newNode(nil, "2", nil)
		nd.newLeftNode("1", nil)
		nd.newRightNode("7", nil)
		nd.right.newLeftNode("4", nil)
		nd.right.newRightNode("9", nil)
		nd.right.left.newRightNode("5", nil)

		nd = nd.balance()
		assert.Equal(t, "4", nd.key)
		assert.Equal(t, "2", nd.left.key)
		assert.Equal(t, "1", nd.left.left.key)
		assert.Equal(t, "7", nd.right.key)
		assert.Equal(t, "5", nd.right.left.key)
		assert.Equal(t, "9", nd.right.right.key)
	}
}

// Some test cases loosely based on https://www.youtube.com/watch?v=7m94k2Qhg68
func TestNodeUpsert(t *testing.T) {
	/*
	 *    2             2
	 *   / \           / \
	 *  1   7    -->  1   7    -->      4
	 *     / \           / \           / \
	 *    4   9         4   9         2   7
	 *                   \           /   / \
	 *                   *5*        1   5   9
	 */
	{
		nd := newNode(nil, "2", nil)
		nd.newLeftNode("1", nil)
		nd.newRightNode("7", nil)
		nd.right.newLeftNode("4", nil)
		nd.right.newRightNode("9", nil)

		nd = nd.upsert("5", nil)
		assert.Equal(t, "4", nd.key)
		assert.Equal(t, "2", nd.left.key)
		assert.Equal(t, "1", nd.left.left.key)
		assert.Equal(t, "7", nd.right.key)
		assert.Equal(t, "5", nd.right.left.key)
		assert.Equal(t, "9", nd.right.right.key)

		/*
		 *      4               4
		 *     / \             / \
		 *    2   7   ->      2   7
		 *   /   / \         /   / \
		 *  1   5   9       1   5   9
		 *                           \
		 *                            91
		 */
		nd = nd.upsert("91", nil)
		assert.Equal(t, "4", nd.key)
		assert.Equal(t, "2", nd.left.key)
		assert.Equal(t, "1", nd.left.left.key)
		assert.Equal(t, "7", nd.right.key)
		assert.Equal(t, "5", nd.right.left.key)
		assert.Equal(t, "9", nd.right.right.key)
		assert.Equal(t, "91", nd.right.right.right.key)

		/*
		 *      4                    7
		 *     / \                 /   \
		 *    2   7       ->      4     90
		 *   /   / \             / \   / \
		 *  1   5   9           2   5 9   91
		 *           \         /
		 *            91      1
		 */
		nd = nd.upsert("90", nil)
		assert.Equal(t, "7", nd.key)
		assert.Equal(t, "4", nd.left.key)
		assert.Equal(t, "90", nd.right.key)
		assert.Equal(t, "2", nd.left.left.key)
		assert.Equal(t, "5", nd.left.right.key)
		assert.Equal(t, "9", nd.right.left.key)
		assert.Equal(t, "91", nd.right.right.key)
		assert.Equal(t, "1", nd.left.left.left.key)
	}
	// change value
	{
		nd := newNode(nil, "foo", 1337)

		nd.upsert("foo", 1338)
		assert.Equal(t, 1338, nd.value)
	}
}

func TestNodeLookup(t *testing.T) {
	// root
	{
		nd := newNode(nil, "2", 2)
		nd.newLeftNode("1", 1)
		nd.newRightNode("3", 3)

		value, err := nd.lookup("2")
		assert.Equal(t, nil, err)
		assert.Equal(t, 2, value)

		value, err = nd.lookup("1")
		assert.Equal(t, nil, err)
		assert.Equal(t, 1, value)

		value, err = nd.lookup("3")
		assert.Equal(t, nil, err)
		assert.Equal(t, 3, value)

		_, err = nd.lookup("bar")
		assert.Equal(t, ErrorNotFound, err)
	}
}
