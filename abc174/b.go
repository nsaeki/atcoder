package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	var n, d int
	fmt.Fscan(r, &n, &d)

	dist := func(x, y int) float64 {
		return math.Sqrt(float64(x*x + y*y))
	}

	var ans int
	for i := 0; i < n; i++ {
		var x, y int
		fmt.Fscan(r, &x, &y)
		if dist(x, y) <= float64(d) {
			ans++
		}
	}
	fmt.Println(ans)
}
