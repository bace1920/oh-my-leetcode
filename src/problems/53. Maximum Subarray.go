package main

import "fmt"

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	// nums := []int{3, 2, 1, 0, 4}
	// nums := []int{1, 1}

	fmt.Println(maxSubArray(nums))
}

func maxSubArray(nums []int) int {
	s, e, max := dc(nums, 0, len(nums))
	fmt.Println(s, e)
	return max
}

// todo
func dc(nums []int, s, e int) (int, int, int) {
	fmt.Println(s, e)
	if s+1 == e {
		return s, e, nums[s]
	}

	mid := (s + e) / 2
	ls, le, lsum := dc(nums, s, mid)
	rs, re, rsum := dc(nums, mid, e)
	if le == rs && lsum+rsum >= lsum && lsum+rsum >= rsum {
		return ls, re, lsum + rsum
	} else {
		if lsum > rsum {
			return ls, le, lsum
		} else {
			return rs, re, rsum
		}
	}
}

func maxSubArrayDP(nums []int) int {
	l := len(nums)
	iter := nums[0]
	sum := nums[0]
	for i := 1; i < l; i++ {
		iter = max(iter+nums[i], nums[i])
		sum = max(iter, sum)
	}
	return sum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
