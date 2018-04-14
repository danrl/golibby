package sort

import "testing"

func TestInsertion(t *testing.T) {
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			list := make([]int, len(tc.in), len(tc.in))
			copy(list, tc.in)
			Insertion(list)
			if !equal(tc.expected, list) {
				t.Errorf("result not a sorted list. expected `%v` got `%v`",
					tc.expected, list)
			}
		})
	}
}
