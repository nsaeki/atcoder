package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	var n, m int
	var k int64
	fmt.Fscan(r, &n, &m, &k)
	a := make([]int64, n+1)
	b := make([]int64, m+1)
	for i := 0; i < n; i++ {
		var tmp int64
		fmt.Fscan(r, &tmp)
		a[i+1] = a[i] + tmp
	}
	for i := 0; i < m; i++ {
		var tmp int64
		fmt.Fscan(r, &tmp)
		b[i+1] = b[i] + tmp
	}
	var ans int
	for i := 0; i <= n; i++ {
		if a[i] > k {
			continue
		}
		j := sort.Search(len(b), func(j int) bool { return a[i]+b[j] > k })
		if i+j-1 > ans {
			ans = i + j - 1
		}
	}
	fmt.Println(ans)
}
