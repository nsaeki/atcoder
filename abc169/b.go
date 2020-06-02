package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int64, n)
	var zero bool
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
		if a[i] == 0 {
			zero = true
		}
	}

	if zero {
		fmt.Println(0)
		return
	}

	max := int64(1e18)
	ans := int64(1)
	for i := 0; i < n; i++ {
		if ans > max/a[i] {
			ans = -1
			break
		}
		ans *= a[i]
		if ans > max {
			ans = -1
			break
		}
	}
	fmt.Println(ans)
}
