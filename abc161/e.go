package main

import (
	"fmt"
)

func main() {
	var n, k, c int
	var s string
	fmt.Scan(&n, &k, &c, &s)

	var i, j int
	a := make([]int, k)
	for i < n && j < k {
		if s[i] == 'o' {
			a[j] = i
			i += c
			j++
		}
		i++
	}

	b := make([]int, k)
	i, j = n-1, k-1
	for i >= 0 && j >= 0 {
		if s[i] == 'o' {
			b[j] = i
			i -= c
			j--
		}
		i--
	}

	for i := 0; i < k; i++ {
		if a[i] == b[i] {
			fmt.Println(a[i] + 1)
		}
	}
}
