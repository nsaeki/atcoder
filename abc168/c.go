package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, h, m int
	fmt.Scan(&a, &b, &h, &m)

	radH := (float64(h) + float64(m)/60) * 2 * math.Pi / 12
	radM := float64(m) * 2 * math.Pi / 60
	rad := math.Abs(radH - radM)
	ans := float64(a*a) + float64(b*b) - 2*float64(a*b)*math.Cos(rad)
	ans = math.Sqrt(ans)
	fmt.Println(ans)
}
