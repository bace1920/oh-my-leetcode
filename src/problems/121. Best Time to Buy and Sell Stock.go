package main

import "fmt"

func main() {
	// nums := []int{1, 2, 3, 1}
	// nums := []int{2, 3, -2, 4}
	// nums := []int{0, 1, -2, -3, -4}
	// nums := []int{-16, 0, -5, 2, 2, -13, 11, 8}
	nums := []int{7, 1, 5, 3, 6, 4}
	// nums := []int{-1, 2}
	// nums := []int{2, 3, -2, 4}
	// nums := []int{1, 1}
	fmt.Println(maxProfit(nums))
}

func maxProfit(values []int) int {
	maxSoFar, maxp := 0, 0
	for i := 1; i < len(values); i++ {
		maxp += values[i] - values[i-1]
		maxp = max(maxp, 0)
		maxSoFar = max(maxSoFar, maxp)
	}
	return maxSoFar
}

func maxProfitOld(values []int) int {
	profit, dp_max := 0, values[len(values)-1]
	dp_min := make([]int, len(values))

	dp_min[0] = values[0]
	for i := 1; i < len(values); i++ {
		dp_min[i] = min(values[i], dp_min[i-1])
	}

	for i := len(values) - 1; i >= 1; i-- {
		dp_max = max(values[i], dp_max)
		profit = max(profit, dp_max-dp_min[i-1])
	}
	return profit
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
