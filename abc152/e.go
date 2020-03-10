package main

import (
	"bufio"
	"fmt"
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

func genPrimes(n int) []int {
	p := make([]int, n)
	for i := 2; i < n; i++ {
		if p[i] == 0 {
			p[i] = i
		}
		for j := i * i; j < n; j += i {
			if p[j] == 0 {
				p[j] = i
			}
		}
	}
	return p
}

func pow(p, n, m int) int {
	if n == 0 {
		return 1
	}

	r := pow(p, n/2, m) % m
	r *= r
	r %= m
	if n&1 == 1 {
		r *= p
		r %= m
	}
	return r
}

func main() {
	const m int = 1000000007
	sc := newScanner()
	n := scanInt(sc)
	a := make([]int, n)
	maxA := 0
	for i := 0; i < n; i++ {
		a[i] = scanInt(sc)
		if a[i] > maxA {
			maxA = a[i]
		}
	}

	p := genPrimes(maxA + 1)
	f := make(map[int]int)
	for i := 0; i < n; i++ {
		fa := make(map[int]int)
		x := a[i]
		for x > 1 && p[x] > 0 {
			fa[p[x]]++
			x /= p[x]
		}
		for k, v := range fa {
			if f[k] < v {
				f[k] = v
			}
		}
	}

	lcm := 1
	for k, v := range f {
		lcm *= pow(k, v, m)
		lcm %= m
	}

	res := 0
	for i := 0; i < n; i++ {
		res += lcm * pow(a[i], m-2, m) % m
		res %= m
	}

	fmt.Println(res)
}
