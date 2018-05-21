package metric

// PopulationCount returns the bit sum of a byte using Brian Kernighan's
// algorithm with a time complexity of O(log n)
func PopulationCount(x byte) int {
	count := 0
	for x != 0 {
		count++
		x &= (x - 1)
	}
	return count
}

// PopulationCounts returns the bit sum of all bytes in a byte slice
func PopulationCounts(xs []byte) int {
	count := 0
	for _, x := range xs {
		count += PopulationCount(x)
	}
	return count
}
