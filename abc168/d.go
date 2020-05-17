package main

import (
	"bufio"
	"fmt"
	"os"
)

var n, m int
var e [][]int
var stats []stat
var inf int = int(1e9)

type stat struct {
	prev, count int
}

func bfs() {
	for i := 0; i < n; i++ {
		stats[i] = stat{-1, inf}
	}
	stats[0] = stat{-1, 0}

	var q [][]int
	i := 0
	q = append(q, []int{0, 0})
	for i != len(q) {
		v, cnt := q[i][0], q[i][1]
		i++
		if stats[v].count != cnt {
			continue
		}
		for _, u := range e[v] {
			if cnt+1 < stats[u].count {
				stats[u] = stat{v, cnt + 1}
				q = append(q, []int{u, cnt + 1})
			}
		}
	}
}

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	fmt.Fscan(r, &n, &m)
	e = make([][]int, n)
	for i := 0; i < m; i++ {
		var u, v int
		fmt.Fscan(r, &u, &v)
		u--
		v--
		e[u] = append(e[u], v)
		e[v] = append(e[v], u)
	}

	stats = make([]stat, n)
	bfs()
	for i := 0; i < n; i++ {
		if stats[i].count == inf {
			fmt.Fprintln(w, "No")
			return
		}
	}
	fmt.Fprintln(w, "Yes")
	for i := 1; i < n; i++ {
		fmt.Fprintln(w, stats[i].prev+1)
	}
}
