package main

import "fmt"

func main() {
	var k, q int
	fmt.Scan(&k, &q)

	d := make([]int, k)
	for i := 0; i < k; i++ {
		fmt.Scan(&d[i])
	}

	for i := 0; i < q; i++ {
		var n, x, m int
		fmt.Scan(&n, &x, &m)

		q := (n - 1) / k
		r := (n - 1) % k
		nz := 0
		al := x

		for j := 0; j < k; j++ {
			md := d[j] % m
			if md == 0 {
				nz += q
				if j < r {
					nz++
				}
			} else {
				al += q * md
				if j < r {
					al += md
				}
			}
		}

		ret := (n - 1) - nz - (al/m - x/m)
		fmt.Println(ret)
	}
}
