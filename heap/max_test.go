package heap

import "testing"

func TestMaxHeapLen(t *testing.T) {
	t.Run("empty heap", func(t *testing.T) {
		mh := MaxHeap{}
		if got := mh.Len(); got != 0 {
			t.Errorf("expected `%v`, got `%v`", 0, got)
		}
	})
	t.Run("single item heap", func(t *testing.T) {
		mh := MaxHeap{
			data: []int{1},
		}
		if got := mh.Len(); got != 1 {
			t.Errorf("expected `%v`, got `%v`", 1, got)
		}
	})
	t.Run("larger heap", func(t *testing.T) {
		mh := MaxHeap{
			data: []int{1, 1, 1, 1, 1, 1},
		}
		if got := mh.Len(); got != 6 {
			t.Errorf("expected `%v`, got `%v`", 6, got)
		}
	})
}

func TestMaxHeapPeek(t *testing.T) {
	t.Run("peek on empty heap", func(t *testing.T) {
		mh := MaxHeap{}
		_, err := mh.Peek()
		if err != ErrorNoData {
			t.Errorf("expected `%v`, got `%v`", ErrorNoData, err)
		}
	})
	t.Run("peek on heap with data", func(t *testing.T) {
		mh := MaxHeap{
			data: []int{3, 2, 1},
		}
		item, err := mh.Peek()
		if err != nil {
			t.Errorf("expected `%v`, got `%v`", nil, err)
		}
		if item != 3 {
			t.Errorf("expected `%v`, got `%v`", 3, item)
		}
	})
}

func TestMaxHeapInsert(t *testing.T) {
	t.Run("insert into empty heap", func(t *testing.T) {
		mh := MaxHeap{}
		mh.Insert(3)
		if got := len(mh.data); got != 1 {
			t.Fatalf("expected data slice of length `%v` got `%v`", 1, got)
		}
		if got := mh.data[0]; got != 3 {
			t.Errorf("expected max value `%v` in data slice, got `%v`", 3, got)
		}
	})
	t.Run("multiple inserts", func(t *testing.T) {
		mh := MaxHeap{}
		mh.Insert(3)
		mh.Insert(2)
		mh.Insert(10)
		mh.Insert(4)
		mh.Insert(9)
		expected := []int{10, 9, 3, 2, 4}
		if !equal(mh.data, expected) {
			t.Errorf("expected data slice `%v` got `%v`", expected, mh.data)
		}
	})
}

func TestMaxHeapPop(t *testing.T) {
	t.Run("pop from empty heap", func(t *testing.T) {
		mh := MaxHeap{}
		_, err := mh.Pop()
		if err != ErrorNoData {
			t.Errorf("expected `%v`, got `%v`", ErrorNoData, err)
		}
	})
	t.Run("pop from heap with data", func(t *testing.T) {
		mh := MaxHeap{
			data: []int{3, 2, 1},
		}
		item, err := mh.Pop()
		if err != nil {
			t.Errorf("expected `%v`, got `%v`", nil, err)
		}
		if item != 3 {
			t.Errorf("expected `%v`, got `%v`", 3, item)
		}
		if got := len(mh.data); got != 2 {
			t.Fatalf("expected data slice of length `%v` got `%v`", 2, got)
		}
		if got := mh.data[0]; got != 2 {
			t.Errorf("expected max value `%v` in data slice, got `%v`", 2, got)
		}
	})
	t.Run("pop from heap and rebalance", func(t *testing.T) {
		mh := MaxHeap{
			data: []int{89, 75, 72, 36, 63, 70, 65, 21, 18, 15},
		}
		for _, expected := range []int{89, 75, 72, 70, 65, 63, 36, 21, 18, 15} {
			got, _ := mh.Pop()
			if got != expected {
				t.Errorf("expected value `%v` got `%v`", expected, got)
			}
		}
	})
}
