package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)
	ans := 0
	for x >= 500 {
		x -= 500
		ans += 1000
	}
	for x >= 5 {
		x -= 5
		ans += 5
	}
	fmt.Println(ans)
}
