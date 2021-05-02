package bitmatrix

// New creates a size x size matrix of bits
func New(size int) BitMatrix {
	if size < 9 {
		return NewBitMatrixInt(size)
	} else {
		return NewBitMatrixArray(size)
	}
}
