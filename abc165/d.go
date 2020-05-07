package main
import "fmt"

func main() {
	var a, b, n int64
	fmt.Scan(&a, &b, &n)

	var x int64
	if n < b {
		x = n
	} else {
		x = b-1
	}

	ans := a*x/b - a*(x/b)
	fmt.Println(ans)
}