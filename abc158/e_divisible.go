package main

import "fmt"

func main() {
	var n, p int
	var s string
	fmt.Scan(&n, &p, &s)

	if p == 2 || p == 5 {
		res := 0
		for i := 0; i < n; i++ {
			if int(s[i]-'0')%p == 0 {
				res += i + 1
			}
		}
		fmt.Println(res)
		return
	}

	r := make([]int, n+1)
	d := 1
	for i := n - 1; i >= 0; i-- {
		x := int(s[i]-'0') * d % p
		r[i] = (x + r[i+1]) % p
		d *= 10
		d %= p
	}

	res := 0
	cnt := make(map[int]int)
	for i := n; i >= 0; i-- {
		res += cnt[r[i]]
		cnt[r[i]]++
	}

	fmt.Println(res)
}
