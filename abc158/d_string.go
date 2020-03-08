package main

import "fmt"

func reverse(s []byte) {
	n := len(s)
	for i := 0; i < n/2; i++ {
		s[i], s[n-i-1] = s[n-i-1], s[i]
	}
}

func main() {
	var s string
	var q int
	fmt.Scan(&s, &q)

	head := []byte{}
	tail := []byte(s)

	for ; q > 0; q-- {
		var t int
		fmt.Scan(&t)
		if t == 1 {
			head, tail = tail, head
		} else {
			var f int
			fmt.Scan(&f, &s)

			if f == 1 {
				head = append(head, s[0])
			} else {
				tail = append(tail, s[0])
			}
		}
	}

	reverse(head)
	fmt.Print(string(head), string(tail))
	fmt.Println()
}
