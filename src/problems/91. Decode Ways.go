package main

import (
	"fmt"
	"strings"
)

func main() {
	// s := "26"
	s := "123123"
	// s := "226"
	// s := "2106"
	// s := "1123"
	fmt.Println(numDecodings(s))
}

func numDecodings(s string) int {
	if s[0] == '0' {
		return 0
	} else if len(s) == 1 {
		return 1
	}

	dp := make([]int, len(s))
	dp[0] = 1

	if s[1] == '0' {
		if strings.Contains("12", string(s[0])) {
			dp[1] = 1
		} else {
			return 0
		}
	} else if toInt(s[0])*10+toInt(s[1]) > 26 {
		dp[1] = 1
	} else {
		dp[1] = 2
	}

	for i := 2; i < len(s); i++ {
		if s[i] == '0' {
			if strings.Contains("12", string(s[i-1])) {
				dp[i] = dp[i-2]
			} else {
				return 0
			}
		} else if s[i-1] != '0' && toInt(s[i-1])*10+toInt(s[i]) <= 26 {
			dp[i] = dp[i-2]*2 + dp[i-1] - dp[i-2]
		} else if s[i-1] != '0' {
			dp[i] = dp[i-1]
		} else {
			dp[i] = dp[i-1]
		}
	}

	return dp[len(s)-1]
}

func toInt(ch uint8) int {
	return int(ch - 48)
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
