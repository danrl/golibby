package metric

import "fmt"

// HammingDistance returns the Hamming distance of two strings of equal lengths
func HammingDistance(a, b string) (int, error) {
	ar := []rune(a)
	br := []rune(b)

	if len(ar) != len(br) {
		return 0, fmt.Errorf("input length mismatch")
	}
	dist := 0
	for i := range ar {
		if ar[i] != br[i] {
			dist++
		}
	}
	return dist, nil
}
