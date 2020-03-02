package main

import "fmt"

func dfs(x, y, h, w int, graph [][]byte, visited [][]bool) bool {
	// fmt.Println(x, y)
	if x < 0 || x >= h || y < 0 || y >= w {
		return false
	}

	if visited[x][y] {
		return false
	}
	visited[x][y] = true

	n := graph[x][y]

	if n == 'g' {
		return true
	}

	if n == '#' {
		return false
	}

	for _, a := range [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
		nx := x + a[0]
		ny := y + a[1]
		if dfs(nx, ny, h, w, graph, visited) {
			return true
		}
	}
	return false
}

func main() {
	var h, w int
	fmt.Scan(&h, &w)

	var sx, sy int
	graph := make([][]byte, h)
	visited := make([][]bool, h)

	for i := 0; i < h; i++ {
		graph[i] = make([]byte, w)
		visited[i] = make([]bool, w)
		var s []byte
		fmt.Scan(&s)
		for j := 0; j < w; j++ {
			graph[i][j] = s[j]
			if s[j] == 's' {
				sx = i
				sy = j
			}
		}
	}

	// fmt.Println(sx, sy)
	ret := dfs(sx, sy, h, w, graph, visited)
	if ret {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
