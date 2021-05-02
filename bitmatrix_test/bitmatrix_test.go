package bitmatrix_test

import (
	"testing"

	. "github.com/dranidis/bitmatrix"
)

type test struct {
	expected BitMatrix
	actual   BitMatrix
}

func TestUpExample(t *testing.T) {
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
		{NewBitMatrixInt(size).Read(up), NewBitMatrixInt(size).Read(board).Up()},
		{NewBitMatrixArray(size).Read(up), NewBitMatrixArray(size).Read(board).Up()},
	}

	for i, cas := range cases {
		if !cas.expected.Equal(cas.actual) {
			t.Errorf("Case: %d Not equal. expected \n%v actual \n%v", i, cas.expected, cas.actual)
		}
	}
}

func TestDnExample(t *testing.T) {
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
		{NewBitMatrixInt(size).Read(dn), NewBitMatrixInt(size).Read(board).Down()},
		{NewBitMatrixArray(size).Read(dn), NewBitMatrixArray(size).Read(board).Down()},
	}

	for i, cas := range cases {
		if !cas.expected.Equal(cas.actual) {
			t.Errorf("Case: %d Not equal. expected \n%v actual \n%v", i, cas.expected, cas.actual)
		}
	}
}

func TestLeftExample(t *testing.T) {
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
		{NewBitMatrixInt(size).Read(left), NewBitMatrixInt(size).Read(board).Left()},
		{NewBitMatrixArray(size).Read(left), NewBitMatrixArray(size).Read(board).Left()},
	}
	for i, cas := range cases {
		if !cas.expected.Equal(cas.actual) {
			t.Errorf("Case: %d Not equal. expected \n%v actual \n%v", i, cas.expected, cas.actual)
		}
	}
}

func TestRightExample(t *testing.T) {
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
		{NewBitMatrixInt(size).Read(right), NewBitMatrixInt(size).Read(board).Right()},
		{NewBitMatrixArray(size).Read(right), NewBitMatrixArray(size).Read(board).Right()},
	}
	for i, cas := range cases {
		if !cas.expected.Equal(cas.actual) {
			t.Errorf("Case: %d Not equal. expected \n%v actual \n%v", i, cas.expected, cas.actual)
		}
	}
}

func TestNoneGetAll(t *testing.T) {
	size := 3
	var b BitMatrix = NewBitMatrixArray(size)

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

func TestSet1GetAll(t *testing.T) {
	size := 3
	var b BitMatrix = NewBitMatrixArray(size)

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

func TestSet8GetAll(t *testing.T) {
	size := 3
	var b BitMatrix = NewBitMatrixArray(size)

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

func TestSetAllGetAll(t *testing.T) {
	size := 3
	var b BitMatrix = NewBitMatrixArray(size)

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

func TestSetGet(t *testing.T) {
	size := 3
	var b BitMatrix = NewBitMatrixArray(size)

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
