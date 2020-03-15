package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	x := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Scan(&x[i])
	}
	sort.Ints(x)

	d := make([]int, m-1)
	for i := 0; i < m-1; i++ {
		d[i] = x[i+1] - x[i]
	}
	sort.Ints(d)

	ans := 0
	for i := 0; i < len(x)-n; i++ {
		ans += d[i]
	}
	fmt.Println(ans)
}
