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
	sum := 0
	for i := 1; i < len(values); i++ {
		sum += max(values[i]-values[i-1], 0)
	}
	return sum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
