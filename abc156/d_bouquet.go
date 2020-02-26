package main

import (
	"fmt"
	"math"
)

func repeatedSquaring(n, p, m int) int {
	if p == 0 {
		return 1
	}
	if p%2 == 0 {
		t := repeatedSquaring(n, p/2, m)
		return t * t % m
	}
	return n * repeatedSquaring(n, p-1, m) % m
}

func combination(n, a, m int) int {
	x := 1
	y := 1
	for i := 0; i < a; i++ {
		x = x * (n - i) % m
		y = y * (i + 1) % m
	}

	return x * repeatedSquaring(y, m-2, m) % m
}

func main() {
	var n, a, b int
	fmt.Scan(&n, &a, &b)

	divisor := int(math.Pow10(9) + 7)
	all := repeatedSquaring(2, n, divisor) - 1
	all -= combination(n, a, divisor)
	all += divisor
	all %= divisor
	all -= combination(n, b, divisor)
	all += divisor
	all %= divisor
	fmt.Println(all)
}
