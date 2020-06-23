package main

import "fmt"

func main() {
	var a string
	fmt.Scan(&a)
	if a[0] >= 'A' && a[0] <= 'Z' {
		fmt.Println("A")
	} else {
		fmt.Println("a")
	}
}
