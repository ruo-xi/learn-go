package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

func euler(g complex128) {
	println(cmplx.Abs(g))
	println(cmplx.Exp(1i*math.Pi) + 1)
	println(cmplx.Pow(math.E, 1i*math.Pi) + 1)
}

func triangle() {
	var (
		a = 3
		b = 4
	)
	fmt.Println(calcTriangle(a, b))
}

func calcTriangle(a, b int) int {
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	return c
}

func testBuiltInVariable() {
	var (
		a bool       = true
		b string     = "b"
		c int        = 1
		d uint       = 1
		e int8       = 2
		f uint8      = 2
		h byte       = 1
		i rune       = '草'
		g complex128 = 3 + 4i
	)

	println(a, b, c, d, e, f, g, h, i)
	euler(g)
}

func main() {
	testBuiltInVariable()
}
