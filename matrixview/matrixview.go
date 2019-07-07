// Copyright 2009 Geraldo Augusto Massahud Rodrigues dos Santos. All rights reserved.
// This code is licensed under APACHE-2 and MIT, the respective licenses can be found
// at LICENSE-APACHE and LICENSE-MIT files.

// package to create simple views of slice matrixes.
//
// The intention of this package is to reuse a single big matrix without needing to
// allocate more matrixes when you need different matrix sizes. e.g. calculate edit distance of lots of pairs of strings
//
// The functions of this package allocate only the rows when executing, reusing the column slices.
package matrixview

import "fmt"

// FromInterface receives a buffer matrix and returns a view of it
// setting with the specified number of rows and columns
// The matrix buffer must have the dimensions of rows and columns at least equal
// to rows and cols, or else it will error with index out of range
func FromInterface(buffer [][]interface{}, startRow, startCol, rows, cols int) ([][]interface{}, error) {
	view := make([][]interface{}, rows)
	if len(buffer) < startRow+rows {
		return nil, fmt.Errorf("matrix has less rows than asked: %d x %d", len(buffer), startRow+rows)
	}
	for i := range view {
		if len(buffer[startRow+i]) < startCol+cols {
			return nil, fmt.Errorf("row %d has less cols than asked: %d x %d", i, len(buffer[startRow+i]), startCol+cols)
		}
		view[i] = buffer[startRow+i][startCol : startCol+cols : startCol+cols]
	}
	return view, nil
}

// FromByte receives a buffer matrix and returns a view of it
// setting with the specified number of rows and columns
// The matrix buffer must have the dimensions of rows and columns at least equal
// to rows and cols, or else it will error with index out of range
func FromByte(buffer [][]byte, startRow, startCol, rows, cols int) ([][]byte, error) {
	view := make([][]byte, rows)
	if len(buffer) < startRow+rows {
		return nil, fmt.Errorf("matrix has less rows than asked: %d x %d", len(buffer), startRow+rows)
	}
	for i := range view {
		if len(buffer[startRow+i]) < startCol+cols {
			return nil, fmt.Errorf("row %d has less cols than asked: %d x %d", i, len(buffer[startRow+i]), startCol+cols)
		}
		view[i] = buffer[startRow+i][startCol : startCol+cols : startCol+cols]
	}
	return view, nil
}

// FromInt receives a buffer matrix and returns a view of it
// setting with the specified number of rows and columns
// The matrix buffer must have the dimensions of rows and columns at least equal
// to rows and cols, or else it will error with index out of range
func FromInt(buffer [][]int, startRow, startCol, rows, cols int) ([][]int, error) {
	view := make([][]int, rows)
	if len(buffer) < startRow+rows {
		return nil, fmt.Errorf("matrix has less rows than asked: %d x %d", len(buffer), startRow+rows)
	}
	for i := range view {
		if len(buffer[startRow+i]) < startCol+cols {
			return nil, fmt.Errorf("row %d has less cols than asked: %d x %d", i, len(buffer[startRow+i]), startCol+cols)
		}
		view[i] = buffer[startRow+i][startCol : startCol+cols : startCol+cols ]
	}
	return view, nil
}

// FromInt8 receives a buffer matrix and returns a view of it
// setting with the specified number of rows and columns
// The matrix buffer must have the dimensions of rows and columns at least equal
// to rows and cols, or else it will error with index out of range
func FromInt8(buffer [][]int8, startRow, startCol, rows, cols int) ([][]int8, error) {
	view := make([][]int8, rows)
	if len(buffer) < startRow+rows {
		return nil, fmt.Errorf("matrix has less rows than asked: %d x %d", len(buffer), startRow+rows)
	}
	for i := range view {
		if len(buffer[startRow+i]) < startCol+cols {
			return nil, fmt.Errorf("row %d has less cols than asked: %d x %d", i, len(buffer[startRow+i]), startCol+cols)
		}
		view[i] = buffer[startRow+i][startCol : startCol+cols : startCol+cols]
	}
	return view, nil
}

// FromInt16 receives a buffer matrix and returns a view of it
// setting with the specified number of rows and columns
// The matrix buffer must have the dimensions of rows and columns at least equal
// to rows and cols, or else it will error with index out of range
func FromInt16(buffer [][]int16, startRow, startCol, rows, cols int) ([][]int16, error) {
	view := make([][]int16, rows)
	if len(buffer) < startRow+rows {
		return nil, fmt.Errorf("matrix has less rows than asked: %d x %d", len(buffer), startRow+rows)
	}
	for i := range view {
		if len(buffer[startRow+i]) < startCol+cols {
			return nil, fmt.Errorf("row %d has less cols than asked: %d x %d", i, len(buffer[startRow+i]), startCol+cols)
		}
		view[i] = buffer[startRow+i][startCol : startCol+cols : startCol+cols]
	}
	return view, nil
}

// FromInt32 receives a buffer matrix and returns a view of it
// setting with the specified number of rows and columns
// The matrix buffer must have the dimensions of rows and columns at least equal
// to rows and cols, or else it will error with index out of range
func FromInt32(buffer [][]int32, startRow, startCol, rows, cols int) ([][]int32, error) {
	view := make([][]int32, rows)
	if len(buffer) < startRow+rows {
		return nil, fmt.Errorf("matrix has less rows than asked: %d x %d", len(buffer), startRow+rows)
	}
	for i := range view {
		if len(buffer[startRow+i]) < startCol+cols {
			return nil, fmt.Errorf("row %d has less cols than asked: %d x %d", i, len(buffer[startRow+i]), startCol+cols)
		}
		view[i] = buffer[startRow+i][startCol : startCol+cols : startCol+cols]
	}
	return view, nil
}

// FromInt64 receives a buffer matrix and returns a view of it
// setting with the specified number of rows and columns
// The matrix buffer must have the dimensions of rows and columns at least equal
// to rows and cols, or else it will error with index out of range
func FromInt64(buffer [][]int64, startRow, startCol, rows, cols int) ([][]int64, error) {
	view := make([][]int64, rows)
	if len(buffer) < startRow+rows {
		return nil, fmt.Errorf("matrix has less rows than asked: %d x %d", len(buffer), startRow+rows)
	}
	for i := range view {
		if len(buffer[startRow+i]) < startCol+cols {
			return nil, fmt.Errorf("row %d has less cols than asked: %d x %d", i, len(buffer[startRow+i]), startCol+cols)
		}
		view[i] = buffer[startRow+i][startCol : startCol+cols : startCol+cols]
	}
	return view, nil
}

// FromUint receives a buffer matrix and returns a view of it
// setting with the specified number of rows and columns
// The matrix buffer must have the dimensions of rows and columns at least equal
// to rows and cols, or else it will error with index out of range
func FromUint(buffer [][]uint, startRow, startCol, rows, cols int) ([][]uint, error) {
	view := make([][]uint, rows)
	if len(buffer) < startRow+rows {
		return nil, fmt.Errorf("matrix has less rows than asked: %d x %d", len(buffer), startRow+rows)
	}
	for i := range view {
		if len(buffer[startRow+i]) < startCol+cols {
			return nil, fmt.Errorf("row %d has less cols than asked: %d x %d", i, len(buffer[startRow+i]), startCol+cols)
		}
		view[i] = buffer[startRow+i][startCol : startCol+cols : startCol+cols]
	}
	return view, nil
}

// FromUint8 receives a buffer matrix and returns a view of it
// setting with the specified number of rows and columns
// The matrix buffer must have the dimensions of rows and columns at least equal
// to rows and cols, or else it will error with index out of range
func FromUint8(buffer [][]uint8, startRow, startCol, rows, cols int) ([][]uint8, error) {
	view := make([][]uint8, rows)
	if len(buffer) < startRow+rows {
		return nil, fmt.Errorf("matrix has less rows than asked: %d x %d", len(buffer), startRow+rows)
	}
	for i := range view {
		if len(buffer[startRow+i]) < startCol+cols {
			return nil, fmt.Errorf("row %d has less cols than asked: %d x %d", i, len(buffer[startRow+i]), startCol+cols)
		}
		view[i] = buffer[startRow+i][startCol : startCol+cols : startCol+cols]
	}
	return view, nil
}

// FromUint16 receives a buffer matrix and returns a view of it
// setting with the specified number of rows and columns
// The matrix buffer must have the dimensions of rows and columns at least equal
// to rows and cols, or else it will error with index out of range
func FromUint16(buffer [][]uint16, startRow, startCol, rows, cols int) ([][]uint16, error) {
	view := make([][]uint16, rows)
	if len(buffer) < startRow+rows {
		return nil, fmt.Errorf("matrix has less rows than asked: %d x %d", len(buffer), startRow+rows)
	}
	for i := range view {
		if len(buffer[startRow+i]) < startCol+cols {
			return nil, fmt.Errorf("row %d has less cols than asked: %d x %d", i, len(buffer[startRow+i]), startCol+cols)
		}
		view[i] = buffer[startRow+i][startCol : startCol+cols : startCol+cols]
	}
	return view, nil
}

// FromUint32 receives a buffer matrix and returns a view of it
// setting with the specified number of rows and columns
// The matrix buffer must have the dimensions of rows and columns at least equal
// to rows and cols, or else it will error with index out of range
func FromUint32(buffer [][]uint32, startRow, startCol, rows, cols int) ([][]uint32, error) {
	view := make([][]uint32, rows)
	if len(buffer) < startRow+rows {
		return nil, fmt.Errorf("matrix has less rows than asked: %d x %d", len(buffer), startRow+rows)
	}
	for i := range view {
		if len(buffer[startRow+i]) < startCol+cols {
			return nil, fmt.Errorf("row %d has less cols than asked: %d x %d", i, len(buffer[startRow+i]), startCol+cols)
		}
		view[i] = buffer[startRow+i][startCol : startCol+cols : startCol+cols]
	}
	return view, nil
}

// FromUint64 receives a buffer matrix and returns a view of it
// setting with the specified number of rows and columns
// The matrix buffer must have the dimensions of rows and columns at least equal
// to rows and cols, or else it will error with index out of range
func FromUint64(buffer [][]uint64, startRow, startCol, rows, cols int) ([][]uint64, error) {
	view := make([][]uint64, rows)
	if len(buffer) < startRow+rows {
		return nil, fmt.Errorf("matrix has less rows than asked: %d x %d", len(buffer), startRow+rows)
	}
	for i := range view {
		if len(buffer[startRow+i]) < startCol+cols {
			return nil, fmt.Errorf("row %d has less cols than asked: %d x %d", i, len(buffer[startRow+i]), startCol+cols)
		}
		view[i] = buffer[startRow+i][startCol : startCol+cols : startCol+cols]
	}
	return view, nil
}

// FromFloat32 receives a buffer matrix and returns a view of it
// setting with the specified number of rows and columns
// The matrix buffer must have the dimensions of rows and columns at least equal
// to rows and cols, or else it will error with index out of range
func FromFloat32(buffer [][]float32, startRow, startCol, rows, cols int) ([][]float32, error) {
	view := make([][]float32, rows)
	if len(buffer) < startRow+rows {
		return nil, fmt.Errorf("matrix has less rows than asked: %d x %d", len(buffer), startRow+rows)
	}
	for i := range view {
		if len(buffer[startRow+i]) < startCol+cols {
			return nil, fmt.Errorf("row %d has less cols than asked: %d x %d", i, len(buffer[startRow+i]), startCol+cols)
		}
		view[i] = buffer[startRow+i][startCol : startCol+cols : startCol+cols]
	}
	return view, nil
}

// FromFloat64 receives a buffer matrix and returns a view of it
// setting with the specified number of rows and columns
// The matrix buffer must have the dimensions of rows and columns at least equal
// to rows and cols, or else it will error with index out of range
func FromFloat64(buffer [][]float64, startRow, startCol, rows, cols int) ([][]float64, error) {
	view := make([][]float64, rows)
	if len(buffer) < startRow+rows {
		return nil, fmt.Errorf("matrix has less rows than asked: %d x %d", len(buffer), startRow+rows)
	}
	for i := range view {
		if len(buffer[startRow+i]) < startCol+cols {
			return nil, fmt.Errorf("row %d has less cols than asked: %d x %d", i, len(buffer[startRow+i]), startCol+cols)
		}
		view[i] = buffer[startRow+i][startCol : startCol+cols : startCol+cols]
	}
	return view, nil
}

// FromBool receives a buffer matrix and returns a view of it
// setting with the specified number of rows and columns
// The matrix buffer must have the dimensions of rows and columns at least equal
// to rows and cols, or else it will error with index out of range
func FromBool(buffer [][]bool, startRow, startCol, rows, cols int) ([][]bool, error) {
	view := make([][]bool, rows)
	if len(buffer) < startRow+rows {
		return nil, fmt.Errorf("matrix has less rows than asked: %d x %d", len(buffer), startRow+rows)
	}
	for i := range view {
		if len(buffer[startRow+i]) < startCol+cols {
			return nil, fmt.Errorf("row %d has less cols than asked: %d x %d", i, len(buffer[startRow+i]), startCol+cols)
		}
		view[i] = buffer[startRow+i][startCol : startCol+cols : startCol+cols]
	}
	return view, nil
}

// FromComplex64 receives a buffer matrix and returns a view of it
// setting with the specified number of rows and columns
// The matrix buffer must have the dimensions of rows and columns at least equal
// to rows and cols, or else it will error with index out of range
func FromComplex64(buffer [][]complex64, startRow, startCol, rows, cols int) ([][]complex64, error) {
	view := make([][]complex64, rows)
	if len(buffer) < startRow+rows {
		return nil, fmt.Errorf("matrix has less rows than asked: %d x %d", len(buffer), startRow+rows)
	}
	for i := range view {
		if len(buffer[startRow+i]) < startCol+cols {
			return nil, fmt.Errorf("row %d has less cols than asked: %d x %d", i, len(buffer[startRow+i]), startCol+cols)
		}
		view[i] = buffer[startRow+i][startCol : startCol+cols : startCol+cols]
	}
	return view, nil
}

// FromComplex128 receives a buffer matrix and returns a view of it
// setting with the specified number of rows and columns
// The matrix buffer must have the dimensions of rows and columns at least equal
// to rows and cols, or else it will error with index out of range
func FromComplex128(buffer [][]complex128, startRow, startCol, rows, cols int) ([][]complex128, error) {
	view := make([][]complex128, rows)
	if len(buffer) < startRow+rows {
		return nil, fmt.Errorf("matrix has less rows than asked: %d x %d", len(buffer), startRow+rows)
	}
	for i := range view {
		if len(buffer[startRow+i]) < startCol+cols {
			return nil, fmt.Errorf("row %d has less cols than asked: %d x %d", i, len(buffer[startRow+i]), startCol+cols)
		}
		view[i] = buffer[startRow+i][startCol : startCol+cols : startCol+cols]
	}
	return view, nil
}

// FromComplex128 receives a buffer matrix and returns a view of it
// setting with the specified number of rows and columns
// The matrix buffer must have the dimensions of rows and columns at least equal
// to rows and cols, or else it will error with index out of range
func FromString(buffer [][]string, startRow, startCol, rows, cols int) ([][]string, error) {
	view := make([][]string, rows)
	if len(buffer) < startRow+rows {
		return nil, fmt.Errorf("matrix has less rows than asked: %d x %d", len(buffer), startRow+rows)
	}
	for i := range view {
		if len(buffer[startRow+i]) < startCol+cols {
			return nil, fmt.Errorf("row %d has less cols than asked: %d x %d", i, len(buffer[startRow+i]), startCol+cols)
		}
		view[i] = buffer[startRow+i][startCol : startCol+cols : startCol+cols]
	}
	return view, nil
}

// FromRune receives a buffer matrix and returns a view of it
// setting with the specified number of rows and columns
// The matrix buffer must have the dimensions of rows and columns at least equal
// to rows and cols, or else it will error with index out of range
func FromRune(buffer [][]rune, startRow, startCol, rows, cols int) ([][]rune, error) {
	view := make([][]rune, rows)
	if len(buffer) < startRow+rows {
		return nil, fmt.Errorf("matrix has less rows than asked: %d x %d", len(buffer), startRow+rows)
	}
	for i := range view {
		if len(buffer[startRow+i]) < startCol+cols {
			return nil, fmt.Errorf("row %d has less cols than asked: %d x %d", i, len(buffer[startRow+i]), startCol+cols)
		}
		view[i] = buffer[startRow+i][startCol : startCol+cols : startCol+cols]
	}
	return view, nil
}

// FromBytePtr receives a buffer matrix and returns a view of it
// setting with the specified number of rows and columns
// The matrix buffer must have the dimensions of rows and columns at least equal
// to rows and cols, or else it will error with index out of range
func FromBytePtr(buffer [][]*byte, startRow, startCol, rows, cols int) ([][]*byte, error) {
	view := make([][]*byte, rows)
	if len(buffer) < startRow+rows {
		return nil, fmt.Errorf("matrix has less rows than asked: %d x %d", len(buffer), startRow+rows)
	}
	for i := range view {
		if len(buffer[startRow+i]) < startCol+cols {
			return nil, fmt.Errorf("row %d has less cols than asked: %d x %d", i, len(buffer[startRow+i]), startCol+cols)
		}
		view[i] = buffer[startRow+i][startCol : startCol+cols : startCol+cols]
	}
	return view, nil
}

// FromIntPtr receives a buffer matrix and returns a view of it
// setting with the specified number of rows and columns
// The matrix buffer must have the dimensions of rows and columns at least equal
// to rows and cols, or else it will error with index out of range
func FromIntPtr(buffer [][]*int, startRow, startCol, rows, cols int) ([][]*int, error) {
	view := make([][]*int, rows)
	if len(buffer) < startRow+rows {
		return nil, fmt.Errorf("matrix has less rows than asked: %d x %d", len(buffer), startRow+rows)
	}
	for i := range view {
		if len(buffer[startRow+i]) < startCol+cols {
			return nil, fmt.Errorf("row %d has less cols than asked: %d x %d", i, len(buffer[startRow+i]), startCol+cols)
		}
		view[i] = buffer[startRow+i][startCol : startCol+cols : startCol+cols]
	}
	return view, nil
}

// FromInt8Ptr receives a buffer matrix and returns a view of it
// setting with the specified number of rows and columns
// The matrix buffer must have the dimensions of rows and columns at least equal
// to rows and cols, or else it will error with index out of range
func FromInt8Ptr(buffer [][]*int8, startRow, startCol, rows, cols int) ([][]*int8, error) {
	view := make([][]*int8, rows)
	if len(buffer) < startRow+rows {
		return nil, fmt.Errorf("matrix has less rows than asked: %d x %d", len(buffer), startRow+rows)
	}
	for i := range view {
		if len(buffer[startRow+i]) < startCol+cols {
			return nil, fmt.Errorf("row %d has less cols than asked: %d x %d", i, len(buffer[startRow+i]), startCol+cols)
		}
		view[i] = buffer[startRow+i][startCol : startCol+cols : startCol+cols]
	}
	return view, nil
}

// FromInt16Ptr receives a buffer matrix and returns a view of it
// setting with the specified number of rows and columns
// The matrix buffer must have the dimensions of rows and columns at least equal
// to rows and cols, or else it will error with index out of range
func FromInt16Ptr(buffer [][]*int16, startRow, startCol, rows, cols int) ([][]*int16, error) {
	view := make([][]*int16, rows)
	if len(buffer) < startRow+rows {
		return nil, fmt.Errorf("matrix has less rows than asked: %d x %d", len(buffer), startRow+rows)
	}
	for i := range view {
		if len(buffer[startRow+i]) < startCol+cols {
			return nil, fmt.Errorf("row %d has less cols than asked: %d x %d", i, len(buffer[startRow+i]), startCol+cols)
		}
		view[i] = buffer[startRow+i][startCol : startCol+cols : startCol+cols]
	}
	return view, nil
}

// FromInt32Ptr receives a buffer matrix and returns a view of it
// setting with the specified number of rows and columns
// The matrix buffer must have the dimensions of rows and columns at least equal
// to rows and cols, or else it will error with index out of range
func FromInt32Ptr(buffer [][]*int32, startRow, startCol, rows, cols int) ([][]*int32, error) {
	view := make([][]*int32, rows)
	if len(buffer) < startRow+rows {
		return nil, fmt.Errorf("matrix has less rows than asked: %d x %d", len(buffer), startRow+rows)
	}
	for i := range view {
		if len(buffer[startRow+i]) < startCol+cols {
			return nil, fmt.Errorf("row %d has less cols than asked: %d x %d", i, len(buffer[startRow+i]), startCol+cols)
		}
		view[i] = buffer[startRow+i][startCol : startCol+cols : startCol+cols]
	}
	return view, nil
}

// FromInt64Ptr receives a buffer matrix and returns a view of it
// setting with the specified number of rows and columns
// The matrix buffer must have the dimensions of rows and columns at least equal
// to rows and cols, or else it will error with index out of range
func FromInt64Ptr(buffer [][]*int64, startRow, startCol, rows, cols int) ([][]*int64, error) {
	view := make([][]*int64, rows)
	if len(buffer) < startRow+rows {
		return nil, fmt.Errorf("matrix has less rows than asked: %d x %d", len(buffer), startRow+rows)
	}
	for i := range view {
		if len(buffer[startRow+i]) < startCol+cols {
			return nil, fmt.Errorf("row %d has less cols than asked: %d x %d", i, len(buffer[startRow+i]), startCol+cols)
		}
		view[i] = buffer[startRow+i][startCol : startCol+cols : startCol+cols]
	}
	return view, nil
}

// FromUintPtr receives a buffer matrix and returns a view of it
// setting with the specified number of rows and columns
// The matrix buffer must have the dimensions of rows and columns at least equal
// to rows and cols, or else it will error with index out of range
func FromUintPtr(buffer [][]*uint, startRow, startCol, rows, cols int) ([][]*uint, error) {
	view := make([][]*uint, rows)
	if len(buffer) < startRow+rows {
		return nil, fmt.Errorf("matrix has less rows than asked: %d x %d", len(buffer), startRow+rows)
	}
	for i := range view {
		if len(buffer[startRow+i]) < startCol+cols {
			return nil, fmt.Errorf("row %d has less cols than asked: %d x %d", i, len(buffer[startRow+i]), startCol+cols)
		}
		view[i] = buffer[startRow+i][startCol : startCol+cols : startCol+cols]
	}
	return view, nil
}

// FromUint8Ptr receives a buffer matrix and returns a view of it
// setting with the specified number of rows and columns
// The matrix buffer must have the dimensions of rows and columns at least equal
// to rows and cols, or else it will error with index out of range
func FromUint8Ptr(buffer [][]*uint8, startRow, startCol, rows, cols int) ([][]*uint8, error) {
	view := make([][]*uint8, rows)
	if len(buffer) < startRow+rows {
		return nil, fmt.Errorf("matrix has less rows than asked: %d x %d", len(buffer), startRow+rows)
	}
	for i := range view {
		if len(buffer[startRow+i]) < startCol+cols {
			return nil, fmt.Errorf("row %d has less cols than asked: %d x %d", i, len(buffer[startRow+i]), startCol+cols)
		}
		view[i] = buffer[startRow+i][startCol : startCol+cols : startCol+cols]
	}
	return view, nil
}

// FromUint16Ptr receives a buffer matrix and returns a view of it
// setting with the specified number of rows and columns
// The matrix buffer must have the dimensions of rows and columns at least equal
// to rows and cols, or else it will error with index out of range
func FromUint16Ptr(buffer [][]*uint16, startRow, startCol, rows, cols int) ([][]*uint16, error) {
	view := make([][]*uint16, rows)
	if len(buffer) < startRow+rows {
		return nil, fmt.Errorf("matrix has less rows than asked: %d x %d", len(buffer), startRow+rows)
	}
	for i := range view {
		if len(buffer[startRow+i]) < startCol+cols {
			return nil, fmt.Errorf("row %d has less cols than asked: %d x %d", i, len(buffer[startRow+i]), startCol+cols)
		}
		view[i] = buffer[startRow+i][startCol : startCol+cols : startCol+cols]
	}
	return view, nil
}

// FromUint32Ptr receives a buffer matrix and returns a view of it
// setting with the specified number of rows and columns
// The matrix buffer must have the dimensions of rows and columns at least equal
// to rows and cols, or else it will error with index out of range
func FromUint32Ptr(buffer [][]*uint32, startRow, startCol, rows, cols int) ([][]*uint32, error) {
	view := make([][]*uint32, rows)
	if len(buffer) < startRow+rows {
		return nil, fmt.Errorf("matrix has less rows than asked: %d x %d", len(buffer), startRow+rows)
	}
	for i := range view {
		if len(buffer[startRow+i]) < startCol+cols {
			return nil, fmt.Errorf("row %d has less cols than asked: %d x %d", i, len(buffer[startRow+i]), startCol+cols)
		}
		view[i] = buffer[startRow+i][startCol : startCol+cols : startCol+cols]
	}
	return view, nil
}

// FromUint64Ptr receives a buffer matrix and returns a view of it
// setting with the specified number of rows and columns
// The matrix buffer must have the dimensions of rows and columns at least equal
// to rows and cols, or else it will error with index out of range
func FromUint64Ptr(buffer [][]*uint64, startRow, startCol, rows, cols int) ([][]*uint64, error) {
	view := make([][]*uint64, rows)
	if len(buffer) < startRow+rows {
		return nil, fmt.Errorf("matrix has less rows than asked: %d x %d", len(buffer), startRow+rows)
	}
	for i := range view {
		if len(buffer[startRow+i]) < startCol+cols {
			return nil, fmt.Errorf("row %d has less cols than asked: %d x %d", i, len(buffer[startRow+i]), startCol+cols)
		}
		view[i] = buffer[startRow+i][startCol : startCol+cols : startCol+cols]
	}
	return view, nil
}

// FromFloat32Ptr receives a buffer matrix and returns a view of it
// setting with the specified number of rows and columns
// The matrix buffer must have the dimensions of rows and columns at least equal
// to rows and cols, or else it will error with index out of range
func FromFloat32Ptr(buffer [][]*float32, startRow, startCol, rows, cols int) ([][]*float32, error) {
	view := make([][]*float32, rows)
	if len(buffer) < startRow+rows {
		return nil, fmt.Errorf("matrix has less rows than asked: %d x %d", len(buffer), startRow+rows)
	}
	for i := range view {
		if len(buffer[startRow+i]) < startCol+cols {
			return nil, fmt.Errorf("row %d has less cols than asked: %d x %d", i, len(buffer[startRow+i]), startCol+cols)
		}
		view[i] = buffer[startRow+i][startCol : startCol+cols : startCol+cols]
	}
	return view, nil
}

// FromFloat64Ptr receives a buffer matrix and returns a view of it
// setting with the specified number of rows and columns
// The matrix buffer must have the dimensions of rows and columns at least equal
// to rows and cols, or else it will error with index out of range
func FromFloat64Ptr(buffer [][]*float64, startRow, startCol, rows, cols int) ([][]*float64, error) {
	view := make([][]*float64, rows)
	if len(buffer) < startRow+rows {
		return nil, fmt.Errorf("matrix has less rows than asked: %d x %d", len(buffer), startRow+rows)
	}
	for i := range view {
		if len(buffer[startRow+i]) < startCol+cols {
			return nil, fmt.Errorf("row %d has less cols than asked: %d x %d", i, len(buffer[startRow+i]), startCol+cols)
		}
		view[i] = buffer[startRow+i][startCol : startCol+cols : startCol+cols]
	}
	return view, nil
}

// FromBoolPtr receives a buffer matrix and returns a view of it
// setting with the specified number of rows and columns
// The matrix buffer must have the dimensions of rows and columns at least equal
// to rows and cols, or else it will error with index out of range
func FromBoolPtr(buffer [][]*bool, startRow, startCol, rows, cols int) ([][]*bool, error) {
	view := make([][]*bool, rows)
	if len(buffer) < startRow+rows {
		return nil, fmt.Errorf("matrix has less rows than asked: %d x %d", len(buffer), startRow+rows)
	}
	for i := range view {
		if len(buffer[startRow+i]) < startCol+cols {
			return nil, fmt.Errorf("row %d has less cols than asked: %d x %d", i, len(buffer[startRow+i]), startCol+cols)
		}
		view[i] = buffer[startRow+i][startCol : startCol+cols : startCol+cols]
	}
	return view, nil
}

// FromComplex64Ptr receives a buffer matrix and returns a view of it
// setting with the specified number of rows and columns
// The matrix buffer must have the dimensions of rows and columns at least equal
// to rows and cols, or else it will error with index out of range
func FromComplex64Ptr(buffer [][]*complex64, startRow, startCol, rows, cols int) ([][]*complex64, error) {
	view := make([][]*complex64, rows)
	if len(buffer) < startRow+rows {
		return nil, fmt.Errorf("matrix has less rows than asked: %d x %d", len(buffer), startRow+rows)
	}
	for i := range view {
		if len(buffer[startRow+i]) < startCol+cols {
			return nil, fmt.Errorf("row %d has less cols than asked: %d x %d", i, len(buffer[startRow+i]), startCol+cols)
		}
		view[i] = buffer[startRow+i][startCol : startCol+cols : startCol+cols]
	}
	return view, nil
}

// FromComplex128Ptr receives a buffer matrix and returns a view of it
// setting with the specified number of rows and columns
// The matrix buffer must have the dimensions of rows and columns at least equal
// to rows and cols, or else it will error with index out of range
func FromComplex128Ptr(buffer [][]*complex128, startRow, startCol, rows, cols int) ([][]*complex128, error) {
	view := make([][]*complex128, rows)
	if len(buffer) < startRow+rows {
		return nil, fmt.Errorf("matrix has less rows than asked: %d x %d", len(buffer), startRow+rows)
	}
	for i := range view {
		if len(buffer[startRow+i]) < startCol+cols {
			return nil, fmt.Errorf("row %d has less cols than asked: %d x %d", i, len(buffer[startRow+i]), startCol+cols)
		}
		view[i] = buffer[startRow+i][startCol : startCol+cols : startCol+cols]
	}
	return view, nil
}

// FromComplex128Ptr receives a buffer matrix and returns a view of it
// setting with the specified number of rows and columns
// The matrix buffer must have the dimensions of rows and columns at least equal
// to rows and cols, or else it will error with index out of range
func FromStringPtr(buffer [][]*string, startRow, startCol, rows, cols int) ([][]*string, error) {
	view := make([][]*string, rows)
	if len(buffer) < startRow+rows {
		return nil, fmt.Errorf("matrix has less rows than asked: %d x %d", len(buffer), startRow+rows)
	}
	for i := range view {
		if len(buffer[startRow+i]) < startCol+cols {
			return nil, fmt.Errorf("row %d has less cols than asked: %d x %d", i, len(buffer[startRow+i]), startCol+cols)
		}
		view[i] = buffer[startRow+i][startCol : startCol+cols : startCol+cols]
	}
	return view, nil
}

// FromRunePtr receives a buffer matrix and returns a view of it
// setting with the specified number of rows and columns
// The matrix buffer must have the dimensions of rows and columns at least equal
// to rows and cols, or else it will error with index out of range
func FromRunePtr(buffer [][]*rune, startRow, startCol, rows, cols int) ([][]*rune, error) {
	view := make([][]*rune, rows)
	if len(buffer) < startRow+rows {
		return nil, fmt.Errorf("matrix has less rows than asked: %d x %d", len(buffer), startRow+rows)
	}
	for i := range view {
		if len(buffer[startRow+i]) < startCol+cols {
			return nil, fmt.Errorf("row %d has less cols than asked: %d x %d", i, len(buffer[startRow+i]), startCol+cols)
		}
		view[i] = buffer[startRow+i][startCol : startCol+cols : startCol+cols]
	}
	return view, nil
}
