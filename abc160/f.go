package main

import (
	"bufio"
	"fmt"
	"os"
)

type Combination struct {
	n int
	f *Factorial
}

func NewCombination(n, m int) *Combination {
	f := NewFactorial(n, m)
	return &Combination{n, f}
}

func (c *Combination) Combination(n, k int) int64 {
	return c.f.F[n] * c.f.Inv[k] % c.f.Mod * c.f.Inv[n-k] % c.f.Mod
}

func (c *Combination) Permutation(n, k int) int64 {
	return c.Combination(n+k-1, n-1)
}

func (c *Combination) Factorial(n int) int64 {
	return c.f.F[n]
}

type Factorial struct {
	n      int
	Mod    int64
	F, Inv []int64
}

func NewFactorial(n, m int) *Factorial {
	f := Factorial{n, int64(m), nil, nil}
	f.F = make([]int64, n+1)
	f.Inv = make([]int64, n+1)
	f.F[0] = 1
	for i := 1; i <= n; i++ {
		f.F[i] = f.F[i-1] * int64(i) % f.Mod
	}
	f.Inv[n] = pow(f.F[n], f.Mod-2, f.Mod)
	for i := n; i > 0; i-- {
		f.Inv[i-1] = f.Inv[i] * int64(i) % f.Mod
	}
	return &f
}

func pow(x, n, m int64) int64 {
	if n == 0 {
		return 1
	}
	if n&1 == 0 {
		r := pow(x, n/2, m)
		return r * r % m
	}
	return x * pow(x, n-1, m) % m
}

func div(x, n, m int64) int64 {
	return x * pow(n, m-2, m) % m
}

const m = 1000000007

var (
	e   [][]int
	dp  []int64
	cnt []int
	res []int64
	c   *Combination
)

func merge(i, j int) {
	dp[i] *= dp[j]
	dp[i] %= m
	dp[i] *= c.Combination(cnt[i]+cnt[j]+1, cnt[i])
	dp[i] %= m
	cnt[i] += cnt[j] + 1
}

func subtract(i, j int) {
	cnt[i] -= cnt[j] + 1
	dp[i] = div(dp[i], c.Combination(cnt[i]+cnt[j]+1, cnt[i]), m)
	dp[i] = div(dp[i], dp[j], m)
}

func dfs(i, p int) {
	dp[i] = int64(1)
	for _, j := range e[i] {
		if j == p {
			continue
		}
		dfs(j, i)
		merge(i, j)
	}
}

func dfs2(i, p int) {
	for _, j := range e[i] {
		if j == p {
			continue
		}
		tempDp, tempCnt := dp[i], cnt[i]
		subtract(i, j)
		merge(j, i)
		dp[i], cnt[i] = tempDp, tempCnt
		dfs2(j, i)
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)
	c = NewCombination(n, m)

	e = make([][]int, n+1)
	for i := 0; i < n-1; i++ {
		var a, b int
		fmt.Fscan(in, &a, &b)
		e[a] = append(e[a], b)
		e[b] = append(e[b], a)
	}

	cnt = make([]int, n+1)
	dp = make([]int64, n+1)
	dfs(1, 0)

	res = make([]int64, n+1)
	dfs2(1, 0)
	for i := 1; i <= n; i++ {
		fmt.Fprintln(out, dp[i])
	}
}
