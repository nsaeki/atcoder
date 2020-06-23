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
	var n int
	fmt.Fscan(r, &n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &a[i])
	}

	var s int
	for i := 0; i < n; i++ {
		s ^= a[i]
	}

	for i := 0; i < n; i++ {
		fmt.Fprint(w, s^a[i])
		if i < n-1 {
			fmt.Fprint(w, " ")
		} else {
			fmt.Fprintln(w)
		}
	}
}
