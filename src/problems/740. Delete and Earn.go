package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{10, 8, 4, 2, 1, 3, 4, 8, 2, 9, 10, 4, 8, 5, 9, 1, 5, 1, 6, 8, 1, 1, 6, 7, 8, 9, 1, 7, 6, 8, 4, 5, 4, 1, 5, 9, 8, 6, 10, 6, 4, 3, 8, 4, 10, 8, 8, 10, 6, 4, 4, 4, 9, 6, 9, 10, 7, 1, 5, 3, 4, 4, 8, 1, 1, 2, 1, 4, 1, 1, 4, 9, 4, 7, 1, 5, 1, 10, 3, 5, 10, 3, 10, 2, 1, 10, 4, 1, 1, 4, 1, 2, 10, 9, 7, 10, 1, 2, 7, 5}
	// nums := []int{1, 1}

	fmt.Println(deleteAndEarn(nums))
}

func deleteAndEarn(nums []int) int {
	d := make(map[int]int)
	for i := range nums {
		if _, ok := d[nums[i]]; ok {
			d[nums[i]] += nums[i]
		} else {
			d[nums[i]] = nums[i]
		}
	}

	keys := make([]int, len(d))
	i := 0
	for k := range d {
		keys[i] = k
		i++
	}
	sort.Ints(keys)

	if len(keys) == 1 {
		return d[keys[0]]
	}
	k1 := keys[0]
	k2 := keys[1]
	v1 := d[k1]
	var v2 int
	if k2-k1 == 1 {
		v2 = d[k2]
	} else {
		v2 = v1 + d[k2]
	}

	for i := 2; i < len(keys); i++ {
		k := keys[i]
		v := d[k]
		// keep k>k2>k1
		// if k == k2+1 && k2 == k1+1 {
		if k == k1+2 {
			// 这里由于要记录k和v无法直接用max函数
			if v1+v > v2 && v2 > v1 { // 这里是兼容k1k2k3k4取k2 k4 的情况
				k1, k2 = k2, k
				v1, v2 = v2, v1+v
			} else if v1+v > v2 && v2 <= v1 { // 要注意兼容到k1 k2 k3 k4取 k1,k4的情况
				k1, k2 = k1, k
				v1, v2 = v1, v1+v
			}

		} else if k == k2+1 && k2 != k1+1 { // 都不相连直接push就好
			k1, k2 = k2, k
			v1, v2 = v2, v1+v
		} else if k != k2+1 { // 前置2个相连，那么只要选一个和当前k结合，但是要注意保留大那个
			if v2 > v1 {
				k1, k2 = k2, k
				v1, v2 = v2, v2+v
			} else {
				k1, k2 = k1, k
				v1, v2 = v1, v1+v
			}
		}
	}

	return max(v1, v2)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// func arr_max(arr []int) int {
// 	max := math.MinInt
// 	for i := range arr {
// 		if arr[i] > max {
// 			max = arr[i]
// 		}
// 	}
// 	return max
// }
