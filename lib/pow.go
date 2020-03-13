func pow(p, n, m int) int {
	if n == 0 {
		return 1
	}
	r := pow(p, n/2, m) % m
	r = r * r % m
	if n&1 == 1 {
		r = r * p % m
	}
	return r
}

func div(p, n, m int) int {
	return p * pow(n, m-2, m) % m
}