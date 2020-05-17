package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	d := n % 10
	var ans string
	if d == 3 {
		ans = "bon"
	} else if d == 0 || d == 1 || d == 6 || d == 8 {
		ans = "pon"
	} else {
		ans = "hon"
	}
	fmt.Println(ans)
}
