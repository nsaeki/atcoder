package main

import "fmt"

func main() {
	var n int64
	fmt.Scan(&n)
	var ans []byte
	for n > 0 {
		b := byte(n%26) + 'a' - 1
		if b < 'a' {
			b = 'z'
			n -= 26
		}
		ans = append(ans, b)
		n /= 26
	}
	for i := len(ans) - 1; i >= 0; i-- {
		fmt.Printf("%c", ans[i])
	}
	fmt.Println()
}
