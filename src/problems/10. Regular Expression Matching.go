package main

import "fmt"

func main() {
	//s := "mississippi"
	//p := "mis*is*p*."
	s := "aa"
	p := "a*"
	fmt.Println(isMatch(s, p))
}

func isMatch(s string, p string) bool {
	return do(s, p, 0, 0)
}

func do(s string, p string, sl int, pl int) bool {
	slen, plen := len(s), len(p)
	if pl >= plen && sl >= slen {
		return true
	}

	firstMatch := false
	if sl < slen && pl < plen && (s[sl] == p[pl] || p[pl] == '.') {
		firstMatch = true
	}
	if plen-pl >= 2 && p[pl+1] == '*' {
		return (firstMatch && do(s, p, sl+1, pl)) || (do(s, p, sl, pl+2))
	} else {
		return firstMatch && do(s, p, sl+1, pl+1)
	}
}
