package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	max := 0
	sum := 0
	for i := 0; i < n; i++ {
		var l int
		fmt.Scan(&l)
		if max < l {
			max = l
		}
		sum += l
	}
	if sum-max > max {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}

}
