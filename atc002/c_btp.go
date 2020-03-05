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

// Hu-Tucker alogorithm for building optimal alphabetic binary search trees
// https://scholarworks.rit.edu/theses/6484/
// https://www.cs.rit.edu/~std3246/thesis/node19.html

func HuTucker(w []int) int {
	n := len(w)

	// Huffman Priority Queues (HPQ)
	hpq := make([]*Node, n-1)

	// Left indices of Master Priority Queue Node
	lef := make([]int, n)

	// Right indices of Master Priority Queue (MPQ) Node
	rig := make([]int, n) // comment

	// Weights of Master Priority Queue Node
	cst := make([]int, n)

	// Master Priority Queue (MPQ)
	// Each element in the queue holds a pair value; sum of weight and HPQ Root index
	mpq := make(PriorityQueue, n-1)

	for i := 0; i < n-1; i++ {
		lef[i] = i - 1
		rig[i] = i + 1
		cst[i] = w[i] + w[i+1]
		mpq[i] = &Elem{cst[i], i}
	}
	heap.Init(&mpq)

	ans := 0
	for k := 0; k < n-1; k++ {

		// Find the first element in MPQ which can merges two HPQ nodes.
		e := heap.Pop(&mpq).(*Elem)
		c := e.cost
		i := e.index
		for rig[i] == -1 || cst[i] != c {
			e = heap.Pop(&mpq).(*Elem)
			c = e.cost
			i = e.index
		}

		// ?
		var ml, mr bool
		if w[i]+fst(hpq[i]) == c {
			hpq[i] = Merge(hpq[i].left, hpq[i].right)
			ml = true
		} else if w[i]+w[rig[i]] == c {
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
			w[i] = math.MaxInt32
		}
		if mr {
			w[rig[i]] = math.MaxInt32
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

		// Min of these:
		// w[i] + w[rig[i]]
		// min(w[i] , w[rig[i]]) + hqp[i] (if hpq[i] is not null)
		// hqp[i] + min(hpq[i].left, hpq[i].right) if hpq[i] has at least one child node
		cst[i] = w[i] + w[rig[i]]
		cst[i] = min(cst[i], min(w[i], w[rig[i]])+fst(hpq[i]))
		cst[i] = min(cst[i], fst(hpq[i])+snd(hpq[i]))
		heap.Push(&mpq, &Elem{cst[i], i})
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
