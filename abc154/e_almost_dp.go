package main

import (
	"fmt"
)

func main() {
	var s string
	var k int

	fmt.Scan(&s, &k)

	n := len(s)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = int(s[i] - '0')
	}

	/*
		dp[i][j][ok]
		i桁目まで確定
		j個の非0を使用済み
		ok: この時点でN以下であることが確定しているかどうか
	*/
	dp := [101][4][2]int{}
	dp[0][0][0] = 1

	for i := 0; i < n; i++ {
		for j := 0; j < 4; j++ {
			for ok := 0; ok < 2; ok++ {
				for d := 0; d <= 9; d++ {
					ni, nj, nok := i+1, j, ok
					if d != 0 {
						nj++
					}
					if nj > k {
						continue
					}
					if ok == 0 {
						if d > a[i] {
							continue
						}
						if d < a[i] {
							nok = 1
						}
					}
					dp[ni][nj][nok] += dp[i][j][ok]
				}
			}
		}
	}
	fmt.Println(dp[n][k][0] + dp[n][k][1])
}
