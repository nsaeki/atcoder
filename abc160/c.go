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

var sc = newScanner()

func main() {
	k, n := scanInt(), scanInt()
	a := scanInts(n)
	sort.Ints(a)
	a = append(a, a[0]+k)
	res := k
	for i := 0; i < n; i++ {
		x := k - (a[i+1] - a[i])
		if res > x {
			res = x
		}
	}
	fmt.Println(res)
}
