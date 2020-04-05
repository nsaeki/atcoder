package main

import "fmt"

func abs(a int64) int64 {
	if a < 0 {
		a = -a
	}
	return a
}

func main() {
	var n, k int64
	fmt.Scan(&n, &k)

	n %= k
	d := abs(n - k)
	if n >= d {
		n = d
	}
	fmt.Println(n)
}
