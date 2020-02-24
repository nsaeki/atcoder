package main

import (
	"fmt"
	"sort"
)

type Pair struct {
	X int
	Y int
}

type PairArray []Pair
type ByX []Pair

func (a ByX) Len() int           { return len(a) }
func (a ByX) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByX) Less(i, j int) bool { return a[i].X < a[j].X }

var (
	x    []int
	g    []PairArray
	used []bool
	ans  []int
)

func dfs(v int) int {
	used[v] = true
	res := x[v]
	for _, e := range g[v] {
		if used[e.X] {
			continue
		}
		r := dfs(e.X)
		if r == 1 {
			ans = append(ans, e.Y)
		}
		res ^= r
	}
	return res
}

func main() {
	var n, m int
	a := []Pair{}

	fmt.Scan(&n, &m)
	for i := 0; i < n; i++ {
		var x, y int
		fmt.Scan(&x, &y)
		a = append(a, Pair{x, y})
	}

	sort.Sort(ByX(a))

	// 差分列を作る
	x = make([]int, n+1)
	x[0] = a[0].Y
	for i := 0; i < n-1; i++ {
		x[i+1] = a[i].Y ^ a[i+1].Y
	}
	x[n] = a[n-1].Y

	// グラフを作成
	// g[n+1][]: 差分列の点からつながっている差分列の別の点。ペアの2つ目はコードの番号
	g = make([]PairArray, n+1)
	for i := 0; i < m; i++ {
		var l, r int
		fmt.Scan(&l, &r)

		l = sort.Search(n, func(j int) bool { return a[j].X >= l })
		r = sort.Search(n, func(j int) bool { return a[j].X > r })

		// fmt.Println(l, r)
		g[l] = append(g[l], Pair{r, i + 1})
		g[r] = append(g[r], Pair{l, i + 1})
	}

	used = make([]bool, n+1)

	for i := 0; i < n; i++ {
		if used[i] {
			continue
		}
		var r int
		r = dfs(i)
		if r != 0 {
			fmt.Println(-1)
			return
		}
	}

	fmt.Println(len(ans))
	sort.Ints(ans)
	for _, x := range ans {
		fmt.Print(x, " ")
	}
	fmt.Println()
}
