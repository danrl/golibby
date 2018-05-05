package movavg

import "testing"

func TestNew(t *testing.T) {
	ma := New(5)
	if got := ma.index; got != 0 {
		t.Error("expected index 0, got", got)
	}
	if got := ma.sum; got != 0.0 {
		t.Error("expected sum 0.0, got", got)
	}
	if got := ma.cached; got != false {
		t.Error("expected caching disabled, got enabled")
	}
	if got := len(ma.ring); got != 5 {
		t.Error("expected ring size 5, got", got)
	}
}

func TestNewCached(t *testing.T) {
	ma := NewCached(7)
	if got := ma.index; got != 0 {
		t.Error("expected index 0, got", got)
	}
	if got := ma.sum; got != 0.0 {
		t.Error("expected sum 0.0, got", got)
	}
	if got := ma.cached; got != true {
		t.Error("expected caching enabled, got disabled")
	}
	if got := len(ma.ring); got != 7 {
		t.Error("expected ring size 5, got", got)
	}
}

func TestAdd(t *testing.T) {
	steps := []struct {
		in  float64
		out float64
	}{
		{in: 3, out: 1},
		{in: 3, out: 2},
		{in: 9, out: 5},
		{in: 0, out: 4},
		{in: 0, out: 3},
		{in: 0, out: 0},
		{in: 60, out: 20},
		{in: 3, out: 21},
		{in: 1, out: 64.0 / 3.0},
	}
	t.Run("non-cached moving average", func(t *testing.T) {
		ma := New(3)
		for _, step := range steps {
			if got := ma.Add(step.in); got != step.out {
				t.Errorf("add %v: expected %v got %v", step.in, step.out, got)
			}
		}
	})
	t.Run("cached moving average", func(t *testing.T) {
		ma := NewCached(3)
		for _, step := range steps {
			if got := ma.Add(step.in); got != step.out {
				t.Errorf("add %v: expected %v got %v", step.in, step.out, got)
			}
		}
	})
}
