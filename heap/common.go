// Package heap implements a min and a max heap.
package heap

import "fmt"

var (
	// ErrorNoData indicates an invalid operation on an empty heap
	ErrorNoData = fmt.Errorf("no data")
)

func parent(i int) int {
	return (i - 1) / 2
}

func leftChild(i int) int {
	return (i * 2) + 1
}

func rightChild(i int) int {
	return (i * 2) + 2
}
