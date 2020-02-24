package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Alphabetic []string

func (a Alphabetic) Len() int {
	return len(a)
}

func (a Alphabetic) Less(i, j int) bool {
	return a[i] < a[j]
}

func (a Alphabetic) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func main() {
	n := 2 * 100000
	m := make(map[string]int)

	fmt.Scan(&n)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		s := scanner.Text()
		if _, ok := m[s]; ok {
			m[s]++
		} else {
			m[s] = 1
		}
	}

	maxCount := 0
	result := []string{}
	for k, v := range m {
		if maxCount > v {
			continue
		}

		if maxCount == v {
			result = append(result, k)
		} else {
			maxCount = v
			result = []string{k}
		}
	}

	sort.Sort(Alphabetic(result))
	for _, s := range result {
		fmt.Println(s)
	}
}
