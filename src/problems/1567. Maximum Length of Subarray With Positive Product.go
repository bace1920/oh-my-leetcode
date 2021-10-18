package main

import "fmt"

func main() {
	// nums := []int{1, 2, 3, 1}
	// nums := []int{2, 3, -2, 4}
	// nums := []int{0, 1, -2, -3, -4}
	// nums := []int{-16, 0, -5, 2, 2, -13, 11, 8}
	nums := []int{-7, -10, -7, 21, 20, -12, -34, 26, 2}
	// nums := []int{-1, 2}
	// nums := []int{2, 3, -2, 4}
	// nums := []int{1, 1}
	fmt.Println(getMaxLen(nums))
}

func getMaxLen(nums []int) int {
	firstNegative := -1
	zeroPosition := -1
	sum := 0
	maxl := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] < 0 {
			sum++
			if firstNegative == -1 {
				firstNegative = i
			}
		}
		if nums[i] == 0 {
			firstNegative, zeroPosition, sum = -1, i, 0
		} else {
			if sum%2 == 0 {
				maxl = max(i-zeroPosition, maxl)
			} else {
				maxl = max(i-firstNegative, maxl)
			}
		}
	}

	return maxl
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
