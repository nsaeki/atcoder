package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	if n%2 == 0 {
		r := 1
		for i := 0; i < m; i++ {
			l := r + i + 1
			fmt.Println(r, l)
			r = l + 1
		}
	} else {
		for i := 1; i <= m; i++ {
			fmt.Println(i, n-i)
		}
	}
}
