package metrics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// LevenshteinDistance returns the Levenshtein distance of two strings
func TestLevenshteinDistance(t *testing.T) {
	assert.Equal(t, 0, LevenshteinDistance("", ""))
	assert.Equal(t, 0, LevenshteinDistance("🍺", "🍺"))
	assert.Equal(t, 1, LevenshteinDistance("beer", "bear"))
	assert.Equal(t, 1, LevenshteinDistance("beer", "bee"))
	assert.Equal(t, 1, LevenshteinDistance("beer", "beer🍺"))
	assert.Equal(t, 3, LevenshteinDistance("beer", "water"))
	assert.Equal(t, 5, LevenshteinDistance("java", "golang"))
}
