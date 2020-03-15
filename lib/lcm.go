func divisors(n int) map[int]int {
	for i := 2; i <= n; i++ {
		if n%i == 0 {
			d := divisors(n / i)
			d[i]++
			return d
		}
	}
	return make(map[int]int)
}

func lcm(a, b int) int {
	if a > b {
		a, b = b, a
	}
	if b%a == 0 {
		return b
	}

	d := divisors(a)
	for k, v := range divisors(b) {
		if v > d[k] {
			d[k] = v
		}
	}

	lcm := 1
	for k, v := range d {
		for i := 0; i < v; i++ {
			lcm *= k
		}
	}
	return lcm
}