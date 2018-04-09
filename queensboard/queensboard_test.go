package queensboard

import "testing"

func TestNew(t *testing.T) {
	b := New()
	if b == nil {
		t.Errorf("no board created")
	}
	if b.height != 8 {
		t.Errorf("got board height `%v` expected `8`", b.height)
	}
	if b.width != 8 {
		t.Errorf("got board width `%v` expected `8`", b.width)
	}
}

func TestNewCustom(t *testing.T) {
	t.Run("valid dimension", func(t *testing.T) {
		b, err := NewCustom(13, 17)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if b == nil {
			t.Errorf("no board created")
		}
		if b.height != 17 {
			t.Errorf("got board height `%v` expected `17`", b.height)
		}
		if b.width != 13 {
			t.Errorf("got board width `%v` expected `13`", b.width)
		}
	})
	t.Run("invalid dimension", func(t *testing.T) {
		_, err := NewCustom(-5, 0)
		if err == nil {
			t.Errorf("got `nil` expected error")
		}
	})
}

func TestWithinBounds(t *testing.T) {
	t.Run("within bounds", func(t *testing.T) {
		b := New()
		if ret := b.withinBounds(4, 5); !ret {
			t.Errorf("got `%v` expected `true`", ret)
		}
	})
	t.Run("out of bounds", func(t *testing.T) {
		b := New()
		if ret := b.withinBounds(-3, 5); ret {
			t.Errorf("got `%v` expected `false`", ret)
		}
		if ret := b.withinBounds(4, 99); ret {
			t.Errorf("got `%v` expected `false`", ret)
		}
	})
}

func TestPlaceQueen(t *testing.T) {
	t.Run("out of bounds", func(t *testing.T) {
		b := New()
		err := b.PlaceQueen(Coordinates{X: 10, Y: 10})
		if err != ErrorOutOfBounds {
			t.Errorf("bounds check failed")
		}
	})
	t.Run("queen already set", func(t *testing.T) {
		b := New()
		b.fields[1][1].queen = true
		err := b.PlaceQueen(Coordinates{X: 1, Y: 1})
		if err != ErrorOperationNotPermitted {
			t.Errorf("expected operation to fail")
		}
	})
	t.Run("field under attack", func(t *testing.T) {
		b := New()
		b.fields[1][1].attacks = 1
		err := b.PlaceQueen(Coordinates{X: 1, Y: 1})
		if err != ErrorOperationNotPermitted {
			t.Errorf("expected operation to fail")
		}
	})
	t.Run("successful placement", func(t *testing.T) {
		b := New()
		err := b.PlaceQueen(Coordinates{X: 3, Y: 3})
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		expectedBoard := [][]int{
			{1, 0, 0, 1, 0, 0, 1, 0},
			{0, 1, 0, 1, 0, 1, 0, 0},
			{0, 0, 1, 1, 1, 0, 0, 0},
			{1, 1, 1, 0, 1, 1, 1, 1},
			{0, 0, 1, 1, 1, 0, 0, 0},
			{0, 1, 0, 1, 0, 1, 0, 0},
			{1, 0, 0, 1, 0, 0, 1, 0},
			{0, 0, 0, 1, 0, 0, 0, 1},
		}
		for y := 0; y < 8; y++ {
			for x := 0; x < 8; x++ {
				expected := expectedBoard[y][x]
				got := b.fields[y][x].attacks
				if got != expected {
					t.Errorf("attack counter x=%v y=%v expected `%v` got `%v`",
						x, y, expected, got)
				}
			}
		}
		if queen := b.fields[3][3].queen; !queen {
			t.Errorf("queen boolean: expected `true` got `%v`", queen)
		}

		// another placement
		err = b.PlaceQueen(Coordinates{X: 1, Y: 0})
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		expectedBoard = [][]int{
			{2, 0, 1, 2, 1, 1, 2, 1},
			{1, 2, 1, 1, 0, 1, 0, 0},
			{0, 1, 1, 2, 1, 0, 0, 0},
			{1, 2, 1, 0, 2, 1, 1, 1},
			{0, 1, 1, 1, 1, 1, 0, 0},
			{0, 2, 0, 1, 0, 1, 1, 0},
			{1, 1, 0, 1, 0, 0, 1, 1},
			{0, 1, 0, 1, 0, 0, 0, 1},
		}
		for y := 0; y < 8; y++ {
			for x := 0; x < 8; x++ {
				expected := expectedBoard[y][x]
				got := b.fields[y][x].attacks
				if got != expected {
					t.Errorf("attack counter x=%v y=%v expected `%v` got `%v`",
						x, y, expected, got)
				}
			}
		}
		if queen := b.fields[0][1].queen; !queen {
			t.Errorf("queen boolean: expected `true` got `%v`", queen)
		}
	})
}

func TestRemoveQueen(t *testing.T) {
	t.Run("out of bounds", func(t *testing.T) {
		b := New()
		err := b.RemoveQueen(Coordinates{X: 10, Y: 10})
		if err != ErrorOutOfBounds {
			t.Errorf("bounds check failed")
		}
	})
	t.Run("queen not set", func(t *testing.T) {
		b := New()
		err := b.RemoveQueen(Coordinates{X: 1, Y: 1})
		if err != ErrorOperationNotPermitted {
			t.Errorf("expected operation to fail")
		}
	})
	t.Run("successful removal", func(t *testing.T) {
		b := New()
		_ = b.PlaceQueen(Coordinates{X: 3, Y: 3})
		err := b.RemoveQueen(Coordinates{X: 3, Y: 3})
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		expectedBoard := [][]int{
			{0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0},
		}
		for y := 0; y < 8; y++ {
			for x := 0; x < 8; x++ {
				expected := expectedBoard[y][x]
				got := b.fields[y][x].attacks
				if got != expected {
					t.Errorf("attack counter x=%v y=%v expected `%v` got `%v`",
						x, y, expected, got)
				}
			}
		}
		if queen := b.fields[3][3].queen; queen {
			t.Errorf("queen boolean: expected `false` got `%v`", queen)
		}
	})
}

func TestQueens(t *testing.T) {
	b := New()
	_ = b.PlaceQueen(Coordinates{X: 0, Y: 0})
	if got := b.Queens(); got != 1 {
		t.Errorf("queens on board: expected `%v` got `%v`", 1, got)
	}
	_ = b.PlaceQueen(Coordinates{X: 3, Y: 2})
	if got := b.Queens(); got != 2 {
		t.Errorf("queens on board: expected `%v` got `%v`", 2, got)
	}
	_ = b.RemoveQueen(Coordinates{X: 3, Y: 2})
	if got := b.Queens(); got != 1 {
		t.Errorf("queens on board: expected `%v` got `%v`", 1, got)
	}
}

func TestAvailableFields(t *testing.T) {
	b := New()
	f := b.AvailableFields()
	if len(f) != 64 {
		t.Errorf("expected length `%v` got length `%v`", 64, len(f))
	}

	_ = b.PlaceQueen(Coordinates{X: 3, Y: 3})
	f = b.AvailableFields()
	if len(f) != 36 {
		t.Errorf("expected length `%v` got length `%v`", 36, len(f))
	}

	_ = b.PlaceQueen(Coordinates{X: 1, Y: 0})
	f = b.AvailableFields()
	if len(f) != 22 {
		t.Errorf("expected length `%v` got length `%v`", 22, len(f))
	}
}
