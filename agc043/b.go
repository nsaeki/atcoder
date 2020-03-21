package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	s, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(strings.TrimSpace(s))
	x := make([]int, n)
	for i := 0; i < n; i++ {
		b, _ := reader.ReadByte()
		x[i] = int(b - '0')
	}
	for len(x) > 1 {
		for i := 0; i < len(x)-1; i++ {
			y := x[i] - x[i+1]
			if y < 0 {
				y = -y
			}
			x[i] = y
		}
		x = x[:len(x)-1]
	}
	fmt.Println(x[0])
}
