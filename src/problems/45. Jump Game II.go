package main

import "fmt"

func main() {
	nums := []int{2, 3, 1, 1, 4}
	// nums := []int{1, 1, 1, 1}
	// nums := []int{3, 2, 1, 0, 4}
	// nums := []int{1, 1}

	fmt.Println(jump(nums))
}

func jump(nums []int) int {
	length := len(nums)

	current_step := 0
	farest1 := 0
	farest2 := 0
	for i := 0; i < length ; i++ {
		if i > farest1 {
			farest1 = farest2
			current_step += 1
		}
		if i+nums[i] > farest2 {
			farest2 = i + nums[i]
		}
	}

	return current_step
}

// 可以发现其实一旦在某个点更新到了farest，那[0, farest]里的内容已经是最优解
// 换句话说从起点开始到 [i+1, farest] 所用的步数其实是一样的
// 所以我们继续搜索从[i+1, farest]出发，最远能到达的farest2是多少
// 显而易见从起点到[farest, farest2]的距离d2等于起点到[i+1, farest]的距离d1+1
// 那么我们只需要记录farest1， farest2即可，于是进一步删除了cache
func jump3(nums []int) int {
	length := len(nums)

	current_step := 0
	farest1 := 0 // current_step 所能到达的最远位置
	farest2 := 0 // // current_step + 1 所能到达的最远位置
	for i := 0; i < length; i++ {
		if i > farest1 {
			farest1 = farest2
			current_step += 1
		}
		if i+nums[i] > farest2 {
			farest2 = i + nums[i]
		}
	}

	return current_step
}

// 根据jump1的注释，这里去掉了不必要的初始值复制，也不用再比较大小了
func jump2(nums []int) int {
	length := len(nums)

	count := make([]int, length)

	count[0] = 0

	farest := 0
	for i := 0; i < length; i++ {
		if i+nums[i] > farest {
			for q := farest + 1; q <= i+nums[i] && q < length; q++ {
				count[q] = count[i] + 1
			}
			farest = i + nums[i]
		}
	}

	return count[length-1]
}

func jump1(nums []int) int {
	length := len(nums)

	count := make([]int, length)

	for i := range count {
		count[i] = 10001
	}
	count[0] = 0

	farest := 0
	// 实际上count[i]一定是大于等于count[0:i]任意一点的，
	// 所以如果i能到达的位置没有超过已经更新的最远位置farest的最大值，
	// 那他更新的距离只能是大于等于[i,farest]的任意一点
	// 就不必更新了
	// 或者说已经更新过的点肯定都已经是最小值了
	for i := 0; i < length; i++ {
		if i+nums[i] > farest {
			for q := farest + 1; q <= i+nums[i] && q < length; q++ {
				count[q] = min(count[i]+1, count[q])
			}
			farest = i + nums[i]
		}
	}

	return count[length-1]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
