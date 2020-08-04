package main

import (
	"bufio"
	"fmt"
	"os"
)

// BIT (A Binary indexed tree or Fenwick tree)
type BIT struct {
	n int
	d []int
}

func NewBIT(n int, a ...int) *BIT {
	b := BIT{}
	b.n = n
	b.d = make([]int, n+1)
	for i := 0; i < len(a); i++ {
		b.Add(i, a[i])
	}
	return &b
}

func (b BIT) Add(i, x int) {
	for i++; i <= b.n; i += i & -i {
		b.d[i] += x
	}
}

func (b BIT) Sum(i int) int {
	x := 0
	for i++; i > 0; i -= i & -i {
		x += b.d[i]
	}
	return x
}

type query struct {
	l, r, i int
}

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var n, q int
	fmt.Fscan(r, &n, &q)
	c := make([]int, n+1)
	qm := make(map[int][]query)

	for i := 1; i <= n; i++ {
		fmt.Fscan(r, &c[i])
	}
	for i := 0; i < q; i++ {
		var x, y int
		fmt.Fscan(r, &x, &y)
		qm[y] = append(qm[y], query{x, y, i})
	}

	ans := make([]int, q)
	ci := make([]int, n+1)
	b := NewBIT(n + 1)
	for i := 1; i <= n; i++ {
		if ci[c[i]] > 0 {
			b.Add(ci[c[i]], -1)
		}
		ci[c[i]] = i
		b.Add(i, 1)
		for _, v := range qm[i] {
			ans[v.i] = b.Sum(v.r) - b.Sum(v.l-1)
		}
	}
	for i := 0; i < q; i++ {
		fmt.Fprintln(w, ans[i])
	}
}
