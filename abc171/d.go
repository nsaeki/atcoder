package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	var n, q int
	fmt.Fscan(r, &n)
	a := make([]int64, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &a[i])
	}

	m := make(map[int64]int)
	var sum int64
	for i := 0; i < n; i++ {
		m[a[i]]++
		sum += a[i]
	}

	fmt.Fscan(r, &q)
	for ; q > 0; q-- {
		var b, c int64
		fmt.Fscan(r, &b, &c)
		sum += int64(m[b]) * (c - b)
		m[b], m[c] = 0, m[b]+m[c]
		fmt.Println(sum)
	}
}
