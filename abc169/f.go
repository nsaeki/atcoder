package main

import "fmt"

func main() {
	var n, s int
	fmt.Scan(&n, &s)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	var m int64 = 998244353
	dp := make([][]int64, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int64, s+1)
	}

	dp[0][0] = 1
	for i := 0; i < n; i++ {
		for j := 0; j <= s; j++ {
			if j+a[i] <= s {
				dp[i+1][j+a[i]] += dp[i][j]
				dp[i+1][j+a[i]] %= m
			}
			dp[i+1][j] += dp[i][j] * 2 % m
			dp[i+1][j] %= m
		}
	}
	fmt.Println(dp[n][s])
}
