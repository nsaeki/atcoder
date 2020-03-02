package main

import "fmt"

type UnionFind struct {
	root []int
	rank []int
}

func NewUnionFind(x int) *UnionFind {
	u := UnionFind{make([]int, x), make([]int, x)}
	for i := 0; i < x; i++ {
		u.root[i] = i
		u.rank[i] = 0
	}
	return &u
}

func (u UnionFind) Root(x int) int {
	if u.root[x] == x {
		return x
	}
	u.root[x] = u.Root(u.root[x])
	return u.root[x]
}

func (u UnionFind) Same(x, y int) bool {
	return u.Root(x) == u.Root(y)
}

func (u UnionFind) Unite(x, y int) {
	x = u.Root(x)
	y = u.Root(y)
	if x == y {
		return
	}
	if u.rank[x] < u.rank[y] {
		u.root[x] = y
	} else {
		u.root[y] = x
		if u.rank[x] == u.rank[y] {
			u.rank[x]++
		}
	}
}

func main() {
	var n, q int
	fmt.Scan(&n, &q)

	uf := NewUnionFind(n)

	for i := 0; i < q; i++ {
		var p, a, b int
		fmt.Scan(&p, &a, &b)
		if p == 0 {
			uf.Unite(a, b)
		} else {
			ret := uf.Same(a, b)
			if ret {
				fmt.Println("Yes")
			} else {
				fmt.Println("No")
			}
		}

	}
}
