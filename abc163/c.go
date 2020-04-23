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
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	n := scanInt()
	a := scanInts(n-1)
	c := make([]int, n)
	for i := 0; i < n-1; i++ {
		c[a[i]-1]++
	}
	for i := 0; i < n; i++ {
		fmt.Fprintln(w, c[i])
	}
}