package main

import (
	"container/list"
	"fmt"
	"sort"
)

type node struct {
	x, d int
	c    []int
}

type robots []*node

func (r robots) Len() int           { return len(r) }
func (r robots) Less(i, j int) bool { return r[i].x < r[j].x }
func (r robots) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }

func dfs(i, m int, r []*node) int {
	res := 1
	for _, j := range r[i].c {
		res *= dfs(j, m, r)
		res %= m
	}
	return (res + 1) % m
}

func main() {
	const m int = 998244353
	var n int
	fmt.Scan(&n)

	r := make(robots, n)
	for i := 0; i < n; i++ {
		var x, d int
		fmt.Scan(&x, &d)
		r[i] = &node{x, d, []int{}}
	}
	sort.Sort(r)

	l := list.New()
	for i := n - 1; i >= 0; i-- {
		e := l.Front()
		for e != nil {
			a, _ := e.Value.([]int)
			x := a[0]
			j := a[1]
			if r[i].x+r[i].d > x {
				r[i].c = append(r[i].c, j)
				e1 := e
				e = e1.Next()
				l.Remove(e1)
			} else {
				break
			}
		}
		l.PushFront([]int{r[i].x, i})
	}

	res := 1
	for e := l.Front(); e != nil; e = e.Next() {
		a, _ := e.Value.([]int)
		i := a[1]
		res *= dfs(i, m, r)
		res %= m
	}

	fmt.Println(res)
}
