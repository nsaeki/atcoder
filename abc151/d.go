package main

import (
	"container/list"
	"fmt"
)

const maxCost = 401

func bfs(i, j int, m [][]byte, visited [][]bool) int {
	queue := list.New()
	next := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	queue.PushBack([]int{i, j, 0})
	reset(visited)

	maxCost := -1
	for queue.Len() > 0 {
		e := queue.Front()
		c := e.Value.([]int)
		queue.Remove(e)
		if m[c[0]][c[1]] == '#' || visited[c[0]][c[1]] {
			continue
		}

		visited[c[0]][c[1]] = true
		if maxCost < c[2] {
			maxCost = c[2]
		}

		for _, n := range next {
			ni, nj, nc := c[0]+n[0], c[1]+n[1], c[2]+1
			if ni < 0 || ni >= len(m) || nj < 0 || nj >= len(m[0]) {
				continue
			}
			queue.PushBack([]int{ni, nj, nc})
		}

	}
	return maxCost
}

func reset(v [][]bool) {
	for i := 0; i < len(v); i++ {
		for j := 0; j < len(v[0]); j++ {
			v[i][j] = false
		}
	}
}

func main() {
	var h, w int
	fmt.Scan(&h, &w)

	m := make([][]byte, h)
	v := make([][]bool, h)
	for i := 0; i < h; i++ {
		var s string
		fmt.Scan(&s)
		m[i] = make([]byte, w)
		v[i] = make([]bool, w)
		for j := 0; j < w; j++ {
			m[i][j] = s[j]
		}
	}

	ans := 0
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			c := bfs(i, j, m, v)
			if ans < c {
				ans = c
			}
		}
	}
	fmt.Println(ans)
}
