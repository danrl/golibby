package util

// Min returns the smaller of two or more given integers
func Min(a int, x ...int) int {
	min := a
	for _, v := range x {
		if v < min {
			min = v
		}
	}
	return min
}

// Max returns the larger of two or more given integers
func Max(a int, x ...int) int {
	max := a
	for _, v := range x {
		if v > max {
			max = v
		}
	}
	return max
}
