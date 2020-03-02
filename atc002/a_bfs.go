package main

import (
	"container/list"
	"fmt"
)

func bfs(sx, sy, gx, gy, c, r int, maze [][]byte, steps [][]int) int {
	queue := list.New()
	queue.PushBack([]int{sx, sy, 0})

	for queue.Len() > 0 {
		var x, y, s int
		e := queue.Front()
		queue.Remove(e)
		if v, ok := e.Value.([]int); ok {
			x = v[0]
			y = v[1]
			s = v[2]
		}

		if x == gx && y == gy {
			return s
		}

		if x < 0 || x >= c || y < 0 || y >= r {
			continue
		}

		if maze[y][x] == '#' {
			continue
		}

		if s >= steps[y][x] {
			continue
		}

		steps[y][x] = s

		for _, a := range [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
			queue.PushBack([]int{x + a[0], y + a[1], s + 1})
		}
	}

	return steps[gy][gx]
}

func main() {
	var r, c int
	fmt.Scan(&r, &c)

	var sy, sx, gy, gx int
	fmt.Scan(&sy, &sx, &gy, &gx)

	maze := make([][]byte, r)
	steps := make([][]int, r)
	max := r * c

	for i := 0; i < r; i++ {
		maze[i] = make([]byte, c)
		steps[i] = make([]int, c)
		var s []byte
		fmt.Scan(&s)
		for j := 0; j < c; j++ {
			maze[i][j] = s[j]
			steps[i][j] = max
		}
	}

	ret := bfs(sx-1, sy-1, gx-1, gy-1, c, r, maze, steps)
	fmt.Println(ret)
}
