package main

import "fmt"

func main() {
	var s, t, u string
	var a, b int

	fmt.Scan(&s, &t)
	fmt.Scan(&a, &b)
	fmt.Scan(&u)

	if u == s {
		fmt.Println(a-1, b)
	} else {
		fmt.Println(a, b-1)
	}
}
