package main

import (
	"fmt"
)

func main() {
	// s := "26"
	s := 5
	// s := "226"
	// s := "2106"
	// s := "1123"
	fmt.Println(getRow(s))
}
func getRow(rowIndex int) []int {
	last := make([]int, 1)
	last[0] = 1
	if rowIndex == 0 {
		return last
	}

	for i := 1; i <= rowIndex; i++ {
		new := make([]int, i+1)
		new[0], new[i] = 1, 1
		for q := 1; q < i; q++ {
			new[q] = last[q-1] + last[q]
		}
		last = new
	}
	return last
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
