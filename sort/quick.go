package sort

// QuickLomuto implements a quick sort algorithm using the Lomuto partitioningg
// scheme
func QuickLomuto(list []int) {
	if len(list) > 1 {
		// partition
		p := 0
		hi := len(list) - 1
		pivot := list[hi]
		for j := 0; j < hi; j++ {
			if list[j] < pivot {
				list[p], list[j] = list[j], list[p]
				p++
			}
		}
		list[p], list[hi] = list[hi], list[p]

		// recursion
		QuickLomuto(list[:p])
		QuickLomuto(list[p+1:])
	}
}

// Quick implements a quick sort algorithm using the original Hoare partitioning
// scheme
func Quick(list []int) {
	if len(list) > 1 {
		// partition
		pivot := list[0]
		lo := -1
		hi := len(list)
		for {
			for lo++; list[lo] < pivot; lo++ {
			}
			for hi--; list[hi] > pivot; hi-- {
			}
			if lo >= hi {
				hi++
				break
			}
			// swap
			list[lo], list[hi] = list[hi], list[lo]
		}
		// recursion
		Quick(list[:hi])
		Quick(list[hi:])
	}
}
