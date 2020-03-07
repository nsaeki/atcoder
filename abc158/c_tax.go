package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	xmin := a * 125
	xmax := (a + 1) * 125
	ymin := b * 100
	ymax := (b + 1) * 100

	for x := xmin; x < xmax; x++ {
		if x >= ymin && x < ymax {
			res := x / 10
			if x%10 > 0 {
				res++
			}
			fmt.Println(res)
			return
		}
	}
	fmt.Println(-1)
}
