import "fmt"

type KMP struct {
	w string
	t []int
}

func NewKMP(w string) *KMP {
	t := make([]int, len(w)+1)
	t[0] = -1
	i, j := 1, 0
	for i < len(w) {
		if w[i] == w[j] {
			t[i] = t[j]
		} else {
			t[i] = j
			j = t[j]
			for j >= 0 && w[i] != w[j] {
				j = t[j]
			}
		}
		i++
		j++
	}
	t[i] = j
	return &KMP{w, t}
}

func (k *KMP) Search(s string) []int {
	var ret []int
	var i, j int
	for j < len(s) {
		if k.w[i] == s[j] {
			i++
			j++
			if i == len(k.w) {
				// If only first occurrence is needed, return j-i
				ret = append(ret, j-i)
				i = k.t[i]
				fmt.Println(i, k.t)
			}
		} else {
			i = k.t[i]
			if i < 0 {
				i++
				j++
			}
		}
	}
	return ret
}

func KMPSearch(s, w string) []int {
	return NewKMP(w).Search(s)
}