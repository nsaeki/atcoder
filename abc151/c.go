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
	// buf := make([]byte, 10000)
	// scanner.Buffer(buf, 1000000)
	return scanner
}

func scanInt() int {
	if sc.Scan() {
		t := sc.Text()
		if v, err := strconv.Atoi(t); err == nil {
			return v
		}
		panic(fmt.Sprintln("Could't scan int from input", t))
	}
	panic(sc.Err())
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

var sc *bufio.Scanner = newScanner()

func main() {
	n, m := scanInt(), scanInt()
	ac := make(map[int]int)
	wa := make(map[int]int)

	for i := 0; i < m; i++ {
		p, s := scanInt(), scanString()
		if s == "AC" {
			ac[p] = 1
		} else if ac[p] == 0 {
			wa[p]++
		}
	}

	ans := [2]int{}
	for i := 1; i <= n; i++ {
		ans[0] += ac[i]
		if ac[i] == 1 {
			ans[1] += wa[i]
		}
	}
	fmt.Println(ans[0], ans[1])
}
