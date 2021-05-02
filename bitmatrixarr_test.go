package bitmatrix

import (
	"testing"
)

func TestBM2AllSetCount(t *testing.T) {
	b := NewBitMatrixArray(9)
	if count := b.All().Count(); count != 9*9 {
		t.Errorf("Wrong number of 1 bits: %d", count)
	}
}

func TestBM2RightMask(t *testing.T) {
	b := NewBitMatrixArray(9)
	if count := b.RightMask().Count(); count != 9*(9-1) {
		t.Errorf("Wrong number of 1 bits: %d", count)
	}
}

func TestBM2LeftMask(t *testing.T) {
	b := NewBitMatrixArray(9)
	if count := b.LeftMask().Count(); count != 9*(9-1) {
		t.Errorf("Wrong number of 1 bits: %d", count)
	}
}

func TestBM2LeftRightMask(t *testing.T) {
	b := NewBitMatrixArray(9)
	b1 := NewBitMatrixArray(9)

	// fmt.Printf("%064b\n", b.board1)
	// fmt.Printf("%064b\n", b.board2)
	if count := b.LeftMask().And(b1.RightMask()).Count(); count != 9*(9-2) {
		t.Errorf("Wrong number of 1 bits: %d", count)
	}
}

func TestBM2CachedAll(t *testing.T) {
	size := 9
	b1 := NewBitMatrixArray(size)
	b2 := NewBitMatrixArray(size)
	b3 := NewBitMatrixArray(size)
	if count := b1.All().Count(); count != size*size {
		t.Errorf("Wrong number of 1 bits: %d", count)
	}
	if count := b2.All().Count(); count != size*size {
		t.Errorf("Wrong number of 1 bits: %d", count)
	}
	if count := b3.All().Count(); count != size*size {
		t.Errorf("Wrong number of 1 bits: %d", count)
	}
}

func TestBM2SetGetLow(t *testing.T) {
	b := NewBitMatrixArray(9)
	for pos := 0; pos < b.size*b.size; pos++ {
		if b.Is(pos) {
			t.Errorf("Not all zero at beginning")
		}
	}
	b.Set(63)
	if !b.Is(63) {
		t.Errorf("Not Set/Get")
	}
}

func TestBM2SetGetHi(t *testing.T) {
	b := NewBitMatrixArray(9)
	for pos := 0; pos < b.size*b.size; pos++ {
		if b.Is(0) {
			t.Errorf("Not all zero at beginning")
		}
	}
	b.Set(80)
	if !b.Is(80) {
		t.Errorf("Not Set/Get")
	}
}

func TestBM2InverseInversesAllBits(t *testing.T) {
	b := NewBitMatrixArray(9)
	if count := b.Inverse().Count(); count != 81 {
		t.Errorf("Wrong number of 1 bits: %d", count)
	}
}
