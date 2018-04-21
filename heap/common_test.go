package heap

import "testing"

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestParent(t *testing.T) {
	tt := []struct {
		name     string
		idx      int
		expected int
	}{
		{
			name:     "root node",
			idx:      0,
			expected: 0,
		},
		{
			name:     "1st level node (1/2)",
			idx:      1,
			expected: 0,
		},
		{
			name:     "1st level node (2/2)",
			idx:      2,
			expected: 0,
		},
		{
			name:     "3rd level node",
			idx:      7,
			expected: 3,
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			if got := parent(tc.idx); got != tc.expected {
				t.Errorf("expected `%v` got `%v`", tc.expected, got)
			}
		})
	}
}

func TestLeftChild(t *testing.T) {
	tt := []struct {
		name     string
		idx      int
		expected int
	}{
		{
			name:     "root node",
			idx:      0,
			expected: 1,
		},
		{
			name:     "1st level node (1/2)",
			idx:      1,
			expected: 3,
		},
		{
			name:     "1st level node (2/2)",
			idx:      2,
			expected: 5,
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			if got := leftChild(tc.idx); got != tc.expected {
				t.Errorf("expected `%v` got `%v`", tc.expected, got)
			}
		})
	}
}

func TestRightChild(t *testing.T) {
	tt := []struct {
		name     string
		idx      int
		expected int
	}{
		{
			name:     "root node",
			idx:      0,
			expected: 2,
		},
		{
			name:     "1st level node (1/2)",
			idx:      1,
			expected: 4,
		},
		{
			name:     "1st level node (2/2)",
			idx:      2,
			expected: 6,
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			if got := rightChild(tc.idx); got != tc.expected {
				t.Errorf("expected `%v` got `%v`", tc.expected, got)
			}
		})
	}
}
