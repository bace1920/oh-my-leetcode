package main

import "fmt"

func main() {
	memo := map[int]int{0:0, 1:1}
	fmt.Println(memo[])
	fmt.Println(fib(3))
}

func fib(n int) int {
	fib_list := make([]int, 31)
	for i := range fib_list {
		fib_list[i] = -1
	}
	fib_list[0] = 0
	fib_list[1] = 1
	if n <= 1 {
		return fib_list[n]
	}
	return do_fib(fib_list, n)
}

func do_fib(cache []int, n int) int {
	if cache[n] != -1 {
		return cache[n]
	} else {
		cache[n] = do_fib(cache, n-1) + do_fib(cache, n-2)
		return cache[n]
	}
}
