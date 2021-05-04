package bitmatrix

import (
	"testing"
)

type test struct {
	expected BitMatrix
	actual   BitMatrix
}

func TestBothUpExample(t *testing.T) {
	size := 5
	board := [][]int{
		{0, 1, 0, 1, 0},
		{0, 1, 0, 1, 0},
		{0, 1, 0, 1, 1},
		{0, 0, 1, 0, 0},
		{0, 0, 0, 1, 1},
	}
	up := [][]int{
		{0, 1, 0, 1, 0},
		{0, 1, 0, 1, 1},
		{0, 0, 1, 0, 0},
		{0, 0, 0, 1, 1},
		{0, 0, 0, 0, 0},
	}
	cases := []test{
		{newBitMatrixInt(size).Read(up), newBitMatrixInt(size).Read(board).Up()},
		{newBitMatrixArray(size).Read(up), newBitMatrixArray(size).Read(board).Up()},
	}

	for i, cas := range cases {
		if !cas.expected.Equal(cas.actual) {
			t.Errorf("Case: %d Not equal. expected \n%v actual \n%v", i, cas.expected, cas.actual)
		}
	}
}

func TestBothDnExample(t *testing.T) {
	size := 5
	board := [][]int{
		{0, 1, 0, 1, 0},
		{0, 1, 0, 1, 0},
		{0, 1, 0, 1, 1},
		{0, 0, 1, 0, 0},
		{0, 0, 0, 1, 1},
	}
	dn := [][]int{
		{0, 0, 0, 0, 0},
		{0, 1, 0, 1, 0},
		{0, 1, 0, 1, 0},
		{0, 1, 0, 1, 1},
		{0, 0, 1, 0, 0},
	}

	cases := []test{
		{newBitMatrixInt(size).Read(dn), newBitMatrixInt(size).Read(board).Down()},
		{newBitMatrixArray(size).Read(dn), newBitMatrixArray(size).Read(board).Down()},
	}

	for i, cas := range cases {
		if !cas.expected.Equal(cas.actual) {
			t.Errorf("Case: %d Not equal. expected \n%v actual \n%v", i, cas.expected, cas.actual)
		}
	}
}

func TestBothLeftExample(t *testing.T) {
	size := 5
	board := [][]int{
		{0, 1, 0, 1, 0},
		{0, 1, 0, 1, 0},
		{0, 1, 0, 1, 1},
		{0, 0, 1, 0, 0},
		{0, 0, 0, 1, 1},
	}
	left := [][]int{
		{1, 0, 1, 0, 0},
		{1, 0, 1, 0, 0},
		{1, 0, 1, 1, 0},
		{0, 1, 0, 0, 0},
		{0, 0, 1, 1, 0},
	}
	cases := []test{
		{newBitMatrixInt(size).Read(left), newBitMatrixInt(size).Read(board).Left()},
		{newBitMatrixArray(size).Read(left), newBitMatrixArray(size).Read(board).Left()},
	}
	for i, cas := range cases {
		if !cas.expected.Equal(cas.actual) {
			t.Errorf("Case: %d Not equal. expected \n%v actual \n%v", i, cas.expected, cas.actual)
		}
	}
}

func TestBothRightExample(t *testing.T) {
	size := 5
	board := [][]int{
		{0, 1, 0, 1, 0},
		{0, 1, 0, 1, 0},
		{0, 1, 0, 1, 1},
		{0, 0, 1, 0, 0},
		{0, 0, 0, 1, 1},
	}
	right := [][]int{
		{0, 0, 1, 0, 1},
		{0, 0, 1, 0, 1},
		{0, 0, 1, 0, 1},
		{0, 0, 0, 1, 0},
		{0, 0, 0, 0, 1},
	}
	cases := []test{
		{newBitMatrixInt(size).Read(right), newBitMatrixInt(size).Read(board).Right()},
		{newBitMatrixArray(size).Read(right), newBitMatrixArray(size).Read(board).Right()},
	}
	for i, cas := range cases {
		if !cas.expected.Equal(cas.actual) {
			t.Errorf("Case: %d Not equal. expected \n%v actual \n%v", i, cas.expected, cas.actual)
		}
	}
}

func TestBothNoneGetAll(t *testing.T) {
	size := 3
	var b BitMatrix = newBitMatrixArray(size)

	b = b.None()
	// fmt.Println(b)
	for i := 0; i < size; i++ {
		actual := b.Is(i)
		expected := false
		if actual != expected {
			t.Errorf("should be %v is: %v", expected, actual)
		}
	}
}

func TestBothSet1GetAll(t *testing.T) {
	size := 3
	var b BitMatrix = newBitMatrixArray(size)

	b = b.None().Set(0)
	// fmt.Println(b)
	for i := 0; i < size; i++ {
		expected := false
		if i == 0 {
			expected = true
		}
		actual := b.Is(i)
		if actual != expected {
			t.Errorf("should be %v is: %v", expected, actual)
		}
	}
}

func TestBothSet8GetAll(t *testing.T) {
	size := 3
	var b BitMatrix = newBitMatrixArray(size)

	b = b.None().Set(8)
	// fmt.Println(b)
	for i := 0; i < size*size; i++ {
		expected := false
		if i == 8 {
			expected = true
		}
		actual := b.Is(i)
		if actual != expected {
			t.Errorf("should be %v is: %v", expected, actual)
		}
	}
}

func TestBothSetAllGetAll(t *testing.T) {
	size := 3
	var b BitMatrix = newBitMatrixArray(size)

	b = b.None()
	for i := 0; i < size*size; i++ {
		b = b.Set(i)
	}

	// fmt.Println(b)
	for i := 0; i < size*size; i++ {
		expected := true
		actual := b.Is(i)
		if actual != expected {
			t.Errorf("should be %v is: %v", expected, actual)
		}
	}
}

func TestBothSetGet(t *testing.T) {
	size := 3
	var b BitMatrix = newBitMatrixArray(size)

	pos := []int{8, 0, 2, 4}
	for _, p := range pos {
		b = b.Set(p)
	}
	// fmt.Println(b)
	for _, p := range pos {
		if !b.Is(p) {
			t.Errorf("should be set %d", pos)
		}
	}
}
