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

var sc *bufio.Scanner = newScanner()

func main() {
	n, m, q := scanInt(), scanInt(), scanInt()
	x := [501][501]int{}

	for i := 0; i < m; i++ {
		l, r := scanInt(), scanInt()
		x[l][r]++
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			x[i][j] += x[i][j-1]
		}
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			x[j][i] += x[j-1][i]
		}
	}

	for i := 0; i < q; i++ {
		p, q := scanInt(), scanInt()
		ans := x[q][q] - x[q][p-1] - x[p-1][q] + x[p-1][p-1]
		fmt.Println(ans)
	}
}
