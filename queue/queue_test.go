package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLen(t *testing.T) {
	q := Queue{}
	assert.Equal(t, 0, q.Len())

	q.Add(1)
	assert.Equal(t, 1, q.Len())

	q.Add(23)
	assert.Equal(t, 2, q.Len())

	_, _ = q.Peek()
	assert.Equal(t, 2, q.Len())

	_, _ = q.Remove()
	_, _ = q.Remove()
	assert.Equal(t, 0, q.Len())

	_, _ = q.Remove()
	assert.Equal(t, 0, q.Len())
}

func TestAddRemove(t *testing.T) {
	q := Queue{}
	q.Add(1337)
	item, err := q.Remove()
	assert.Equal(t, nil, err)
	assert.Equal(t, 1337, item)

	_, err = q.Remove()
	assert.Equal(t, ErrorEmpty, err)
}

func TestPeek(t *testing.T) {
	q := Queue{}

	_, err := q.Peek()
	assert.Equal(t, ErrorEmpty, err)

	q.Add(1337)
	item, err := q.Peek()
	assert.Equal(t, nil, err)
	assert.Equal(t, 1337, item)
	assert.Equal(t, 1, q.Len())
}
