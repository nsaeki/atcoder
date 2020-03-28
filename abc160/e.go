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
	x, y, a, b, c := scanInt(), scanInt(), scanInt(), scanInt(), scanInt()
	p := scanInts(a)
	q := scanInts(b)
	r := scanInts(c)
	sort.Sort(sort.Reverse(sort.IntSlice(p)))
	sort.Sort(sort.Reverse(sort.IntSlice(q)))
	sort.Sort(sort.Reverse(sort.IntSlice(r)))

	d := make([]int, 0, x+y+c)
	d = append(d, p[:x]...)
	d = append(d, q[:y]...)
	d = append(d, r...)
	sort.Sort(sort.Reverse(sort.IntSlice(d)))

	ans := int64(0)
	for i := 0; i < x+y; i++ {
		ans += int64(d[i])
	}
	fmt.Println(ans)
}
