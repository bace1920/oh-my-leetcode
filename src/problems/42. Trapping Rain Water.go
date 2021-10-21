package main

import "fmt"

func main() {
	// d := []string{"car", "ca", "rs"}
	// fmt.Println(wordBreak("cars", d))
	d := []string{"aaaa", "aa"}
	fmt.Println(wordBreak("aaaaaaa", d))
	// d := []string{"leet", "code"}
	// fmt.Println(wordBreak("leetcode", d))
}
func wordBreak(s string, wordDict []string) bool {
	dp := make([]bool, len(s)+1)
	dp[0] = true

	for i := 1; i <= len(s); i++ {
		for j := i - 1; j >= 0; j-- {
			if dp[j] && contain(s, j, i, wordDict) {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)]
}

func contain(s string, start, end int, wordDict []string) bool {
	for i := 0; i < len(wordDict); i++ {
		word := wordDict[i]
		mapped := true

		if end-start != len(word) {
			continue
		}

		for q := range word {
			if word[q] != s[start+q] {
				mapped = false
				break
			}
		}

		if mapped {
			return true
		}
	}

	return false
}

// 超时
func wordBreakBrute(s string, wordDict []string) bool {
	return do(s, wordDict, 0)
}

func do(s string, wordDict []string, startAt int) bool {
	if startAt == len(s) {
		return true
	}

	for i := 0; i < len(wordDict); i++ {
		word := wordDict[i]
		mapped := true

		// 匹配开头
		for q := range word {
			if q+startAt >= len(s) || word[q] != s[q+startAt] {
				mapped = false
				break
			}
		}

		if mapped && do(s, wordDict, startAt+len(word)) {
			return true
		}
	}

	return false
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
