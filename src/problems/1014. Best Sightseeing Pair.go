package main

import "fmt"

func main() {
	// nums := []int{1, 2, 3, 1}
	// nums := []int{2, 3, -2, 4}
	// nums := []int{0, 1, -2, -3, -4}
	// nums := []int{-16, 0, -5, 2, 2, -13, 11, 8}
	// nums := []int{8, 1, 5, 2, 6}
	// nums := []int{1, 3, 5}
	nums := []int{2, 7, 7, 2, 1, 7, 10, 4, 3, 3}
	// nums := []int{1, 1}
	fmt.Println(maxScoreSightseeingPair(nums))
}

// 不断记录前i个中对当前i最大的配对
func maxScoreSightseeingPair(values []int) int {
	big, cur := 0, 0
	for i := range values {
		big = max(cur+values[i], big)
		cur = max(cur, values[i]) - 1
	}
	return big
}

// 直接超时了
func maxScoreSightseeingPairBrute(values []int) int {
	score := 0
	for i := 0; i < len(values); i++ {
		for q := i + 1; q < len(values); q++ {
			score = max(values[q]+values[i]-q+i, score)
		}
	}
	return score
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
