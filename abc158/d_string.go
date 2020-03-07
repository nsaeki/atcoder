package main

import "fmt"

func reverseCopy(r, s []byte) {
	n := len(s)
	for i := 0; i < n; i++ {
		r[i] = s[n-i-1]
	}
}

func main() {
	var s string
	var q int
	fmt.Scan(&s, &q)

	rev := false
	head := []byte{}
	tail := []byte{}
	body := []byte(s)

	for ; q > 0; q-- {
		var t int
		fmt.Scan(&t)
		if t == 1 {
			head, tail = tail, head
			rev = !rev
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

	// fmt.Println(string(head), string(body), string(tail))
	res := make([]byte, len(head))
	reverseCopy(res, head)
	if rev {
		tmp := make([]byte, len(body))
		reverseCopy(tmp, body)
		res = append(res, tmp...)
	} else {
		res = append(res, body...)
	}
	res = append(res, tail...)
	fmt.Println(string(res))
}
