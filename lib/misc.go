func pow(p, n, m int) int {
	if n == 0 {
		return 1
	}

	r := pow(p, n/2, m) % m
	r *= r
	r %= m
	if n&1 == 1 {
		r *= p
		r %= m
	}
	return r
}

// Sieve of Eratosthenes
func primes(n int) []int {
	p := make([]int, n+1)
	for i := 2; i <= n; i++ {
		if p[i] == 0 {
			p[i] = i
		}
		if i > n/i {
			continue
		}
		for j := i * i; j <= n; j += i {
			if p[j] == 0 {
				p[j] = i
			}
		}
	}
	return p
}

func factorization(n int, p []int) []int {
	ret := []int{}
	for n > 1 {
		r = append(r, p[n])
		n /= p[n]
	}
	return ret
}