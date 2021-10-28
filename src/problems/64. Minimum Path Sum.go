package main

import "fmt"

func main() {
	grid := [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}
	// grid := [][]int{{1, 2}, {1, 1}}
	// grid := [][]int{{0, 1}, {0, 0}}
	// grid := [][]int{{0, 0}, {0, 1}}
	// grid := [][]int{{1}}
	fmt.Println(minPathSum(grid))
}

func minPathSum(grid [][]int) int {
	height := len(grid)
	width := len(grid[0])
	cur := make([]int, width)

	cur[0] = grid[0][0]
	for q := 1; q < width; q++ {
		cur[q] = cur[q-1] + grid[0][q]
	}

	for i := 1; i < height; i++ {
		cur[0] = cur[0] + grid[i][0]
		for q := 1; q < width; q++ {
			cur[q] = min(cur[q-1]+grid[i][q], cur[q]+grid[i][q])
		}
	}
	return cur[width-1]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
