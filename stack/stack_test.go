package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLen(t *testing.T) {
	s := New()
	assert.Equal(t, 0, s.Len())

	s.Push(23)
	assert.Equal(t, 1, s.Len())

	s.Push(23)
	assert.Equal(t, 2, s.Len())

	_, _ = s.Peek()
	assert.Equal(t, 2, s.Len())

	_, _ = s.Pop()
	_, _ = s.Pop()
	assert.Equal(t, 0, s.Len())

	_, _ = s.Pop()
	assert.Equal(t, 0, s.Len())
}

func TestPushPop(t *testing.T) {
	s := New()
	s.Push(1337)

	x, err := s.Pop()
	assert.Equal(t, nil, err)
	assert.Equal(t, 1337, x)

	_, err = s.Pop()
	assert.Equal(t, ErrorEmptyStack, err)
}

func TestPeek(t *testing.T) {
	s := New()

	_, err := s.Peek()
	assert.Equal(t, ErrorEmptyStack, err)

	s.Push(1337)
	x, err := s.Peek()
	assert.Equal(t, nil, err)
	assert.Equal(t, 1337, x)
	assert.Equal(t, 1, s.Len())
}
