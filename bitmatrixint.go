package bitmatrix

import (
	"fmt"
	"math/bits"

	"github.com/fatih/color"
)

// bitMatrixInt implements the BitMatrix by using
// a single 64bit integer
// This is for matrices of max size 8x8.
type bitMatrixInt struct {
	size  int
	board uint64 // 8x8
}

var lmask = []uint64{
	0,
	0,
	10,
	438,
	61166,
	32472030,
	67628691390,
	558517276622718,
	18374403900871474942,
}

var rmask = []uint64{
	0,
	0,
	5,
	219,
	30583,
	16236015,
	33814345695,
	279258638311359,
	9187201950435737471,
}
var max = []uint64{
	0,
	1,
	15,
	511,
	65535,
	33554431,
	68719476735,
	562949953421311,
	18446744073709551615,
}

//
// Basic operations
//

func NewBitMatrixInt(size int) *bitMatrixInt {
	var b bitMatrixInt
	b.board = 0
	b.size = size
	return &b
}

func (b *bitMatrixInt) New(size int) BitMatrix {
	b = NewBitMatrixInt(size)
	return b
}

func (b *bitMatrixInt) Size() int {
	return b.size
}

func (b *bitMatrixInt) Equal(b1 BitMatrix) bool {
	return b.size == b1.(*bitMatrixInt).size && b.board == b1.(*bitMatrixInt).board
}

func (b *bitMatrixInt) Clone() BitMatrix {
	var clone = NewBitMatrixInt(b.size)
	clone.board = b.board
	return clone
}

func (b *bitMatrixInt) Read(board [][]int) BitMatrix {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == 1 {
				b.Set(EncodeWidth(i, j, b.size))
			}
		}
	}
	return b
}

func (b *bitMatrixInt) None() BitMatrix {
	b.board = 0
	return b
}

func (b *bitMatrixInt) Count() int {
	return bits.OnesCount64(b.board)
}

func (b *bitMatrixInt) String() string {
	col := color.New(color.FgRed)
	s := ""
	for i := 0; i < b.size*b.size; i++ {
		if i != 0 && i%b.size == 0 {
			s += "\n"
		}
		d := b.get(i)
		if d == 1 {
			s += col.Sprintf(fmt.Sprintf("%3d", b.get(i)))
		} else {
			s += fmt.Sprintf("%3d", b.get(i))
		}
	}
	return s + "\n"
}

func (b *bitMatrixInt) Set(index int) BitMatrix {
	b.board = b.board | (1 << index)
	return b
}

func (b *bitMatrixInt) get(index int) uint64 {
	return (b.board & (1 << index)) >> index
}

func (b *bitMatrixInt) Is(index int) bool {
	return b.get(index) == 1
}

func (b *bitMatrixInt) And(w BitMatrix) BitMatrix {
	b.board &= w.(*bitMatrixInt).board
	return b
}

func (b *bitMatrixInt) Or(w BitMatrix) BitMatrix {
	b.board |= w.(*bitMatrixInt).board
	return b
}

func (b *bitMatrixInt) Xor(w BitMatrix) BitMatrix {
	b.board ^= w.(*bitMatrixInt).board
	return b
}

func (b *bitMatrixInt) Inverse() BitMatrix {
	b.board = ^b.board
	return b
}

func (b *bitMatrixInt) Minus(w BitMatrix) BitMatrix {
	b.board = b.board & ^w.(*bitMatrixInt).board
	return b
}

// COMPLEX operations

func (b *bitMatrixInt) Up() BitMatrix {
	b.board = (b.board >> b.size) & max[b.size]
	return b
}

func (b *bitMatrixInt) Down() BitMatrix {
	b.board = (b.board << b.size) & max[b.size]
	return b
}

func (b *bitMatrixInt) Left() BitMatrix {
	b.board = ((b.board & lmask[b.size]) >> 1) & max[b.size]
	return b
}

func (b *bitMatrixInt) Right() BitMatrix {
	b.board = ((b.board & rmask[b.size]) << 1) & max[b.size]
	return b
}
