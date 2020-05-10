package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func newScanner() *bufio.Scanner {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	return scanner
}

func scanInt() int {
	sc.Scan()
	v, _ := strconv.Atoi(sc.Text())
	return v
}

func scanInts(n int) []int {
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = scanInt()
	}
	return a
}

func scanString() string {
	if sc.Scan() {
		return sc.Text()
	}
	panic(sc.Err())
}

var sc = newScanner()

func main() {
	n, k := scanInt(), scanInt()
	a := scanInts(n)
	for i := 0; i < n; i++ {
		a[i]--
	}

	visited := make([]bool, n)
	var s int
	for i := 0; !visited[i]; {
		visited[i] = true
		i, s = a[i], i
	}

	path := make([]int, 0, n)
	for i := 0; i != s; i = a[i] {
		path = append(path, i)
	}

	loop := make([]int, 0, n)
	loop = append(loop, s)
	for i := a[s]; i != s; i = a[i] {
		loop = append(loop, i)
	}

	var ans int
	if k < len(path) {
		ans = path[k]
	} else {
		k -= len(path)
		k %= len(loop)
		ans = loop[k]
	}
	fmt.Println(ans + 1)
}
