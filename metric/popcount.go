package metric

// PopulationCount returns the bit sum of a byte
func PopulationCount(x byte) int {
	s := 0
	for x > 0 {
		if x&0x1 == 1 {
			s++
		}
		x = x >> 1
	}
	return s
}

// PopulationCounts returns the bit sum of all bytes in a byte slice
func PopulationCounts(xs []byte) int {
	s := 0
	for _, x := range xs {
		s += PopulationCount(x)
	}
	return s
}
