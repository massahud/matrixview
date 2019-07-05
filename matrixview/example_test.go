// Copyright 2009 Geraldo Augusto Massahud Rodrigues dos Santos. All rights reserved.
// This code is licensed under APACHE-2 and MIT, the respective licenses can be found
// at LICENSE-APACHE and LICENSE-MIT files.

package matrixview_test

import (
	"fmt"
	"github.com/massahud/matrixview/matrixview"
)


func Example() {
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

		view := matrixview.FromInt(buf, 0, 0, len(from)+1, len(to)+1)
		if err := fillLevenshteinMatrix(view, from, to); err != nil {
			fmt.Printf("ERROR: %s\n", err.Error())
			return
		}
		fmt.Printf("From: %s To: %s Distance: %d Ratio: %.2f\n", from, to, distance(view), ratio(view))
	}

	//Output:
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

func distance(matrix [][]int) int {
	return matrix[len(matrix)-1][len(matrix[len(matrix)-1])-1]
}

func ratio(matrix [][]int) float64 {
	sum := float64(len(matrix) + len(matrix[0]) - 2)
	if sum == 0 {
		return 0
	}
	return 1 - float64(distance(matrix))/sum
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
