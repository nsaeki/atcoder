package main

import "fmt"

func main() {
	var h, w, k int
	fmt.Scan(&h, &w, &k)
	var c [6][6]bool

	for i := 0; i < h; i++ {
		var s string
		fmt.Scan(&s)
		for j := 0; j < w; j++ {
			if s[j] == '#' {
				c[i][j] = true
			}
		}
	}

	ans := 0
	for i := 0; i < (1 << h); i++ {
		for j := 0; j < (1 << w); j++ {
			cnt := 0
			for x := 0; x < h; x++ {
				if (i>>x)&1 == 1 {
					continue
				}
				for y := 0; y < w; y++ {
					if c[x][y] && (j>>y)&1 == 0 {
						cnt++
					}
				}
			}
			if cnt == k {
				ans++
			}
		}
	}
	fmt.Println(ans)
}
