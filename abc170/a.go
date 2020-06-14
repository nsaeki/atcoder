package main

import "fmt"

func main() {
	var ans int
	for i := 0; i < 5; i++ {
		var x int
		fmt.Scan(&x)
		if x == 0 {
			ans = i + 1
		}
	}
	fmt.Println(ans)

}
