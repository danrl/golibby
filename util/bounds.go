package util

// LowerBound restricts value to given lower bound
func LowerBound(value, bound int) int {
	if value < bound {
		return bound
	}
	return value
}

// UpperBound restricts value to given upper bound
func UpperBound(value, bound int) int {
	if value > bound {
		return bound
	}
	return value
}
