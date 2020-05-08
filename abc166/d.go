package main

import (
	"fmt"
)

func pow(n int) int64 {
	m := int64(n)
	return m * m * m * m * m
}

func f(a, b int) int64 {
	return pow(a) - pow(b)
}

func main() {
	var x int64
	fmt.Scan(&x)

	var maxInd int
	for maxInd = 1; f(maxInd, maxInd-1) <= x; maxInd++ {
	}
	for a := -maxInd; a < maxInd; a++ {
		for b := -maxInd; b < maxInd; b++ {
			if f(a, b) == x {
				fmt.Println(a, b)
				return
			}
		}
	}
}
