package main
import "fmt"
func main() {
	var a, b,c, d int
	fmt.Scan(&a, &b, &c, &d) 
	for a > 0 {
		c -= b
		if c <= 0 {
			fmt.Println("Yes")
			return
		}
		a -= d
	}
	fmt.Println("No")
}