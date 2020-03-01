package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	s := make([]int, m)
	c := make([]int, m)

	for i := 0; i < m; i++ {
		fmt.Scan(&s[i], &c[i])
	}

	ret := make([]int, n+1)
	for i := 0; i < n; i++ {
		ret[i] = -1
	}

	for i := 0; i < m; i++ {
		if n >= 2 && s[i] == 1 && c[i] == 0 {
			fmt.Println(-1)
			return
		}
		if ret[s[i]] > 0 && ret[s[i]] != c[i] {
			fmt.Println(-1)
			return
		}
		ret[s[i]] = c[i]
	}

	for i := 1; i < n+1; i++ {
		if ret[i] < 0 {
			if i == 1 {
				ret[i] = 1
			} else {
				ret[i] = 0
			}
		}
		fmt.Print(ret[i])
	}
}
