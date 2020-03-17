package main

import (
	"fmt"
)

func min(a ...int) int {
	r := a[0]
	for i := 1; i < len(a); i++ {
		if r > a[i] {
			r = a[i]
		}
	}
	return r
}

func max(a ...int) int {
	r := a[0]
	for i := 1; i < len(a); i++ {
		if r < a[i] {
			r = a[i]
		}
	}
	return r
}

func createMemo(s1, s2 string) []bool {
	n := len(s1)
	ret := make([]bool, n)
	for i := 0; i < n; i++ {
		ok := true
		for j := i; j < n; j++ {
			if j-i >= len(s2) {
				break
			}
			if s1[j] == '?' || s2[j-i] == '?' {
				continue
			}
			if s1[j] != s2[j-i] {
				ok = false
				break
			}
		}
		ret[i] = ok
	}
	return ret
}

func minStr(s ...string) int {
	const mx = 2001
	memo := [3][]bool{}
	memo[0] = createMemo(s[0], s[1])
	memo[1] = createMemo(s[1], s[2])
	memo[2] = createMemo(s[0], s[2])

	ret := mx * 3
	for i := 0; i <= mx; i++ {
		if i < len(s[0]) && !memo[0][i] {
			continue
		}
		for j := 0; j <= mx; j++ {
			if j < len(s[1]) && !memo[1][j] {
				continue
			}
			if i+j < len(s[0]) && !memo[2][i+j] {
				continue
			}
			l := max(len(s[0]), i+len(s[1]), i+j+len(s[2]))
			if ret > l {
				ret = l
			}
		}
	}
	return ret
}

func main() {
	var a, b, c string
	fmt.Scan(&a, &b, &c)
	fmt.Println(min(
		minStr(a, b, c),
		minStr(a, c, b),
		minStr(b, a, c),
		minStr(b, c, a),
		minStr(c, a, b),
		minStr(c, b, a)))
}
