package main

import (
	"bufio"
	"fmt"
	"os"

	rbt "github.com/emirpasic/gods/trees/redblacktree"
)

func put(t *rbt.Tree, k interface{}) {
	if v, e := t.Get(k); e {
		t.Put(k, v.(int)+1)
	} else {
		t.Put(k, 1)
	}
}

func remove(t *rbt.Tree, k interface{}) {
	if v, e := t.Get(k); e {
		if v.(int) == 1 {
			t.Remove(k)
		} else {
			t.Put(k, v.(int)-1)
		}
	}
}

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
		put(t[b[i]], a[i])
	}

	mt := rbt.NewWithIntComparator()
	for i := 0; i < len(t); i++ {
		if e := t[i].Right(); e != nil {
			put(mt, e.Key)
		}
	}

	for ; q > 0; q-- {
		var c, d int
		fmt.Fscan(r, &c, &d)
		c--

		prev := b[c]
		if e := t[prev].Right(); e != nil {
			remove(mt, e.Key)
		}
		if e := t[d].Right(); e != nil {
			remove(mt, e.Key)
		}

		remove(t[prev], a[c])
		b[c] = d
		put(t[d], a[c])

		if e := t[prev].Right(); e != nil {
			put(mt, e.Key)
		}
		if e := t[d].Right(); e != nil {
			put(mt, e.Key)
		}

		fmt.Fprintln(w, mt.Left().Key)
	}
}
