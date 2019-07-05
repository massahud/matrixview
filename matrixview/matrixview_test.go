// Copyright 2009 Geraldo Augusto Massahud Rodrigues dos Santos. All rights reserved.
// This code is licensed under APACHE-2 and MIT, the respective licenses can be found
// at LICENSE-APACHE and LICENSE-MIT files.

package matrixview

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFromInterface(t *testing.T) {

	type args struct {
		buffer   [][]interface{}
		startRow int
		startCol int
		rows     int
		cols     int
	}
	tests := []struct {
		name string
		args args
		want [][]interface{}
	}{
		{"Should create a view with less rows starting at 0,0",
			args{[][]interface{}{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			}, 0, 0, 2, 3},
			[][]interface{}{
				{1, 2, 3},
				{4, 5, 6},
			},
		},
		{"Should create a view with less columns starting at 0,0",
			args{[][]interface{}{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			}, 0, 0, 3, 2},
			[][]interface{}{
				{1, 2},
				{4, 5},
				{7, 8},
			},
		},
		{
			"Should create a view with less rows and columns starting at 0,0",
			args{[][]interface{}{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			}, 0, 0, 2, 2},
			[][]interface{}{
				{1, 2},
				{4, 5},
			},
		},
		{
			"Should create a view with less rows and columns starting at 1,2",
			args{[][]interface{}{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			}, 1, 2, 2, 1},
			[][]interface{}{
				{6},
				{9},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FromInterface(tt.args.buffer, tt.args.startRow, tt.args.startCol, tt.args.rows, tt.args.cols); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromInterface() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFromByte(t *testing.T) {
	type args struct {
		buffer   [][]byte
		startRow int
		startCol int
		rows     int
		cols     int
	}
	tests := []struct {
		name string
		args args
		want [][]byte
	}{
		{"Should create a view with less rows starting at 0,0",
			args{[][]byte{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			}, 0, 0, 2, 3},
			[][]byte{
				{1, 2, 3},
				{4, 5, 6},
			},
		},
		{"Should create a view with less columns starting at 0,0",
			args{[][]byte{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			}, 0, 0, 3, 2},
			[][]byte{
				{1, 2},
				{4, 5},
				{7, 8},
			},
		},
		{
			"Should create a view with less rows and columns starting at 0,0",
			args{[][]byte{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			}, 0, 0, 2, 2},
			[][]byte{
				{1, 2},
				{4, 5},
			},
		},
		{
			"Should create a view with less rows and columns starting at 1,2",
			args{[][]byte{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			}, 1, 2, 2, 1},
			[][]byte{
				{6},
				{9},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FromByte(tt.args.buffer, tt.args.startRow, tt.args.startCol, tt.args.rows, tt.args.cols); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromByte() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFromInt(t *testing.T) {
	type args struct {
		buffer   [][]int
		startRow int
		startCol int
		rows     int
		cols     int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"Should create a view with less rows starting at 0,0",
			args{[][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			}, 0, 0, 2, 3},
			[][]int{
				{1, 2, 3},
				{4, 5, 6},
			},
		},
		{"Should create a view with less columns starting at 0,0",
			args{[][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			}, 0, 0, 3, 2},
			[][]int{
				{1, 2},
				{4, 5},
				{7, 8},
			},
		},
		{"Should create a view with less rows and columns starting at 0,0",
			args{[][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			}, 0, 0, 2, 2},
			[][]int{
				{1, 2},
				{4, 5},
			},
		},
		{
			"Should create a view with less rows and columns starting at 1,2",
			args{[][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			}, 1, 2, 2, 1},
			[][]int{
				{6},
				{9},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FromInt(tt.args.buffer, tt.args.startRow, tt.args.startCol, tt.args.rows, tt.args.cols); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFromInt8(t *testing.T) {
	type args struct {
		buffer   [][]int8
		startRow int
		startCol int
		rows     int
		cols     int
	}
	tests := []struct {
		name string
		args args
		want [][]int8
	}{
		{"Should create a view with less rows starting at 0,0",
			args{[][]int8{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			}, 0, 0, 2, 3},
			[][]int8{
				{1, 2, 3},
				{4, 5, 6},
			},
		},
		{"Should create a view with less columns starting at 0,0",
			args{[][]int8{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			}, 0, 0, 3, 2},
			[][]int8{
				{1, 2},
				{4, 5},
				{7, 8},
			},
		},
		{"Should create a view with less rows and columns starting at 0,0",
			args{[][]int8{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			}, 0, 0, 2, 2},
			[][]int8{
				{1, 2},
				{4, 5},
			},
		},
		{
			"Should create a view with less rows and columns starting at 1,2",
			args{[][]int8{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			}, 1, 2, 2, 1},
			[][]int8{
				{6},
				{9},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FromInt8(tt.args.buffer, tt.args.startRow, tt.args.startCol, tt.args.rows, tt.args.cols); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromInt8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFromInt16(t *testing.T) {
	type args struct {
		buffer   [][]int16
		startRow int
		startCol int
		rows     int
		cols     int
	}
	tests := []struct {
		name string
		args args
		want [][]int16
	}{
		{"Should create a view with less rows starting at 0,0",
			args{[][]int16{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			}, 0, 0, 2, 3},
			[][]int16{
				{1, 2, 3},
				{4, 5, 6},
			},
		},
		{"Should create a view with less columns starting at 0,0",
			args{[][]int16{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			}, 0, 0, 3, 2},
			[][]int16{
				{1, 2},
				{4, 5},
				{7, 8},
			},
		},
		{"Should create a view with less rows and columns starting at 0,0",
			args{[][]int16{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			}, 0, 0, 2, 2},
			[][]int16{
				{1, 2},
				{4, 5},
			},
		},
		{
			"Should create a view with less rows and columns starting at 1,2",
			args{[][]int16{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			}, 1, 2, 2, 1},
			[][]int16{
				{6},
				{9},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FromInt16(tt.args.buffer, tt.args.startRow, tt.args.startCol, tt.args.rows, tt.args.cols); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromInt16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFromInt32(t *testing.T) {
	type args struct {
		buffer   [][]int32
		startRow int
		startCol int
		rows     int
		cols     int
	}
	tests := []struct {
		name string
		args args
		want [][]int32
	}{
		{"Should create a view with less rows starting at 0,0",
			args{[][]int32{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			}, 0, 0, 2, 3},
			[][]int32{
				{1, 2, 3},
				{4, 5, 6},
			},
		},
		{"Should create a view with less columns starting at 0,0",
			args{[][]int32{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			}, 0, 0, 3, 2},
			[][]int32{
				{1, 2},
				{4, 5},
				{7, 8},
			},
		},
		{"Should create a view with less rows and columns starting at 0,0",
			args{[][]int32{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			}, 0, 0, 2, 2},
			[][]int32{
				{1, 2},
				{4, 5},
			},
		},
		{
			"Should create a view with less rows and columns starting at 1,2",
			args{[][]int32{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			}, 1, 2, 2, 1},
			[][]int32{
				{6},
				{9},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FromInt32(tt.args.buffer, tt.args.startRow, tt.args.startCol, tt.args.rows, tt.args.cols); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromInt32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFromInt64(t *testing.T) {
	type args struct {
		buffer   [][]int64
		startRow int
		startCol int
		rows     int
		cols     int
	}
	tests := []struct {
		name string
		args args
		want [][]int64
	}{
		{"Should create a view with less rows starting at 0,0",
			args{[][]int64{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			}, 0, 0, 2, 3},
			[][]int64{
				{1, 2, 3},
				{4, 5, 6},
			},
		},
		{"Should create a view with less columns starting at 0,0",
			args{[][]int64{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			}, 0, 0, 3, 2},
			[][]int64{
				{1, 2},
				{4, 5},
				{7, 8},
			},
		},
		{"Should create a view with less rows and columns starting at 0,0",
			args{[][]int64{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			}, 0, 0, 2, 2},
			[][]int64{
				{1, 2},
				{4, 5},
			},
		},
		{
			"Should create a view with less rows and columns starting at 1,2",
			args{[][]int64{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			}, 1, 2, 2, 1},
			[][]int64{
				{6},
				{9},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FromInt64(tt.args.buffer, tt.args.startRow, tt.args.startCol, tt.args.rows, tt.args.cols); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFromUint(t *testing.T) {
	type args struct {
		buffer   [][]uint
		startRow int
		startCol int
		rows     int
		cols     int
	}
	tests := []struct {
		name string
		args args
		want [][]uint
	}{
		{"Should create a view with less rows starting at 0,0",
			args{[][]uint{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			}, 0, 0, 2, 3},
			[][]uint{
				{1, 2, 3},
				{4, 5, 6},
			},
		},
		{"Should create a view with less columns starting at 0,0",
			args{[][]uint{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			}, 0, 0, 3, 2},
			[][]uint{
				{1, 2},
				{4, 5},
				{7, 8},
			},
		},
		{"Should create a view with less rows and columns starting at 0,0",
			args{[][]uint{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			}, 0, 0, 2, 2},
			[][]uint{
				{1, 2},
				{4, 5},
			},
		},
		{
			"Should create a view with less rows and columns starting at 1,2",
			args{[][]uint{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			}, 1, 2, 2, 1},
			[][]uint{
				{6},
				{9},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FromUint(tt.args.buffer, tt.args.startRow, tt.args.startCol, tt.args.rows, tt.args.cols); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromUint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFromUint8(t *testing.T) {
	type args struct {
		buffer   [][]uint8
		startRow int
		startCol int
		rows     int
		cols     int
	}
	tests := []struct {
		name string
		args args
		want [][]uint8
	}{
		{"Should create a view with less rows starting at 0,0",
			args{[][]uint8{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			}, 0, 0, 2, 3},
			[][]uint8{
				{1, 2, 3},
				{4, 5, 6},
			},
		},
		{"Should create a view with less columns starting at 0,0",
			args{[][]uint8{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			}, 0, 0, 3, 2},
			[][]uint8{
				{1, 2},
				{4, 5},
				{7, 8},
			},
		},
		{"Should create a view with less rows and columns starting at 0,0",
			args{[][]uint8{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			}, 0, 0, 2, 2},
			[][]uint8{
				{1, 2},
				{4, 5},
			},
		},
		{
			"Should create a view with less rows and columns starting at 1,2",
			args{[][]uint8{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			}, 1, 2, 2, 1},
			[][]uint8{
				{6},
				{9},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FromUint8(tt.args.buffer, tt.args.startRow, tt.args.startCol, tt.args.rows, tt.args.cols); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromUint8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFromUint16(t *testing.T) {
	type args struct {
		buffer   [][]uint16
		startRow int
		startCol int
		rows     int
		cols     int
	}
	tests := []struct {
		name string
		args args
		want [][]uint16
	}{
		{"Should create a view with less rows starting at 0,0",
			args{[][]uint16{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			}, 0, 0, 2, 3},
			[][]uint16{
				{1, 2, 3},
				{4, 5, 6},
			},
		},
		{"Should create a view with less columns starting at 0,0",
			args{[][]uint16{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			}, 0, 0, 3, 2},
			[][]uint16{
				{1, 2},
				{4, 5},
				{7, 8},
			},
		},
		{"Should create a view with less rows and columns starting at 0,0",
			args{[][]uint16{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			}, 0, 0, 2, 2},
			[][]uint16{
				{1, 2},
				{4, 5},
			},
		},
		{
			"Should create a view with less rows and columns starting at 1,2",
			args{[][]uint16{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			}, 1, 2, 2, 1},
			[][]uint16{
				{6},
				{9},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FromUint16(tt.args.buffer, tt.args.startRow, tt.args.startCol, tt.args.rows, tt.args.cols); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromUint16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFromUint32(t *testing.T) {
	type args struct {
		buffer   [][]uint32
		startRow int
		startCol int
		rows     int
		cols     int
	}
	tests := []struct {
		name string
		args args
		want [][]uint32
	}{
		{"Should create a view with less rows starting at 0,0",
			args{[][]uint32{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			}, 0, 0, 2, 3},
			[][]uint32{
				{1, 2, 3},
				{4, 5, 6},
			},
		},
		{"Should create a view with less columns starting at 0,0",
			args{[][]uint32{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			}, 0, 0, 3, 2},
			[][]uint32{
				{1, 2},
				{4, 5},
				{7, 8},
			},
		},
		{"Should create a view with less rows and columns starting at 0,0",
			args{[][]uint32{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			}, 0, 0, 2, 2},
			[][]uint32{
				{1, 2},
				{4, 5},
			},
		},
		{
			"Should create a view with less rows and columns starting at 1,2",
			args{[][]uint32{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			}, 1, 2, 2, 1},
			[][]uint32{
				{6},
				{9},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FromUint32(tt.args.buffer, tt.args.startRow, tt.args.startCol, tt.args.rows, tt.args.cols); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromUint32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFromUint64(t *testing.T) {
	type args struct {
		buffer   [][]uint64
		startRow int
		startCol int
		rows     int
		cols     int
	}
	tests := []struct {
		name string
		args args
		want [][]uint64
	}{
		{"Should create a view with less rows starting at 0,0",
			args{[][]uint64{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			}, 0, 0, 2, 3},
			[][]uint64{
				{1, 2, 3},
				{4, 5, 6},
			},
		},
		{"Should create a view with less columns starting at 0,0",
			args{[][]uint64{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			}, 0, 0, 3, 2},
			[][]uint64{
				{1, 2},
				{4, 5},
				{7, 8},
			},
		},
		{"Should create a view with less rows and columns starting at 0,0",
			args{[][]uint64{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			}, 0, 0, 2, 2},
			[][]uint64{
				{1, 2},
				{4, 5},
			},
		},
		{
			"Should create a view with less rows and columns starting at 1,2",
			args{[][]uint64{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			}, 1, 2, 2, 1},
			[][]uint64{
				{6},
				{9},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FromUint64(tt.args.buffer, tt.args.startRow, tt.args.startCol, tt.args.rows, tt.args.cols); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromUint64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFromFloat32(t *testing.T) {
	type args struct {
		buffer   [][]float32
		startRow int
		startCol int
		rows     int
		cols     int
	}
	tests := []struct {
		name string
		args args
		want [][]float32
	}{
		{"Should create a view with less rows starting at 0,0",
			args{[][]float32{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			}, 0, 0, 2, 3},
			[][]float32{
				{1, 2, 3},
				{4, 5, 6},
			},
		},
		{"Should create a view with less columns starting at 0,0",
			args{[][]float32{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			}, 0, 0, 3, 2},
			[][]float32{
				{1, 2},
				{4, 5},
				{7, 8},
			},
		},
		{"Should create a view with less rows and columns starting at 0,0",
			args{[][]float32{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			}, 0, 0, 2, 2},
			[][]float32{
				{1, 2},
				{4, 5},
			},
		},
		{
			"Should create a view with less rows and columns starting at 1,2",
			args{[][]float32{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			}, 1, 2, 2, 1},
			[][]float32{
				{6},
				{9},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FromFloat32(tt.args.buffer, tt.args.startRow, tt.args.startCol, tt.args.rows, tt.args.cols); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromFloat32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFromFloat64(t *testing.T) {
	type args struct {
		buffer   [][]float64
		startRow int
		startCol int
		rows     int
		cols     int
	}
	tests := []struct {
		name string
		args args
		want [][]float64
	}{
		{"Should create a view with less rows starting at 0,0",
			args{[][]float64{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			}, 0, 0, 2, 3},
			[][]float64{
				{1, 2, 3},
				{4, 5, 6},
			},
		},
		{"Should create a view with less columns starting at 0,0",
			args{[][]float64{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			}, 0, 0, 3, 2},
			[][]float64{
				{1, 2},
				{4, 5},
				{7, 8},
			},
		},
		{"Should create a view with less rows and columns starting at 0,0",
			args{[][]float64{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			}, 0, 0, 2, 2},
			[][]float64{
				{1, 2},
				{4, 5},
			},
		},
		{
			"Should create a view with less rows and columns starting at 1,2",
			args{[][]float64{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			}, 1, 2, 2, 1},
			[][]float64{
				{6},
				{9},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FromFloat64(tt.args.buffer, tt.args.startRow, tt.args.startCol, tt.args.rows, tt.args.cols); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromFloat64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFromBool(t *testing.T) {
	type args struct {
		buffer   [][]bool
		startRow int
		startCol int
		rows     int
		cols     int
	}
	tests := []struct {
		name string
		args args
		want [][]bool
	}{
		{"Should create a view with less rows starting at 0,0",
			args{[][]bool{
				{true, false, true},
				{true, true, false},
				{false, false, false},
			}, 0, 0, 2, 3},
			[][]bool{
				{true, false, true},
				{true, true, false},
			},
		},
		{"Should create a view with less columns starting at 0,0",
			args{[][]bool{
				{true, false, true},
				{true, true, false},
				{false, false, false},
			}, 0, 0, 3, 2},
			[][]bool{
				{true, false},
				{true, true},
				{false, false},
			},
		},
		{"Should create a view with less rows and columns starting at 0,0",
			args{[][]bool{
				{true, false, true},
				{true, true, false},
				{false, false, false},
			}, 0, 0, 2, 2},
			[][]bool{
				{true, false},
				{true, true},
			},
		},
		{
			"Should create a view with less rows and columns starting at 1,2",
			args{[][]bool{
				{true, false, true},
				{true, true, false},
				{false, false, false},
			}, 1, 2, 2, 1},
			[][]bool{
				{false},
				{false},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FromBool(tt.args.buffer, tt.args.startRow, tt.args.startCol, tt.args.rows, tt.args.cols); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromBool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFromComplex64(t *testing.T) {
	type args struct {
		buffer   [][]complex64
		startRow int
		startCol int
		rows     int
		cols     int
	}
	tests := []struct {
		name string
		args args
		want [][]complex64
	}{
		{"Should create a view with less rows starting at 0,0",
			args{[][]complex64{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			}, 0, 0, 2, 3},
			[][]complex64{
				{1, 2, 3},
				{4, 5, 6},
			},
		},
		{"Should create a view with less columns starting at 0,0",
			args{[][]complex64{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			}, 0, 0, 3, 2},
			[][]complex64{
				{1, 2},
				{4, 5},
				{7, 8},
			},
		},
		{"Should create a view with less rows and columns starting at 0,0",
			args{[][]complex64{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			}, 0, 0, 2, 2},
			[][]complex64{
				{1, 2},
				{4, 5},
			},
		},
		{
			"Should create a view with less rows and columns starting at 1,2",
			args{[][]complex64{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			}, 1, 2, 2, 1},
			[][]complex64{
				{6},
				{9},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FromComplex64(tt.args.buffer, tt.args.startRow, tt.args.startCol, tt.args.rows, tt.args.cols); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromComplex64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFromComplex128(t *testing.T) {
	type args struct {
		buffer   [][]complex128
		startRow int
		startCol int
		rows     int
		cols     int
	}
	tests := []struct {
		name string
		args args
		want [][]complex128
	}{
		{"Should create a view with less rows starting at 0,0",
			args{[][]complex128{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			}, 0, 0, 2, 3},
			[][]complex128{
				{1, 2, 3},
				{4, 5, 6},
			},
		},
		{"Should create a view with less columns starting at 0,0",
			args{[][]complex128{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			}, 0, 0, 3, 2},
			[][]complex128{
				{1, 2},
				{4, 5},
				{7, 8},
			},
		},
		{"Should create a view with less rows and columns starting at 0,0",
			args{[][]complex128{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			}, 0, 0, 2, 2},
			[][]complex128{
				{1, 2},
				{4, 5},
			},
		},
		{
			"Should create a view with less rows and columns starting at 1,2",
			args{[][]complex128{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			}, 1, 2, 2, 1},
			[][]complex128{
				{6},
				{9},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FromComplex128(tt.args.buffer, tt.args.startRow, tt.args.startCol, tt.args.rows, tt.args.cols); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromComplex128() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFromString(t *testing.T) {
	type args struct {
		buffer   [][]string
		startRow int
		startCol int
		rows     int
		cols     int
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{"Should create a view with less rows starting at 0,0",
			args{[][]string{
				{"1", "2", "3"},
				{"4", "5", "6"},
				{"7", "8", "9"},
			}, 0, 0, 2, 3},
			[][]string{
				{"1", "2", "3"},
				{"4", "5", "6"},
			},
		},
		{"Should create a view with less columns starting at 0,0",
			args{[][]string{
				{"1", "2", "3"},
				{"4", "5", "6"},
				{"7", "8", "9"},
			}, 0, 0, 3, 2},
			[][]string{
				{"1", "2"},
				{"4", "5"},
				{"7", "8"},
			},
		},
		{"Should create a view with less rows and columns starting at 0,0",
			args{[][]string{
				{"1", "2", "3"},
				{"4", "5", "6"},
				{"7", "8", "9"},
			}, 0, 0, 2, 2},
			[][]string{
				{"1", "2"},
				{"4", "5"},
			},
		},
		{
			"Should create a view with less rows and columns starting at 1,2",
			args{[][]string{
				{"1", "2", "3"},
				{"4", "5", "6"},
				{"7", "8", "9"},
			}, 1, 2, 2, 1},
			[][]string{
				{"6"},
				{"9"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FromString(tt.args.buffer, tt.args.startRow, tt.args.startCol, tt.args.rows, tt.args.cols); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFromRune(t *testing.T) {
	type args struct {
		buffer   [][]rune
		startRow int
		startCol int
		rows     int
		cols     int
	}
	tests := []struct {
		name string
		args args
		want [][]rune
	}{
		{"Should create a view with less rows starting at 0,0",
			args{[][]rune{
				{'1', '2', '3'},
				{'4', '5', '6'},
				{'7', '8', '9'},
			}, 0, 0, 2, 3},
			[][]rune{
				{'1', '2', '3'},
				{'4', '5', '6'},
			},
		},
		{"Should create a view with less columns starting at 0,0",
			args{[][]rune{
				{'1', '2', '3'},
				{'4', '5', '6'},
				{'7', '8', '9'},
			}, 0, 0, 3, 2},
			[][]rune{
				{'1', '2'},
				{'4', '5'},
				{'7', '8'},
			},
		},
		{"Should create a view with less rows and columns starting at 0,0",
			args{[][]rune{
				{'1', '2', '3'},
				{'4', '5', '6'},
				{'7', '8', '9'},
			}, 0, 0, 2, 2},
			[][]rune{
				{'1', '2'},
				{'4', '5'},
			},
		},
		{
			"Should create a view with less rows and columns starting at 1,2",
			args{[][]rune{
				{'1', '2', '3'},
				{'4', '5', '6'},
				{'7', '8', '9'},
			}, 1, 2, 2, 1},
			[][]rune{
				{'6'},
				{'9'},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FromRune(tt.args.buffer, tt.args.startRow, tt.args.startCol, tt.args.rows, tt.args.cols); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromRune() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFromBytePtr(t *testing.T) {

	var (
		one   byte = 1
		two   byte = 2
		three byte = 3
		four  byte = 4
		five  byte = 5
		six   byte = 6
		seven byte = 7
		eight byte = 8
		nine  byte = 9
	)
	type args struct {
		buffer   [][]*byte
		startRow int
		startCol int
		rows     int
		cols     int
	}
	tests := []struct {
		name string
		args args
		want [][]*byte
	}{
		{"Should create a view with less rows starting at 0,0",
			args{[][]*byte{
				{&one, &two, &three},
				{&four, &five, &six},
				{&seven, &eight, &nine},
			},
				0, 0, 2, 3},
			[][]*byte{
				{&one, &two, &three},
				{&four, &five, &six},
			},
		},
		{"Should create a view with less columns starting at 0,0",
			args{[][]*byte{
				{&one, &two, &three},
				{&four, &five, &six},
				{&seven, &eight, &nine},
			},
				0, 0, 3, 2},
			[][]*byte{
				{&one, &two},
				{&four, &five},
				{&seven, &eight},
			},
		},
		{
			"Should create a view with less rows and columns starting at 0,0",
			args{[][]*byte{
				{&one, &two, &three},
				{&four, &five, &six},
				{&seven, &eight, &nine},
			},
				0, 0, 2, 2},
			[][]*byte{
				{&one, &two},
				{&four, &five},
			},
		},
		{
			"Should create a view with less rows and columns starting at 1,2",
			args{[][]*byte{
				{&one, &two, &three},
				{&four, &five, &six},
				{&seven, &eight, &nine},
			},
				1, 2, 2, 1},
			[][]*byte{
				{&six},
				{&nine},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FromBytePtr(tt.args.buffer, tt.args.startRow, tt.args.startCol, tt.args.rows, tt.args.cols); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromBytePtr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFromIntPtr(t *testing.T) {
	var (
		one   = 1
		two   = 2
		three = 3
		four  = 4
		five  = 5
		six   = 6
		seven = 7
		eight = 8
		nine  = 9
	)

	type args struct {
		buffer   [][]*int
		startRow int
		startCol int
		rows     int
		cols     int
	}
	tests := []struct {
		name string
		args args
		want [][]*int
	}{
		{"Should create a view with less rows starting at 0,0",
			args{[][]*int{
				{&one, &two, &three},
				{&four, &five, &six},
				{&seven, &eight, &nine},
			},
				0, 0, 2, 3},
			[][]*int{
				{&one, &two, &three},
				{&four, &five, &six},
			},
		},
		{"Should create a view with less columns starting at 0,0",
			args{[][]*int{
				{&one, &two, &three},
				{&four, &five, &six},
				{&seven, &eight, &nine},
			},
				0, 0, 3, 2},
			[][]*int{
				{&one, &two},
				{&four, &five},
				{&seven, &eight},
			},
		},
		{"Should create a view with less rows and columns starting at 0,0",
			args{[][]*int{
				{&one, &two, &three},
				{&four, &five, &six},
				{&seven, &eight, &nine},
			},
				0, 0, 2, 2},
			[][]*int{
				{&one, &two},
				{&four, &five},
			},
		},
		{
			"Should create a view with less rows and columns starting at 1,2",
			args{[][]*int{
				{&one, &two, &three},
				{&four, &five, &six},
				{&seven, &eight, &nine},
			},
				1, 2, 2, 1},
			[][]*int{
				{&six},
				{&nine},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FromIntPtr(tt.args.buffer, tt.args.startRow, tt.args.startCol, tt.args.rows, tt.args.cols); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromIntPtr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFromInt8Ptr(t *testing.T) {
	var (
		one   int8 = 1
		two   int8 = 2
		three int8 = 3
		four  int8 = 4
		five  int8 = 5
		six   int8 = 6
		seven int8 = 7
		eight int8 = 8
		nine  int8 = 9
	)

	type args struct {
		buffer   [][]*int8
		startRow int
		startCol int
		rows     int
		cols     int
	}
	tests := []struct {
		name string
		args args
		want [][]*int8
	}{
		{"Should create a view with less rows starting at 0,0",
			args{[][]*int8{
				{&one, &two, &three},
				{&four, &five, &six},
				{&seven, &eight, &nine},
			},
				0, 0, 2, 3},
			[][]*int8{
				{&one, &two, &three},
				{&four, &five, &six},
			},
		},
		{"Should create a view with less columns starting at 0,0",
			args{[][]*int8{
				{&one, &two, &three},
				{&four, &five, &six},
				{&seven, &eight, &nine},
			},
				0, 0, 3, 2},
			[][]*int8{
				{&one, &two},
				{&four, &five},
				{&seven, &eight},
			},
		},
		{"Should create a view with less rows and columns starting at 0,0",
			args{[][]*int8{
				{&one, &two, &three},
				{&four, &five, &six},
				{&seven, &eight, &nine},
			},
				0, 0, 2, 2},
			[][]*int8{
				{&one, &two},
				{&four, &five},
			},
		},
		{
			"Should create a view with less rows and columns starting at 1,2",
			args{[][]*int8{
				{&one, &two, &three},
				{&four, &five, &six},
				{&seven, &eight, &nine},
			},
				1, 2, 2, 1},
			[][]*int8{
				{&six},
				{&nine},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FromInt8Ptr(tt.args.buffer, tt.args.startRow, tt.args.startCol, tt.args.rows, tt.args.cols); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromInt8Ptr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFromInt16Ptr(t *testing.T) {
	var (
		one   int16 = 1
		two   int16 = 2
		three int16 = 3
		four  int16 = 4
		five  int16 = 5
		six   int16 = 6
		seven int16 = 7
		eight int16 = 8
		nine  int16 = 9
	)

	type args struct {
		buffer   [][]*int16
		startRow int
		startCol int
		rows     int
		cols     int
	}
	tests := []struct {
		name string
		args args
		want [][]*int16
	}{
		{"Should create a view with less rows starting at 0,0",
			args{[][]*int16{
				{&one, &two, &three},
				{&four, &five, &six},
				{&seven, &eight, &nine},
			},
				0, 0, 2, 3},
			[][]*int16{
				{&one, &two, &three},
				{&four, &five, &six},
			},
		},
		{"Should create a view with less columns starting at 0,0",
			args{[][]*int16{
				{&one, &two, &three},
				{&four, &five, &six},
				{&seven, &eight, &nine},
			},
				0, 0, 3, 2},
			[][]*int16{
				{&one, &two},
				{&four, &five},
				{&seven, &eight},
			},
		},
		{"Should create a view with less rows and columns starting at 0,0",
			args{[][]*int16{
				{&one, &two, &three},
				{&four, &five, &six},
				{&seven, &eight, &nine},
			},
				0, 0, 2, 2},
			[][]*int16{
				{&one, &two},
				{&four, &five},
			},
		},
		{
			"Should create a view with less rows and columns starting at 1,2",
			args{[][]*int16{
				{&one, &two, &three},
				{&four, &five, &six},
				{&seven, &eight, &nine},
			},
				1, 2, 2, 1},
			[][]*int16{
				{&six},
				{&nine},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FromInt16Ptr(tt.args.buffer, tt.args.startRow, tt.args.startCol, tt.args.rows, tt.args.cols); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromInt16Ptr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFromInt32Ptr(t *testing.T) {
	var (
		one   int32 = 1
		two   int32 = 2
		three int32 = 3
		four  int32 = 4
		five  int32 = 5
		six   int32 = 6
		seven int32 = 7
		eight int32 = 8
		nine  int32 = 9
	)

	type args struct {
		buffer   [][]*int32
		startRow int
		startCol int
		rows     int
		cols     int
	}
	tests := []struct {
		name string
		args args
		want [][]*int32
	}{
		{"Should create a view with less rows starting at 0,0",
			args{[][]*int32{
				{&one, &two, &three},
				{&four, &five, &six},
				{&seven, &eight, &nine},
			},
				0, 0, 2, 3},
			[][]*int32{
				{&one, &two, &three},
				{&four, &five, &six},
			},
		},
		{"Should create a view with less columns starting at 0,0",
			args{[][]*int32{
				{&one, &two, &three},
				{&four, &five, &six},
				{&seven, &eight, &nine},
			},
				0, 0, 3, 2},
			[][]*int32{
				{&one, &two},
				{&four, &five},
				{&seven, &eight},
			},
		},
		{"Should create a view with less rows and columns starting at 0,0",
			args{[][]*int32{
				{&one, &two, &three},
				{&four, &five, &six},
				{&seven, &eight, &nine},
			},
				0, 0, 2, 2},
			[][]*int32{
				{&one, &two},
				{&four, &five},
			},
		},
		{
			"Should create a view with less rows and columns starting at 1,2",
			args{[][]*int32{
				{&one, &two, &three},
				{&four, &five, &six},
				{&seven, &eight, &nine},
			},
				1, 2, 2, 1},
			[][]*int32{
				{&six},
				{&nine},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FromInt32Ptr(tt.args.buffer, tt.args.startRow, tt.args.startCol, tt.args.rows, tt.args.cols); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromInt32Ptr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFromInt64Ptr(t *testing.T) {
	var (
		one   int64 = 1
		two   int64 = 2
		three int64 = 3
		four  int64 = 4
		five  int64 = 5
		six   int64 = 6
		seven int64 = 7
		eight int64 = 8
		nine  int64 = 9
	)

	type args struct {
		buffer   [][]*int64
		startRow int
		startCol int
		rows     int
		cols     int
	}
	tests := []struct {
		name string
		args args
		want [][]*int64
	}{
		{"Should create a view with less rows starting at 0,0",
			args{[][]*int64{
				{&one, &two, &three},
				{&four, &five, &six},
				{&seven, &eight, &nine},
			},
				0, 0, 2, 3},
			[][]*int64{
				{&one, &two, &three},
				{&four, &five, &six},
			},
		},
		{"Should create a view with less columns starting at 0,0",
			args{[][]*int64{
				{&one, &two, &three},
				{&four, &five, &six},
				{&seven, &eight, &nine},
			},
				0, 0, 3, 2},
			[][]*int64{
				{&one, &two},
				{&four, &five},
				{&seven, &eight},
			},
		},
		{"Should create a view with less rows and columns starting at 0,0",
			args{[][]*int64{
				{&one, &two, &three},
				{&four, &five, &six},
				{&seven, &eight, &nine},
			},
				0, 0, 2, 2},
			[][]*int64{
				{&one, &two},
				{&four, &five},
			},
		},
		{
			"Should create a view with less rows and columns starting at 1,2",
			args{[][]*int64{
				{&one, &two, &three},
				{&four, &five, &six},
				{&seven, &eight, &nine},
			},
				1, 2, 2, 1},
			[][]*int64{
				{&six},
				{&nine},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FromInt64Ptr(tt.args.buffer, tt.args.startRow, tt.args.startCol, tt.args.rows, tt.args.cols); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromInt64Ptr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFromUintPtr(t *testing.T) {
	var (
		one   uint = 1
		two   uint = 2
		three uint = 3
		four  uint = 4
		five  uint = 5
		six   uint = 6
		seven uint = 7
		eight uint = 8
		nine  uint = 9
	)

	type args struct {
		buffer   [][]*uint
		startRow int
		startCol int
		rows     int
		cols     int
	}
	tests := []struct {
		name string
		args args
		want [][]*uint
	}{
		{"Should create a view with less rows starting at 0,0",
			args{[][]*uint{
				{&one, &two, &three},
				{&four, &five, &six},
				{&seven, &eight, &nine},
			},
				0, 0, 2, 3},
			[][]*uint{
				{&one, &two, &three},
				{&four, &five, &six},
			},
		},
		{"Should create a view with less columns starting at 0,0",
			args{[][]*uint{
				{&one, &two, &three},
				{&four, &five, &six},
				{&seven, &eight, &nine},
			},
				0, 0, 3, 2},
			[][]*uint{
				{&one, &two},
				{&four, &five},
				{&seven, &eight},
			},
		},
		{"Should create a view with less rows and columns starting at 0,0",
			args{[][]*uint{
				{&one, &two, &three},
				{&four, &five, &six},
				{&seven, &eight, &nine},
			},
				0, 0, 2, 2},
			[][]*uint{
				{&one, &two},
				{&four, &five},
			},
		},
		{
			"Should create a view with less rows and columns starting at 1,2",
			args{[][]*uint{
				{&one, &two, &three},
				{&four, &five, &six},
				{&seven, &eight, &nine},
			},
				1, 2, 2, 1},
			[][]*uint{
				{&six},
				{&nine},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FromUintPtr(tt.args.buffer, tt.args.startRow, tt.args.startCol, tt.args.rows, tt.args.cols); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromUintPtr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFromUint8Ptr(t *testing.T) {
	var (
		one   uint8 = 1
		two   uint8 = 2
		three uint8 = 3
		four  uint8 = 4
		five  uint8 = 5
		six   uint8 = 6
		seven uint8 = 7
		eight uint8 = 8
		nine  uint8 = 9
	)

	type args struct {
		buffer   [][]*uint8
		startRow int
		startCol int
		rows     int
		cols     int
	}
	tests := []struct {
		name string
		args args
		want [][]*uint8
	}{
		{"Should create a view with less rows starting at 0,0",
			args{[][]*uint8{
				{&one, &two, &three},
				{&four, &five, &six},
				{&seven, &eight, &nine},
			},
				0, 0, 2, 3},
			[][]*uint8{
				{&one, &two, &three},
				{&four, &five, &six},
			},
		},
		{"Should create a view with less columns starting at 0,0",
			args{[][]*uint8{
				{&one, &two, &three},
				{&four, &five, &six},
				{&seven, &eight, &nine},
			},
				0, 0, 3, 2},
			[][]*uint8{
				{&one, &two},
				{&four, &five},
				{&seven, &eight},
			},
		},
		{"Should create a view with less rows and columns starting at 0,0",
			args{[][]*uint8{
				{&one, &two, &three},
				{&four, &five, &six},
				{&seven, &eight, &nine},
			},
				0, 0, 2, 2},
			[][]*uint8{
				{&one, &two},
				{&four, &five},
			},
		},
		{
			"Should create a view with less rows and columns starting at 1,2",
			args{[][]*uint8{
				{&one, &two, &three},
				{&four, &five, &six},
				{&seven, &eight, &nine},
			},
				1, 2, 2, 1},
			[][]*uint8{
				{&six},
				{&nine},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FromUint8Ptr(tt.args.buffer, tt.args.startRow, tt.args.startCol, tt.args.rows, tt.args.cols); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromUint8Ptr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFromUint16Ptr(t *testing.T) {
	var (
		one   uint16 = 1
		two   uint16 = 2
		three uint16 = 3
		four  uint16 = 4
		five  uint16 = 5
		six   uint16 = 6
		seven uint16 = 7
		eight uint16 = 8
		nine  uint16 = 9
	)

	type args struct {
		buffer   [][]*uint16
		startRow int
		startCol int
		rows     int
		cols     int
	}
	tests := []struct {
		name string
		args args
		want [][]*uint16
	}{
		{"Should create a view with less rows starting at 0,0",
			args{[][]*uint16{
				{&one, &two, &three},
				{&four, &five, &six},
				{&seven, &eight, &nine},
			},
				0, 0, 2, 3},
			[][]*uint16{
				{&one, &two, &three},
				{&four, &five, &six},
			},
		},
		{"Should create a view with less columns starting at 0,0",
			args{[][]*uint16{
				{&one, &two, &three},
				{&four, &five, &six},
				{&seven, &eight, &nine},
			},
				0, 0, 3, 2},
			[][]*uint16{
				{&one, &two},
				{&four, &five},
				{&seven, &eight},
			},
		},
		{"Should create a view with less rows and columns starting at 0,0",
			args{[][]*uint16{
				{&one, &two, &three},
				{&four, &five, &six},
				{&seven, &eight, &nine},
			},
				0, 0, 2, 2},
			[][]*uint16{
				{&one, &two},
				{&four, &five},
			},
		},
		{
			"Should create a view with less rows and columns starting at 1,2",
			args{[][]*uint16{
				{&one, &two, &three},
				{&four, &five, &six},
				{&seven, &eight, &nine},
			},
				1, 2, 2, 1},
			[][]*uint16{
				{&six},
				{&nine},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FromUint16Ptr(tt.args.buffer, tt.args.startRow, tt.args.startCol, tt.args.rows, tt.args.cols); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromUint16Ptr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFromUint32Ptr(t *testing.T) {
	var (
		one   uint32 = 1
		two   uint32 = 2
		three uint32 = 3
		four  uint32 = 4
		five  uint32 = 5
		six   uint32 = 6
		seven uint32 = 7
		eight uint32 = 8
		nine  uint32 = 9
	)

	type args struct {
		buffer   [][]*uint32
		startRow int
		startCol int
		rows     int
		cols     int
	}
	tests := []struct {
		name string
		args args
		want [][]*uint32
	}{
		{"Should create a view with less rows starting at 0,0",
			args{[][]*uint32{
				{&one, &two, &three},
				{&four, &five, &six},
				{&seven, &eight, &nine},
			},
				0, 0, 2, 3},
			[][]*uint32{
				{&one, &two, &three},
				{&four, &five, &six},
			},
		},
		{"Should create a view with less columns starting at 0,0",
			args{[][]*uint32{
				{&one, &two, &three},
				{&four, &five, &six},
				{&seven, &eight, &nine},
			},
				0, 0, 3, 2},
			[][]*uint32{
				{&one, &two},
				{&four, &five},
				{&seven, &eight},
			},
		},
		{"Should create a view with less rows and columns starting at 0,0",
			args{[][]*uint32{
				{&one, &two, &three},
				{&four, &five, &six},
				{&seven, &eight, &nine},
			},
				0, 0, 2, 2},
			[][]*uint32{
				{&one, &two},
				{&four, &five},
			},
		},
		{
			"Should create a view with less rows and columns starting at 1,2",
			args{[][]*uint32{
				{&one, &two, &three},
				{&four, &five, &six},
				{&seven, &eight, &nine},
			},
				1, 2, 2, 1},
			[][]*uint32{
				{&six},
				{&nine},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FromUint32Ptr(tt.args.buffer, tt.args.startRow, tt.args.startCol, tt.args.rows, tt.args.cols); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromUint32Ptr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFromUint64Ptr(t *testing.T) {
	var (
		one   uint64 = 1
		two   uint64 = 2
		three uint64 = 3
		four  uint64 = 4
		five  uint64 = 5
		six   uint64 = 6
		seven uint64 = 7
		eight uint64 = 8
		nine  uint64 = 9
	)

	type args struct {
		buffer   [][]*uint64
		startRow int
		startCol int
		rows     int
		cols     int
	}
	tests := []struct {
		name string
		args args
		want [][]*uint64
	}{
		{"Should create a view with less rows starting at 0,0",
			args{[][]*uint64{
				{&one, &two, &three},
				{&four, &five, &six},
				{&seven, &eight, &nine},
			},
				0, 0, 2, 3},
			[][]*uint64{
				{&one, &two, &three},
				{&four, &five, &six},
			},
		},
		{"Should create a view with less columns starting at 0,0",
			args{[][]*uint64{
				{&one, &two, &three},
				{&four, &five, &six},
				{&seven, &eight, &nine},
			},
				0, 0, 3, 2},
			[][]*uint64{
				{&one, &two},
				{&four, &five},
				{&seven, &eight},
			},
		},
		{"Should create a view with less rows and columns starting at 0,0",
			args{[][]*uint64{
				{&one, &two, &three},
				{&four, &five, &six},
				{&seven, &eight, &nine},
			},
				0, 0, 2, 2},
			[][]*uint64{
				{&one, &two},
				{&four, &five},
			},
		},
		{
			"Should create a view with less rows and columns starting at 1,2",
			args{[][]*uint64{
				{&one, &two, &three},
				{&four, &five, &six},
				{&seven, &eight, &nine},
			},
				1, 2, 2, 1},
			[][]*uint64{
				{&six},
				{&nine},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FromUint64Ptr(tt.args.buffer, tt.args.startRow, tt.args.startCol, tt.args.rows, tt.args.cols); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromUint64Ptr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFromBoolPtr(t *testing.T) {

	var (
		one   = true
		two   = false
		three = true
		four  = true
		five  = true
		six   = false
		seven = false
		eight = false
		nine  = false
	)

	type args struct {
		buffer   [][]*bool
		startRow int
		startCol int
		rows     int
		cols     int
	}
	tests := []struct {
		name string
		args args
		want [][]*bool
	}{
		{"Should create a view with less rows starting at 0,0",
			args{[][]*bool{
				{&one, &two, &three},
				{&four, &five, &six},
				{&seven, &eight, &nine},
			},
				0, 0, 2, 3},
			[][]*bool{
				{&one, &two, &three},
				{&four, &five, &six},
			},
		},
		{"Should create a view with less columns starting at 0,0",
			args{[][]*bool{
				{&one, &two, &three},
				{&four, &five, &six},
				{&seven, &eight, &nine},
			},
				0, 0, 3, 2},
			[][]*bool{
				{&one, &two},
				{&four, &five},
				{&seven, &eight},
			},
		},
		{"Should create a view with less rows and columns starting at 0,0",
			args{[][]*bool{
				{&one, &two, &three},
				{&four, &five, &six},
				{&seven, &eight, &nine},
			},
				0, 0, 2, 2},
			[][]*bool{
				{&one, &two},
				{&four, &five},
			},
		},
		{
			"Should create a view with less rows and columns starting at 1,2",
			args{[][]*bool{
				{&one, &two, &three},
				{&four, &five, &six},
				{&seven, &eight, &nine},
			},
				1, 2, 2, 1},
			[][]*bool{
				{&six},
				{&nine},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FromBoolPtr(tt.args.buffer, tt.args.startRow, tt.args.startCol, tt.args.rows, tt.args.cols); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromBoolPtr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFromFloat32Ptr(t *testing.T) {
	var (
		one   float32 = 1
		two   float32 = 2
		three float32 = 3
		four  float32 = 4
		five  float32 = 5
		six   float32 = 6
		seven float32 = 7
		eight float32 = 8
		nine  float32 = 9
	)

	type args struct {
		buffer   [][]*float32
		startRow int
		startCol int
		rows     int
		cols     int
	}
	tests := []struct {
		name string
		args args
		want [][]*float32
	}{
		{"Should create a view with less rows starting at 0,0",
			args{[][]*float32{
				{&one, &two, &three},
				{&four, &five, &six},
				{&seven, &eight, &nine},
			},
				0, 0, 2, 3},
			[][]*float32{
				{&one, &two, &three},
				{&four, &five, &six},
			},
		},
		{"Should create a view with less columns starting at 0,0",
			args{[][]*float32{
				{&one, &two, &three},
				{&four, &five, &six},
				{&seven, &eight, &nine},
			},
				0, 0, 3, 2},
			[][]*float32{
				{&one, &two},
				{&four, &five},
				{&seven, &eight},
			},
		},
		{"Should create a view with less rows and columns starting at 0,0",
			args{[][]*float32{
				{&one, &two, &three},
				{&four, &five, &six},
				{&seven, &eight, &nine},
			},
				0, 0, 2, 2},
			[][]*float32{
				{&one, &two},
				{&four, &five},
			},
		},
		{
			"Should create a view with less rows and columns starting at 1,2",
			args{[][]*float32{
				{&one, &two, &three},
				{&four, &five, &six},
				{&seven, &eight, &nine},
			},
				1, 2, 2, 1},
			[][]*float32{
				{&six},
				{&nine},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FromFloat32Ptr(tt.args.buffer, tt.args.startRow, tt.args.startCol, tt.args.rows, tt.args.cols); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromFloat32Ptr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFromFloat64Ptr(t *testing.T) {
	var (
		one   float64 = 1
		two   float64 = 2
		three float64 = 3
		four  float64 = 4
		five  float64 = 5
		six   float64 = 6
		seven float64 = 7
		eight float64 = 8
		nine  float64 = 9
	)

	type args struct {
		buffer   [][]*float64
		startRow int
		startCol int
		rows     int
		cols     int
	}
	tests := []struct {
		name string
		args args
		want [][]*float64
	}{
		{"Should create a view with less rows starting at 0,0",
			args{[][]*float64{
				{&one, &two, &three},
				{&four, &five, &six},
				{&seven, &eight, &nine},
			},
				0, 0, 2, 3},
			[][]*float64{
				{&one, &two, &three},
				{&four, &five, &six},
			},
		},
		{"Should create a view with less columns starting at 0,0",
			args{[][]*float64{
				{&one, &two, &three},
				{&four, &five, &six},
				{&seven, &eight, &nine},
			},
				0, 0, 3, 2},
			[][]*float64{
				{&one, &two},
				{&four, &five},
				{&seven, &eight},
			},
		},
		{"Should create a view with less rows and columns starting at 0,0",
			args{[][]*float64{
				{&one, &two, &three},
				{&four, &five, &six},
				{&seven, &eight, &nine},
			},
				0, 0, 2, 2},
			[][]*float64{
				{&one, &two},
				{&four, &five},
			},
		},
		{
			"Should create a view with less rows and columns starting at 1,2",
			args{[][]*float64{
				{&one, &two, &three},
				{&four, &five, &six},
				{&seven, &eight, &nine},
			},
				1, 2, 2, 1},
			[][]*float64{
				{&six},
				{&nine},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FromFloat64Ptr(tt.args.buffer, tt.args.startRow, tt.args.startCol, tt.args.rows, tt.args.cols); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromFloat64Ptr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFromComplex64Ptr(t *testing.T) {

	var (
		one   complex64 = 1
		two   complex64 = 2
		three complex64 = 3
		four  complex64 = 4
		five  complex64 = 5
		six   complex64 = 6
		seven complex64 = 7
		eight complex64 = 8
		nine  complex64 = 9
	)

	type args struct {
		buffer   [][]*complex64
		startRow int
		startCol int
		rows     int
		cols     int
	}
	tests := []struct {
		name string
		args args
		want [][]*complex64
	}{
		{"Should create a view with less rows starting at 0,0",
			args{[][]*complex64{
				{&one, &two, &three},
				{&four, &five, &six},
				{&seven, &eight, &nine},
			},
				0, 0, 2, 3},
			[][]*complex64{
				{&one, &two, &three},
				{&four, &five, &six},
			},
		},
		{"Should create a view with less columns starting at 0,0",
			args{[][]*complex64{
				{&one, &two, &three},
				{&four, &five, &six},
				{&seven, &eight, &nine},
			},
				0, 0, 3, 2},
			[][]*complex64{
				{&one, &two},
				{&four, &five},
				{&seven, &eight},
			},
		},
		{"Should create a view with less rows and columns starting at 0,0",
			args{[][]*complex64{
				{&one, &two, &three},
				{&four, &five, &six},
				{&seven, &eight, &nine},
			},
				0, 0, 2, 2},
			[][]*complex64{
				{&one, &two},
				{&four, &five},
			},
		},
		{
			"Should create a view with less rows and columns starting at 1,2",
			args{[][]*complex64{
				{&one, &two, &three},
				{&four, &five, &six},
				{&seven, &eight, &nine},
			},
				1, 2, 2, 1},
			[][]*complex64{
				{&six},
				{&nine},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FromComplex64Ptr(tt.args.buffer, tt.args.startRow, tt.args.startCol, tt.args.rows, tt.args.cols); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromComplexPtr64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFromFromComplex128Ptr(t *testing.T) {
	var (
		one   complex128 = 1
		two   complex128 = 2
		three complex128 = 3
		four  complex128 = 4
		five  complex128 = 5
		six   complex128 = 6
		seven complex128 = 7
		eight complex128 = 8
		nine  complex128 = 9
	)

	type args struct {
		buffer   [][]*complex128
		startRow int
		startCol int
		rows     int
		cols     int
	}
	tests := []struct {
		name string
		args args
		want [][]*complex128
	}{
		{"Should create a view with less rows starting at 0,0",
			args{[][]*complex128{
				{&one, &two, &three},
				{&four, &five, &six},
				{&seven, &eight, &nine},
			},
				0, 0, 2, 3},
			[][]*complex128{
				{&one, &two, &three},
				{&four, &five, &six},
			},
		},
		{"Should create a view with less columns starting at 0,0",
			args{[][]*complex128{
				{&one, &two, &three},
				{&four, &five, &six},
				{&seven, &eight, &nine},
			},
				0, 0, 3, 2},
			[][]*complex128{
				{&one, &two},
				{&four, &five},
				{&seven, &eight},
			},
		},
		{"Should create a view with less rows and columns starting at 0,0",
			args{[][]*complex128{
				{&one, &two, &three},
				{&four, &five, &six},
				{&seven, &eight, &nine},
			},
				0, 0, 2, 2},
			[][]*complex128{
				{&one, &two},
				{&four, &five},
			},
		},
		{
			"Should create a view with less rows and columns starting at 1,2",
			args{[][]*complex128{
				{&one, &two, &three},
				{&four, &five, &six},
				{&seven, &eight, &nine},
			},
				1, 2, 2, 1},
			[][]*complex128{
				{&six},
				{&nine},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FromComplex128Ptr(tt.args.buffer, tt.args.startRow, tt.args.startCol, tt.args.rows, tt.args.cols); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromComplexPtr128() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFromStringPtr(t *testing.T) {

	var (
		one   = "1"
		two   = "2"
		three = "3"
		four  = "4"
		five  = "5"
		six   = "6"
		seven = "7"
		eight = "8"
		nine  = "9"
	)

	type args struct {
		buffer   [][]*string
		startRow int
		startCol int
		rows     int
		cols     int
	}
	tests := []struct {
		name string
		args args
		want [][]*string
	}{
		{"Should create a view with less rows starting at 0,0",
			args{[][]*string{
				{&one, &two, &three},
				{&four, &five, &six},
				{&seven, &eight, &nine},
			},
				0, 0, 2, 3},
			[][]*string{
				{&one, &two, &three},
				{&four, &five, &six},
			},
		},
		{"Should create a view with less columns starting at 0,0",
			args{[][]*string{
				{&one, &two, &three},
				{&four, &five, &six},
				{&seven, &eight, &nine},
			},
				0, 0, 3, 2},
			[][]*string{
				{&one, &two},
				{&four, &five},
				{&seven, &eight},
			},
		},
		{"Should create a view with less rows and columns starting at 0,0",
			args{[][]*string{
				{&one, &two, &three},
				{&four, &five, &six},
				{&seven, &eight, &nine},
			},
				0, 0, 2, 2},
			[][]*string{
				{&one, &two},
				{&four, &five},
			},
		},
		{
			"Should create a view with less rows and columns starting at 1,2",
			args{[][]*string{
				{&one, &two, &three},
				{&four, &five, &six},
				{&seven, &eight, &nine},
			},
				1, 2, 2, 1},
			[][]*string{
				{&six},
				{&nine},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FromStringPtr(tt.args.buffer, tt.args.startRow, tt.args.startCol, tt.args.rows, tt.args.cols); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromStringPtr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFromRunePtr(t *testing.T) {

	var (
		one   = '1'
		two   = '2'
		three = '3'
		four  = '4'
		five  = '5'
		six   = '6'
		seven = '7'
		eight = '8'
		nine  = '9'
	)

	type args struct {
		buffer   [][]*rune
		startRow int
		startCol int
		rows     int
		cols     int
	}
	tests := []struct {
		name string
		args args
		want [][]*rune
	}{
		{"Should create a view with less rows starting at 0,0",
			args{[][]*rune{
				{&one, &two, &three},
				{&four, &five, &six},
				{&seven, &eight, &nine},
			},
				0, 0, 2, 3},
			[][]*rune{
				{&one, &two, &three},
				{&four, &five, &six},
			},
		},
		{"Should create a view with less columns starting at 0,0",
			args{[][]*rune{
				{&one, &two, &three},
				{&four, &five, &six},
				{&seven, &eight, &nine},
			},
				0, 0, 3, 2},
			[][]*rune{
				{&one, &two},
				{&four, &five},
				{&seven, &eight},
			},
		},
		{"Should create a view with less rows and columns starting at 0,0",
			args{[][]*rune{
				{&one, &two, &three},
				{&four, &five, &six},
				{&seven, &eight, &nine},
			},
				0, 0, 2, 2},
			[][]*rune{
				{&one, &two},
				{&four, &five},
			},
		},
		{
			"Should create a view with less rows and columns starting at 1,2",
			args{[][]*rune{
				{&one, &two, &three},
				{&four, &five, &six},
				{&seven, &eight, &nine},
			},
				1, 2, 2, 1},
			[][]*rune{
				{&six},
				{&nine},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FromRunePtr(tt.args.buffer, tt.args.startRow, tt.args.startCol, tt.args.rows, tt.args.cols); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromRunePtr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleEditDistance() {
	buf := make([][]int, 100)
	for i := range buf {
		buf[i] = make([]int, 100)
	}


	work := [][2]string{
		{"sitting", "kitten"},
		{"Sunday", "Saturday"},
	}

	for _, words := range work {
		from := words[0]
		to := words[1]

		view := FromInt(buf, 0, 0, len(from)+1, len(to)+1)
		if err := fillLevenshteinMatrix(view, from, to); err != nil {
			fmt.Printf("ERROR: %s\n", err.Error())
			return
		}
		fmt.Printf("From: %s To: %s Distance: %d Ratio: %.2f\n", from, to, calculateDistance(view), calculateRatio(view))
	}

	// Output:
	//From: sitting To: kitten Distance: 5 Ratio: 0.62
	//From: Sunday To: Saturday Distance: 4 Ratio: 0.71
}

func fillLevenshteinMatrix(matrix [][]int, from, to string) error {

	rows := len(from) + 1
	cols := len(to) + 1

	if len(matrix) != rows {
		return fmt.Errorf("matrix must have len(from) + 1 = %d rows, but has %d rows.", rows, len(matrix))
	}

	for row := range matrix {
		if len(matrix[row]) != cols {
			return fmt.Errorf("matrix must have len(cols) + 1 = %d cols, but row %d has %d cols.", cols, row, len(matrix[row]))
		}
		matrix[row][0] = row
	}

	for col := range matrix[0] {
		matrix[0][col] = col
	}

	for row := 1; row < len(matrix); row++ {
		for col := 1; col < len(matrix[row]); col++ {
			var substitutionCost int
			if from[row-1] != to[col-1] {
				substitutionCost = 2
			}

			matrix[row][col] = min(
				matrix[row-1][col]+1,                  // deletion
				matrix[row][col-1]+1,                  // insertion
				matrix[row-1][col-1]+substitutionCost, // substitution
			)
		}
	}
	return nil
}

func calculateDistance(matrix [][]int) int {
	return matrix[len(matrix)-1][len(matrix[len(matrix)-1])-1]
}

func calculateRatio(matrix [][]int) float64 {
	sum := float64(len(matrix) + len(matrix[0]) - 2)
	if sum == 0 {
		return 0
	}
	return 1 - float64(calculateDistance(matrix))/sum
}

func min(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		}
		return c
	}
	if b < c {
		return b
	}
	return c
}
