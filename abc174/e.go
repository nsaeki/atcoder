package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	var n, k int
	fmt.Fscan(r, &n, &k)
	a := make([]int, n)
	var max int
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &a[i])
		if max < a[i] {
			max = a[i]
		}
	}

	ans := sort.Search(max, func(x int) bool {
		if x == 0 {
			return false
		}
		var cnt int
		for i := 0; i < n; i++ {
			if a[i] <= x {
				continue
			}
			cnt += a[i] / x
		}
		return cnt <= k
	})
	fmt.Println(ans)
}
