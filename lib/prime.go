// Sieve of Eratosthenes
func primes(n int) []int {
	p := make([]int, n+1)
	for i := 2; i <= n; i++ {
		if p[i] != 0 {
			continue
		}
		p[i] = i
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

func factor(n int, p []int) map[int]int {
	ret := make(map[int]int)
	for n > 1 {
		ret[p[n]]++
		n /= p[n]
	}
	return ret
}