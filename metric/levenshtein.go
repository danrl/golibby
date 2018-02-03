package metric

import "github.com/danrl/golibby/util"

// LevenshteinDistance returns the Levenshtein distance of two strings
func LevenshteinDistance(a, b string) int {
	ar := []rune(a)
	br := []rune(b)

	v0 := make([]int, len(br)+1)
	v1 := make([]int, len(br)+1)

	// initialize start vector
	for i := 0; i < len(v0); i++ {
		v0[i] = i
	}

	for i := 0; i < len(ar); i++ {
		v1[0] = i + 1
		for j := 0; j < len(br); j++ {
			delCost := v0[j+1] + 1
			insCost := v1[j] + 1
			subCost := v0[j] + 1
			if ar[i] == br[j] {
				subCost = v0[j]
			}
			v1[j+1] = util.Min(delCost, insCost, subCost)
		}
		for i := range v0 {
			v0[i] = v1[i]
		}
	}
	return v0[len(br)]
}
