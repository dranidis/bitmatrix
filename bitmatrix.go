package bitmatrix

// New creates a size x size matrix of bits
func New(size int) BitMatrix {
	if size < 9 {
		return newBitMatrixInt(size)
	} else {
		return newBitMatrixArray(size)
	}
}
