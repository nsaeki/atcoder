package main
import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	rd := bufio.NewReader(os.Stdin)
	s, _ := rd.ReadString('\n')
	n := len(s)
	m := make(map[int]int)
	r := 0
	k := 1
	ans := 0
	for i := n-1; i >= 0; i-- {
		r += int(s[i] - '0') * k
		r %= 2019
		if m[r] > 0 {
			ans += m[r]
		}
		m[r]++
		k *= 10
		k %= 2019

	}
	fmt.Println(ans)
}