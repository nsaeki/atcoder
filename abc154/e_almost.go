package main

import (
	"fmt"
)

func combination(n, k int) int {
	if n < k {
		return 0
	}
	x := 1
	y := 1
	for i := 0; i < k; i++ {
		x *= (n - i)
		y *= i + 1
	}
	return x / y
}

func main() {
	var s string
	var k int

	fmt.Scan(&s, &k)

	a := []int{}
	for i := 0; i < len(s); i++ {
		a = append(a, int(s[i]-'0'))
	}

	ret := 0

	if k == 1 {
		ret = (len(a) - 1) * 9
		ret += a[0]
	} else if k == 2 {
		ret = combination(len(a)-1, k) * 9 * 9
		ret += (a[0] - 1) * (len(a) - 1) * 9

		for i := 1; i < len(a); i++ {
			if a[i] > 0 {
				ret += a[i]
				ret += (len(a) - i - 1) * 9
				break
			}
		}
	} else if k == 3 {
		ret = combination(len(a)-1, k) * 9 * 9 * 9
		ret += (a[0] - 1) * combination(len(a)-1, k-1) * 9 * 9
		for i := 1; i < len(a); i++ {
			if a[i] > 0 {
				// i桁目が0
				ret += combination(len(a)-i-1, k-1) * 9 * 9
				// i桁目が最大値以下
				ret += (a[i] - 1) * combination(len(a)-i-1, k-2) * 9
				// i桁目が最大値
				for j := i + 1; j < len(a); j++ {
					if a[j] > 0 {
						ret += a[j]
						ret += (len(a) - j - 1) * 9
						break
					}
				}
				break
			}
		}
	}

	fmt.Println(ret)
}
