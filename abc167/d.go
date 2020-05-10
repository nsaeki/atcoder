package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n int
	var k int64
	rd := bufio.NewReader(os.Stdin)
	fmt.Fscan(rd, &n, &k)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(rd, &a[i])
		a[i]--
	}

	v := make([]bool, n)
	path := make([]int, 0, n)
	var s int
	for !v[s] {
		v[s] = true
		path = append(path, s)
		s = a[s]
	}

	var l int
	for i := 0; i < len(path); i++ {
		if path[i] == s {
			break
		}
		l++
	}

	var ans int
	if k < int64(l) {
		ans = path[int(k)]
	} else {
		k -= int64(l)
		k %= int64(len(path) - l)
		ans = path[l+int(k)]
	}
	fmt.Println(ans + 1)
}
