package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	if len(s)&1 == 1 {
		fmt.Println("No")
		return
	}

	for i := 0; i < len(s)/2; i += 2 {
		if s[i] != 'h' || s[i+1] != 'i' {
			fmt.Println("No")
			return
		}
	}
	fmt.Println("Yes")
	return
}
