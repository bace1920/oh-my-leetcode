package main

import (
	"fmt"
)

func main() {
	// s := "26"
	s := 15
	// s := "226"
	// s := "2106"
	// s := "1123"
	fmt.Println(generate(s))
}
func generate(numRows int) [][]int {
	rows := make([][]int, numRows)
	for i := range rows {
		rows[i] = make([]int, i+1)
		rows[i][0], rows[i][i] = 1, 1
		for q := 1; q < i; q++ {
			rows[i][q] = rows[i-1][q-1] + rows[i-1][q]
		}
	}
	return rows
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
