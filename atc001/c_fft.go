package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

func fft(x []complex128) {
	n := len(x)
	if n == 1 {
		return
	}

	x0 := make([]complex128, n/2)
	x1 := make([]complex128, n/2)
	for i := 0; i < n/2; i++ {
		x0[i] = x[2*i]
		x1[i] = x[2*i+1]
	}

	fft(x0)
	fft(x1)

	zeta := complex(math.Cos(2*math.Pi/float64(n)), math.Sin(2*math.Pi/float64(n)))
	powZeta := complex(1, 0)

	for i := 0; i < n; i++ {
		x[i] = x0[i%(n/2)] + powZeta*x1[i%(n/2)]
		powZeta *= zeta
	}
}

func ifft(x []complex128) {
	for i := 0; i < len(x); i++ {
		x[i] = cmplx.Conj(x[i])
	}
	fft(x)
	for i := 0; i < len(x); i++ {
		x[i] = cmplx.Conj(x[i]) / complex(float64(len(x)), 0)
	}
}

func main() {
	var n int
	fmt.Scan(&n)

	n2 := 1
	for n2 < 2*n+1 {
		n2 <<= 1
	}

	a := make([]complex128, n2)
	b := make([]complex128, n2)
	for i := 0; i < n; i++ {
		var ai, bi int
		fmt.Scan(&ai, &bi)
		a[i+1] = complex(float64(ai), 0)
		b[i+1] = complex(float64(bi), 0)
	}

	fft(a)
	fft(b)

	f := make([]complex128, n2)
	for i := 0; i < n2; i++ {
		f[i] = a[i] * b[i]
	}

	ifft(f)
	for i := 1; i < 2*n+1; i++ {
		fmt.Println(int(real(f[i]) + 0.5))
	}
}
