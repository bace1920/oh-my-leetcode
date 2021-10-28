package main

import "fmt"

func main() {
	// grid := [][]byte{{'1', '0', '1', '0', '0'}, {'1', '0', '1', '1', '1'}, {'1', '1', '1', '1', '1'}, {'1', '0', '0', '1', '0'}}
	grid := [][]byte{{'0', '0', '0', '1'}, {'1', '1', '0', '1'}, {'1', '1', '1', '1'}, {'0', '1', '1', '1'}, {'0', '1', '1', '1'}}
	// grid := [][]int{{1, 2}, {1, 1}}
	// grid := [][]int{{0, 1}, {0, 0}}
	// grid := [][]int{{0, 0}, {0, 1}}
	// grid := [][]int{{1}}
	fmt.Println(maximalSquare(grid))
}

// dp[i][q] 代表以i,q为右下角的最大矩阵对角线长度
func maximalSquare(matrix [][]byte) int {
	height := len(matrix)
	width := len(matrix[0])
	dp := make([][]int, height)
	for i := 0; i < height; i++ {
		dp[i] = make([]int, width)
	}

	res := 0
	for i := 0; i < height; i++ {
		for q := 0; q < width; q++ {
			if matrix[i][q] == '0' || i == 0 || q == 0 {
				dp[i][q] = int(matrix[i][q] - '0')
			} else {
				dp[i][q] = min(dp[i][q-1], min(dp[i-1][q-1], dp[i-1][q])) + 1
			}

			res = max(res, dp[i][q])
		}
	}
	return res * res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
