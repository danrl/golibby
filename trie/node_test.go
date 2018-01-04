package trie

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNodeNew(t *testing.T) {
	nd := new()
	assert.NotEqual(t, nil, nd)
	assert.Equal(t, 0, len(nd.keys))
}

func TestNodeNode(t *testing.T) {
	{
		nd := new()

		_, err := nd.node([]string{"foo"}, false)
		assert.Equal(t, 0, len(nd.keys))
		assert.Equal(t, ErrorNotFound, err)
	}
	{
		nd := new()

		_, err := nd.node([]string{"foo"}, true)
		assert.Equal(t, 1, len(nd.keys))
		assert.Equal(t, nil, err)
	}
}

func TestNodeUpsert(t *testing.T) {
	{
		nd := new()

		nd.upsert([]string{"foo"}, 23)
		assert.NotEqual(t, (nil), nd.keys["foo"])
	}
}

func TestNodeData(t *testing.T) {
	{
		nd := new()

		_, err := nd.data([]string{"foo"})
		assert.Equal(t, ErrorNotFound, err)
	}
	{
		nd := new()

		nd.upsert([]string{"foo", "bar"}, 42)
		_, err := nd.data([]string{"foo"})
		assert.Equal(t, ErrorNoData, err)
	}
	{
		nd := new()

		nd.upsert([]string{"foo"}, 23)
		value, err := nd.data([]string{"foo"})
		assert.Equal(t, 23, value)
		assert.Equal(t, nil, err)
	}
}

func TestNodeDelete(t *testing.T) {
	{
		nd := new()

		_, err := nd.delete([]string{"foo"})
		assert.Equal(t, ErrorNotFound, err)
	}
	{
		nd := new()
		nd.upsert([]string{"foo"}, 23)

		_, err := nd.delete([]string{"foo"})
		assert.Equal(t, nil, err)
	}
	{
		nd := new()
		nd.upsert([]string{"foo"}, 23)

		_, err := nd.delete([]string{"foo", "bar"})
		assert.Equal(t, ErrorNotFound, err)
	}
}
