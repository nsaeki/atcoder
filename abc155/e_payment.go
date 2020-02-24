package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	n := []int{}

	for i := 0; i < len(input)-1; i++ {
		n = append(n, int(input[i]-'0'))
	}
	// fmt.Println(n)

	res := 0
	carryOver := false

	for i := len(n) - 1; i >= 0; i-- {
		d := n[i]
		if carryOver {
			d++
		}
		carryOver = false

		if d > 5 || (d == 5 && i > 0 && n[i-1] > 4) {
			res += 10 - d
			carryOver = true
		} else {
			res += d
		}
	}
	if carryOver {
		res++
	}

	fmt.Println(res)
}
