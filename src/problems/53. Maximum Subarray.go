package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	//nums := []int{-2, 1}
	fmt.Println(maxSubArrayOn(nums))
}

func maxSubArrayOn(nums []int) int {
	length := len(nums)
	dp := make([]int, length)
	dp[0] = nums[0]
	for i := 1; i < length; i++ {
		if dp[i-1] > 0 {
			dp[i] = dp[i-1] + nums[i]
		} else {
			dp[i] = nums[i]
		}
	}

	max := dp[0]
	for _, sum := range dp {
		if sum > max {
			max = sum
		}
	}
	return max
}

func maxSubArray(nums []int) int {
	length := len(nums)
	maxSum := nums[0]
	for i := 0; i < length; i++ {
		sum := foo(nums, length, nums[i], i)
		if sum > maxSum {
			maxSum = sum
		}
	}
	return maxSum
}

func foo(nums []int, length int, sum int, pos int) int {
	if pos == length-1 {
		return sum
	} else {
		return int(math.Max(float64(sum), float64(foo(nums, length, sum+nums[pos+1], pos+1))))
	}
}
