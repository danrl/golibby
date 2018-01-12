package metrics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHammingDistance(t *testing.T) {
	{
		_, err := HammingDistance("foo", "")
		assert.NotEqual(t, nil, err)
	}
	{
		_, err := HammingDistance("", "bar")
		assert.NotEqual(t, nil, err)
	}
	{
		dist, err := HammingDistance("", "")
		assert.Equal(t, nil, err)
		assert.Equal(t, 0, dist)
	}
	{
		dist, err := HammingDistance("foo", "bar")
		assert.Equal(t, nil, err)
		assert.Equal(t, 3, dist)
	}
	{
		dist, err := HammingDistance("beer", "tear")
		assert.Equal(t, nil, err)
		assert.Equal(t, 2, dist)
	}
	{
		dist, err := HammingDistance("beer", "be🍺r")
		assert.Equal(t, nil, err)
		assert.Equal(t, 1, dist)
	}
}
