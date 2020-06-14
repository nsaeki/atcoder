package main

import (
	"bufio"
	"fmt"
	"os"

	rbt "github.com/emirpasic/gods/trees/redblacktree"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var n, q int
	fmt.Fscan(r, &n, &q)
	a := make([]int, n)
	b := make([]int, n)
	var t [200001]*rbt.Tree
	for i := 0; i < len(t); i++ {
		t[i] = rbt.NewWithIntComparator()
	}
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &a[i], &b[i])
		b[i]--
		t[b[i]].Put(a[i], i)
	}

	mi := make([]bool, n)
	mt := rbt.NewWithIntComparator()
	for i := 0; i < len(t); i++ {
		if e := t[i].Right(); e != nil {
			mi[e.Value.(int)] = true
			mt.Put(e.Key, e.Value)
		}
	}

	for ; q > 0; q-- {
		var c, d int
		fmt.Fscan(r, &c, &d)
		c--
		d--

		prev := b[c]
		if e := t[prev].Right(); e != nil {
			mi[e.Value.(int)] = false
			mt.Remove(e.Key)
		}
		if e := t[d].Right(); e != nil {
			mi[e.Value.(int)] = false
			mt.Remove(e.Key)
		}

		t[prev].Remove(a[c])
		b[c] = d
		t[d].Put(a[c], c)

		if e := t[prev].Right(); e != nil {
			mi[e.Value.(int)] = true
			mt.Put(e.Key, e.Value)
		}
		if e := t[d].Right(); e != nil {
			mi[e.Value.(int)] = true
			mt.Put(e.Key, e.Value)
		}

		fmt.Fprintln(w, mt.Left().Key)
	}
}
