// Segment Trees
// https://codeforces.com/blog/entry/18051

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