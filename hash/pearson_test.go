package hash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPearson(t *testing.T) {
	assert.Equal(t, uint8(0x0), Pearson([]byte{}))
	assert.Equal(t, uint8(0x1b), Pearson([]byte{'f', 'o', 'o'}))
	assert.Equal(t, uint8(0x4a), Pearson([]byte{'b', 'a', 'r'}))
}

func TestPearson16(t *testing.T) {
	assert.Equal(t, uint16(0x0), Pearson16([]byte{}))
	assert.Equal(t, uint16(0x1bc5), Pearson16([]byte{'f', 'o', 'o'}))
	assert.Equal(t, uint16(0x4a61), Pearson16([]byte{'b', 'a', 'r'}))
}
