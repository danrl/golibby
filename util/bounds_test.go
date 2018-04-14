package util

import "testing"

func TestLowerBound(t *testing.T) {
	tt := []struct {
		name     string
		value    int
		bound    int
		expected int
	}{
		{
			name:     "zeros",
			value:    0,
			bound:    0,
			expected: 0,
		},
		{
			name:     "same",
			value:    5,
			bound:    5,
			expected: 5,
		},
		{
			name:     "bound",
			value:    3,
			bound:    5,
			expected: 5,
		},
		{
			name:     "unbound",
			value:    3,
			bound:    5,
			expected: 5,
		},
		{
			name:     "negative bound",
			value:    -7,
			bound:    -5,
			expected: -5,
		},
		{
			name:     "negative unbound",
			value:    -7,
			bound:    -10,
			expected: -7,
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			if got := LowerBound(tc.value, tc.bound); got != tc.expected {
				t.Errorf("expected `%v` got `%v`", tc.expected, got)
			}
		})
	}
}

func TestUpperBound(t *testing.T) {
	tt := []struct {
		name     string
		value    int
		bound    int
		expected int
	}{
		{
			name:     "zeros",
			value:    0,
			bound:    0,
			expected: 0,
		},
		{
			name:     "same",
			value:    5,
			bound:    5,
			expected: 5,
		},
		{
			name:     "bound",
			value:    5,
			bound:    3,
			expected: 3,
		},
		{
			name:     "unbound",
			value:    3,
			bound:    5,
			expected: 3,
		},
		{
			name:     "negative bound",
			value:    -5,
			bound:    -7,
			expected: -7,
		},
		{
			name:     "negative unbound",
			value:    -10,
			bound:    -7,
			expected: -10,
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			if got := UpperBound(tc.value, tc.bound); got != tc.expected {
				t.Errorf("expected `%v` got `%v`", tc.expected, got)
			}
		})
	}
}
