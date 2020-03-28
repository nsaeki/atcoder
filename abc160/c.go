package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	k, n := scanInt(), scanInt()
	a := scanInts(n)
	sort.Ints(a)
	a = append(a, a[0])
	res := k
	for i := 1; i <= n; i++ {
		p := a[i-1]
		if p < a[i] {
			p += k
		}
		x := p - a[i]
		if res > x {
			res = x
		}
	}
	fmt.Println(res)
}
