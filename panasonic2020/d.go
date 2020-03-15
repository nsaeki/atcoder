package main

import (
	"fmt"
)

func dfs(m byte, n int, s []byte) {
	if len(s) == n {
		fmt.Println(string(s))
		return
	}
	for i := byte('a'); i <= m; i++ {
		dfs(m, n, append(s, i))
	}
	m++
	dfs(m, n, append(s, m))
}

func main() {
	var n int
	fmt.Scan(&n)
	s := make([]byte, 0, n)
	s = append(s, 'a')
	dfs('a', n, s)
}
