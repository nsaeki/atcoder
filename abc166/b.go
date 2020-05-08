package main

import "fmt"

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	s := make([]bool, n)
	for i := 0; i < k; i++ {
		var d int
		fmt.Scan(&d)
		for j := 0; j < d; j++ {
			var a int
			fmt.Scan(&a)
			a--
			s[a] = true
		}
	}

	var ans int
	for i := 0; i < n; i++ {
		if !s[i] {
			ans++
		}
	}

	fmt.Println(ans)
}
