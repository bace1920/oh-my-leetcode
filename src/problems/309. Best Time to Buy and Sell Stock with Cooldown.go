package main

import (
	"fmt"
)

func main() {
	// nums := []int{6, 1, 6, 4, 3, 0, 2}
	nums := []int{6, 1, 3, 2, 4, 7}
	// nums := []int{0, 1, -2, -3, -4}
	// nums := []int{-16, 0, -5, 2, 2, -13, 11, 8}
	// nums := []int{1, 2, 3, 0, 2}
	// nums := []int{3, 2, 6, 5, 0, 3}
	// nums := []int{2, 3, -2, 4}
	// nums := []int{1, 2, 4}
	fmt.Println(maxProfit(nums))
}

// 状态机
func maxProfit(prices []int) int {
	if len(prices) == 1 {
		return 0
	}

	// s0 静息态, s1持仓态，s2刚刚平仓
	s0 := make([]int, len(prices)) //rest
	s1 := make([]int, len(prices)) // buy
	s2 := make([]int, len(prices)) // sell

	s0[0] = 0
	s1[0] = -prices[0]
	s2[0] = -5001

	m := 0
	for i := 1; i < len(prices); i++ {
		s0[i] = max(s0[i-1], s2[i-1])
		s1[i] = max(s1[i-1], s0[i-1]-prices[i])
		s2[i] = s1[i-1] + prices[i]
		m = max(max(s0[i], s2[i]), m)
	}

	return m
}

// 超时了
func maxProfitRecursion(prices []int) int {
	if len(prices) == 1 {
		return 0
	}
	dp := make([]int, len(prices))
	for i := 1; i < len(dp); i++ {
		dp[i] = prices[i] - prices[i-1]
	}

	return do(dp, 0, 1, false)
}

func do(dp []int, curProfit int, i int, last bool) int {
	if i >= len(dp) {
		return curProfit
	}
	if len(dp)-i == 1 {
		return max(curProfit, curProfit+dp[i])
	} else {
		var a, b, c int
		if last {
			a = do(dp, curProfit, i+2, false)
			b = do(dp, curProfit+dp[i], i+1, true)
		} else {
			a = do(dp, curProfit, i+1, false)
			b = do(dp, curProfit+dp[i], i+1, true)
			c = do(dp, curProfit+dp[i+1], i+2, true)
		}
		return max(max(a, b), c)
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
