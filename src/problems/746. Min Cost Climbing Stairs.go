package main

import (
	"fmt"
)

func main() {
	cost := []int{10, 15, 20}
	fmt.Println(minCostClimbingStairs(cost))
}

func minCostClimbingStairs(cost []int) int {
	cache1 := 0
	cache2 := 0

	for i := 2; i <= len(cost); i++ {
		cache1 = min(cache1+cost[i-2], cache2+cost[i-1])
		cache1, cache2 = cache2, cache1
	}

	return cache2
}

// 自顶向下搜索
func minCostClimbingStairsTopDown(cost []int) int {
	var climb func([]int, int) int

	cache := make([]int, len(cost)+1)
	for i := range cache {
		cache[i] = -1
	}

	cache[0] = 0
	cache[1] = 0

	climb = func(cost []int, target int) int {
		if cache[target] == -1 {
			cache[target] = min(climb(cost, target-1)+cost[target-1], climb(cost, target-2)+cost[target-2])
		}
		return cache[target]
	}
	return climb(cost, len(cost))
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
