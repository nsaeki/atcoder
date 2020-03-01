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

func combinationFunc(n, m int) func(int, int) int {
	factorial := make([]int, n+1)
	factorialInv := make([]int, n+1)

	factorial[0] = 1
	for i := 1; i < n+1; i++ {
		factorial[i] = i * factorial[i-1] % m
	}

	factorialInv[n] = repeatedSquaring(factorial[n], m-2, m)
	for i := n; i > 0; i-- {
		factorialInv[i-1] = i * factorialInv[i] % m
	}

	return func(n, k int) int {
		if k < 0 || k > n {
			return 0
		}
		return (factorial[n] * factorialInv[k] % m) * factorialInv[n-k] % m
	}
}

func g(r, c int, k func(int, int) int) int {
	return k(r+c+2, c+1)
}

func main() {
	m := int(math.Pow10(9) + 7)
	n := 200003
	combination := combinationFunc(n, m)

	var r1, r2, c1, c2 int
	fmt.Scan(&r1, &c1, &r2, &c2)

	ret := g(r2, c2, combination) % m
	ret -= g(r1-1, c2, combination)
	ret -= g(r2, c1-1, combination)
	ret += g(r1-1, c1-1, combination)
	ret += m
	ret %= m
	fmt.Println(ret)
}
