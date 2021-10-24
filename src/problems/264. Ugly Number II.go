package main

import (
	"fmt"
)

func main() {
	// s := "26"
	s := 10
	// s := "226"
	// s := "2106"
	// s := "1123"
	fmt.Println(nthUglyNumber(s))
}

// 下一个数必须是已知所有数于2，3，5相乘后，最小的，却又未知的数

func nthUglyNumber(n int) int {
	if n == 1 {
		return 1
	}

	dp := make([]int, n)
	dp[0] = 1
	p2, p3, p5 := 0, 0, 0

	for i := 1; i < n; i++ {
		dp[i] = min(min(dp[p2]*2, dp[p3]*3), dp[p5]*5)
		if dp[i] == dp[p2]*2 { // 我们已经考虑过 dp[p2]*2 所以我们可以push一位，不必担心重复考虑
			p2++
		}
		if dp[i] == dp[p3]*3 {
			p3++
		}
		if dp[i] == dp[p5]*5 {
			p5++
		}
	}
	return dp[n-1]
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
