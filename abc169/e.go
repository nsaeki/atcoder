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
	b := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &a[i], &b[i])
	}

	sort.Ints(a)
	sort.Ints(b)
	if n%2 == 0 {
		ans := (b[n/2] + b[n/2-1]) - (a[n/2] + a[n/2-1]) + 1
		fmt.Println(ans)
	} else {
		fmt.Println(b[n/2] - a[n/2] + 1)
	}
}
