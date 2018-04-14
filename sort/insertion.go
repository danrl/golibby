package sort

// Insertion implements a insertion sort algorithm
func Insertion(list []int) {
	for i := 1; i < len(list); i++ {
		tmp := list[i]
		j := i
		for ; j > 0 && list[j-1] > tmp; j-- {
			list[j] = list[j-1]
		}
		list[j] = tmp
	}
}
