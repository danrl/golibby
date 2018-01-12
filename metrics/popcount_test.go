package metrics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPopulationCount(t *testing.T) {
	assert.Equal(t, 0, PopulationCount(0x0))
	assert.Equal(t, 1, PopulationCount(0x1))
	assert.Equal(t, 1, PopulationCount(0x2))
	assert.Equal(t, 2, PopulationCount(0x3))
	assert.Equal(t, 4, PopulationCount(0xf0))
}

func TestPopulationCounts(t *testing.T) {
	assert.Equal(t, 0, PopulationCounts([]byte{0x0}))
	assert.Equal(t, 0, PopulationCounts([]byte{0x0, 0x0, 0x0}))
	assert.Equal(t, 1, PopulationCounts([]byte{0x0, 0x4, 0x0}))
	assert.Equal(t, 16, PopulationCounts([]byte{0xff, 0x0f, 0xf0}))
}
