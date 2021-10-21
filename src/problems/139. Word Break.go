package main

import "fmt"

func main() {
	d := []string{"leet", "code"}
	fmt.Println(wordBreak("leetcode", d))
}
func wordBreak(s string, wordDict []string) bool {

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
