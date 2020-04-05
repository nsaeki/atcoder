package main

import "fmt"

func divisors(n int64) []int64 {
	d := []int64{}
	for i := int64(1); i <= n/i; i++ {
		if n%i == 0 {
			d = append(d, i)
			if i != n/i {
				d = append(d, n/i)
			}
		}
	}
	return d
}

func f(n, k int64) bool {
	for n%k == 0 {
		n /= k
	}
	return n%k == 1
}

func main() {
	var n int64
	fmt.Scan(&n)
	ans := len(divisors(n-1)) - 1

	for _, x := range divisors(n) {
		if x == 1 {
			continue
		}
		if f(n, x) {
			ans++
		}
	}

	fmt.Println(ans)
}
