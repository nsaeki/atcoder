package main

import "fmt"

func dfs(i, rank int, g []int, r []int) {
	r[i] = append(r[i], i)
	for _, j := range g[i] {
		dfs(j, rank+1, g, r)
	}
}

func main() {
	var n int
	fmt.Scan(&n)

	g := make([][]int, n+1)
	//	c := make([][]int, n)

	for i := 0; i < n; i++ {
		var a, b int
		fmt.Scan(&a, &b)
		if a > b {
			a, b = b, a
		}
		g[a] = append(g[a], b)
	}

	r := make(map[int][]int)
	dfs(1, 1, g, r)

}
