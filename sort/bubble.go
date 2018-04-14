package sort

// Bubble implements a bubble sort algorithm
func Bubble(list []int) {
	swap := true
	for swap {
		swap = false
		for i := 0; i < len(list)-1; i++ {
			if list[i] > list[i+1] {
				list[i], list[i+1] = list[i+1], list[i]
				swap = true
			}
		}
	}
}
