package main

import "fmt"

func main() {
	var h, w int
	fmt.Scan(&h, &w)
	m := make([][]bool, h)
	dp := make([][]int, h)

	for i := 0; i < h; i++ {
		dp[i] = make([]int, w)
		m[i] = make([]bool, w)
		var s string
		fmt.Scan(&s)
		for j := 0; j < w; j++ {
			if s[j] == '#' {
				m[i][j] = true
				dp[i][j] = 1
			}
		}
	}

	for i := 1; i < w; i++ {
		dp[0][i] += dp[0][i-1]
		if m[0][i] && m[0][i-1] {
			dp[0][i]--
		}
	}
	for i := 1; i < h; i++ {
		dp[i][0] += dp[i-1][0]
		if m[i][0] && m[i-1][0] {
			dp[i][0]--
		}
	}
	for i := 1; i < h; i++ {
		for j := 1; j < w; j++ {
			x := dp[i][j] + dp[i-1][j]
			if m[i][j] && m[i-1][j] {
				x--
			}
			y := dp[i][j] + dp[i][j-1]
			if m[i][j] && m[i][j-1] {
				y--
			}
			if x > y {
				dp[i][j] = y
			} else {
				dp[i][j] = x
			}
		}
	}

	fmt.Println(dp[h-1][w-1])
}
