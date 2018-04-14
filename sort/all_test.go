package sort

var tt = []struct {
	name     string
	in       []int
	expected []int
}{
	{
		name:     "empty list",
		in:       []int{},
		expected: []int{},
	},
	{
		name:     "one element in list",
		in:       []int{1337},
		expected: []int{1337},
	},
	{
		name:     "already sorted list",
		in:       []int{1, 2, 3, 4, 5, 6, 7},
		expected: []int{1, 2, 3, 4, 5, 6, 7},
	},
	{
		name:     "reversed sorted list",
		in:       []int{7, 6, 5, 4, 3, 2, 1},
		expected: []int{1, 2, 3, 4, 5, 6, 7},
	},
	{
		name:     "odd number of elements",
		in:       []int{1, 15, 4, 99, 0},
		expected: []int{0, 1, 4, 15, 99},
	},
	{
		name:     "even number of elements",
		in:       []int{6, 100, 83, 99, 13, 5},
		expected: []int{5, 6, 13, 83, 99, 100},
	},
	{
		name:     "duplicate elements",
		in:       []int{6, 100, 83, 99, 6, 5},
		expected: []int{5, 6, 6, 83, 99, 100},
	},
	{
		name:     "only duplicates (odd)",
		in:       []int{6, 6, 6, 6, 6},
		expected: []int{6, 6, 6, 6, 6},
	},
	{
		name:     "only duplicates (even)",
		in:       []int{6, 6, 6, 6, 6, 6},
		expected: []int{6, 6, 6, 6, 6, 6},
	},
	{
		name:     "first half sorted",
		in:       []int{1, 2, 3, 6, 9, 4},
		expected: []int{1, 2, 3, 4, 6, 9},
	},
	{
		name:     "second half sorted",
		in:       []int{2, 1, 0, 9, 10, 11},
		expected: []int{0, 1, 2, 9, 10, 11},
	},
	{
		name:     "second half has lower sorted elements",
		in:       []int{100, 99, 98, 97, 1, 2, 3, 4},
		expected: []int{1, 2, 3, 4, 97, 98, 99, 100},
	},
	{
		name:     "x",
		in:       []int{100, 1, 4},
		expected: []int{1, 4, 100},
	},
}

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
