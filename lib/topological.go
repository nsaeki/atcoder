import "fmt"

func TopologicalSort(edge [][]int) []int {
	n := len(edge)
	cnt := make([]int, n)
	queue := make([]int, 0, n)
	for i := 0; i < n; i++ {
		for _, t := range edge[i] {
			cnt[t]++
		}
	}
	fmt.Println(cnt)
	for i := 0; i < n; i++ {
		if cnt[i] == 0 {
			queue = append(queue, i)
		}
	}

	var i int
	for i < len(queue) {
		f := queue[i]
		i++
		for _, t := range edge[f] {
			cnt[t]--
			if cnt[t] == 0 {
				queue = append(queue, t)
			}
		}
	}
	fmt.Println(cnt)

	return queue
}