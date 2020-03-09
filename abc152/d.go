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
		v, err := strconv.Atoi(t)
		if err == nil {
			return v
		}
		panic(err)
	} else {
		panic(s.Err())
	}
}

func main() {
	sc := newScanner()
	n := scanInt(sc)
	m := [10][10]int{}

	res := 0
	for i := 1; i <= n; i++ {
		e := i % 10
		if e == 0 {
			continue
		}

		s := i
		for s >= 10 {
			s /= 10
		}
		if s == 0 {
			continue
		}

		res += m[s][e] * 2
		if s == e {
			res++
		}
		m[e][s]++
	}

	fmt.Println(res)
}
