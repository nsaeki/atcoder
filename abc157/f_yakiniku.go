package main

import (
	"fmt"
	"math"
)

const EPSILON float64 = 1e-9

func equals(a, b float64) bool {
	return (a-b) < EPSILON && (b-a) < EPSILON
}

type Vec struct {
	x, y float64
}

func vecAdd(a, b Vec) Vec {
	return Vec{a.x + b.x, a.y + b.y}
}

func vecSub(a, b Vec) Vec {
	return Vec{a.x - b.x, a.y - b.y}
}

func (v *Vec) Mul(float64 w) Vec {
	return Vec{v.x * w, v.y * w}
}

func (v *Vec) norm() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

func (v *Vec) rotate90() Vec {
	return Vec{y, -x}
}

type Circle struct {
	v Vec
	r float64
}

func xp(a, b Circle) []Vec {
	v := b.v - a.v
	l := norm(v)
	if equals(l, 0) {
		return []Vec{}
	}
	if equals(l+a.r+b.r, math.Max(l, a.r, b.r)*2) {
		return []Vec{vecAdd(a.v, v.Mul(r/l))}
	}
}

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	if k == 1 {
		fmt.Println(0)
		return
	}

	c := make([]Circle, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&c[i].v.x, &c[i].v.y, &c[i].r)
	}

	l := 0
	r := 400000
	for it := 0; it < 100; it++ {
		x := (1 + r) / 2
		var ps []Vec

		for i := 0; i < n; i++ {
			ps = append(ps, c[i].v)
		}
		for i := 0; i < n; i++ {
			for j := 0; j < i; j++ {
				a := c[i]
				b := c[j]
				a.r = x / a.r
				b.r = x / b.r
				ps = append(ps, xp(a, b)...)
			}
		}
	}
}
