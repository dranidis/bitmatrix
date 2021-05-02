package bitmatrix

type BitMatrix interface {
	// New returns a new BitMatrix. It does not change the receiver.
	New(size int) BitMatrix

	// None sets all the bits of the receiver matrix to 0.
	// It returns the receiver matrix to support chaining.
	None() BitMatrix

	// Size returns the size of the matrix size x size.
	// It does not change the receiver matrix.
	Size() int

	// Set sets the index position in the receiver matrix to 1.
	// It returns the receiver matrix to support chaining.
	Set(index int) BitMatrix

	// Is returns true if the bit at the index is 1.
	// It does not change the receiver matrix.
	Is(index int) bool

	// Equal compares the receiver matrix with the argument matrix.
	// It does not change the receiver matrix.
	Equal(matrix BitMatrix) bool

	// Count returns the number of 1 bits of the receiver matrix.
	// It does not change the receiver matrix.
	Count() int

	// Clones creates and returns a clone of the receiver matrix.
	// It does not change the receiver matrix.
	Clone() BitMatrix

	// Up shifts all the bits of the receiver matrix up.
	// The bottom row is filled with 0s.
	// It returns the receiver matrix to support chaining.
	Up() BitMatrix

	// Down shifts all the bits of the receiver matrix down.
	// The top row is filled with 0s.
	// It returns the receiver matrix to support chaining.
	Down() BitMatrix

	// Left shifts all the bits of the receiver matrix to the left.
	// The leftmost column is filled with 0s.
	// It returns the receiver matrix to support chaining.
	Left() BitMatrix

	// Right shifts all the bits of the receiver matrix to the right.
	// The rightmost column is filled with 0s.
	// It returns the receiver matrix to support chaining.
	Right() BitMatrix

	// Inverse inverts all bits of receiver matrix.
	// It returns the receiver matrix to support chaining.
	Inverse() BitMatrix

	// Or performs a boolean OR operation of the receiver and the argument matrix.
	// The receiver is changed to the result of the operation.
	// It returns the receiver matrix to support chaining.
	Or(BitMatrix) BitMatrix

	// And performs a boolean AND operation of the receiver and the argument matrix.
	// The receiver is changed to the result of the operation.
	// It returns the receiver matrix to support chaining.
	And(BitMatrix) BitMatrix

	// Xor performs a boolean XOR operation of the receiver and the argument matrix.
	// The receiver is changed to the result of the operation.
	// It returns the receiver matrix to support chaining.
	Xor(BitMatrix) BitMatrix

	// Minuns performs a boolean MINUS operation of the receiver and the argument matrix.
	// The receiver is changed to the result of the operation.
	// It returns the receiver matrix to support chaining.
	Minus(BitMatrix) BitMatrix

	// Read reads the board 2 dimensional slice and sets all the bits of the receiver matrix.
	// It returns the receiver matrix to support chaining.
	Read(board [][]int) BitMatrix
}
