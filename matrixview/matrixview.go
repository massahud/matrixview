// Copyright 2009 Geraldo Augusto Massahud Rodrigues dos Santos. All rights reserved.
// This code is licensed under APACHE-2 and MIT, the respective licenses can be found
// at LICENSE-APACHE and LICENSE-MIT files.

// library to create simple views of slice matrixes.
//
//The intention of this package is to reuse a single big matrix without needing to
//allocate more matrixes when you need different matrix sizes.
//e.g. calculate edit distance of lots of pairs of strings
package matrixview

// FromInterface receives a buffer matrix and returns a view of it
// setting with the specified number of rows and columns
// The matrix buffer must have the dimensions of rows and columns at least equal
// to rows and cols, or else it will error with index out of range
func FromInterface(buffer [][]interface{}, startRow, startCol, rows, cols int) [][]interface{} {
	view := buffer[startRow : startRow+rows : startRow+rows]
	for i := range view {
		view[i] = view[i][startCol : startCol+cols]
	}
	return view
}

// FromByte receives a buffer matrix and returns a view of it
// setting with the specified number of rows and columns
// The matrix buffer must have the dimensions of rows and columns at least equal
// to rows and cols, or else it will error with index out of range
func FromByte(buffer [][]byte, startRow, startCol, rows, cols int) [][]byte {
	view := buffer[startRow : startRow+rows : startRow+rows]
	for i := range view {
		view[i] = view[i][startCol : startCol+cols]
	}
	return view
}

// FromInt receives a buffer matrix and returns a view of it
// setting with the specified number of rows and columns
// The matrix buffer must have the dimensions of rows and columns at least equal
// to rows and cols, or else it will error with index out of range
func FromInt(buffer [][]int, startRow, startCol, rows, cols int) [][]int {
	view := make([][]int, rows)
	copy(view, buffer[startRow : startRow+rows : startRow+rows])
	for i := range view {
		view[i] = view[i][startCol : startCol+cols]
	}
	return view
}

// FromInt8 receives a buffer matrix and returns a view of it
// setting with the specified number of rows and columns
// The matrix buffer must have the dimensions of rows and columns at least equal
// to rows and cols, or else it will error with index out of range
func FromInt8(buffer [][]int8, startRow, startCol, rows, cols int) [][]int8 {
	view := buffer[startRow : startRow+rows : startRow+rows]
	for i := range view {
		view[i] = view[i][startCol : startCol+cols]
	}
	return view
}

// FromInt16 receives a buffer matrix and returns a view of it
// setting with the specified number of rows and columns
// The matrix buffer must have the dimensions of rows and columns at least equal
// to rows and cols, or else it will error with index out of range
func FromInt16(buffer [][]int16, startRow, startCol, rows, cols int) [][]int16 {
	view := buffer[startRow : startRow+rows : startRow+rows]
	for i := range view {
		view[i] = view[i][startCol : startCol+cols]
	}
	return view
}

// FromInt32 receives a buffer matrix and returns a view of it
// setting with the specified number of rows and columns
// The matrix buffer must have the dimensions of rows and columns at least equal
// to rows and cols, or else it will error with index out of range
func FromInt32(buffer [][]int32, startRow, startCol, rows, cols int) [][]int32 {
	view := buffer[startRow : startRow+rows : startRow+rows]
	for i := range view {
		view[i] = view[i][startCol : startCol+cols]
	}
	return view
}

// FromInt64 receives a buffer matrix and returns a view of it
// setting with the specified number of rows and columns
// The matrix buffer must have the dimensions of rows and columns at least equal
// to rows and cols, or else it will error with index out of range
func FromInt64(buffer [][]int64, startRow, startCol, rows, cols int) [][]int64 {
	view := buffer[startRow : startRow+rows : startRow+rows]
	for i := range view {
		view[i] = view[i][startCol : startCol+cols]
	}
	return view
}

// FromUint receives a buffer matrix and returns a view of it
// setting with the specified number of rows and columns
// The matrix buffer must have the dimensions of rows and columns at least equal
// to rows and cols, or else it will error with index out of range
func FromUint(buffer [][]uint, startRow, startCol, rows, cols int) [][]uint {
	view := buffer[startRow : startRow+rows : startRow+rows]
	for i := range view {
		view[i] = view[i][startCol : startCol+cols]
	}
	return view
}

// FromUint8 receives a buffer matrix and returns a view of it
// setting with the specified number of rows and columns
// The matrix buffer must have the dimensions of rows and columns at least equal
// to rows and cols, or else it will error with index out of range
func FromUint8(buffer [][]uint8, startRow, startCol, rows, cols int) [][]uint8 {
	view := buffer[startRow : startRow+rows : startRow+rows]
	for i := range view {
		view[i] = view[i][startCol : startCol+cols]
	}
	return view
}

// FromUint16 receives a buffer matrix and returns a view of it
// setting with the specified number of rows and columns
// The matrix buffer must have the dimensions of rows and columns at least equal
// to rows and cols, or else it will error with index out of range
func FromUint16(buffer [][]uint16, startRow, startCol, rows, cols int) [][]uint16 {
	view := buffer[startRow : startRow+rows : startRow+rows]
	for i := range view {
		view[i] = view[i][startCol : startCol+cols]
	}
	return view
}

// FromUint32 receives a buffer matrix and returns a view of it
// setting with the specified number of rows and columns
// The matrix buffer must have the dimensions of rows and columns at least equal
// to rows and cols, or else it will error with index out of range
func FromUint32(buffer [][]uint32, startRow, startCol, rows, cols int) [][]uint32 {
	view := buffer[startRow : startRow+rows : startRow+rows]
	for i := range view {
		view[i] = view[i][startCol : startCol+cols]
	}
	return view
}

// FromUint64 receives a buffer matrix and returns a view of it
// setting with the specified number of rows and columns
// The matrix buffer must have the dimensions of rows and columns at least equal
// to rows and cols, or else it will error with index out of range
func FromUint64(buffer [][]uint64, startRow, startCol, rows, cols int) [][]uint64 {
	view := buffer[startRow : startRow+rows : startRow+rows]
	for i := range view {
		view[i] = view[i][startCol : startCol+cols]
	}
	return view
}

// FromFloat32 receives a buffer matrix and returns a view of it
// setting with the specified number of rows and columns
// The matrix buffer must have the dimensions of rows and columns at least equal
// to rows and cols, or else it will error with index out of range
func FromFloat32(buffer [][]float32, startRow, startCol, rows, cols int) [][]float32 {
	view := buffer[startRow : startRow+rows : startRow+rows]
	for i := range view {
		view[i] = view[i][startCol : startCol+cols]
	}
	return view
}

// FromFloat64 receives a buffer matrix and returns a view of it
// setting with the specified number of rows and columns
// The matrix buffer must have the dimensions of rows and columns at least equal
// to rows and cols, or else it will error with index out of range
func FromFloat64(buffer [][]float64, startRow, startCol, rows, cols int) [][]float64 {
	view := buffer[startRow : startRow+rows : startRow+rows]
	for i := range view {
		view[i] = view[i][startCol : startCol+cols]
	}
	return view
}

// FromBool receives a buffer matrix and returns a view of it
// setting with the specified number of rows and columns
// The matrix buffer must have the dimensions of rows and columns at least equal
// to rows and cols, or else it will error with index out of range
func FromBool(buffer [][]bool, startRow, startCol, rows, cols int) [][]bool {
	view := buffer[startRow : startRow+rows : startRow+rows]
	for i := range view {
		view[i] = view[i][startCol : startCol+cols]
	}
	return view
}

// FromComplex64 receives a buffer matrix and returns a view of it
// setting with the specified number of rows and columns
// The matrix buffer must have the dimensions of rows and columns at least equal
// to rows and cols, or else it will error with index out of range
func FromComplex64(buffer [][]complex64, startRow, startCol, rows, cols int) [][]complex64 {
	view := buffer[startRow : startRow+rows : startRow+rows]
	for i := range view {
		view[i] = view[i][startCol : startCol+cols]
	}
	return view
}

// FromComplex128 receives a buffer matrix and returns a view of it
// setting with the specified number of rows and columns
// The matrix buffer must have the dimensions of rows and columns at least equal
// to rows and cols, or else it will error with index out of range
func FromComplex128(buffer [][]complex128, startRow, startCol, rows, cols int) [][]complex128 {
	view := buffer[startRow : startRow+rows : startRow+rows]
	for i := range view {
		view[i] = view[i][startCol : startCol+cols]
	}
	return view
}

// FromComplex128 receives a buffer matrix and returns a view of it
// setting with the specified number of rows and columns
// The matrix buffer must have the dimensions of rows and columns at least equal
// to rows and cols, or else it will error with index out of range
func FromString(buffer [][]string, startRow, startCol, rows, cols int) [][]string {
	view := buffer[startRow : startRow+rows : startRow+rows]
	for i := range view {
		view[i] = view[i][startCol : startCol+cols]
	}
	return view
}

// FromRune receives a buffer matrix and returns a view of it
// setting with the specified number of rows and columns
// The matrix buffer must have the dimensions of rows and columns at least equal
// to rows and cols, or else it will error with index out of range
func FromRune(buffer [][]rune, startRow, startCol, rows, cols int) [][]rune {
	view := buffer[startRow : startRow+rows : startRow+rows]
	for i := range view {
		view[i] = view[i][startCol : startCol+cols]
	}
	return view
}

// FromBytePtr receives a buffer matrix and returns a view of it
// setting with the specified number of rows and columns
// The matrix buffer must have the dimensions of rows and columns at least equal
// to rows and cols, or else it will error with index out of range
func FromBytePtr(buffer [][]*byte, startRow, startCol, rows, cols int) [][]*byte {
	view := buffer[startRow : startRow+rows : startRow+rows]
	for i := range view {
		view[i] = view[i][startCol : startCol+cols]
	}
	return view
}

// FromIntPtr receives a buffer matrix and returns a view of it
// setting with the specified number of rows and columns
// The matrix buffer must have the dimensions of rows and columns at least equal
// to rows and cols, or else it will error with index out of range
func FromIntPtr(buffer [][]*int, startRow, startCol, rows, cols int) [][]*int {
	view := buffer[startRow : startRow+rows : startRow+rows]
	for i := range view {
		view[i] = view[i][startCol : startCol+cols]
	}
	return view
}

// FromInt8Ptr receives a buffer matrix and returns a view of it
// setting with the specified number of rows and columns
// The matrix buffer must have the dimensions of rows and columns at least equal
// to rows and cols, or else it will error with index out of range
func FromInt8Ptr(buffer [][]*int8, startRow, startCol, rows, cols int) [][]*int8 {
	view := buffer[startRow : startRow+rows : startRow+rows]
	for i := range view {
		view[i] = view[i][startCol : startCol+cols]
	}
	return view
}

// FromInt16Ptr receives a buffer matrix and returns a view of it
// setting with the specified number of rows and columns
// The matrix buffer must have the dimensions of rows and columns at least equal
// to rows and cols, or else it will error with index out of range
func FromInt16Ptr(buffer [][]*int16, startRow, startCol, rows, cols int) [][]*int16 {
	view := buffer[startRow : startRow+rows : startRow+rows]
	for i := range view {
		view[i] = view[i][startCol : startCol+cols]
	}
	return view
}

// FromInt32Ptr receives a buffer matrix and returns a view of it
// setting with the specified number of rows and columns
// The matrix buffer must have the dimensions of rows and columns at least equal
// to rows and cols, or else it will error with index out of range
func FromInt32Ptr(buffer [][]*int32, startRow, startCol, rows, cols int) [][]*int32 {
	view := buffer[startRow : startRow+rows : startRow+rows]
	for i := range view {
		view[i] = view[i][startCol : startCol+cols]
	}
	return view
}

// FromInt64Ptr receives a buffer matrix and returns a view of it
// setting with the specified number of rows and columns
// The matrix buffer must have the dimensions of rows and columns at least equal
// to rows and cols, or else it will error with index out of range
func FromInt64Ptr(buffer [][]*int64, startRow, startCol, rows, cols int) [][]*int64 {
	view := buffer[startRow : startRow+rows : startRow+rows]
	for i := range view {
		view[i] = view[i][startCol : startCol+cols]
	}
	return view
}

// FromUintPtr receives a buffer matrix and returns a view of it
// setting with the specified number of rows and columns
// The matrix buffer must have the dimensions of rows and columns at least equal
// to rows and cols, or else it will error with index out of range
func FromUintPtr(buffer [][]*uint, startRow, startCol, rows, cols int) [][]*uint {
	view := buffer[startRow : startRow+rows : startRow+rows]
	for i := range view {
		view[i] = view[i][startCol : startCol+cols]
	}
	return view
}

// FromUint8Ptr receives a buffer matrix and returns a view of it
// setting with the specified number of rows and columns
// The matrix buffer must have the dimensions of rows and columns at least equal
// to rows and cols, or else it will error with index out of range
func FromUint8Ptr(buffer [][]*uint8, startRow, startCol, rows, cols int) [][]*uint8 {
	view := buffer[startRow : startRow+rows : startRow+rows]
	for i := range view {
		view[i] = view[i][startCol : startCol+cols]
	}
	return view
}

// FromUint16Ptr receives a buffer matrix and returns a view of it
// setting with the specified number of rows and columns
// The matrix buffer must have the dimensions of rows and columns at least equal
// to rows and cols, or else it will error with index out of range
func FromUint16Ptr(buffer [][]*uint16, startRow, startCol, rows, cols int) [][]*uint16 {
	view := buffer[startRow : startRow+rows : startRow+rows]
	for i := range view {
		view[i] = view[i][startCol : startCol+cols]
	}
	return view
}

// FromUint32Ptr receives a buffer matrix and returns a view of it
// setting with the specified number of rows and columns
// The matrix buffer must have the dimensions of rows and columns at least equal
// to rows and cols, or else it will error with index out of range
func FromUint32Ptr(buffer [][]*uint32, startRow, startCol, rows, cols int) [][]*uint32 {
	view := buffer[startRow : startRow+rows : startRow+rows]
	for i := range view {
		view[i] = view[i][startCol : startCol+cols]
	}
	return view
}

// FromUint64Ptr receives a buffer matrix and returns a view of it
// setting with the specified number of rows and columns
// The matrix buffer must have the dimensions of rows and columns at least equal
// to rows and cols, or else it will error with index out of range
func FromUint64Ptr(buffer [][]*uint64, startRow, startCol, rows, cols int) [][]*uint64 {
	view := buffer[startRow : startRow+rows : startRow+rows]
	for i := range view {
		view[i] = view[i][startCol : startCol+cols]
	}
	return view
}

// FromFloat32Ptr receives a buffer matrix and returns a view of it
// setting with the specified number of rows and columns
// The matrix buffer must have the dimensions of rows and columns at least equal
// to rows and cols, or else it will error with index out of range
func FromFloat32Ptr(buffer [][]*float32, startRow, startCol, rows, cols int) [][]*float32 {
	view := buffer[startRow : startRow+rows : startRow+rows]
	for i := range view {
		view[i] = view[i][startCol : startCol+cols]
	}
	return view
}

// FromFloat64Ptr receives a buffer matrix and returns a view of it
// setting with the specified number of rows and columns
// The matrix buffer must have the dimensions of rows and columns at least equal
// to rows and cols, or else it will error with index out of range
func FromFloat64Ptr(buffer [][]*float64, startRow, startCol, rows, cols int) [][]*float64 {
	view := buffer[startRow : startRow+rows : startRow+rows]
	for i := range view {
		view[i] = view[i][startCol : startCol+cols]
	}
	return view
}

// FromBoolPtr receives a buffer matrix and returns a view of it
// setting with the specified number of rows and columns
// The matrix buffer must have the dimensions of rows and columns at least equal
// to rows and cols, or else it will error with index out of range
func FromBoolPtr(buffer [][]*bool, startRow, startCol, rows, cols int) [][]*bool {
	view := buffer[startRow : startRow+rows : startRow+rows]
	for i := range view {
		view[i] = view[i][startCol : startCol+cols]
	}
	return view
}

// FromComplex64Ptr receives a buffer matrix and returns a view of it
// setting with the specified number of rows and columns
// The matrix buffer must have the dimensions of rows and columns at least equal
// to rows and cols, or else it will error with index out of range
func FromComplex64Ptr(buffer [][]*complex64, startRow, startCol, rows, cols int) [][]*complex64 {
	view := buffer[startRow : startRow+rows : startRow+rows]
	for i := range view {
		view[i] = view[i][startCol : startCol+cols]
	}
	return view
}

// FromComplex128Ptr receives a buffer matrix and returns a view of it
// setting with the specified number of rows and columns
// The matrix buffer must have the dimensions of rows and columns at least equal
// to rows and cols, or else it will error with index out of range
func FromComplex128Ptr(buffer [][]*complex128, startRow, startCol, rows, cols int) [][]*complex128 {
	view := buffer[startRow : startRow+rows : startRow+rows]
	for i := range view {
		view[i] = view[i][startCol : startCol+cols]
	}
	return view
}

// FromComplex128Ptr receives a buffer matrix and returns a view of it
// setting with the specified number of rows and columns
// The matrix buffer must have the dimensions of rows and columns at least equal
// to rows and cols, or else it will error with index out of range
func FromStringPtr(buffer [][]*string, startRow, startCol, rows, cols int) [][]*string {
	view := buffer[startRow : startRow+rows : startRow+rows]
	for i := range view {
		view[i] = view[i][startCol : startCol+cols]
	}
	return view
}

// FromRunePtr receives a buffer matrix and returns a view of it
// setting with the specified number of rows and columns
// The matrix buffer must have the dimensions of rows and columns at least equal
// to rows and cols, or else it will error with index out of range
func FromRunePtr(buffer [][]*rune, startRow, startCol, rows, cols int) [][]*rune {
	view := buffer[startRow : startRow+rows : startRow+rows]
	for i := range view {
		view[i] = view[i][startCol : startCol+cols]
	}
	return view
}
