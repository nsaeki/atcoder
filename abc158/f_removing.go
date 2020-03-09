package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func newScanner() *bufio.Scanner {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	// buf := make([]byte, 10000)
	// scanner.Buffer(buf, 1000000)
	return scanner
}

func scanInt(s *bufio.Scanner) int {
	if s.Scan() {
		t := s.Text()
		v, err := strconv.Atoi(t)
		if err == nil {
			return v
		}
		panic(err)
	} else {
		panic(s.Err())
	}
}

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
	sc := newScanner()
	const m int = 998244353
	n := scanInt(sc)

	r := make(robots, n)
	for i := 0; i < n; i++ {
		x, d := scanInt(sc), scanInt(sc)
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
