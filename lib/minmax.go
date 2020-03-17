func min(a ...int) int {
	ret := a[0]
	for i := 1; i < len(a); i++ {
		if ret > a[i] {
			ret = a[i]
		}
	}
	return ret
}

func max(a ...int) int {
	ret := a[0]
	for i := 1; i < len(a); i++ {
		if ret < a[i] {
			ret = a[i]
		}
	}
	return ret
}

func minmax(a ...int) (int, int) {
	min, max := a[0], a[0]
	for i := 1; i < len(a); i++ {
		if min > a[i] {
			min = a[i]
		}
		if max < a[i] {
			max = a[i]
		}
	}
	return min, max
}