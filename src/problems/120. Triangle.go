package main

import (
	"fmt"
)

func main() {
	triangle := [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}
	fmt.Println(minimumTotal(triangle))
}
func minimumTotal(triangle [][]int) int {
	if len(triangle) == 1 {
		return triangle[0][0]
	}
	last := make([]int, len(triangle[len(triangle)-1]))
	now := make([]int, len(triangle[len(triangle)-1]))

	last[0] = triangle[0][0] + triangle[1][0]
	last[1] = triangle[0][0] + triangle[1][1]
	for i := 2; i < len(triangle); i++ {
		now[0] = last[0] + triangle[i][0]
		now[i] = last[i-1] + triangle[i][i]
		for q := 1; q < i; q++ {
			now[q] = min(last[q], last[q-1]) + triangle[i][q]
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
