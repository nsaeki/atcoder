type KMP struct {
	w []int
	t []int
}

func NewKMP(w []int) *KMP {
	var t = make([]int, len(w))
	t[0], t[1] = -1, 0
	i, j := 2, 0
	for i < len(w) {
		if w[i-1] == w[j] {
			t[i] = j + 1
			i++
			j++
		} else if j > 0 {
			j = t[j]
		} else {
			t[i] = 0
			i++
		}
	}
	return &KMP{w, t}
}

func (k *KMP) FindAll(s []int) []int {
	ret := []int{}
	var m, i int
	for m+i < len(s) {
		if k.w[i] == s[m+i] {
			i++
			if i == len(w) {
				ret = append(ret, m)
			}
		} else {
			m += i - t[i]
			if i > 0 {
				i = t[i]
			}
		}
	}
	return ret
}

func kmp(s, w []int) int {
	return NewKMP(w).search(s)
}
