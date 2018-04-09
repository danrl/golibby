// Package queensboard provides a data structure and method similar to a chess
// board. Queens can be placed on the board, but not within the attack range of
// another queen. The purpose of this package is to provide the underlying tools
// for demonstrating simple backtracking algorithms.
package queensboard

import (
	"fmt"
	"sync"
)

var (
	// ErrorOperationNotPermitted indicates that an operation is not permitted
	// on a particular field. E.g. placing a queen on an attacked field or
	// removing a queen from a field that holds no queen
	ErrorOperationNotPermitted = fmt.Errorf("operation not permitted")
	// ErrorOutOfBounds is returned when coordinates point to a field outside
	// of the board
	ErrorOutOfBounds = fmt.Errorf("out of bounds coordinates")
	// ErrorInvalidBoardDimensions is returned in case of invalid width or
	// height values for creating a new board
	ErrorInvalidBoardDimensions = fmt.Errorf("invalid board dimensions")
)

// Coordinates holds the x and y values for addressing a field on the board
type Coordinates struct {
	X, Y int
}

type field struct {
	queen   bool // true if a queen is placed on this field
	attacks int  // number of times the field is under attack by a queen
}

// Board holds a concurrency-safe Queens Board
type Board struct {
	lock   sync.RWMutex
	fields [][]field
	queens int
	width  int
	height int
}

// New returns an empty standard-sized (8x8) fields board
func New() *Board {
	b, _ := NewCustom(8, 8)
	return b
}

// NewCustom returns a board with custom width and height
func NewCustom(width, height int) (*Board, error) {
	if width < 1 || height < 1 {
		return nil, ErrorInvalidBoardDimensions
	}
	b := &Board{
		width:  width,
		height: height,
	}
	b.fields = make([][]field, b.height, b.height)
	for i := range b.fields {
		b.fields[i] = make([]field, b.width, b.width)
	}
	return b, nil
}

func (b *Board) withinBounds(x, y int) bool {
	if x >= 0 && y >= 0 && x < b.width && y < b.height {
		return true
	}
	return false
}

// PlaceQueen places a queen on the board and updates the attack counters of all
// affected fields.
func (b *Board) PlaceQueen(c Coordinates) error {
	if !b.withinBounds(c.X, c.Y) {
		return ErrorOutOfBounds
	}
	b.lock.Lock()
	defer b.lock.Unlock()

	if b.fields[c.Y][c.X].queen || b.fields[c.Y][c.X].attacks > 0 {
		return ErrorOperationNotPermitted
	}

	// increment fields attack counter: horizontally
	for j := 0; b.withinBounds(j, c.Y); j++ {
		b.fields[c.Y][j].attacks++
	}
	// increment fields attack counter: vertically
	for i := 0; b.withinBounds(c.X, i); i++ {
		b.fields[i][c.X].attacks++
	}
	// increment fields attack counter: up-left
	for j, i := c.X-1, c.Y-1; b.withinBounds(j, i); {
		b.fields[i][j].attacks++
		j--
		i--
	}
	// increment fields attack counter: up-right
	for j, i := c.X+1, c.Y-1; b.withinBounds(j, i); {
		b.fields[i][j].attacks++
		j++
		i--
	}
	// increment fields attack counter: down-left
	for j, i := c.X-1, c.Y+1; b.withinBounds(j, i); {
		b.fields[i][j].attacks++
		j--
		i++
	}
	// increment fields attack counter: down-right
	for j, i := c.X+1, c.Y+1; b.withinBounds(j, i); {
		b.fields[i][j].attacks++
		j++
		i++
	}
	// place queen
	b.fields[c.Y][c.X].attacks = 0
	b.fields[c.Y][c.X].queen = true
	b.queens++
	return nil
}

// RemoveQueen removes a queen from the board and updates the attack counters of
// all affected fields.
func (b *Board) RemoveQueen(c Coordinates) error {
	if !b.withinBounds(c.X, c.Y) {
		return ErrorOutOfBounds
	}
	b.lock.Lock()
	defer b.lock.Unlock()

	if !b.fields[c.Y][c.X].queen {
		return ErrorOperationNotPermitted
	}

	// decrement fields attack counter: horizontally
	for j := 0; b.withinBounds(j, c.Y); j++ {
		b.fields[c.Y][j].attacks--
	}
	// decrement fields attack counter: vertically
	for i := 0; b.withinBounds(c.X, i); i++ {
		b.fields[i][c.X].attacks--
	}
	// decrement fields attack counter: up-left
	for j, i := c.X-1, c.Y-1; b.withinBounds(j, i); {
		b.fields[i][j].attacks--
		j--
		i--
	}
	// decrement fields attack counter: up-right
	for j, i := c.X+1, c.Y-1; b.withinBounds(j, i); {
		b.fields[i][j].attacks--
		j++
		i--
	}
	// decrement fields attack counter: down-left
	for j, i := c.X-1, c.Y+1; b.withinBounds(j, i); {
		b.fields[i][j].attacks--
		j--
		i++
	}
	// decrement fields attack counter: down-right
	for j, i := c.X+1, c.Y+1; b.withinBounds(j, i); {
		b.fields[i][j].attacks--
		j++
		i++
	}
	// remove queen
	b.fields[c.Y][c.X].attacks = 0
	b.fields[c.Y][c.X].queen = false
	b.queens--
	return nil
}

// Queens returns the number of queens on the board
func (b *Board) Queens() int {
	b.lock.RLock()
	defer b.lock.RUnlock()
	return b.queens
}

// AvailableFields returns a list of fields available for queen placement
func (b *Board) AvailableFields() []Coordinates {
	var cs []Coordinates
	b.lock.RLock()
	defer b.lock.RUnlock()
	for i := range b.fields {
		for j := range b.fields[i] {
			if b.fields[i][j].attacks > 0 {
				continue
			}
			if b.fields[i][j].queen {
				continue
			}
			cs = append(cs, Coordinates{X: j, Y: i})
		}
	}
	return cs
}
