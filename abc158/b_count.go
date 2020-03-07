package main

import "fmt"

func main() {
	var n, a, b int
	fmt.Scan(&n, &a, &b)

	res := n / (a + b) * a
	r := n % (a + b)
	if r > a {
		res += a
	} else {
		res += r
	}
	fmt.Println(res)
}
