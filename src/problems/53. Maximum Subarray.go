package main

import (
	"fmt"
	"math"
)

func main() {
	//nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	nums := []int{-1}
	fmt.Println(maxSubArray(nums))
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
