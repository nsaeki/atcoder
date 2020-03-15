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

type city struct {
	p, id, year int
}

type cities []*city

func (p cities) Len() int { return len(p) }
func (p cities) Less(i, j int) bool {
	return p[i].p < p[j].p || (p[i].p == p[j].p && p[i].year < p[j].year)
}
func (p cities) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

func main() {
	_, m := scanInt(), scanInt()
	c := make(cities, m)

	for i := 0; i < m; i++ {
		p, y := scanInt(), scanInt()
		c[i] = &city{p, 0, y}
	}

	s := make(cities, m)
	copy(s, c)
	sort.Sort(c)

	c[0].id = 1
	for i := 1; i < m; i++ {
		if c[i-1].p == c[i].p {
			c[i].id = c[i-1].id + 1
		} else {
			c[i].id = 1
		}
	}

	for i := 0; i < m; i++ {
		fmt.Printf("%06d%06d\n", s[i].p, s[i].id)
	}
}
