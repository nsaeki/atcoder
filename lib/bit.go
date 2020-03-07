// A Binary indexed tree (Fenwick tree)

type BIT struct {
	var n int
	var d []int
}

func NewBIT(n int) &BIT {
	b := BIT{}
	b.n = n
	b.d = make([]int, n+1)
	for i := 0; i < n; i++ {
		b.d[i] = -1
	}
	return &b
}

func (b BIT) add(i, x int) {
	for i++; i <= b.n; i += i&-i {
		d[i] += x
	}
}

func (b BIT) sum(i int) int {
	x := 0
	for i++; i > 0; i -= i&-i {
		x += d[i]
	}
	return x
}