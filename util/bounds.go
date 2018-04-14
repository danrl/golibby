package util

func UpperBound(value, bound int) int {
	if value > bound {
		return bound
	}
	return value
}

func LowerBound(value, bound int) int {
	if value < bound {
		return bound
	}
	return value
}
