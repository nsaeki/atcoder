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

var sc = newScanner()

func main() {
	n, k := int64(scanInt()), int64(scanInt())
	m := int64(1000000007)
	ans := int64(0)
	for i := k; i <= n+1; i++ {
		min, max := i*(i-1)/2, n*(n+1)/2-(n-i)*(n-i+1)/2
		ans += max - min + 1
		ans %= m
	}
	fmt.Println(ans)
}

