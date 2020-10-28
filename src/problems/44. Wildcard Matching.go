package main

import "fmt"

func main() {
	//s := "mississippi"
	//p := "mis*is*p*."
	s := "aab"
	p := "c*a*b"
	fmt.Println(isMatch(s, p))
}

func isMatch(s string, p string) bool {
	slen, plen := len(s), len(p)

	startIndex, matchTo := -1, -1

	i, q := 0, 0
	for i < slen {
		if q < plen && (p[q] == '?' || s[i] == p[q]) {
			i++
			q++
		} else if q < plen && p[q] == '*' {
			startIndex = q
			matchTo = i
			q++
		} else if startIndex != -1 {
			q = startIndex + 1
			matchTo++
			i = matchTo

		} else {
			return false
		}
	}
	for ;q < plen;q++ {
		if p[q] != '*' {
			return false
		}
	}
	return true
}

func isMatchDp(s string, p string) bool {
	/*
	 dp[i][q] means if s[:i] match p[:q]
	*/
	slen, plen := len(s), len(p)
	if slen == 0 && plen == 0 {
		return true
	} else if plen == 0 {
		return false
	}

	dp := make([][]bool, slen+1)
	for i, _ := range dp {
		dp[i] = make([]bool, plen+1)
	}
	dp[0][0] = true

	if p[0] == '*' {
		dp[0][1] = true
	}
	for i:=1;i<plen;i++ {
		if p[i] == '*' && dp[0][i] {
			dp[0][i+1] = true
		}
	}



	for i := 0; i < slen; i++ {
		for q := 0; q < plen; q++ {
			if p[q] == '?' || p[q] == s[i] {
				dp[i+1][q+1] = dp[i][q]
			}

			if p[q] == '*' {
				if dp[i+1][q] {
					dp[i+1][q+1] = true
				} else {
					dp[i+1][q+1] = dp[i][q+1]
				}
			}
		}
	}

	return dp[slen][plen]
}