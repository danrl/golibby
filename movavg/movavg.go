// Package movavg provides functions to calculate moving averages
package movavg

import "sync"

// MovAvg represents the moving average data type
type MovAvg struct {
	lock   sync.Mutex
	length uint
	ring   []float64
	index  uint
	cached bool
	sum    float64
}

// New creates a new moving average type of given length with caching disabled
func New(length uint) *MovAvg {
	return &MovAvg{
		length: length,
		ring:   make([]float64, length),
		index:  0,
		cached: false,
	}
}

// NewCached creates a new moving average type of given length with caching
// enabled. Caching may lead to less precise moving average results for certain
// data series
func NewCached(length uint) *MovAvg {
	return &MovAvg{
		length: length,
		ring:   make([]float64, length),
		index:  0,
		cached: true,
	}
}

// Add adds a value to the series and updates the moving average value for the
// series
func (ma *MovAvg) Add(n float64) float64 {
	ma.lock.Lock()
	defer ma.lock.Unlock()

	ma.sum = ma.sum - ma.ring[ma.index] + n
	ma.ring[ma.index] = n
	ma.index = (ma.index + 1) % ma.length

	if !ma.cached {
		ma.sum = 0.0
		for i := range ma.ring {
			ma.sum += ma.ring[i]
		}
	}
	return ma.sum / float64(ma.length)
}
