package trie

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	path1 = []string{"lvl-1"}
	path2 = []string{"lvl-1", "lvl-2"}
	path3 = []string{"lvl-1", "lvl-2", "lvl-3"}
)

func TestTrieNew(t *testing.T) {
	tr := New()
	assert.NotEqual(t, nil, tr)
}

func TestTrieUpsert(t *testing.T) {
	tr := New()

	assert.NotPanics(t, func() { tr.Upsert(path2, 23) })
}

func TestTrieData(t *testing.T) {
	{
		tr := New()

		_, err := tr.Data(path1)
		assert.Equal(t, ErrorNotFound, err)
	}
	{
		tr := New()
		tr.Upsert(path2, 23)

		value, err := tr.Data(path2)
		assert.Equal(t, 23, value)
		assert.Equal(t, nil, err)
	}
}

func TestTrieDelete(t *testing.T) {
	{
		tr := New()

		err := tr.Delete(path3)
		assert.Equal(t, ErrorNotFound, err)
	}
	{
		tr := New()
		tr.Upsert(path1, 1)
		_ = tr.Delete(path1)
		_, err := tr.Data(path1)
		assert.Equal(t, ErrorNotFound, err)
	}
	{
		tr := New()
		tr.Upsert(path1, 1)
		tr.Upsert(path3, 3)

		// remove leaf node and its data
		err := tr.Delete(path3)
		assert.Equal(t, nil, err)

		// expecting path2 to not exist anymore after leaf node removal
		_, err = tr.Data(path2)
		assert.Equal(t, ErrorNotFound, err)
	}
	{
		tr := New()
		tr.Upsert(path2, 2)
		err := tr.Delete(path3)
		assert.Equal(t, ErrorNotFound, err)
	}
}
