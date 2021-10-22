package main

import "fmt"

func main() {
	// nums := []int{1, 2, 3, 4}
	// nums := []int{1, 2, 3}
	nums := []int{1, 2, 3, 8, 9, 10}
	fmt.Println(numberOfArithmeticSlices(nums))
}

func numberOfArithmeticSlices(nums []int) int {
	if len(nums) <= 2 {
		return 0
	}

	if len(nums) == 3 && nums[2]-nums[1] == nums[1]-nums[0] {
		return 1
	}

	total := 0
	diff := nums[1] - nums[0]
	s, e := 0, 1
	for i := 2; i < len(nums); i++ {
		if nums[i]-nums[i-1] != diff {
			if e-s >= 2 {
				total += (e - s) * (e - s - 1) / 2
			}
			s, e = i-1, i
			diff = nums[i] - nums[i-1]
		} else if i == len(nums)-1 && e-s >= 1 {
			e++
			total += (e - s) * (e - s - 1) / 2
		} else {
			e++
		}
	}
	return total
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
