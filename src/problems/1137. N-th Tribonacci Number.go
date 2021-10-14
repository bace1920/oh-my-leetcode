package main

import "fmt"

func main() {
	// memo := map[int]int{0: 0, 1: 1}
	// fmt.Println(memo[0])
	fmt.Println(tib(4))
}

func tib(n int) int {
	tib_list := make([]int, 38)
	for i := range tib_list {
		tib_list[i] = -1
	}
	tib_list[0] = 0
	tib_list[1] = 1
	tib_list[2] = 1
	if n <= 2 {
		return tib_list[n]
	}
	return do_tib(tib_list, n)
}

func do_tib(cache []int, n int) int {
	if cache[n] != -1 {
		return cache[n]
	} else {
		cache[n] = do_tib(cache, n-1) + do_tib(cache, n-2) + do_tib(cache, n-3)
		return cache[n]
	}
}
