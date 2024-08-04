// Sieve of Eratosthenes
package lib

type Sieve struct {
	n int
	p []int
}

func NewSieve(n int) *Sieve {
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
	return &Sieve{n, p}
}

func (s *Sieve) Primes() []int {
	ret := make([]int, 0, s.n)
	for i := 0; i < s.n; i++ {
		if s.p[i] == i {
			ret = append(ret, i)
		}
	}
	return ret
}

func (s *Sieve) Factors(n int) map[int]int {
	ret := make(map[int]int)
	for n > 1 {
		ret[s.p[n]]++
		n /= s.p[n]
	}
	return ret
}
