package avltree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	assert.NotEqual(t, nil, New())
}

func TestAVLTreeUpsert(t *testing.T) {
	avl := New()

	avl.Upsert("foo", nil)
	assert.Equal(t, nil, avl.root.value)

	avl.Upsert("foo", 1337)
	assert.Equal(t, 1337, avl.root.value)
}

func TestAVLTreeLookup(t *testing.T) {
	avl := New()
	avl.Upsert("foo", 1337)

	value, err := avl.Lookup("foo")
	assert.Equal(t, 1337, value)
	assert.Equal(t, nil, err)
}

func TestAVLTreeDelete(t *testing.T) {
	{
		avl := New()

		err := avl.Delete("foo")
		assert.Equal(t, ErrorNotFound, err)
	}
	{
		avl := New()
		avl.Upsert("foo", 1337)

		err := avl.Delete("foo")
		assert.Equal(t, nil, err)
	}
}
