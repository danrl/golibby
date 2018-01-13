package avltree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAVLTreeUpsert(t *testing.T) {
	avl := AVLTree{}

	avl.Upsert("foo", nil)
	assert.Equal(t, nil, avl.root.value)

	avl.Upsert("foo", 1337)
	assert.Equal(t, 1337, avl.root.value)
}

func TestAVLTreeLookup(t *testing.T) {
	avl := AVLTree{}
	avl.Upsert("foo", 1337)

	value, err := avl.Lookup("foo")
	assert.Equal(t, 1337, value)
	assert.Equal(t, nil, err)
}

func TestAVLTreeDelete(t *testing.T) {
	{
		avl := AVLTree{}

		err := avl.Delete("foo")
		assert.Equal(t, ErrorNotFound, err)
	}
	{
		avl := AVLTree{}
		avl.Upsert("foo", 1337)

		err := avl.Delete("foo")
		assert.Equal(t, nil, err)
	}
}

func TestAVLTreeIter(t *testing.T) {
	avl := AVLTree{}
	avl.Upsert("1", nil)
	avl.Upsert("2", nil)
	avl.Upsert("3", nil)

	var n int
	for range avl.Iter() {
		n++
	}
	assert.Equal(t, 3, n)
}
