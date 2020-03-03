package main

import (
	"container/heap"
	"fmt"
	"math"
)

// Skew heap

type Node struct {
	v           int
	left, right *Node
}

func Merge(x, y *Node) *Node {
	if x == nil {
		return y
	}
	if y == nil {
		return x
	}
	if x.v > y.v {
		x, y = y, x
	}

	x.right = Merge(x.right, y)
	x.right, x.left = x.left, x.right
	return x
}

func Push(x *Node, v int) *Node {
	y := &Node{v, nil, nil}
	return Merge(x, y)
}

func Pop(x *Node) {
	x = Merge(x.left, x.right)
}

func Second(x *Node) *Node {
	if x.right == nil {
		return x.left
	}
	if x.left.v < x.right.v {
		return x.left
	}
	return x.right
}

// Priority queue

type Elem struct {
	cost, index int
}

type PriorityQueue []*Elem

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].cost < pq[j].cost }
func (pq PriorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }

func (pq *PriorityQueue) Push(x interface{}) {
	p := x.(*Elem)
	*pq = append(*pq, p)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(*pq)
	x := old[n-1]
	old[n-1] = nil
	*pq = old[0 : n-1]
	return x
}

// Hu-Tucker

func fst(x *Node) int {
	if x == nil {
		return math.MaxInt32
	}
	return x.v
}

func snd(x *Node) int {
	if x == nil {
		return math.MaxInt32
	}
	return fst(Second(x))
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func HuTucker(x []int) int {
	n := len(x)
	hpq := make([]*Node, n)
	rig := make([]int, n)
	lef := make([]int, n)
	cst := make([]int, n)

	var pq PriorityQueue
	for i := 0; i < n-1; i++ {
		rig[i] = i + 1
		lef[i] = i - 1
		cst[i] = x[i] + x[i+1]
		heap.Push(&pq, &Elem{cst[i], i})
	}

	ans := 0
	for k := 0; k < n-1; k++ {
		e := heap.Pop(&pq).(*Elem)
		c := e.cost
		i := e.index

		for rig[i] == -1 || cst[i] != c {
			e = heap.Pop(&pq).(*Elem)
			c = e.cost
			i = e.index
		}

		var ml, mr bool
		if x[i]+fst(hpq[i]) == c {
			hpq[i] = Merge(hpq[i].left, hpq[i].right)
			ml = true
		} else if x[i]+x[rig[i]] == c {
			ml = true
			mr = true
		} else if fst(hpq[i])+snd(hpq[i]) == c {
			a := Merge(hpq[i].left, hpq[i].right)
			hpq[i] = Merge(a.left, a.right)
		} else {
			hpq[i] = Merge(hpq[i].left, hpq[i].right)
			mr = true
		}

		ans += c
		hpq[i] = Merge(hpq[i], &Node{c, nil, nil})

		if ml {
			x[i] = math.MaxInt32
		}
		if mr {
			x[rig[i]] = math.MaxInt32
		}

		if ml && i > 0 {
			j := lef[i]
			hpq[j] = Merge(hpq[j], hpq[i])
			rig[j] = rig[i]
			rig[i] = -1
			lef[rig[j]] = j
			i = j
		}
		if mr && rig[i]+1 < n {
			j := rig[i]
			hpq[i] = Merge(hpq[i], hpq[j])
			rig[i] = rig[j]
			rig[j] = -1
			lef[rig[i]] = i
		}

		cst[i] = x[i] + x[rig[i]]
		cst[i] = min(cst[i], min(x[i], x[rig[i]])+fst(hpq[i]))
		cst[i] = min(cst[i], fst(hpq[i])+snd(hpq[i]))
		heap.Push(&pq, &Elem{cst[i], i})
	}
	return ans
}

func main() {
	var n int
	fmt.Scan(&n)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	ret := HuTucker(a)
	fmt.Println(ret)
}
