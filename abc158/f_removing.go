package main

import (
	"fmt"
	"sort"
)

var factorialMemo [400000]int

func factorial(n, m int) int {
	if n <= 1 {
		return 1
	}
	if factorialMemo[n] != 0 {
		return factorialMemo[n]
	}
	ret := n * factorial(n-1, m) % m
	factorialMemo[n] = ret
	return ret
}

func pow(n, p, m int) int {
	if p == 0 {
		return 1
	}
	if p%2 == 0 {
		t := pow(n, p/2, m)
		return t * t % m
	}
	return n * pow(n, p-1, m) % m
}

func choose(n, k, m int) int {
	x := factorial(n, m) % m
	y := factorial(k, m) % m
	y *= factorial(n-k, m) % m
	y %= m
	return x * pow(y, m-2, m) % m
}

type pair struct {
	x, d, n int
	c       []*pair
}

type robots []*pair

func (r robots) Len() int           { return len(r) }
func (r robots) Less(i, j int) bool { return r[i].x < r[j].x }
func (r robots) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }

func size(r *pair) int {
	ret := 1
	for i := 0; i < r.n-1; i++ {
		ret += size(r.c[i])
	}
	return ret
}

func main() {
	const m int = 998244353
	var n int
	fmt.Scan(&n)

	r := make(robots, n)
	for i := 0; i < n; i++ {
		var rx, rd int
		fmt.Scan(&rx, &rd)
		r[i] = &pair{rx, rd, 1, []*pair{}}
	}

	sort.Sort(r)
	c := make([]int, n)
	for i := 0; i < n; i++ {
		c[i]++
		for j := i + 1; j < n; j++ {
			if r[i].x+r[i].d > r[j].x {
				r[i].c = append(r[i].c, r[j])
			}
		}
	}

	for i := 0; i < n; i++ {
		c[i] = size(r[i])
	}
	fmt.Println(c)

	res := 0
	for i := 0; i <= n; i++ {
		res += choose(n, i, m)
		res %= m
	}

	fmt.Println(res)
	for i := 0; i < n; i++ {
		for j := 0; j <= c[i]; j++ {
			res -= choose(c[i], j, m)
			res %= m
		}
	}

	fmt.Println(res)
}
