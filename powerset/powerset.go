package powerset

// Iterative creates the powerset of the input set
func Iterative(in []int) [][]int {
	ps := make([][]int, 1)
	for _, item := range in {
		for i := range ps {
			ps = append(ps, append(ps[i], item))
		}
	}
	return ps
}

// recursive creates the powerset of the input set
func recursive(ps, new []int) [][]int {
	if len(ps) == 0 {
		return [][]int{new}
	}
	res := [][]int{}
	for _, set := range recursive(ps[1:], new) {
		res = append(res, set)
	}
	for _, set := range recursive(ps[1:], append(new, ps[0])) {
		res = append(res, set)
	}
	return res
}

// Recursive creates the powerset of the input set
func Recursive(in []int) [][]int {
	return recursive(in, []int{})
}
