package tensor

import (
	"testing"
)

func TestTensor(t *testing.T) {
	ts := NewTensor(WithRank(1))

	ts.Set(NewIndex(0), 1.0)
	ts.Set(NewIndex(1), 2.0)
	ts.Set(NewIndex(2), 3.0)

	if ts.Get(NewIndex(0)).Float() != 1.0 {
		t.Error("expected 1.0")
	}

	if ts.Get(NewIndex(1)).Float() != 2.0 {
		t.Error("expected 2.0")
	}

	if ts.Get(NewIndex(2)).Float() != 3.0 {
		t.Error("expected 3.0")
	}
}

func TestFromArray(t *testing.T) {
	ts := NewTensor(FromArray([]float64{1.0, 2.0, 3.0}))

	if ts.Get(NewIndex(0)).Float() != 1.0 {
		t.Error("expected 1.0")
	}

	if ts.Get(NewIndex(1)).Float() != 2.0 {
		t.Error("expected 2.0")
	}

	if ts.Get(NewIndex(2)).Float() != 3.0 {
		t.Error("expected 3.0")
	}
}

func TestIndexes(t *testing.T) {
	ts := NewTensor(WithRank(2))

	ts.Set(NewIndex(0, 2), 10.0)

	if ts.At(0, 2).Float() != 10.0 {
		t.Error("expected 10.0, got", ts.At(0, 2))
	}

	ts.Set(NewIndex(5, 3), 3.0)

	if ts.At(5, 3).Float() != 3.0 {
		t.Error("expected 3.0, got", ts.At(5))
	}
}

func TestMultiply(t *testing.T) {
	t1 := NewTensor(WithRank(2))

	t1.Set(NewIndex(0, 0), 2.0)
	t1.Set(NewIndex(1, 0), 3.0)
	t1.Set(NewIndex(0, 1), 10.0)
	t1.Set(NewIndex(1, 1), 8.0)

	t2 := NewTensor(WithRank(2))

	t2.Set(NewIndex(0, 0), 1.0)
	t2.Set(NewIndex(0, 1), 4.0)

	result := t1.Multiply(t2)

	if result.At(0, 0).Float() != 14.0 {
		t.Error("expected 14.0, got", result.At(0, 1))
	}

	if result.At(0, 1).Float() != 42.0 {
		t.Error("expected 42.0, got", result.At(0, 1))
	}
}

func TestFlip(t *testing.T) {
	ts := NewTensor(FromArray([]float64{1.0, 2.0, 3.0}))

	flipped := ts.Flip()
	flipped.PrintDebug()

	if flipped.At(0, 0).Float() != 1.0 {
		t.Error("expected 1.0, got", flipped.At(0, 0))
	}

	if flipped.At(0, 1).Float() != 2.0 {
		t.Error("expected 2.0, got", flipped.At(0, 1))
	}

	if flipped.At(0, 2).Float() != 3.0 {
		t.Error("expected 3.0, got", flipped.At(0, 2))
	}
}
