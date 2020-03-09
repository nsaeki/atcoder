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
	p := make([]int, n)

	for i := 0; i < n; i++ {
		p[i] = scanInt(sc)
	}

	res := 0
	i := 0
	for j := 0; j < n; j++ {
		if p[i] >= p[j] {
			res++
			i = j
		}
	}

	fmt.Println(res)
}
