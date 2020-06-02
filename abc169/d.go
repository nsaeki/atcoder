package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int64
	fmt.Scan(&n)

	m := make(map[int64]int)
	for i := int64(2); i*i <= n; i++ {
		if n == 0 {
			break
		}
		for n%i == 0 {
			m[i]++
			n /= i
		}
	}
	if n > 1 {
		m[n]++
	}

	var ans int
	for _, v := range m {
		x := sort.Search(50, func(i int) bool {
			return i*(i+1)/2 > v
		})
		ans += x - 1
	}
	fmt.Println(ans)
}
