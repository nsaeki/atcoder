var factorialMemo []int64

func factorial(n, m int) int64 {
	if n <= 1 {
		return 1
	}
	if factorialMemo[n] != 0 {
		return factorialMemo[n]
	}
	ret := int64(n) * factorial(n-1, m) % int64(m)
	factorialMemo[n] = ret
	return ret
}

func combination(n, k, m int) int64 {
	mm := int64(m)
	x := factorial(n, m) % mm
	y := factorial(k, m) % mm
	y = y * factorial(n-k, m) % mm
	return x * pow(y, m-2, m) % mm
}

func permutation(n, k, m int) int64 {
	return combination(n+k-1, n-1, m)
}

func pow(p, n, m int64) int64 {
	if n == 0 {
		return 1
	}
	if n&1 == 0 {
		r := pow(p, n/2, m)
		return r * r % int64(m)
	}
	return int64(p) * pow(p, n-1, m) % int64(m)
}
