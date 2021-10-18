package main

import "fmt"

func main() {
	// nums := []int{1, 2, 3, 1}
	// nums := []int{2, 3, -2, 4}
	nums := []int{4, 0, 2}
	// nums := []int{-2, 3, -4}
	// nums := []int{2, 3, -2, 4}
	// nums := []int{1, 1}
	fmt.Println(maxProduct(nums))
}

func maxProduct(nums []int) int {
	pmax := nums[0]
	pmin := nums[0]
	prod := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < 0 {
			pmax, pmin = pmin, pmax
		}
		pmax = max(nums[i], pmax*nums[i])
		pmin = min(nums[i], pmin*nums[i])
		prod = max(prod, pmax)
	}
	return prod
}

// 考虑0分割，当负数个数为偶数时只要看0分割后的子串
// 负数为奇数个时只能从左乘到最右边的负数，或反之
func maxProductOld(nums []int) int {

	return do(nums, 0, len(nums))
}

func do(nums []int, s, e int) int {
	if s == e {
		return 0
	}
	if s+1 == e {
		return nums[s]
	}

	for i := s; i < e; i++ {
		if nums[i] == 0 {
			return max(0, max(do(nums, s, i), do(nums, i+1, e)))
		}
	}

	prod := nums[s]
	for i := s + 1; i < e; i++ {
		prod *= nums[i]
	}

	if prod < 0 {
		pl := prod
		pr := prod
		for i := e - 1; i >= s; i-- {
			pr /= nums[i]
			if pr > 0 {
				break
			}
		}
		for i := s; i < e; i++ {
			pl /= nums[i]
			if pl > 0 {
				break
			}
		}
		return max(pr, pl)
	} else {
		return prod
	}
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
