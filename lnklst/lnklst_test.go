package lnklst

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAppend(t *testing.T) {
	ll := LinkedList{}
	ll.Append("foo")
	assert.NotEqual(t, nil, ll.head)
	assert.Equal(t, "foo", ll.head.val)

	ll.Append("bar")
	assert.NotEqual(t, nil, ll.head.next)
	assert.Equal(t, "bar", ll.head.next.val)

	ll.Append("lorem")
	assert.NotEqual(t, nil, ll.head.next.next)
	assert.Equal(t, "lorem", ll.head.next.next.val)
}

func TestRemove(t *testing.T) {
	// beginning
	{
		ll := LinkedList{}
		ll.Append("foo")
		ll.Append("bar")
		ll.Append("lorem")

		ll.Remove("foo")
		assert.Equal(t, "bar", ll.head.val)
		assert.NotEqual(t, nil, ll.head.next)
	}
	// middle
	{
		ll := LinkedList{}
		ll.Append("foo")
		ll.Append("bar")
		ll.Append("lorem")

		ll.Remove("bar")
		assert.Equal(t, "foo", ll.head.val)
		assert.NotEqual(t, nil, ll.head.next)
		assert.Equal(t, "lorem", ll.head.next.val)
		assert.Equal(t, (*item)(nil), ll.head.next.next)
	}
	// end
	{
		ll := LinkedList{}
		ll.Append("foo")
		ll.Append("bar")
		ll.Append("lorem")

		ll.Remove("lorem")
		assert.Equal(t, (*item)(nil), ll.head.next.next)
	}
	// not found
	{
		ll := LinkedList{}
		ll.Append("foo")
		ll.Append("bar")

		err := ll.Remove("lorem")
		assert.Equal(t, ErrorNotFound, err)
	}
}

func TestIter(t *testing.T) {
	ll := LinkedList{}
	ll.Append("foo")
	ll.Append("bar")
	ll.Append("lorem")
	var i int
	for val := range ll.Iter() {
		switch i {
		case 0:
			assert.Equal(t, "foo", val)
		case 1:
			assert.Equal(t, "bar", val)
		case 2:
			assert.Equal(t, "lorem", val)
		}
		i++
	}
}

func TestLen(t *testing.T) {
	ll := LinkedList{}
	assert.Equal(t, 0, ll.Len())

	ll.Append("foo")
	assert.Equal(t, 1, ll.Len())
	ll.Append("bar")
	assert.Equal(t, 2, ll.Len())
	_ = ll.Remove("foo")
	assert.Equal(t, 1, ll.Len())
}
