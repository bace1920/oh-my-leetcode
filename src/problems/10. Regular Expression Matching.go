package main

import "fmt"

func main() {
	//s := "mississippi"
	//p := "mis*is*p*."
	s := "aab"
	p := "c*a*b"
	fmt.Println(dp(s, p))
}

func isMatch(s string, p string) bool {
	return doRecursion(s, p, 0, 0)
}

func doRecursion(s string, p string, sl int, pl int) bool {
	slen, plen := len(s), len(p)
	if pl >= plen && sl >= slen {
		return true
	}

	firstMatch := false
	if sl < slen && pl < plen && (s[sl] == p[pl] || p[pl] == '.') {
		firstMatch = true
	}
	if plen-pl >= 2 && p[pl+1] == '*' {
		return (firstMatch && doRecursion(s, p, sl+1, pl)) || (doRecursion(s, p, sl, pl+2))
	} else {
		return firstMatch && doRecursion(s, p, sl+1, pl+1)
	}
}

func dp(s string, p string) bool {
	slen, plen := len(s), len(p)
	if slen == 0 && plen == 0 {
		return true
	}

	dp := make([][]bool, slen+1)
	for i, _ := range dp {
		dp[i] = make([]bool, plen+1)
	}
	dp[0][0] = true

	for i, char := range p {
		if char == '*' && dp[0][i-1] {
			dp[0][i+1] = true
		}
	}

	for i := 0; i < slen; i++ {
		for q := 0; q < plen; q++ {
			if p[q] == '.' || p[q] == s[i] {
				dp[i+1][q+1] = dp[i][q]
			}

			if p[q] == '*' {
				if p[q-1] != s[i] && p[q-1] != '.' {
					dp[i+1][q+1] = dp[i+1][q-1]
				} else {
					dp[i+1][q+1] = dp[i+1][q] || dp[i][q+1] || dp[i+1][q-1]
				}
			}
		}
	}

	return dp[slen][plen]
}
