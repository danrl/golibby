package avltree

import (
	"testing"

	"github.com/danrl/golibby/util"
	"github.com/stretchr/testify/assert"
)

func testhelperRecursiveHeightsUpdate(nd *node) int {
	if nd.left == nil {
		nd.leftHeight = 0
	} else {
		nd.leftHeight = 1 + testhelperRecursiveHeightsUpdate(nd.left)
	}
	if nd.right == nil {
		nd.rightHeight = 0
	} else {
		nd.rightHeight = 1 + testhelperRecursiveHeightsUpdate(nd.right)
	}
	return util.Max(nd.leftHeight, nd.rightHeight)
}

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

func TestNodeHasParent(t *testing.T) {
	nd := newNode(nil, "", nil)
	assert.Equal(t, false, nd.hasParent())

	nd.newLeftNode("", nil)
	assert.Equal(t, true, nd.left.hasParent())
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
}

func TestUpdateHeights(t *testing.T) {
	{
		assert.NotPanics(t, func() { (*node)(nil).updateHeights() })
	}
	{
		nd := newNode(nil, "", nil)
		nd.updateHeights()
		assert.Equal(t, 0, nd.leftHeight)
		assert.Equal(t, 0, nd.rightHeight)

		nd.newLeftNode("", nil)
		nd.left.leftHeight = 3
		nd.newRightNode("", nil)
		nd.right.rightHeight = 7

		nd.updateHeights()
		assert.Equal(t, 4, nd.leftHeight)
		assert.Equal(t, 8, nd.rightHeight)
	}
}

func TestNodeLeftRotate(t *testing.T) {
	{
		nd := newNode(nil, "", nil)
		assert.Panics(t, func() { nd.leftRotate() })
	}
	{
		//
		//    8 <= n
		//     \
		//      9     --> 9 <= nd
		//               /
		//              8
		//
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
		//
		//   4 <= n
		//  / \
		// 2   9    -->      9 <= nd
		//    / \           / \
		//   8   12        4   12
		//                / \
		//               2   8
		//
		nd := newNode(nil, "4", nil)
		nd.newLeftNode("2", nil)
		nd.newRightNode("9", nil)
		nd.right.newLeftNode("8", nil)
		nd.right.newRightNode("12", nil)
		_ = testhelperRecursiveHeightsUpdate(nd)

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
		nd := newNode(nil, "", nil)
		assert.Panics(t, func() { nd.rightRotate() })
	}
	{
		//
		//     8 <= n
		//    /
		//   6         -->  6 <= nd
		//                   \
		//                    8
		//
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
		//
		//     8 <= n
		//    / \
		//   4   9  -->   4 <= nd
		//  / \          / \
		// 2   5        2   8
		//                 / \
		//                5   9
		//
		nd := newNode(nil, "8", nil)
		nd.newLeftNode("4", nil)
		nd.left.newLeftNode("2", nil)
		nd.left.newRightNode("5", nil)
		nd.newRightNode("9", nil)
		_ = testhelperRecursiveHeightsUpdate(nd)

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
		//
		//      8
		//     /
		//    6    ->    6
		//   /          / \
		//  4          4   8
		//
		nd := newNode(nil, "8", nil)
		nd.newLeftNode("6", nil)
		nd.left.newLeftNode("4", nil)
		_ = testhelperRecursiveHeightsUpdate(nd)

		nd = nd.balance()
		assert.Equal(t, "6", nd.key)
		assert.Equal(t, "4", nd.left.key)
		assert.Equal(t, "8", nd.right.key)
	}
	// right imbalance
	{
		//
		//  4
		//   \
		//    6    ->    6
		//     \        / \
		//      8      4   8
		//
		nd := newNode(nil, "4", nil)
		nd.newRightNode("6", nil)
		nd.right.newRightNode("8", nil)
		_ = testhelperRecursiveHeightsUpdate(nd)

		nd = nd.balance()
		assert.Equal(t, "6", nd.key)
		assert.Equal(t, "4", nd.left.key)
		assert.Equal(t, "8", nd.right.key)
	}
	// left-right imbalance
	{
		//
		//  4
		//   \
		//    8    ->     6
		//   /           / \
		//  6           4   8
		//
		nd := newNode(nil, "4", nil)
		nd.newRightNode("8", nil)
		nd.right.newLeftNode("6", nil)
		_ = testhelperRecursiveHeightsUpdate(nd)

		nd = nd.balance()
		assert.Equal(t, "6", nd.key)
		assert.Equal(t, "4", nd.left.key)
		assert.Equal(t, "8", nd.right.key)
	}
	// right-left imbalance
	{
		//
		//     8
		//    /
		//   4    ->    6
		//    \        / \
		//     6      4   8
		//
		nd := newNode(nil, "8", nil)
		nd.newLeftNode("4", nil)
		nd.left.newRightNode("6", nil)
		_ = testhelperRecursiveHeightsUpdate(nd)

		nd = nd.balance()
		assert.Equal(t, "6", nd.key)
		assert.Equal(t, "4", nd.left.key)
		assert.Equal(t, "8", nd.right.key)
	}
	// complex imbalance
	{
		//
		//    2
		//   / \
		//  1   7    -->      4
		//     / \           / \
		//    4   9         2   7
		//     \           /   / \
		//      5         1   5   9
		//
		nd := newNode(nil, "2", nil)
		nd.newLeftNode("1", nil)
		nd.newRightNode("7", nil)
		nd.right.newLeftNode("4", nil)
		nd.right.newRightNode("9", nil)
		nd.right.left.newRightNode("5", nil)
		_ = testhelperRecursiveHeightsUpdate(nd)

		nd = nd.balance()
		assert.Equal(t, "4", nd.key)
		assert.Equal(t, "2", nd.left.key)
		assert.Equal(t, "1", nd.left.left.key)
		assert.Equal(t, "7", nd.right.key)
		assert.Equal(t, "5", nd.right.left.key)
		assert.Equal(t, "9", nd.right.right.key)
	}
}

func TestNodeUpsert(t *testing.T) {
	//
	//    2             2
	//   / \           / \
	//  1   7    -->  1   7    -->      4
	//     / \           / \           / \
	//    4   9         4   9         2   7
	//                   \           /   / \
	//                   *5*        1   5   9
	//
	{
		nd := newNode(nil, "2", nil)
		nd.newLeftNode("1", nil)
		nd.newRightNode("7", nil)
		nd.right.newLeftNode("4", nil)
		nd.right.newRightNode("8", nil)
		_ = testhelperRecursiveHeightsUpdate(nd)

		nd = nd.upsert("5", nil)
		assert.Equal(t, 2, nd.leftHeight)
		assert.Equal(t, 2, nd.rightHeight)
		assert.Equal(t, "4", nd.key)
		assert.Equal(t, "2", nd.left.key)
		assert.Equal(t, "1", nd.left.left.key)
		assert.Equal(t, "7", nd.right.key)
		assert.Equal(t, "5", nd.right.left.key)
		assert.Equal(t, "8", nd.right.right.key)

		//
		//      4               4
		//     / \             / \
		//    2   7   ->      2   7
		//   /   / \         /   / \
		//  1   5   8       1   5   8
		//                           \
		//                            91
		//
		nd = nd.upsert("91", nil)
		assert.Equal(t, 2, nd.leftHeight)
		assert.Equal(t, 3, nd.rightHeight)
		assert.Equal(t, "4", nd.key)
		assert.Equal(t, "2", nd.left.key)
		assert.Equal(t, "1", nd.left.left.key)
		assert.Equal(t, "7", nd.right.key)
		assert.Equal(t, "5", nd.right.left.key)
		assert.Equal(t, "8", nd.right.right.key)
		assert.Equal(t, "91", nd.right.right.right.key)

		//
		//      4                    4
		//     / \                 /   \
		//    2   7       ->      2     7
		//   /   / \             /     / \
		//  1   5   8           1     5   90
		//           \                   /  \
		//            91                8   91
		//
		nd = nd.upsert("90", nil)
		assert.Equal(t, 2, nd.leftHeight)
		assert.Equal(t, 3, nd.rightHeight)
		assert.Equal(t, "4", nd.key)
		assert.Equal(t, "2", nd.left.key)
		assert.Equal(t, "7", nd.right.key)
		assert.Equal(t, "1", nd.left.left.key)
		assert.Equal(t, "5", nd.right.left.key)
		assert.Equal(t, "90", nd.right.right.key)
		assert.Equal(t, "8", nd.right.right.left.key)
		assert.Equal(t, "91", nd.right.right.right.key)
	}
	// change value
	{
		nd := newNode(nil, "foo", 1337)

		nd.upsert("foo", 1338)
		assert.Equal(t, 1338, nd.value)
	}
	// balancing upserts
	{
		nd := newNode(nil, "10", nil)

		nd = nd.upsert("20", nil)
		assert.Equal(t, "10", nd.key)
		assert.Equal(t, 0, nd.leftHeight)
		assert.Equal(t, 1, nd.rightHeight)

		nd = nd.upsert("30", nil)
		assert.Equal(t, "20", nd.key)
		assert.Equal(t, 1, nd.leftHeight)
		assert.Equal(t, 1, nd.rightHeight)

		nd = nd.upsert("40", nil)
		assert.Equal(t, "20", nd.key)
		assert.Equal(t, 1, nd.leftHeight)
		assert.Equal(t, 2, nd.rightHeight)

		nd = nd.upsert("05", nil)
		assert.Equal(t, "20", nd.key)
		assert.Equal(t, 2, nd.leftHeight)
		assert.Equal(t, 2, nd.rightHeight)

		nd = nd.upsert("03", nil)
		assert.Equal(t, "20", nd.key)
		assert.Equal(t, 2, nd.leftHeight)
		assert.Equal(t, 2, nd.rightHeight)

		nd = nd.upsert("02", nil)
		assert.Equal(t, "20", nd.key)
		assert.Equal(t, 3, nd.leftHeight)
		assert.Equal(t, 2, nd.rightHeight)
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

func TestNodeDelete(t *testing.T) {
	// delete nonexistent
	{
		_, err := (*node)(nil).delete("")
		assert.Equal(t, ErrorNotFound, err)
	}
	// delete left leaf node from level-1 tree
	{
		//
		//   02        02
		//  /  \  ->     \
		// 01  03        03
		//
		nd := newNode(nil, "01", nil)
		nd = nd.upsert("02", nil)
		nd = nd.upsert("03", nil)

		nd, err := nd.delete("01")
		assert.Equal(t, nil, err)

		assert.Equal(t, "02", nd.key)
		assert.Equal(t, false, nd.hasLeft())
		assert.Equal(t, true, nd.hasRight())
		assert.Equal(t, 0, nd.leftHeight)
		assert.Equal(t, 1, nd.rightHeight)

		assert.Equal(t, "03", nd.right.key)
		assert.Equal(t, false, nd.right.hasLeft())
		assert.Equal(t, false, nd.right.hasRight())
		assert.Equal(t, 0, nd.right.leftHeight)
		assert.Equal(t, 0, nd.right.rightHeight)
	}
	// delete right leaf node from level-1 tree
	{
		//
		//   02        02
		//  /  \  ->  /
		// 01  03    01
		//
		nd := newNode(nil, "01", nil)
		nd = nd.upsert("02", nil)
		nd = nd.upsert("03", nil)

		nd, err := nd.delete("03")
		assert.Equal(t, nil, err)

		assert.Equal(t, "02", nd.key)
		assert.Equal(t, true, nd.hasLeft())
		assert.Equal(t, false, nd.hasRight())
		assert.Equal(t, 1, nd.leftHeight)
		assert.Equal(t, 0, nd.rightHeight)

		assert.Equal(t, "01", nd.left.key)
		assert.Equal(t, false, nd.left.hasLeft())
		assert.Equal(t, false, nd.left.hasRight())
		assert.Equal(t, 0, nd.left.leftHeight)
		assert.Equal(t, 0, nd.left.rightHeight)
	}
	// delete left leaf node from level-2 tree
	{
		//
		//        04                  04
		//     /      \            /      \
		//    02      06    ->    02      06
		//   /  \    /  \           \    /  \
		//  01  03  05  07          03  05  07
		//
		nd := newNode(nil, "01", nil)
		nd = nd.upsert("02", nil)
		nd = nd.upsert("03", nil)
		nd = nd.upsert("04", nil)
		nd = nd.upsert("05", nil)
		nd = nd.upsert("06", nil)
		nd = nd.upsert("07", nil)

		nd, err := nd.delete("01")
		assert.Equal(t, nil, err)

		assert.Equal(t, "04", nd.key)
		assert.Equal(t, true, nd.hasLeft())
		assert.Equal(t, true, nd.hasRight())
		assert.Equal(t, 2, nd.leftHeight)
		assert.Equal(t, 2, nd.rightHeight)

		assert.Equal(t, "02", nd.left.key)
		assert.Equal(t, false, nd.left.hasLeft())
		assert.Equal(t, true, nd.left.hasRight())
		assert.Equal(t, 0, nd.left.leftHeight)
		assert.Equal(t, 1, nd.left.rightHeight)
	}
	// delete right leaf node from level-2 tree
	{
		//
		//        04                  04
		//     /      \            /      \
		//    02      06    ->    02      06
		//   /  \    /  \        /       /  \
		//  01  03  05  07      01      05  07
		//
		nd := newNode(nil, "01", nil)
		nd = nd.upsert("02", nil)
		nd = nd.upsert("03", nil)
		nd = nd.upsert("04", nil)
		nd = nd.upsert("05", nil)
		nd = nd.upsert("06", nil)
		nd = nd.upsert("07", nil)

		nd, err := nd.delete("03")
		assert.Equal(t, nil, err)

		assert.Equal(t, "04", nd.key)
		assert.Equal(t, true, nd.hasLeft())
		assert.Equal(t, true, nd.hasRight())
		assert.Equal(t, 2, nd.leftHeight)
		assert.Equal(t, 2, nd.rightHeight)

		assert.Equal(t, "02", nd.left.key)
		assert.Equal(t, true, nd.left.hasLeft())
		assert.Equal(t, false, nd.left.hasRight())
		assert.Equal(t, 1, nd.left.leftHeight)
		assert.Equal(t, 0, nd.left.rightHeight)
	}
	// delete node with left child only
	{
		//
		//        04                  04
		//     /      \            /      \
		//    02      06    ->    01      06
		//   /
		//  01
		//
		nd := newNode(nil, "04", nil)
		nd = nd.upsert("02", nil)
		nd = nd.upsert("06", nil)
		nd = nd.upsert("01", nil)

		nd, err := nd.delete("02")
		assert.Equal(t, nil, err)

		assert.Equal(t, "04", nd.key)
		assert.Equal(t, true, nd.hasLeft())
		assert.Equal(t, true, nd.hasRight())
		assert.Equal(t, 1, nd.leftHeight)
		assert.Equal(t, 1, nd.rightHeight)

		assert.Equal(t, "01", nd.left.key)
		assert.Equal(t, false, nd.left.hasLeft())
		assert.Equal(t, false, nd.left.hasRight())
		assert.Equal(t, 0, nd.left.leftHeight)
		assert.Equal(t, 0, nd.left.rightHeight)

		assert.Equal(t, "06", nd.right.key)
		assert.Equal(t, false, nd.right.hasLeft())
		assert.Equal(t, false, nd.right.hasRight())
		assert.Equal(t, 0, nd.right.leftHeight)
		assert.Equal(t, 0, nd.right.rightHeight)
	}
	// delete node with right child only
	{
		//
		//        04                  04
		//     /      \            /      \
		//    02      06    ->    03      06
		//      \
		//      03
		//
		nd := newNode(nil, "04", nil)
		nd = nd.upsert("02", nil)
		nd = nd.upsert("06", nil)
		nd = nd.upsert("03", nil)

		nd, err := nd.delete("02")
		assert.Equal(t, nil, err)

		assert.Equal(t, "04", nd.key)
		assert.Equal(t, true, nd.hasLeft())
		assert.Equal(t, true, nd.hasRight())
		assert.Equal(t, 1, nd.leftHeight)
		assert.Equal(t, 1, nd.rightHeight)

		assert.Equal(t, "03", nd.left.key)
		assert.Equal(t, false, nd.left.hasLeft())
		assert.Equal(t, false, nd.left.hasRight())
		assert.Equal(t, 0, nd.left.leftHeight)
		assert.Equal(t, 0, nd.left.rightHeight)

		assert.Equal(t, "06", nd.right.key)
		assert.Equal(t, false, nd.right.hasLeft())
		assert.Equal(t, false, nd.right.hasRight())
		assert.Equal(t, 0, nd.right.leftHeight)
		assert.Equal(t, 0, nd.right.rightHeight)
	}
	// delete node with two children
	{
		//
		//        40                  40                  40
		//     /      \            /      \            /      \
		//    20     *60*   ->    20      60    ->    20     *70*
		//   /  \    /  \        /  \    /  \        /  \    /  \
		//  10  30  50  75      10  30  50  75      10  30  50  75
		//             /                    /
		//            70                  *70*
		//
		nd := newNode(nil, "10", nil)
		nd = nd.upsert("20", nil)
		nd = nd.upsert("30", nil)
		nd = nd.upsert("40", nil)
		nd = nd.upsert("50", nil)
		nd = nd.upsert("60", nil)
		nd = nd.upsert("75", nil)
		nd = nd.upsert("70", nil)

		nd, err := nd.delete("60")
		assert.Equal(t, nil, err)

		assert.Equal(t, "40", nd.key)
		assert.Equal(t, true, nd.hasLeft())
		assert.Equal(t, true, nd.hasRight())
		assert.Equal(t, 2, nd.leftHeight)
		assert.Equal(t, 2, nd.rightHeight)

		assert.Equal(t, "70", nd.right.key)
		assert.Equal(t, true, nd.right.hasLeft())
		assert.Equal(t, true, nd.right.hasRight())
		assert.Equal(t, 1, nd.right.leftHeight)
		assert.Equal(t, 1, nd.right.rightHeight)

		assert.Equal(t, "50", nd.right.left.key)
		assert.Equal(t, false, nd.right.left.hasLeft())
		assert.Equal(t, false, nd.right.left.hasRight())
		assert.Equal(t, 0, nd.right.left.leftHeight)
		assert.Equal(t, 0, nd.right.left.rightHeight)

		assert.Equal(t, "75", nd.right.right.key)
		assert.Equal(t, false, nd.right.right.hasLeft())
		assert.Equal(t, false, nd.right.right.hasRight())
		assert.Equal(t, 0, nd.right.right.leftHeight)
		assert.Equal(t, 0, nd.right.right.rightHeight)
	}
}

func TestNodeIter(t *testing.T) {
	nd := newNode(nil, "01", "value-01")
	nd = nd.upsert("02", "value-02")
	nd = nd.upsert("03", "value-03")
	nd = nd.upsert("04", "value-04")
	nd = nd.upsert("05", "value-05")
	ch := make(chan Item)
	go func() {
		nd.iter(ch)
		close(ch)
	}()

	var n int
	for i := range ch {
		switch n {
		case 0:
			assert.Equal(t, "01", i.Key)
			assert.Equal(t, "value-01", i.Val)
		case 1:
			assert.Equal(t, "02", i.Key)
			assert.Equal(t, "value-02", i.Val)
		case 2:
			assert.Equal(t, "03", i.Key)
			assert.Equal(t, "value-03", i.Val)
		case 3:
			assert.Equal(t, "04", i.Key)
			assert.Equal(t, "value-04", i.Val)
		case 4:
			assert.Equal(t, "05", i.Key)
			assert.Equal(t, "value-05", i.Val)
		default:
			assert.Fail(t, "iteration error")
		}
		n++
	}
	assert.Equal(t, 5, n)
}
