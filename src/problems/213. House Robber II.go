package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 1}
	// nums := []int{1, 1}

	fmt.Println(rob(nums))
}

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}

	foo := func(nums []int) int {	

		cache1 := nums[0]
		cache2 := max(nums[0], nums[1])

		for i := 2; i < len(nums); i++ {
			cache1, cache2 = cache2, max(cache1+nums[i], cache2)
		}

		return cache2
	}
	return max(foo(nums[1:]), foo(nums[:len(nums)-1]))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
