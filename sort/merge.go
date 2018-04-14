package sort

// Merge implements a top-down merge sort algorithm for sorting integer
// slices
func Merge(list []int) {
	// list with zero or one element is always sorted
	if len(list) <= 1 {
		return
	}

	mid := len(list) / 2

	// split list into sublists, recursive calls
	Merge(list[:mid])
	Merge(list[mid:])

	// merge sorted sublists
	merged := []int{}
	li, ri := 0, mid
	for li < mid && ri < len(list) {
		if list[li] <= list[ri] {
			merged = append(merged, list[li])
			li++
		} else {
			merged = append(merged, list[ri])
			ri++
		}
	}
	merged = append(merged, list[li:mid]...)
	merged = append(merged, list[ri:]...)

	copy(list, merged) // not so great
}
