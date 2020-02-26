package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)

	x := make([]int, n)
	sum := 0
	for i := 0; i < n; i++ {
		fmt.Scan(&x[i])
		sum += x[i]
	}

	avg := sum / n
	p := []int{avg - 1, avg, avg + 1}
	res := []int{0, 0, 0}

	for i := 0; i < n; i++ {
		for j := 0; j < len(p); j++ {
			res[j] += (x[i] - p[j]) * (x[i] - p[j])
		}
	}

	sort.Ints(res)
	fmt.Println(res[0])
}
