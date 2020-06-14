package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(r, &n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &a[i])
	}
	sort.Ints(a)
	max := a[n-1]
	b := make(map[int]int)
	for i := 0; i < n; i++ {
		b[a[i]]++
	}

	var ans int
	for i := 0; i < n; i++ {
		if b[a[i]] == 1 {
			ans++
		}
		for j := a[i] * 2; j <= max; j += a[i] {
			b[j] = 0
		}
	}
	fmt.Println(ans)
}
