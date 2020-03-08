package main

import (
	"fmt"
)

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	xmin := a * 125
	xmax := (a + 1) * 125
	ymin := b * 100
	ymax := (b + 1) * 100

	fmt.Println(xmin, xmax, ymin, ymax)

	if xmin >= ymin && xmin < ymax {
		fmt.Println((xmin + 9) / 10)
	} else if ymin >= xmin && ymin < xmax {
		fmt.Println((ymin + 9) / 10)
	} else {
		fmt.Println(-1)
	}
}
