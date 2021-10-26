package main

import (
	"fmt"
)

func main() {
	// s := "26"
	m := [][]int{{2, 1, 3}, {6, 5, 4}, {7, 8, 9}}
	// s := "226"
	// s := "2106"
	// s := "1123"
	fmt.Println(minFallingPathSum(m))
}
func minFallingPathSum(matrix [][]int) int {
	if len(matrix) == 1 {
		return matrix[0][0]
	}

	last := make([]int, len(matrix))
	now := make([]int, len(matrix))

	for i := 0; i < len(matrix); i++ {
		now[0] = min(last[0], last[1]) + matrix[i][0]
		now[len(matrix)-1] = min(last[len(matrix)-1], last[len(matrix)-2]) + matrix[i][len(matrix)-1]

		for q := 1; q < len(matrix)-1; q++ {
			now[q] = min(last[q-1], min(last[q], last[q+1])) + matrix[i][q]
		}

		last, now = now, last
	}

	res := last[0]
	for i := range last {
		res = min(res, last[i])
	}
	return res
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
