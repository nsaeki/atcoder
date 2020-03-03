package main

import "fmt"

func pow(n, p, m int) int {
	if p == 0 {
		return 1
	}
	ret := pow(n, p/2, m)
	ret = ret * ret % m
	if p%2 == 1 {
		ret = ret * n % m
	}
	return ret
}

func main() {
	var n, m, p int
	fmt.Scan(&n, &m, &p)
	fmt.Println(pow(n, p, m))
}
