package util

// Clamp lower limits value to `low` and upper limits value to `high``
func Clamp(value, low, high int) int {
	if value < low {
		return low
	} else if value > high {
		return high
	}
	return value
}
