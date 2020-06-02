package main

import "fmt"

func main() {
	var a, b int64
	var s string
	fmt.Scan(&a, &s)
	for i := 0; i < len(s); i++ {
		if s[i] == '.' {
			continue
		}
		b *= 10
		b += int64(s[i] - '0')
	}
	ans := a * b / 100
	fmt.Println(ans)
}
