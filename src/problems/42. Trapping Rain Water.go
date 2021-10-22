package main

import (
	"fmt"
)

func main() {

	d := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	// d := []int{2, 0, 2}
	fmt.Println(trap(d))

}

func trap(height []int) int {
	total := 0
	s, e := 0, len(height)-1
	left, right := height[0], height[len(height)-1]
	for s < e {
		if height[s] > height[e] {
			e--
			total += max(0, min(left, right)-height[e])
			right = max(right, height[e])
		} else {
			s++
			total += max(0, min(left, right)-height[s])
			left = max(left, height[s])
		}
	}
	return total
}

func trapOld2(height []int) int {
	// 减少空间占用
	dp := make([]int, len(height))
	dp[0] = 0

	for i := 1; i < len(height); i++ {
		dp[i] = max(dp[i-1], height[i-1])

	}

	total := 0
	dp[len(height)-1] = 0
	for i := len(height) - 2; i >= 0; i-- {
		dp[i] = min(max(dp[i+1], height[i+1]), dp[i])
		total += max(dp[i]-height[i], 0)
	}

	return total
}

func trapOld(height []int) int {
	//dp 记录此位置左边最高高度
	dp := make([]int, len(height))
	dp[0] = 0
	//dp2 记录此位置右边最高高度
	dp2 := make([]int, len(height))
	dp2[len(height)-1] = 0

	for i := 1; i < len(height); i++ {
		dp[i] = max(dp[i-1], height[i-1])
		dp2[len(height)-1-i] = max(dp2[len(height)-i], height[len(height)-i])
	}

	total := 0
	for i := 0; i < len(height); i++ {
		total += max(min(dp[i], dp2[i])-height[i], 0)
	}

	return total
}

func trapBrute(height []int) int {
	cache := make([]int, 100000)
	h := 0
	total := 0
	for i := 0; i < len(height); i++ {
		// 已经围起来的部分吃下
		for q := 0; q < height[i]; q++ {
			total += cache[q]
			cache[q] = 0
		}

		// 此柱体上部分左边有墙体的部分算上
		for q := height[i]; q < h; q++ {
			cache[q]++
		}

		h = max(height[i], h)
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
