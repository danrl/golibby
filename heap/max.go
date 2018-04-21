package heap

import (
	"sync"
)

// MaxHeap represents a max heap instance
type MaxHeap struct {
	lock sync.RWMutex
	data []int
}

// Len returns the number of items on the heap
func (mh *MaxHeap) Len() int {
	mh.lock.RLock()
	defer mh.lock.RUnlock()
	length := len(mh.data)
	return length
}

// Peek returns the top item of the heap without removing it
func (mh *MaxHeap) Peek() (int, error) {
	mh.lock.RLock()
	defer mh.lock.RUnlock()
	if len(mh.data) == 0 {
		return 0, ErrorNoData
	}
	item := mh.data[0]
	return item, nil
}

// Insert adds an item to the heap and relbalances the heap if necessary
func (mh *MaxHeap) Insert(item int) {
	mh.lock.Lock()
	defer mh.lock.Unlock()
	// index of new node
	i := len(mh.data)
	// append node to tree
	mh.data = append(mh.data, item)
	// heapify up
	for ; i > 0 && mh.data[i] > mh.data[parent(i)]; i = parent(i) {
		mh.data[i], mh.data[parent(i)] = mh.data[parent(i)], mh.data[i]
	}
}

// Pop retrieves the largest item from the heap and rebalances the heap if
// necessary
func (mh *MaxHeap) Pop() (int, error) {
	mh.lock.Lock()
	defer mh.lock.Unlock()
	if len(mh.data) == 0 {
		return 0, ErrorNoData
	}
	item := mh.data[0]
	// move last node to root
	mh.data[0] = mh.data[len(mh.data)-1]
	mh.data = mh.data[:len(mh.data)-1]
	// heapify down
	for i := 0; i < len(mh.data); {
		li := leftChild(i)
		ri := rightChild(i)
		if li < len(mh.data) && mh.data[li] > mh.data[i] {
			if ri < len(mh.data) && mh.data[ri] > mh.data[li] {
				// right child is even larger than left child, swap right child
				mh.data[ri], mh.data[i] = mh.data[i], mh.data[ri]
				i = ri
			} else {
				// left child is larger than new node, swap left child
				mh.data[li], mh.data[i] = mh.data[i], mh.data[li]
				i = li
			}
		} else if ri < len(mh.data) && mh.data[ri] > mh.data[i] {
			// right child is larger than new new, swap right child
			mh.data[ri], mh.data[i] = mh.data[i], mh.data[ri]
			i = ri
		} else {
			// all good, new node is at correct position
			break
		}
	}
	return item, nil
}
