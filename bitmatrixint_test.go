package bitmatrix

import "testing"

func equalBoards(t *testing.T, expected *bitMatrixInt, actual *bitMatrixInt) {
	if expected.Size() != actual.Size() {
		t.Errorf("BinMatrices have different sizes: %d %d", expected.Size(), actual.Size())
	}
	if expected.board != actual.board {
		t.Errorf("Int boards are different: \n%v \n%v", expected, actual)
	}
	for i := 0; i < expected.Size()*expected.Size(); i++ {
		expectedVal := expected.Is(i)
		actualVal := actual.Is(i)
		if actualVal != expectedVal {
			t.Errorf("%d pos should be %v is: %v", i, expectedVal, actualVal)
		}
	}
}

func TestOne(t *testing.T) {
	size := 5
	one := NewBitMatrixInt(size).Set(0).(*bitMatrixInt)
	// fmt.Println(one)

	if one.board != uint64(1) {
		t.Errorf("Not one")
	}
}

func TestUpExample(t *testing.T) {
	size := 5
	var b = NewBitMatrixInt(size)

	board := [][]int{
		{0, 1, 0, 1, 0},
		{0, 1, 0, 1, 0},
		{0, 1, 0, 1, 1},
		{0, 0, 1, 0, 0},
		{0, 0, 0, 1, 1},
	}
	b.Read(board)
	// fmt.Println(b)
	b = b.Up().(*bitMatrixInt)
	// fmt.Println(b)

	up := [][]int{
		{0, 1, 0, 1, 0},
		{0, 1, 0, 1, 1},
		{0, 0, 1, 0, 0},
		{0, 0, 0, 1, 1},
		{0, 0, 0, 0, 0},
	}
	var bn = NewBitMatrixInt(size)
	bn.Read(up)
	equalBoards(t, b, bn)
}

func TestDnExample(t *testing.T) {
	size := 5
	var b = NewBitMatrixInt(size)

	board := [][]int{
		{0, 1, 0, 1, 0},
		{0, 1, 0, 1, 0},
		{0, 1, 0, 1, 1},
		{0, 0, 1, 0, 0},
		{0, 0, 0, 1, 1},
	}
	b.Read(board)
	// fmt.Println(b)
	b = b.Down().(*bitMatrixInt)
	// fmt.Println(b)

	dn := [][]int{
		{0, 0, 0, 0, 0},
		{0, 1, 0, 1, 0},
		{0, 1, 0, 1, 0},
		{0, 1, 0, 1, 1},
		{0, 0, 1, 0, 0},
	}
	var bn = NewBitMatrixInt(size)
	bn.Read(dn)
	equalBoards(t, b, bn)
}

func TestLeftExample(t *testing.T) {
	size := 5
	var b = NewBitMatrixInt(size)

	board := [][]int{
		{0, 1, 0, 1, 0},
		{0, 1, 0, 1, 0},
		{0, 1, 0, 1, 1},
		{0, 0, 1, 0, 0},
		{0, 0, 0, 1, 1},
	}
	b.Read(board)
	// fmt.Println(b)
	b = b.Left().(*bitMatrixInt)
	// fmt.Println(b)

	left := [][]int{
		{1, 0, 1, 0, 0},
		{1, 0, 1, 0, 0},
		{1, 0, 1, 1, 0},
		{0, 1, 0, 0, 0},
		{0, 0, 1, 1, 0},
	}
	var bn = NewBitMatrixInt(size)
	bn.Read(left)
	equalBoards(t, b, bn)
}

func TestRightExample(t *testing.T) {
	size := 5
	var b = NewBitMatrixInt(size)

	board := [][]int{
		{0, 1, 0, 1, 0},
		{0, 1, 0, 1, 0},
		{0, 1, 0, 1, 1},
		{0, 0, 1, 0, 0},
		{0, 0, 0, 1, 1},
	}
	b.Read(board)
	// fmt.Println(b)
	b = b.Right().(*bitMatrixInt)
	// fmt.Println(b)

	right := [][]int{
		{0, 0, 1, 0, 1},
		{0, 0, 1, 0, 1},
		{0, 0, 1, 0, 1},
		{0, 0, 0, 1, 0},
		{0, 0, 0, 0, 1},
	}
	var bn = NewBitMatrixInt(size)
	bn.Read(right)
	equalBoards(t, b, bn)
}
func TestNoneGetAll(t *testing.T) {
	size := 3
	var b = NewBitMatrixInt(size)

	b = b.None().(*bitMatrixInt)
	// fmt.Println(b)
	for i := 0; i < size; i++ {
		actual := b.get(i)
		expected := uint64(0)
		if actual != expected {
			t.Errorf("should be %d is: %d", expected, actual)
		}
	}
}

func TestSet1GetAll(t *testing.T) {
	size := 3
	var b = NewBitMatrixInt(size)

	b = b.None().Set(0).(*bitMatrixInt)
	// fmt.Println(b)
	for i := 0; i < size; i++ {
		expected := uint64(0)
		if i == 0 {
			expected = 1
		}
		actual := b.get(i)
		if actual != expected {
			t.Errorf("should be %d is: %d", expected, actual)
		}
	}
}

func TestSet8GetAll(t *testing.T) {
	size := 3
	var b = NewBitMatrixInt(size)

	b = b.None().Set(8).(*bitMatrixInt)
	// fmt.Println(b)
	for i := 0; i < size*size; i++ {
		expected := uint64(0)
		if i == 8 {
			expected = uint64(1)
		}
		actual := b.get(i)
		if actual != expected {
			t.Errorf("should be %d is: %d", expected, actual)
		}
	}
}

func TestSetAllGetAll(t *testing.T) {
	size := 3
	var b = NewBitMatrixInt(size)

	b = b.None().(*bitMatrixInt)
	for i := 0; i < size*size; i++ {
		b = b.Set(i).(*bitMatrixInt)
	}

	// fmt.Println(b)
	for i := 0; i < size*size; i++ {
		expected := uint64(1)
		actual := b.get(i)
		if actual != expected {
			t.Errorf("should be %d is: %d", expected, actual)
		}
	}
}

func TestSetGet(t *testing.T) {
	size := 3
	var b = NewBitMatrixInt(size)

	pos := []int{8, 0, 2, 4}
	for _, p := range pos {
		b = b.Set(p).(*bitMatrixInt)
	}
	// fmt.Println(b)
	for _, p := range pos {
		if !b.Is(p) {
			t.Errorf("should be set %d", pos)
		}
	}

}
