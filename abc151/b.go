package main

import "fmt"

func main() {
	var n, k, m int
	fmt.Scan(&n, &k, &m)
	ans := m * n
	for i := 0; i < n-1; i++ {
		var a int
		fmt.Scan(&a)
		ans -= a
	}
	if ans > k {
		fmt.Println(-1)
	} else if ans < 0 {
		fmt.Println(0)
	} else {
		fmt.Println(ans)
	}
}
