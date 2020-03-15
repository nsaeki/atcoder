package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
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
		if v, err := strconv.Atoi(t); err == nil {
			return v
		}
		panic(fmt.Sprintln("Could't scan int from input", t))
	}
	panic(s.Err())
}

type edge struct {
	id, to int
}

func dfs(f, t, p int, g [][]*edge) []int {
	if f == t {
		return []int{}
	}

	for _, e := range g[f] {
		if e.to == p {
			continue
		}
		if r := dfs(e.to, t, f, g); r != nil {
			return append(r, e.id)
		}
	}
	return nil
}

func main() {
	sc := newScanner()
	n := scanInt(sc)
	g := make([][]*edge, n)

	for i := 0; i < n-1; i++ {
		a, b := scanInt(sc), scanInt(sc)
		a--
		b--
		g[a] = append(g[a], &edge{b, i})
		g[b] = append(g[b], &edge{a, i})
	}
	m := scanInt(sc)
	eset := make([]uint64, m)
	for i := 0; i < m; i++ {
		a, b := scanInt(sc), scanInt(sc)
		a--
		b--
		for _, e := range dfs(a, b, -1, g) {
			eset[i] |= uint64(1) << e
		}
	}

	ans := uint64(0)
	for i := 0; i < 1<<m; i++ {
		mask := uint64(0)
		for j := 0; j < m; j++ {
			if i>>j&1 == 1 {
				mask |= eset[j]
			}
		}
		w := bits.OnesCount64(mask)
		now := uint64(1) << (n - 1 - w)
		if bits.OnesCount32(uint32(i))%2 == 0 {
			ans += now
		} else {
			ans -= now
		}
	}
	fmt.Println(ans)
}
