package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	var n int
	var s string
	fmt.Fscan(r, &n, &s)

	wc := make([]int, n+1)
	for i := 1; i <= n; i++ {
		wc[i] = wc[i-1]
		if s[i-1] == 'W' {
			wc[i]++
		}
	}
	red := n - wc[n]
	fmt.Println(wc[red])
}
