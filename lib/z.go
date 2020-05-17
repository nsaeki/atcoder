// ZArray contains max length of prefix for s[0:n] and s[i:n]
func ZArray(s string) []int {
	n := len(s)
	z := make([]int, n)
	var l, r, k int
	for i := 0; i < len(s); i++ {
		if i > r {
			l = i
			r = i
			for r < n && s[r-l] == s[r] {
				r++
			}
			z[i] = r - l
			r--
		} else {
			k = i - l
			if z[k] < r-i+1 {
				z[i] = z[k]
			} else {
				l = i
				for r < n && s[r-l] == s[r] {
					r++
				}
				z[i] = r - l
				r--
			}
		}
	}
	return z
}

// ZSearch returns indices for text which has same prefix for pattern.
// sep must be a character which does not exist in text.
func ZSearch(text, pattern string, sep byte) []int {
	var ret []int
	z := ZArray(pattern + string(sep) + text)
	for i := 0; i < len(z); i++ {
		if z[i] == len(pattern) {
			ret = append(ret, i-len(pattern)-1)
		}
	}
	return ret
}
