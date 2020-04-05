package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	a := make([]int, n)
	sum := 0
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
		sum += a[i]
	}

	sort.Ints(a)
	if a[n-m] >= (sum+4*m-1)/(4*m) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
