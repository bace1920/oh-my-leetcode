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
	obj := Constructor(m)
	fmt.Println(obj.SumRegion(0, 0, 1, 1))
}

type NumMatrix struct {
	dp     [][]int
	matrix [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	height := len(matrix)
	width := len(matrix[0])

	dp := make([][]int, height+1)
	for i := 0; i <= height; i++ {
		dp[i] = make([]int, width+1)
	}
	for i := 0; i < height; i++ {
		for q := 0; q < width; q++ {
			dp[i+1][q+1] = dp[i+1][q] + dp[i][q+1] - dp[i][q] + matrix[i][q]
		}
	}
	return NumMatrix{
		dp:     dp,
		matrix: matrix,
	}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	// r1, c1, r2, c2 := row1+1, col1+1, row2+1, col2+1
	r2, c2 := row2+1, col2+1
	return this.dp[r2][c2] - this.dp[r2][col1] - this.dp[row1][c2] + this.dp[row1][col1]

}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
