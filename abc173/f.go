package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(r, &n)
	e := make([][]int, n-1)
	for i := 0; i < n-1; i++ {
		var u, v int
		fmt.Fscan(r, &u, &v)
		if u > v {
			u, v = v, u
		}
		e[i] = []int{u, v}
	}

	var ans int64
	for i := 1; i <= n; i++ {
		ans += int64(i) * int64(n-i+1)
	}
	for i := 0; i < n-1; i++ {
		ans -= int64(e[i][0]) * int64(n-e[i][1]+1)
	}
	fmt.Println(ans)
}
