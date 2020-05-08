package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var n, a, b, c int
	fmt.Fscan(r, &n, &a, &b, &c)
	v := make(map[byte]int)
	v['A'], v['B'], v['C'] = a, b, c
	s := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &s[i])
	}

	ans := make([]byte, 0, n)

	choose := func(s string, i int) {
		v[s[i]]++
		v[s[i^1]]--
		ans = append(ans, s[i])
	}

	chooseAny := func(s string) {
		if v[s[0]] == 0 {
			choose(s, 0)
		} else {
			choose(s, 1)
		}
	}

	for i := 0; i < n; i++ {
		if v[s[i][0]] == 0 && v[s[i][1]] == 0 {
			fmt.Fprintln(w, "No")
			return
		}
		if i == n-1 || v[s[i][0]] == 0 || v[s[i][1]] == 0 {
			chooseAny(s[i])
		} else {
			if s[i][0] == s[i+1][0] || s[i][0] == s[i+1][1] {
				choose(s[i], 0)
			} else {
				choose(s[i], 1)
			}
		}
	}

	fmt.Fprintln(w, "Yes")
	for i := 0; i < n; i++ {
		fmt.Fprintf(w, "%c\n", ans[i])
	}
}
