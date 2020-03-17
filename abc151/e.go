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

var factorialMemo []int64

func factorial(n, m int) int64 {
	if n <= 1 {
		return 1
	}
	if factorialMemo[n] != 0 {
		return factorialMemo[n]
	}
	ret := int64(n) * factorial(n-1, m) % int64(m)
	factorialMemo[n] = ret
	return ret
}

func combination(n, k, m int) int64 {
	x := factorial(n, m) % int64(m)
	y := factorial(k, m) % int64(m)
	y = y * factorial(n-k, m) % int64(m)
	return x * pow(y, int64(m-2), int64(m)) % int64(m)
}

func permutation(n, k, m int) int64 {
	return combination(n+k-1, n-1, m)
}

func pow(p, n, m int64) int64 {
	if n == 0 {
		return 1
	}
	if n&1 == 0 {
		r := pow(p, n/2, m)
		return r * r % int64(m)
	}
	return int64(p) * pow(p, n-1, m) % int64(m)
}

func main() {
	const m = 1000000007
	n, k := scanInt(), scanInt()
	a := scanInts(n)
	sort.Ints(a)

	factorialMemo = make([]int64, n)
	var sum int64
	for d := 0; d < 2; d++ {
		for i := k - 1; i < n; i++ {
			ret := combination(i, k-1, m)
			if d == 1 {
				ret = -ret
			}
			sum += ret * int64(a[i]) % int64(m)
			sum %= m
		}
		sort.Sort(sort.Reverse(sort.IntSlice(a)))
	}
	fmt.Println(sum)
}
