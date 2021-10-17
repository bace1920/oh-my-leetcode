package main

import "fmt"

func main() {
	// nums := []int{1, -2, 3, -2}
	nums := []int{5, -3, 5}
	// nums := []int{3, 2, 1, 0, 4}
	// nums := []int{1, 1}

	fmt.Println(maxSubarraySumCircular(nums))
}

// 如果存在使用到循环的情况，其实中间没有取的部分等于是minSubarray
func maxSubarraySumCircular(nums []int) int {
	l := len(nums)
	maxIter := nums[0]
	minIter := 0
	maxSum := nums[0]
	minSum := 0
	totalSum := nums[0]
	for i := 1; i < l; i++ {
		totalSum += nums[i]
		maxIter = max(maxIter+nums[i], nums[i])
		maxSum = max(maxIter, maxSum)
		minIter = min(minIter+nums[i], nums[i])
		minSum = min(minIter, minSum)
	}

	return max(maxSum, totalSum-minSum)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
