package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	var ans int64
	g := func(x int) int64 {
		y := int64(n / x)
		return int64(x) * y * (y + 1) / 2
	}
	for i := 1; i <= n; i++ {
		ans += g(i)
	}
	fmt.Println(ans)
}
