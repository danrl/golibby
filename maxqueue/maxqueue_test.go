package maxqueue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	{
		_, err := New(-1337)
		assert.Equal(t, ErrorIllegalLength, err)
	}
	{
		_, err := New(0)
		assert.Equal(t, ErrorIllegalLength, err)
	}
	{
		q, err := New(1)
		assert.NotEqual(t, nil, q)
		assert.Equal(t, nil, err)
	}
	{
		q, err := New(1337)
		assert.NotEqual(t, nil, q)
		assert.Equal(t, nil, err)
	}
}

func TestLen(t *testing.T) {
	q, _ := New(3)
	_ = q.Add(1)
	assert.Equal(t, 1, q.Len())

	_ = q.Add(2)
	assert.Equal(t, 2, q.Len())

	_, _ = q.Remove()
	assert.Equal(t, 1, q.Len())

	_, _ = q.Remove()
	assert.Equal(t, 0, q.Len())

	_, _ = q.Remove()
	assert.Equal(t, 0, q.Len())
}

func TestAddRemove(t *testing.T) {
	q, _ := New(3)

	err := q.Add(1)
	assert.Equal(t, nil, err)
	err = q.Add(2)
	assert.Equal(t, nil, err)
	err = q.Add(3)
	assert.Equal(t, nil, err)
	err = q.Add(4)
	assert.Equal(t, ErrorFull, err)

	item, err := q.Remove()
	assert.Equal(t, 1, item)
	assert.Equal(t, nil, err)

	item, err = q.Remove()
	assert.Equal(t, 2, item)
	assert.Equal(t, nil, err)

	item, err = q.Remove()
	assert.Equal(t, 3, item)
	assert.Equal(t, nil, err)

	_, err = q.Remove()
	assert.Equal(t, ErrorEmpty, err)
}

func TestPeek(t *testing.T) {
	q, _ := New(3)

	_, err := q.Peek()
	assert.Equal(t, ErrorEmpty, err)

	q.Add(1337)
	item, err := q.Peek()
	assert.Equal(t, nil, err)
	assert.Equal(t, 1337, item)
	assert.Equal(t, 1, q.Len())
}
