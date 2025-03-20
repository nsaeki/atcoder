package lib

type UnionFind struct {
	root []int
	size []int
}

func NewUnionFind(x int) *UnionFind {
	u := UnionFind{make([]int, x), make([]int, x)}
	for i := 0; i < x; i++ {
		u.root[i] = i
		u.size[i] = 1
	}
	return &u
}

func (u *UnionFind) Root(x int) int {
	if u.root[x] == x {
		return x
	}
	u.root[x] = u.Root(u.root[x])
	return u.root[x]
}

func (u *UnionFind) Same(x, y int) bool {
	return u.Root(x) == u.Root(y)
}

func (u *UnionFind) Unite(x, y int) {
	x = u.Root(x)
	y = u.Root(y)
	if x == y {
		return
	}
	if u.size[x] > u.size[y] {
		x, y = y, x
	}
	u.size[y] += u.size[x]
	u.root[x] = y
}

func (u *UnionFind) Size(x int) int {
	return u.size[u.Root(x)]
}
