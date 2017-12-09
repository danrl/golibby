package utils

// Min returns the smaller of two given integers
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Max returns the larger of two given integers
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
