package main

import (
	"fmt"
)

func main() {
	triangle := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	// triangle := [][]int{{67, 64, 78}, {99, 98, 38}, {82, 46, 46}, {6, 52, 55}, {55, 99, 45}}
	fmt.Println(matrixBlockSum(triangle, 1))
}

// dp记录(0,0)到(i,j)为对角线的矩形总和
func matrixBlockSum(mat [][]int, k int) [][]int {
	height := len(mat)
	width := len(mat[0])

	dp := make([][]int, height+1)
	for i := 0; i <= height; i++ {
		dp[i] = make([]int, width+1)
	}
	for i := 0; i < height; i++ {
		for q := 0; q < width; q++ {
			dp[i+1][q+1] = dp[i+1][q] + dp[i][q+1] - dp[i][q] + mat[i][q]
		}
	}

	res := make([][]int, height)
	for i := 0; i < height; i++ {
		res[i] = make([]int, width)
		for q := 0; q < width; q++ {
			left, right := max(0, q-k), min(width, q+k+1)
			up, down := max(0, i-k), min(height, i+k+1)
			res[i][q] = dp[down][right] - dp[down][left] - dp[up][right] + dp[up][left]
		}
	}

	return res
}

// dp记录(0,0)到(i,j)总和
// col 记录每一列到第i行的总和
// 计算时用两者去做排除相减
func matrixBlockSumOld(mat [][]int, k int) [][]int {
	height := len(mat)
	width := len(mat[0])

	dp := make([][]int, height)
	for i := range dp {
		dp[i] = make([]int, width)
	}

	dp[0][0] = mat[0][0]
	for i := range dp {
		for q := 1; q < width; q++ {
			dp[i][q] = dp[i][q-1] + mat[i][q]
		}
		if i != height-1 {
			dp[i+1][0] = dp[i][width-1] + mat[i+1][0]
		}
	}

	col := make([][]int, height)
	col[0] = make([]int, width)
	for i := 0; i < width; i++ {
		col[0][i] = mat[0][i]
	}

	for i := 1; i < height; i++ {
		col[i] = make([]int, width)
		for q := 0; q < width; q++ {
			col[i][q] = col[i-1][q] + mat[i][q]
		}
	}

	res := make([][]int, height)
	for i := range dp {
		res[i] = make([]int, width)
		for q := 0; q < width; q++ {
			left, right := max(0, q-k), min(width-1, q+k)
			up, down := max(0, i-k), min(height-1, i+k)
			res[i][q] = dp[down][width-1]
			for m := 0; m < left; m++ {
				res[i][q] -= col[down][m]
			}
			for m := right + 1; m < width; m++ {
				res[i][q] -= col[down][m]
			}

			if up > 0 {
				for m := left; m <= right; m++ {
					res[i][q] -= col[up-1][m]
				}
			}
		}
	}

	return res
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
