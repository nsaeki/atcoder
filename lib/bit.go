// BIT (A Binary indexed tree or Fenwick tree)
type BIT struct {
	n int
	d []int
}

func NewBIT(n int, a ...int) *BIT {
	b := BIT{}
	b.n = n
	b.d = make([]int, n+1)
	for i := 0; i < len(a); i++ {
		b.add(i, a[i])
	}
	return &b
}

func (b BIT) add(i, x int) {
	for i++; i <= b.n; i += i & -i {
		b.d[i] += x
	}
}

func (b BIT) sum(i int) int {
	x := 0
	for i++; i > 0; i -= i & -i {
		x += b.d[i]
	}
	return x
}
