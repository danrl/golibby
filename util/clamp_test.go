package util

import "testing"

func TestClamp(t *testing.T) {
	tt := []struct {
		name                       string
		value, low, high, expected int
	}{
		{
			name:     "lower exceed",
			value:    -4,
			low:      -2,
			high:     99,
			expected: -2,
		},
		{
			name:     "lower bound",
			value:    -2,
			low:      -2,
			high:     99,
			expected: -2,
		},
		{
			name:     "within bounds (negative)",
			value:    -10,
			low:      -20,
			high:     -9,
			expected: -10,
		},
		{
			name:     "within bounds (positive)",
			value:    10,
			low:      3,
			high:     80,
			expected: 10,
		},
		{
			name:     "upper bound",
			value:    99,
			low:      -2,
			high:     99,
			expected: 99,
		},
		{
			name:     "upper exceed",
			value:    101,
			low:      -2,
			high:     99,
			expected: 99,
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			if got := Clamp(tc.value, tc.low, tc.high); got != tc.expected {
				t.Errorf("expected `%v` got `%v`", tc.expected, got)
			}
		})
	}
}
