package main

import (
	"fmt"
)

func main() {
	// nums := []int{6, 1, 6, 4, 3, 0, 2}
	nums := []int{1, 3, 2, 8, 4, 9}
	// nums := []int{0, 1, -2, -3, -4}
	// nums := []int{-16, 0, -5, 2, 2, -13, 11, 8}
	// nums := []int{1, 2, 3, 0, 2}
	// nums := []int{3, 2, 6, 5, 0, 3}
	// nums := []int{2, 3, -2, 4}
	// nums := []int{1, 2, 4}
	fmt.Println(maxProfit(nums, 2))
}

func maxProfit(prices []int, fee int) int {
	if len(prices) == 1 {
		return 0
	}

	// s0 静息态, s1持仓态, 参考309状态机
	s0 := 0
	s1 := -prices[0]

	m := 0
	for i := 1; i < len(prices); i++ {
		s0 = max(s0, s1+prices[i]-fee)
		s1 = max(s1, s0-prices[i])
		m = max(s0, m)
	}

	return m
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
