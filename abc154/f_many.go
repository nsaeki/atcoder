package main

import "fmt"

func main() {
	var r1, c1, r2, c2 int
	fmt.Scan(&r1, &c1, &r2, &c2)

	for i := 0; i < count; i++ {
		countPath(i, i)
	}
}
