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
	var n, k int
	fmt.Fscan(r, &n, &k)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &a[i])
	}

	for i := 0; i < k; i++ {
		b := make([]int, n)
		for j := 0; j < n; j++ {
			for k := -a[j]; k <= a[j]; k++ {
				if j+k < 0 || j+k >= n {
					continue
				}
				b[j+k]++
			}
		}
		copy(a, b)
	}
	for i := 0; i < n-1; i++ {
		fmt.Fprint(w, a[i], " ")
	}
	fmt.Fprintln(w, a[n-1])
}
