package main

import (
	"fmt"
)

func main() {
	// s := "26"
	s := 3
	// s := "226"
	// s := "2106"
	// s := "1123"
	fmt.Println(numTrees(s))
}

func numTrees(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	// n个连续数，从中抽取一个，必然分成2个相同的子问题，分别是大于n的右子树和小于n的左子树，于是可以dp
	for i := 2; i < n+1; i++ {
		for q := 1; q <= i; q++ {
			dp[i] += dp[q-1] * dp[i-q]
		}
	}
	return dp[n]
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
