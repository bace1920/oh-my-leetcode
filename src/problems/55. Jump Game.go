package main

import "fmt"

func main() {
	nums := []int{2, 3, 1, 1, 4}
	// nums := []int{3, 2, 1, 0, 4}
	// nums := []int{1, 1}

	fmt.Println(canJump(nums))
}

// 此题不需要知道路径，只要顺序搜索尝试扩大最远可达位置即可
func canJump(nums []int) bool {
	canReach := 0
	if len(nums) == 1 {
		return true
	}
	for i := 0; i <= canReach; i++ {
		if i+nums[i] >= canReach {
			canReach = i + nums[i]
		}
		if canReach >= len(nums)-1 {
			return true
		}

	}

	return false
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
