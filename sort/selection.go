package sort

// Selection implements a selection sort algorithm
func Selection(list []int) []int {
	for i := range list {
		for j := i + 1; j < len(list); j++ {
			if list[j] < list[i] {
				list[i], list[j] = list[j], list[i]
			}
		}
	}
	return list
}
