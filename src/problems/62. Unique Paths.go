package main

import "fmt"

func main() {
	//s := "mississippi"
	//p := "mis*is*p*."
	fmt.Println(uniquePaths(3, 2))
}

func uniquePaths(m int, n int) int {
	cur := make([]int, n)

	for q := 0; q < n; q++ {
		cur[q] = 1
	}

	for i := 1; i < m; i++ {
		for q := 1; q < n; q++ {
			cur[q] += cur[q-1]
		}
	}
	return cur[n-1]
}

