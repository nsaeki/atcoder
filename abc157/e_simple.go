package main

import (
	"fmt"
)

type SegTree struct {
	n int
	d []int
}

func NewSegTree(n int) *SegTree {
	s := SegTree{}
	s.n = n
	s.d = make([]int, 2*n)
	return &s
}

func (s *SegTree) Modify(i, x int) {
	i += s.n
	s.d[i] = x
	for ; i > 0; i >>= 1 {
		s.d[i>>1] = s.d[i] + s.d[i^1]
	}
}

func (s *SegTree) Query(l, r int) int {
	res := 0
	l += s.n
	r += s.n
	for l < r {
		if l&1 == 1 {
			res += s.d[l]
			l++
		}
		if r&1 == 1 {
			r--
			res += s.d[r]
		}
		l >>= 1
		r >>= 1
	}
	return res
}

func main() {
	var n, q int
	var str string
	fmt.Scan(&n, &str, &q)

	s := []byte(str)
	abc := [26]*SegTree{}

	for i := 0; i < 26; i++ {
		abc[i] = NewSegTree(n)
	}

	for i := 0; i < n; i++ {
		c := s[i] - 'a'
		abc[c].Modify(i, 1)
	}

	for k := 0; k < q; k++ {
		var t int
		fmt.Scan(&t)

		if t == 1 {
			var i int
			var a string
			fmt.Scan(&i, &a)
			i--

			abc[s[i]-'a'].Modify(i, 0)
			s[i] = a[0]
			abc[s[i]-'a'].Modify(i, 1)
		} else {
			var l, r int
			fmt.Scan(&l, &r)
			l--

			res := 0
			for i := 0; i < 26; i++ {
				if abc[i].Query(l, r) > 0 {
					res++
				}
			}
			fmt.Println(res)
		}

	}
}
