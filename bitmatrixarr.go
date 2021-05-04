package bitmatrix

import (
	"fmt"
	"log"

	"github.com/dranidis/bitarray"
	"github.com/fatih/color"
)

type bitMatrixArray struct {
	size  int
	board *bitarray.BitArray
}

var rightMask map[int]*bitarray.BitArray
var leftMask map[int]*bitarray.BitArray

func initRightMask(num int) *bitarray.BitArray {
	rightMask := bitarray.New(num * num)
	for pos := 0; pos < num*num; pos++ {
		if (pos+1)%num != 0 {
			rightMask.Set(pos)
		}
	}
	return rightMask
}

func initLeftMask(num int) *bitarray.BitArray {
	leftMask := bitarray.New(num * num)
	for pos := 0; pos < num*num; pos++ {
		if (pos)%num != 0 {
			leftMask.Set(pos)
		}
	}
	return leftMask
}

// NewBitMatrixArray creates a size x size matrix of bits.
// There are no size restrictions. The underlying data strucure is a bit array.
func NewBitMatrixArray(size int) *bitMatrixArray {
	var b bitMatrixArray
	b.board = bitarray.New(size * size)

	b.size = size

	if rightMask == nil {
		rightMask = make(map[int]*bitarray.BitArray)
	}
	if leftMask == nil {
		leftMask = make(map[int]*bitarray.BitArray)
	}
	_, ok := rightMask[size]
	if !ok {
		rightMask[size] = initRightMask(size)
	}
	_, ok = leftMask[size]
	if !ok {
		leftMask[size] = initLeftMask(size)
	}
	return &b
}

func (b *bitMatrixArray) New(size int) BitMatrix {
	b = NewBitMatrixArray(size)
	return b
}

func (b *bitMatrixArray) Size() int {
	return b.size
}

func (b *bitMatrixArray) Clone() BitMatrix {
	var clone = NewBitMatrixArray(b.size)
	clone.board = b.board.Clone()
	return clone
}

func (b *bitMatrixArray) None() BitMatrix {
	b.board = b.board.None()
	return b
}

func (b *bitMatrixArray) All() BitMatrix {
	b.board = b.board.All()
	return b
}

func (b *bitMatrixArray) rightMask() *bitarray.BitArray {
	return rightMask[b.size].Clone()
}

func (b *bitMatrixArray) leftMask() *bitarray.BitArray {
	return leftMask[b.size].Clone()
}

func (b *bitMatrixArray) Count() int {
	return b.board.Count()
}

func (b *bitMatrixArray) Set(index int) BitMatrix {
	b.board = b.board.Set(index)
	return b
}

func (b *bitMatrixArray) Is(index int) bool {
	return b.board.Is(index)
}

// boolean operators work on all bits.
// Also on the ones outside the board area
func (b *bitMatrixArray) And(w BitMatrix) BitMatrix {
	b.board = b.board.And(w.(*bitMatrixArray).board)
	return b
}

func (b *bitMatrixArray) Or(w BitMatrix) BitMatrix {
	b.board = b.board.Or(w.(*bitMatrixArray).board)
	return b
}

func (b *bitMatrixArray) Xor(w BitMatrix) BitMatrix {
	b.board = b.board.Xor(w.(*bitMatrixArray).board)
	return b
}

func (b *bitMatrixArray) Inverse() BitMatrix {
	b.board = b.board.Inverse()
	return b
}

func (b *bitMatrixArray) Minus(w BitMatrix) BitMatrix {
	b.board = b.board.Minus(w.(*bitMatrixArray).board)
	return b
}

// COPY from BinMatrix1

// COMPLEX operations

func (b *bitMatrixArray) Up() BitMatrix {
	max := bitarray.New(b.size * b.size)
	max = max.All()
	b.board = b.board.ShiftRight(b.size).And(max)
	return b
}

func (b *bitMatrixArray) Down() BitMatrix {
	max := bitarray.New(b.size * b.size)
	max = max.All()
	b.board = b.board.ShiftLeft(b.size).And(max)
	return b
}

func (b *bitMatrixArray) Left() BitMatrix {
	max := bitarray.New(b.size * b.size)
	max = max.All()
	lmask := NewBitMatrixArray(b.size)
	lmaskBinArray := lmask.leftMask()
	b.board = b.board.And(lmaskBinArray).ShiftRight(1).And(max)
	return b
}

func (b *bitMatrixArray) Right() BitMatrix {
	max := bitarray.New(b.size * b.size)
	max = max.All()
	rmask := NewBitMatrixArray(b.size)
	rmaskBinArray := rmask.rightMask()
	b.board = b.board.And(rmaskBinArray).ShiftLeft(1).And(max)
	return b
}

func (b *bitMatrixArray) String() string {
	col := color.New(color.FgRed)
	s := ""
	for i := 0; i < b.size*b.size; i++ {
		if i != 0 && i%b.size == 0 {
			s += "\n"
		}
		d := b.board.Get(i)
		if d == 1 {
			s += col.Sprintf(fmt.Sprintf("%3d", d))
		} else {
			s += fmt.Sprintf("%3d", d)
		}
	}
	return s + "\n"
}

func (b *bitMatrixArray) Equal(w BitMatrix) bool {
	return b.board.Equal(w.(*bitMatrixArray).board)
}

func (b *bitMatrixArray) Read(board [][]int) BitMatrix {
	if b.size != len(board) {
		log.Panicf("BinMatrixArray.Read Cannot read. Different size: \n%v\n", board)
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == 1 {
				xy := EncodeWidth(i, j, b.size)
				b.Set(xy)
			}
		}
	}
	return b
}
