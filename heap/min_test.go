package heap

import (
	"testing"
)

func TestMinHeapLen(t *testing.T) {
	t.Run("empty heap", func(t *testing.T) {
		mh := MinHeap{}
		if got := mh.Len(); got != 0 {
			t.Errorf("expected `%v`, got `%v`", 0, got)
		}
	})
	t.Run("single item heap", func(t *testing.T) {
		mh := MinHeap{
			data: []int{1},
		}
		if got := mh.Len(); got != 1 {
			t.Errorf("expected `%v`, got `%v`", 1, got)
		}
	})
	t.Run("larger heap", func(t *testing.T) {
		mh := MinHeap{
			data: []int{1, 1, 1, 1, 1, 1},
		}
		if got := mh.Len(); got != 6 {
			t.Errorf("expected `%v`, got `%v`", 6, got)
		}
	})
}

func TestMinHeapPeek(t *testing.T) {
	t.Run("peek on empty heap", func(t *testing.T) {
		mh := MinHeap{}
		_, err := mh.Peek()
		if err != ErrorNoData {
			t.Errorf("expected `%v`, got `%v`", ErrorNoData, err)
		}
	})
	t.Run("peek on heap with data", func(t *testing.T) {
		mh := MinHeap{
			data: []int{1, 2, 3},
		}
		item, err := mh.Peek()
		if err != nil {
			t.Errorf("expected `%v`, got `%v`", nil, err)
		}
		if item != 1 {
			t.Errorf("expected `%v`, got `%v`", 1, item)
		}
	})
}

func TestMinHeapInsert(t *testing.T) {
	t.Run("insert into empty heap", func(t *testing.T) {
		mh := MinHeap{}
		mh.Insert(3)
		if got := len(mh.data); got != 1 {
			t.Fatalf("expected data slice of length `%v` got `%v`", 1, got)
		}
		if got := mh.data[0]; got != 3 {
			t.Errorf("expected max value `%v` in data slice, got `%v`", 3, got)
		}
	})
	t.Run("multiple inserts", func(t *testing.T) {
		mh := MinHeap{}
		mh.Insert(22)
		mh.Insert(78)
		mh.Insert(10)
		mh.Insert(56)
		mh.Insert(12)
		mh.Insert(7)
		mh.Insert(9)
		expected := []int{7, 12, 9, 78, 56, 22, 10}
		if !equal(mh.data, expected) {
			t.Errorf("expected data slice `%v` got `%v`", expected, mh.data)
		}
	})
}

func TestMinHeapPop(t *testing.T) {
	t.Run("pop from empty heap", func(t *testing.T) {
		mh := MinHeap{}
		_, err := mh.Pop()
		if err != ErrorNoData {
			t.Errorf("expected `%v`, got `%v`", ErrorNoData, err)
		}
	})
	t.Run("pop from heap with data", func(t *testing.T) {
		mh := MinHeap{
			data: []int{1, 2, 3},
		}
		item, err := mh.Pop()
		if err != nil {
			t.Errorf("expected `%v`, got `%v`", nil, err)
		}
		if item != 1 {
			t.Errorf("expected `%v`, got `%v`", 1, item)
		}
		if got := len(mh.data); got != 2 {
			t.Fatalf("expected data slice of length `%v` got `%v`", 2, got)
		}
		if got := mh.data[0]; got != 2 {
			t.Errorf("expected max value `%v` in data slice, got `%v`", 2, got)
		}
	})
	t.Run("pop from small heap and rebalance", func(t *testing.T) {
		mh := MinHeap{
			data: []int{4, 9, 5, 89, 72, 75},
		}
		for _, expected := range []int{4, 5, 9, 72, 75, 89} {
			got, _ := mh.Pop()
			if got != expected {
				t.Errorf("expected value `%v` got `%v`", expected, got)
			}
		}
	})
	t.Run("pop from heap and rebalance", func(t *testing.T) {
		mh := MinHeap{
			data: []int{15, 18, 63, 36, 21, 75, 65, 89, 70, 72},
		}
		for _, expected := range []int{15, 18, 21, 36, 63, 65, 70, 72, 75, 89} {
			got, _ := mh.Pop()
			if got != expected {
				t.Errorf("expected value `%v` got `%v`", expected, got)
			}
		}
	})
}
