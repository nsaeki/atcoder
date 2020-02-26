package main

import (
	"fmt"
	"math"
)

var factorialMemo [400000]int

func factorial(n, m int) int {
	if n <= 1 {
		return 1
	}
	if factorialMemo[n] != 0 {
		return factorialMemo[n]
	}
	ret := n * factorial(n-1, m) % m
	factorialMemo[n] = ret
	return ret
}

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

func choose(n, k, m int) int {
	x := factorial(n, m) % m
	y := factorial(k, m) % m
	y *= factorial(n-k, m) % m
	y %= m
	return x * repeatedSquaring(y, m-2, m) % m
}

func multiChoose(n, k, m int) int {
	return choose(n+k-1, n-1, m)
}

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	divisor := int(math.Pow10(9) + 7)
	ret := 0
	for i := 0; i < n; i++ {
		if i > k {
			break
		}
		ret += choose(n, i, divisor) * multiChoose(n-i, i, divisor) % divisor
	}
	ret %= divisor
	fmt.Println(ret)
}
