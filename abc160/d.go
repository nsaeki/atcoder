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

var sc = newScanner()

func abs(x int) int {
	if x < 0 {
		x = -x
	}
	return x
}

func min(a ...int) int {
	m := a[0]
	for i := 1; i < len(a); i++ {
		if m > a[i] {
			m = a[i]
		}
	}
	return m
}

func main() {
	n, x, y := scanInt(), scanInt(), scanInt()
	m := make(map[int]int)

	for i := 1; i <= n-1; i++ {
		for j := i + 1; j <= n; j++ {
			d := j - i
			d = min(d, abs(i-x)+1+abs(y-j))
			m[d]++
		}
	}

	for i := 1; i <= n-1; i++ {
		fmt.Println(m[i])
	}
}
