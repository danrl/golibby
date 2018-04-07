package powerset

import "testing"

var (
	tt = []struct {
		name     string
		in       []int
		expected [][]int
	}{
		{
			name:     "empty set",
			in:       []int{},
			expected: [][]int{{}},
		},
		{
			name:     "one item set",
			in:       []int{7},
			expected: [][]int{[]int{}, []int{7}},
		},
		{
			name: "two item set",
			in:   []int{3, 5},
			expected: [][]int{
				[]int{},
				[]int{3},
				[]int{5}, []int{3, 5},
			},
		},
		{
			name: "three item set",
			in:   []int{3, 5, 7},
			expected: [][]int{
				[]int{},
				[]int{3},
				[]int{5}, []int{3, 5},
				[]int{7}, []int{3, 7}, []int{5, 7}, []int{3, 5, 7},
			},
		},
	}
)

// helper function: horribly inefficient slice comparator
func equalSlice(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		found := false
		for j := range b {
			if a[i] == b[j] {
				found = true
			}
		}
		if !found {
			return false
		}
	}
	return true
}

// helper function: horribly inefficient slice of slices comparator
func equalSliceOfSlices(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		found := false
		for j := range b {
			if equalSlice(a[i], b[j]) {
				found = true
			}
		}
		if !found {
			return false
		}
	}
	return true
}

func TestIterative(t *testing.T) {
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got := Iterative(tc.in)
			if !equalSliceOfSlices(got, tc.expected) {
				t.Errorf("expected `%v` got `%v`", tc.expected, got)
			}
		})
	}
}

func TestRecursive(t *testing.T) {
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got := Recursive(tc.in)
			if !equalSliceOfSlices(got, tc.expected) {
				t.Errorf("expected `%v` got `%v`", tc.expected, got)
			}
		})
	}
}
