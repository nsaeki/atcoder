package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Frags [][]int

func (a Frags) Len() int           { return len(a) }
func (a Frags) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Frags) Less(i, j int) bool { return a[i][0] < a[j][0] }

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(reader, &n)

	var sum int
	var l, r Frags
	for i := 0; i < n; i++ {
		var s string
		fmt.Fscan(reader, &s)
		var d, min int
		for j := 0; j < len(s); j++ {
			if s[j] == '(' {
				d++
			} else {
				d--
			}
			if min > d {
				min = d
			}
		}
		if d > 0 {
			l = append(l, []int{min, d})
		} else {
			r = append(r, []int{min - d, -d})
		}
		sum += d
	}

	sort.Sort(sort.Reverse(l))
	sort.Sort(sort.Reverse(r))

	check := func(frags Frags) bool {
		var h int
		for _, f := range frags {
			if h+f[0] < 0 {
				return false
			}
			h += f[1]
		}
		return true
	}

	if sum == 0 && check(l) && check(r) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
