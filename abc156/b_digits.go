package main

import (
	"fmt"
	"math"
)

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	for i := 0; ; i++ {
		if math.Pow(float64(k), float64(i)) > float64(n) {
			fmt.Println(i)
			break
		}
	}
}
