package bstree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	bst := New()
	assert.NotEqual(t, nil, bst)
}

func TestUpsert(t *testing.T) {
	var err error
	bst := New()

	err = bst.Upsert("foo", "bar")
	assert.Equal(t, nil, err)
	assert.Equal(t, "foo", bst.root.key)
	assert.Equal(t, "bar", bst.root.val)

	// left
	err = bst.Upsert("bbb", "bar-b")
	assert.Equal(t, nil, err)
	assert.Equal(t, "bbb", bst.root.left.key)
	assert.Equal(t, "bar-b", bst.root.left.val)
	err = bst.Upsert("aaa", "bar-a")
	assert.Equal(t, nil, err)
	assert.Equal(t, "aaa", bst.root.left.left.key)
	assert.Equal(t, "bar-a", bst.root.left.left.val)

	// right
	err = bst.Upsert("yyy", "bar-y")
	assert.Equal(t, nil, err)
	assert.Equal(t, "yyy", bst.root.right.key)
	assert.Equal(t, "bar-y", bst.root.right.val)
	err = bst.Upsert("zzz", "bar-z")
	assert.Equal(t, nil, err)
	assert.Equal(t, "zzz", bst.root.right.right.key)
	assert.Equal(t, "bar-z", bst.root.right.right.val)

	// overwrite
	err = bst.Upsert("foo", "1337")
	assert.Equal(t, nil, err)
	assert.Equal(t, "foo", bst.root.key)
	assert.Equal(t, "1337", bst.root.val)
}

func TestValue(t *testing.T) {
	var err error
	bst := New()
	_ = bst.Upsert("foo", "bar")

	// root
	val, err := bst.Value("foo")
	assert.Equal(t, "bar", val)
	assert.Equal(t, nil, err)

	// left
	_, err = bst.Value("aaa")
	assert.Equal(t, ErrorNotFound, err)
	_ = bst.Upsert("aaa", "bar-a")
	val, err = bst.Value("aaa")
	assert.Equal(t, "bar-a", val)
	assert.Equal(t, nil, err)

	// right
	_, err = bst.Value("zzz")
	assert.Equal(t, ErrorNotFound, err)
	_ = bst.Upsert("zzz", "bar-z")
	val, err = bst.Value("zzz")
	assert.Equal(t, "bar-z", val)
	assert.Equal(t, nil, err)
}

func TestIsLeafHasLeftHasRight(t *testing.T) {
	var nd *node

	nd = &node{}
	assert.Equal(t, true, nd.isLeaf())
	assert.Equal(t, false, nd.hasLeft())
	assert.Equal(t, false, nd.hasRight())

	nd = &node{
		left: &node{},
	}
	assert.Equal(t, false, nd.isLeaf())
	assert.Equal(t, true, nd.hasLeft())
	assert.Equal(t, false, nd.hasRight())

	nd = &node{
		right: &node{},
	}
	assert.Equal(t, false, nd.isLeaf())
	assert.Equal(t, false, nd.hasLeft())
	assert.Equal(t, true, nd.hasRight())

	nd = &node{
		left:  &node{},
		right: &node{},
	}
	assert.Equal(t, false, nd.isLeaf())
	assert.Equal(t, true, nd.hasLeft())
	assert.Equal(t, true, nd.hasRight())
}

func TestDelete(t *testing.T) {
	// leaf node
	{
		var err error
		bst := New()
		_ = bst.Upsert("foo", "bar")
		_ = bst.Upsert("aaa", "bar-a")
		_ = bst.Upsert("zzz", "bar-z")

		// left
		err = bst.Delete("aaa")
		assert.Equal(t, nil, err)
		err = bst.Delete("aaa")
		assert.Equal(t, ErrorNotFound, err)

		// right
		err = bst.Delete("zzz")
		assert.Equal(t, nil, err)
		err = bst.Delete("zzz")
		assert.Equal(t, ErrorNotFound, err)
	}
	// left child
	{
		var err error
		bst := New()
		_ = bst.Upsert("foo", "bar")
		_ = bst.Upsert("aaa", "bar-a")

		err = bst.Delete("foo")
		assert.Equal(t, nil, err)
		err = bst.Delete("foo")
		assert.Equal(t, ErrorNotFound, err)
	}
	// right child
	{
		var err error
		bst := New()
		_ = bst.Upsert("foo", "bar")
		_ = bst.Upsert("zzz", "bar-z")

		err = bst.Delete("foo")
		assert.Equal(t, nil, err)
		err = bst.Delete("foo")
		assert.Equal(t, ErrorNotFound, err)
	}
	// two children
	{
		var err error
		bst := New()
		_ = bst.Upsert("foo", "bar")
		_ = bst.Upsert("aaa", "bar-a")
		_ = bst.Upsert("zzz", "bar-z")

		err = bst.Delete("foo")
		assert.Equal(t, nil, err)
		err = bst.Delete("foo")
		assert.Equal(t, ErrorNotFound, err)
	}
}

func TestMin(t *testing.T) {
	bst := New()
	_ = bst.Upsert("foo", "bar")
	_ = bst.Upsert("bbb", "bar-b")
	_ = bst.Upsert("aaa", "bar-a")
	assert.Equal(t, "bar-a", bst.root.min().val)
}

func TestHeight(t *testing.T) {
	bst := New()
	assert.Equal(t, 0, bst.Height())

	bst.Upsert("foo", "bar")
	assert.Equal(t, 1, bst.Height())

	bst.Upsert("aaa", "bar-a")
	assert.Equal(t, 2, bst.Height())
	bst.Upsert("zzz", "bar-z")
	assert.Equal(t, 2, bst.Height())
}
